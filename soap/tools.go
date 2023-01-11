package soap

import (
	"encoding/xml"
	"fmt"

	"github.com/fiorix/wsdl2go/soap"
)

// IntPtr is a wrapper to create inline pointer to int
func IntPtr(n int) *int {
	return &n
}

// StringPtr is a wrapper to create inline pointer to string
func StringPtr(s string) *string {
	return &s
}

// ArrToAOS is a wrapper to create inline pointer to Array of String
func ArrToAOS(arr []string) *ArrayOfString {
	arrToPTR := make([]*string, 0, len(arr))
	for k := range arr {
		arrToPTR = append(arrToPTR, &arr[k])
	}
	return &ArrayOfString{
		String: arrToPTR,
	}
}

// ArrToAOI is a wrapper to create inline pointer to Array of Int
func ArrToAOI(arr []int) *ArrayOfInt {
	arrToPTR := make([]*int, 0, len(arr))
	for k := range arr {
		arrToPTR = append(arrToPTR, &arr[k])
	}
	return &ArrayOfInt{
		Integer: arrToPTR,
	}
}

// SoapError used for unmarshaling Service Desk SOAP API errors
type SoapError struct {
	XMLName     xml.Name `xml:"Envelope"`
	FaultCode   string   `xml:"Body>Fault>faultcode"`
	FaultString string   `xml:"Body>Fault>faultstring"`
	FaultActor  string   `xml:"Body>Fault>faultactor"`
	ErrorCode   int      `xml:"Body>Fault>detail>ErrorCode"`
	ErrorMsg    string   `xml:"Body>Fault>detail>ErrorMessage"`
}

func (e *SoapError) Error() string {
	return fmt.Sprintf("%d: %s", e.ErrorCode, e.ErrorMsg)
}

func ParseSoapError(err error) *SoapError {
	switch terr := err.(type) {
	case *soap.HTTPError:
		var soapErr SoapError
		err := xml.Unmarshal([]byte(terr.Msg), &soapErr)
		if err == nil {
			return &soapErr
		}
	}
	return nil
}

// USDObjectList is for
type USDObjectList struct {
	XMLName   xml.Name    `xml:"UDSObjectList" json:"xml_name"`
	UDSObject []USDObject `xml:"UDSObject" json:"uds_object"`
}

// USDObject is for
type USDObject struct {
	Handle     string           `xml:"Handle" json:"persistent_id"`
	Attributes []USDObjectAttrs `xml:"Attributes>Attribute" json:"attr"`
}

// USDObjectAttrs is for
type USDObjectAttrs struct {
	AttrName  string `xml:"AttrName" json:"a"`
	AttrValue string `xml:"AttrValue" json:"v"`
}

func UmarshalError() {}
