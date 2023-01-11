package soap

import (
	"github.com/fiorix/wsdl2go/soap"
)

// SdmSoapInterface implements [soap.Clinet] interface
type SdmSoapInterface interface {
	RoundTrip(in, out soap.Message) error
	RoundTripWithAction(soapAction string, in, out soap.Message) error
	RoundTripSoap12(action string, in, out soap.Message) error
}

// NewSdmSoap is a quick definition of a proper CA Service Desk SOAP WebServices Client
// with default implementation.
func NewSdmSoap(url string) *SdmSoap {
	return NewSdmSoapWithCli(&soap.Client{
		URL:       url,
		Namespace: Namespace,
	})
}

// NewSdmRestWithCli is a quick definition of a proper CA Service Desk SOAP WebServices Client
// with custom implementation.
func NewSdmSoapWithCli(cli SdmSoapInterface) *SdmSoap {
	return &SdmSoap{cli}
}
