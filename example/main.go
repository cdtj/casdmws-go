package main

import (
	"fmt"
	"log"
	"os"

	casdm "github.com/cdtj/casdmws-go"
)

var mockDataSet map[casdm.Method][]casdm.MockData

func main() {
	f, err := os.OpenFile("./stdlog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	/* this is productive client */
	/*
		log.Println("\tProd Client")
		doLogin(&soap.Client{
			URL:       "https://servicedeskattachuat.datacom.co.nz/axis/services/USD_R11_WebService",
			Namespace: casdm.Namespace,
		}, "test", "test")
	*/

	/* here goes mock client */
	mc := &casdm.MockClient{Mocks: mockDataSet}
	log.Println("\tMock Client")

	// logging in
	loginResp, err := doLogin(mc, "admin", "admin")
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("Successful login, sid: ", *loginResp.LoginReturn)

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
	crResp, err := createCr(mc, loginResp.LoginReturn, creatorHandle, &newHandle, &newRefNum, attrValsArr, propValsArr, retValsArr)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Printf("New request [%s] created: %s\n", *crResp.NewRequestNumber, *crResp.NewRequestHandle)
}

func init() {
	var loginErrData = &casdm.ErrData{"soapenv:Client", "Error - invalid login password", "", "Error - invalid login password", "1000"}
	loginErr := casdm.ExtractError(loginErrData)

	mockDataSet = make(map[casdm.Method][]casdm.MockData)
	mockDataSet[casdm.NewMethod("Login")] = []casdm.MockData{
		casdm.MockData{
			Requests: []casdm.Request{
				casdm.Request{Attr: "Username", Value: "admin"},
				casdm.Request{Attr: "Password", Value: "admin"},
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
		casdm.MockData{
			Requests: []casdm.Request{},
			Response: casdm.Response{
				ResponseData: nil,
				Error:        fmt.Errorf("\"500 Internal Server Error\": \"%s\"", loginErr),
			},
		},
	}
	mockDataSet[casdm.NewMethod("CreateRequest")] = []casdm.MockData{
		casdm.MockData{
			Requests: []casdm.Request{
				casdm.Request{Attr: "category", Value: "pcat:1234"},
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
		casdm.MockData{
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

	resp, err := ws.CreateRequest(&casdm.CreateRequest{sid,
		&creatorHandle,
		attrVals,
		propVals,
		nil,
		retVals,
		newHandle,
		newRefNum,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
