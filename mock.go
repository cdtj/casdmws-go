package casdmwsgo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/beevik/etree"
	"github.com/fiorix/wsdl2go/soap"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
)

// ErrData is a struct for mocking errors with a template
type ErrData struct {
	FaultCode    string
	FaultString  string
	FaultActor   string
	ErrorMessage string
	ErrorCode    string
}

const errTpl = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<soapenv:Body>
		<soapenv:Fault>
			<faultcode>{{.FaultCode}}</faultcode>
			<faultstring>{{.FaultString}}</faultstring>
			<faultactor>{{.FaultActor}}</faultactor>
			<detail>
				<ErrorMessage>{{.ErrorMessage}}</ErrorMessage>
				<ErrorCode>{{.ErrorCode}}</ErrorCode>
			</detail>
		</soapenv:Fault>
	</soapenv:Body>
</soapenv:Envelope>`

// Method should match mocked method.
// All lower case
type Method string

type MockDataSet map[Method][]MockData

// MockClient is a fake SOAP client.
type MockClient struct {
	Mocks MockDataSet
}

// NewMockClient is a quick definition of a proper Mocked USD WebServices ClientInterface
func NewMockClient(mocks MockDataSet) ClientInterface {
	return &MockClient{
		Mocks: mocks,
	}
}

// MockData is a set of data that be validated,
// if every element in Requests match incoming request,
// Response part would be processed.
//
// Empty Requests array always trigger Response return,
// use it as switch-like "default" case.
type MockData struct {
	Requests []Request
	Response
}

// Request contains key:value pair
// where key is presented as request attribute
type Request struct {
	Attr  string
	Value string
}

// Response should contain Operation level of
// SOAP response struct in ResponseData,
// ex: OperationLoginResponse{LoginResponse:{LoginReturn:{data}}}
type Response struct {
	ResponseData soap.Message
	Error        error
}

// FakeRoundTrip is endconing input request, comparing with mock data
// and returning decoded response back
func FakeRoundTrip(c *MockClient, action Method, in, out soap.Message) error {
	b, err := encodeIn(in)
	if err != nil {
		return fmt.Errorf(ExtractError(&ErrData{
			ErrorMessage: err.Error(),
		}))
	}
	logrus.Debugln(string(b))

	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		return fmt.Errorf(ExtractError(&ErrData{
			ErrorMessage: err.Error(),
		}))
	}

	root := doc.SelectElement("Envelope")
	logrus.Debugln("ROOT element:", root.Tag)

	if mocks, found := c.Mocks[action]; found {
		for _, m := range mocks {
			if len(m.Requests) == 0 {
				logrus.Debugln("Empty request data. Setting default response.")
				decodeOut(out, m.Response.ResponseData)
				return m.Response.Error
			}

			var wholeMatch bool
			for _, r := range m.Requests {
				var match matchStatus
				wholeMatch = false
				lAttr := strings.ToLower(r.Attr)
				printEl(root, &lAttr, &r.Value, &match)
				switch match {
				case matchStatusNotFound:
					break
				case matchStatusFound:
					wholeMatch = true
					break
				}
				wholeMatch = true
			}
			if wholeMatch {
				logrus.Debugln("Whole request match. Setting response.")
				err := decodeOut(out, m.Response.ResponseData)
				if err != nil {
					return fmt.Errorf(ExtractError(&ErrData{
						ErrorMessage: err.Error(),
					}))
				}
				return m.Response.Error
			}
		}
		return fmt.Errorf(ExtractError(&ErrData{
			ErrorMessage: fmt.Sprintf("'%s' has no matching mocks", action),
		}))
	}
	return fmt.Errorf(ExtractError(&ErrData{
		ErrorMessage: fmt.Sprintf("'%s' is not specified in Mocks map", action),
	}))
}

func encodeIn(msg soap.Message) ([]byte, error) {
	req := &soap.Envelope{
		XSIAttr: soap.XSINamespace,
		Body:    msg,
	}
	var b bytes.Buffer
	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		logrus.Debugln("err:", err)
		return nil, err
	}
	return b.Bytes(), nil
}

func decodeOut(msg soap.Message, msgData soap.Message) error {
	rb, err := encodeIn(msgData)
	if err != nil {
		return err
	}

	marshalStructure := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    soap.Message
	}{Body: msg}

	decoder := xml.NewDecoder(bytes.NewBuffer(rb))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&marshalStructure)
	if err != nil {
		return err
	}
	return nil
}

type matchStatus int

const (
	matchStatusNotFound matchStatus = 0
	matchStatusFound    matchStatus = 1
	matchStatusNextVal  matchStatus = 2
)

func printEl(el *etree.Element, attr, val *string, match *matchStatus) bool {
	if len(el.ChildElements()) == 0 {
		logrus.Debugf("checking [%s / %s]", *attr, *val)
		logrus.Debugln(el.GetPath(), "[", el.Tag, "]:", el.Text())
		if strings.Contains(el.GetPath(), "/string") {
			logrus.Debugln("\tnext")
			// logrus.Debugln("\tnextval:", *attr, "@", *val)
			if el.Text() == *attr {
				*match = matchStatusNextVal
				return false
			}
			if *match == matchStatusNextVal {
				if el.Text() == *val {
					logrus.Debugln("\tmatches")
					// logrus.Debugln("\tmatched:", *attr, "@", *val)
					*match = matchStatusFound
					return true
				}
			}
			return false
		}
		if strings.Contains(el.GetPath(), *attr) {
			logrus.Debugln("\tcontains")
			// logrus.Debugln("\tcontains:", el.GetPath(), "@", *val)
			if el.Text() == *val {
				logrus.Debugln("\tmatches")
				// logrus.Debugln("\tmatched:", *attr, "@", *val)
				*match = matchStatusFound
				return true
			}
		} else {
			logrus.Debugln("\tnot contains:", el.GetPath(), "@", *val)
		}
	}
	for _, elm := range el.ChildElements() {
		if printEl(elm, attr, val, match) {
			logrus.Debugln("breaking loop for:", *attr, "@", *val)
			return true
		}
	}
	return false
}

// RoundTrip is required to match client interface
func (c *MockClient) RoundTrip(in, out soap.Message) error {
	return FakeRoundTrip(c, NewMethod(""), in, out)
}

// RoundTripSoap12 is required to match client interface
func (c *MockClient) RoundTripSoap12(action string, in, out soap.Message) error {
	return FakeRoundTrip(c, NewMethod(action), in, out)
}

// RoundTripWithAction is required to match client interface
func (c *MockClient) RoundTripWithAction(soapAction string, in, out soap.Message) error {
	return FakeRoundTrip(c, NewMethod(soapAction), in, out)
}

// ExtractError is extracting Error struct to SDM SOAP-like template
func ExtractError(errData *ErrData) string {
	tmpl, err := template.New("tmpl").Parse(errTpl)

	if err != nil {
		// this should not fail, if yes just quit
		log.Fatal("err:", err)
	}
	return processErrorTemplate(tmpl, errData)
}

func processErrorTemplate(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		// this should not fail, if yes just quit
		log.Fatal("err:", err)
	}
	return tmplBytes.String()
}

// NewMethod converts string to method
// by lowering case, nothing special
func NewMethod(name string) Method {
	return Method(strings.ToLower(name))
}
