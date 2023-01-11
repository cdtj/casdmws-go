package main

import (
	"fmt"

	casdmwsgo "github.com/cdtj/casdmws-go"
	"github.com/cdtj/casdmws-go/soap"
	"github.com/sirupsen/logrus"
)

// SoapSample also implements auth struct to ease session control.
type SoapSample struct {
	*soap.SdmSoap
	auth soap.LoginIntfc
}

// SdUUID is a test contact UUID from our test env.
const SdUUID = "cnt:B7E3AB990A140C4F8E00338C18C67C8A"

func (s *SoapSample) run() error {
	// Due to low TTL of sid we're obtaining sid manually.
	sid, err := s.LoginCall(s.auth)
	if err != nil {
		if serr := soap.ParseSoapError(err); serr != nil {
			return fmt.Errorf("login error: %s", serr)
		}
		return fmt.Errorf("login error: %s", err)
	}
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Soap", "Run", "LoginCall"),
		"sid":    *sid,
	}).Println("ok")

	// attrVals is array of string where odd:even values are key:value pair,
	// due to SDM SOAP API all values should be strings.
	attrValsArr := []string{
		"type", "R",
		"status", "OP",
		"customer", SdUUID,
		"description", "Hello World",
		"category", "pcat:400001",
	}
	// propValsArr is array of string with property values, properties might be declared by current
	// and/or previous category (depending on prop options), proprty values are stored as string.
	// Props might have validation rule which strictly restricts incoming values by filters.
	propValsArr := []string{"prop1", "prop2"}
	// retValsArr is array of string with attribute names which will be returned as part of response
	// in CreateRequestReturn *string.
	// For example we can return calculated fields if any, like group after autoassignment.
	retValsArr := []string{"persistent_id", "ref_num"}
	crResp, err := s.soapCreateCr(sid, soap.StringPtr(SdUUID), attrValsArr, propValsArr, retValsArr)
	if err != nil {
		if serr := soap.ParseSoapError(err); serr != nil {
			return serr
		}
		return err
	}
	// Response also has 2 fields with ref_num and hanlde, that are enough in most cases.
	logrus.WithFields(logrus.Fields{
		"NewRequestNumber": *crResp.NewRequestNumber,
		"NewRequestHandle": *crResp.NewRequestHandle,
	}).Printf("New request created")
	return nil
}

func (s *SoapSample) soapCreateCr(sid *int,
	creatorHandle *string,
	attrValsArr, propValsArr, retValsArr []string) (*soap.CreateRequestResponse, error) {

	attrVals := soap.ArrToAOS(attrValsArr)
	propVals := soap.ArrToAOS(propValsArr)
	retVals := soap.ArrToAOS(retValsArr)

	resp, err := s.CreateRequest(&soap.CreateRequest{Sid: sid,
		CreatorHandle:    creatorHandle,
		AttrVals:         attrVals,
		PropertyValues:   propVals,
		Template:         soap.StringPtr(""),
		Attributes:       retVals,
		NewRequestHandle: soap.StringPtr(""),
		NewRequestNumber: soap.StringPtr(""),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
