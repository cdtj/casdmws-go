package main

import (
	"github.com/cdtj/casdmws-go/soap"
	soapcli "github.com/fiorix/wsdl2go/soap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SoapMock is a sample CA Service Desk SOAP WebServices mock client where we are:
//   - login using username, password, and soap policy
//
// and:
//   - creating a request with type "I" (Incident)
//   - throwing an error on creation beacuse of incorrect category
func SoapMock() {
	logrus.Println("mock soap cli")
	mc := SoapSample{
		soap.NewSdmSoapWithCli(soap.NewMockClient(soapMockData())),
		soap.NewLoginUserPassword(viper.GetString("soap.username"),
			viper.GetString("soap.password"),
			viper.GetString("soap.policy"),
		),
	}
	if err := mc.run(); err != nil {
		logrus.Errorln("mock cli failed with: ", err)
	}
}

func soapMockData() soap.MockDataSet {
	var loginErrData = &soap.SoapError{
		FaultCode:   "soapenv:Client",
		FaultString: "Error - invalid login password",
		FaultActor:  "",
		ErrorMsg:    "Error - invalid login password",
		ErrorCode:   1000,
	}
	loginErr := soap.ExtractError(loginErrData)
	soapCliErr := &soapcli.HTTPError{
		StatusCode: 500,
		Status:     "",
		Msg:        loginErr,
	}

	mds := make(soap.MockDataSet)
	mds[soap.NewMethod("Login")] = []soap.MockData{
		{
			Requests: []soap.Request{
				{Attr: "Username", Value: "admin"},
				{Attr: "Password", Value: "admin"},
			},
			Response: soap.Response{
				ResponseData: soap.OperationLoginResponse{
					LoginResponse: &soap.LoginResponse{
						LoginReturn: soap.IntPtr(1337),
					},
				},
				Error: nil,
			},
		},
		{
			Requests: []soap.Request{},
			Response: soap.Response{
				ResponseData: nil,
				Error:        soapCliErr,
			},
		},
	}
	mds[soap.NewMethod("LoginService")] = []soap.MockData{
		{
			Requests: []soap.Request{
				{Attr: "Username", Value: "admin"},
				{Attr: "Password", Value: "admin"},
				{Attr: "Policy", Value: "Default"},
			},
			Response: soap.Response{
				ResponseData: soap.OperationLoginServiceResponse{
					LoginServiceResponse: &soap.LoginServiceResponse{
						LoginServiceReturn: soap.IntPtr(1337),
					},
				},
				Error: nil,
			},
		},
		{
			Requests: []soap.Request{},
			Response: soap.Response{
				ResponseData: nil,
				Error:        soapCliErr,
			},
		},
	}
	mds[soap.NewMethod("CreateRequest")] = []soap.MockData{
		{
			Requests: []soap.Request{
				{Attr: "category", Value: "pcat:400001"},
			},
			Response: soap.Response{
				ResponseData: nil,
				Error: &soap.SoapError{
					FaultCode:   "soapenv:Client",
					FaultString: "Error setting attr 'category' on object 'cr:1001' to value 'pcat:400001' NOT FOUND pcat:400001",
					ErrorMsg:    "Error setting attr 'category' on object 'cr:1001' to value 'pcat:400001' NOT FOUND pcat:400001",
					ErrorCode:   1003,
				},
			},
		},
		{
			Requests: []soap.Request{},
			Response: soap.Response{
				ResponseData: &soap.OperationCreateRequestResponse{
					CreateRequestResponse: &soap.CreateRequestResponse{
						CreateRequestReturn: soap.StringPtr("hello"),
						NewRequestHandle:    soap.StringPtr("cr:1001"),
						NewRequestNumber:    soap.StringPtr("1001"),
					},
				},
				Error: nil,
			},
		},
	}
	return mds
}
