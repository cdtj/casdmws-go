package main

import (
	"fmt"

	casdm "github.com/cdtj/casdmws-go"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Println("running with prodClient")
	pc := casdm.NewSoapClient("https://sdm/axis/services/USD_R11_WebService")
	if err := example(pc); err != nil {
		logrus.Errorln("prodClient failed with: ", err)
	}

	logrus.Println("running with mockClient")
	mc := casdm.NewMockClient(mockDataSet())
	if err := example(mc); err != nil {
		logrus.Errorln("mockClient failed with: ", err)
	}
}

func example(cli casdm.ClientInterface) error {
	// logging in
	loginResp, err := doLogin(cli, "admin", "admin")
	if err != nil {
		return fmt.Errorf("login error: %s", err)
	}
	logrus.Println("Successful login, sid: ", *loginResp.LoginReturn)

	// creating incident
	var newHandle, newRefNum string
	creatorHandle := "cnt:UUID_HERE"
	attrValsArr := []string{
		"requestor",
		"cnt:UUID_HERE",
		"description",
		"Hello World",
		"category",
		"pcat:12345",
	}
	propValsArr := []string{"prop1", "prop2"}
	retValsArr := []string{"persistent_id", "ref_num"}
	crResp, err := createCr(cli, loginResp.LoginReturn, creatorHandle, &newHandle, &newRefNum, attrValsArr, propValsArr, retValsArr)
	if err != nil {
		return fmt.Errorf("createCr error: %s", err)
	}
	logrus.Println("New request [%s] created: %s\n", *crResp.NewRequestNumber, *crResp.NewRequestHandle)
	return nil
}

func mockDataSet() casdm.MockDataSet {
	var loginErrData = &casdm.ErrData{
		FaultCode:    "soapenv:Client",
		FaultString:  "Error - invalid login password",
		FaultActor:   "",
		ErrorMessage: "Error - invalid login password",
		ErrorCode:    "1000",
	}
	loginErr := casdm.ExtractError(loginErrData)

	mds := make(casdm.MockDataSet)
	mds[casdm.NewMethod("Login")] = []casdm.MockData{
		{
			Requests: []casdm.Request{
				{Attr: "Username", Value: "admin"},
				{Attr: "Password", Value: "admin"},
			},
			Response: casdm.Response{
				ResponseData: casdm.OperationLoginResponse{
					LoginResponse: &casdm.LoginResponse{
						LoginReturn: casdm.NewInt(1337),
					},
				},
				Error: nil,
			},
		},
		{
			Requests: []casdm.Request{},
			Response: casdm.Response{
				ResponseData: nil,
				Error:        fmt.Errorf("\"500 Internal Server Error\": \"%s\"", loginErr),
			},
		},
	}
	mds[casdm.NewMethod("CreateRequest")] = []casdm.MockData{
		{
			Requests: []casdm.Request{
				{Attr: "category", Value: "pcat:1234"},
			},
			Response: casdm.Response{
				ResponseData: nil,
				Error: fmt.Errorf("\"500 Internal Server Error\": \"%s\"", casdm.ExtractError(&casdm.ErrData{
					FaultCode:    "soapenv:Client",
					FaultString:  "Error fetching: AHD03053: pcat:1234 not found",
					ErrorMessage: "Error fetching: AHD03053: pcat:1234 not found",
					ErrorCode:    "103",
				})),
			},
		},
		{
			Requests: []casdm.Request{},
			Response: casdm.Response{
				ResponseData: &casdm.OperationCreateRequestResponse{
					CreateRequestResponse: &casdm.CreateRequestResponse{
						CreateRequestReturn: casdm.NewString("hello"),
						NewRequestHandle:    casdm.NewString("cr:1001"),
						NewRequestNumber:    casdm.NewString("1001"),
					},
				},
				Error: nil,
			},
		},
	}
	return mds
}

func doLogin(cli casdm.ClientInterface, username, password string) (*casdm.LoginResponse, error) {
	ws := casdm.NewUSD_WebServiceSoap(cli)

	resp, err := ws.Login(&casdm.Login{
		Username: &username,
		Password: &password,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func createCr(cli casdm.ClientInterface, sid *int,
	creatorHandle string, newHandle, newRefNum *string,
	attrValsArr, propValsArr, retValsArr []string) (*casdm.CreateRequestResponse, error) {
	ws := casdm.NewUSD_WebServiceSoap(cli)

	attrVals := casdm.ArrToAOS(attrValsArr)
	propVals := casdm.ArrToAOS(propValsArr)
	retVals := casdm.ArrToAOS(retValsArr)

	resp, err := ws.CreateRequest(&casdm.CreateRequest{Sid: sid,
		CreatorHandle:    &creatorHandle,
		AttrVals:         attrVals,
		PropertyValues:   propVals,
		Template:         nil,
		Attributes:       retVals,
		NewRequestHandle: newHandle,
		NewRequestNumber: newRefNum,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
