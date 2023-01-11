package soap

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

const errTpl = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<soapenv:Body>
		<soapenv:Fault>
			<faultcode>{{.FaultCode}}</faultcode>
			<faultstring>{{.FaultString}}</faultstring>
			<faultactor>{{.FaultActor}}</faultactor>
			<detail>
				<ErrorMessage>{{.ErrorMsg}}</ErrorMessage>
				<ErrorCode>{{.ErrorCode}}</ErrorCode>
			</detail>
		</soapenv:Fault>
	</soapenv:Body>
</soapenv:Envelope>`

// Method should match mocked method.
// All lowercase.
type Method string

// MockDataSet must contain all mocket method with their data.
type MockDataSet map[Method][]MockData

// MockClient is a mock SOAP client.
type MockClient struct {
	Mocks MockDataSet
}

// NewMockClient is a quick definition of a proper Mocked USD WebServices ClientInterface
func NewMockClient(mocks MockDataSet) SdmSoapInterface {
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
		if serr := ParseSoapError(err); serr != nil {
			return serr
		}
		return err
	}
	logrus.WithFields(logrus.Fields{
		"action": action,
		"in":     string(b),
	}).Debugln("FakeRoundTrip")

	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		if serr := ParseSoapError(err); serr != nil {
			return serr
		}
		return err
	}

	root := doc.SelectElement("Envelope")
	logrus.WithFields(logrus.Fields{
		"action": action,
		"root":   root.Tag,
	}).Debugln("FakeRoundTrip")
	if mocks, found := c.Mocks[action]; found {
		for _, m := range mocks {
			if len(m.Requests) == 0 {
				logrus.WithFields(logrus.Fields{
					"action": action,
					"result": "Empty request data. Setting default response.",
				}).Debugln("FakeRoundTrip")
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
				logrus.WithFields(logrus.Fields{
					"action": action,
					"result": "Whole request match. Setting response.",
				}).Debugln("FakeRoundTrip")
				err := decodeOut(out, m.Response.ResponseData)
				if err != nil {
					if serr := ParseSoapError(err); serr != nil {
						return serr
					}
					return err
				}
				return m.Response.Error
			}
		}
		return fmt.Errorf("'%s' has no matching mocks", action)
	}
	return fmt.Errorf("'%s' is not specified in Mocks map", action)
}

func encodeIn(msg soap.Message) ([]byte, error) {
	req := &soap.Envelope{
		XSIAttr: soap.XSINamespace,
		Body:    msg,
	}
	var b bytes.Buffer
	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": msg,
			"err": err,
		}).Error("encodeIn failed")
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
		logrus.WithFields(logrus.Fields{
			"msg": msg,
			"err": err,
		}).Error("decodeOut failed")
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
		logrus.WithFields(logrus.Fields{
			"attr":  *attr,
			"val":   *val,
			"match": *match,
			"path":  el.GetPath(),
			"tag":   el.Tag,
			"text":  el.Text(),
		}).Trace("printEl")

		if strings.Contains(el.GetPath(), "/string") {
			if el.Text() == *attr {
				*match = matchStatusNextVal
				return false
			}
			if *match == matchStatusNextVal {
				if el.Text() == *val {
					*match = matchStatusFound
					return true
				}
			}
			return false
		}
		if strings.Contains(el.GetPath(), *attr) {
			if el.Text() == *val {
				*match = matchStatusFound
				return true
			}
		}
	}
	for _, elm := range el.ChildElements() {
		if printEl(elm, attr, val, match) {
			logrus.WithFields(logrus.Fields{
				"attr":  *attr,
				"val":   *val,
				"match": *match,
			}).Trace("breaking loop")
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
//
// Deprecated.
func ExtractError(errData *SoapError) string {
	tmpl, err := template.New("tmpl").Parse(errTpl)
	if err != nil {
		// this should not fail, if yes just quit
		log.Fatal("err:", err)
	}
	return processErrorTemplate(tmpl, errData)
}

// Deprecated.
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
