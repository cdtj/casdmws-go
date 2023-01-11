package main

import (
	"fmt"

	casdmwsgo "github.com/cdtj/casdmws-go"
	"github.com/cdtj/casdmws-go/rest"
	"github.com/sirupsen/logrus"
)

type RestSample struct {
	*rest.SdmRest
}

// CrsResp matches REST API response
type CrsResp struct {
	Crs struct {
		CommonName  string `json:"@COMMON_NAME"`
		RelAttr     string `json:"@REL_ATTR"`
		ID          string `json:"@id"`
		Active      int    `json:"active"`
		Code        string `json:"code"`
		CrFlag      int    `json:"cr_flag"`
		InFlag      int    `json:"in_flag"`
		PrFlag      int    `json:"pr_flag"`
		Hold        int    `json:"hold"`
		Description string `json:"description"`
		Resolved    int    `json:"resolved"`
		Sym         string `json:"sym"`
	} `json:"crs"`
}

func (c *RestSample) run() error {
	// Logout on any result.
	defer func() {
		if err := c.Logout(); err != nil {
			// Just log an error, nothig serious if we won't know that logout was unsuccessfull.
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": casdmwsgo.Concat(">", "Rest", "Run", "Logout"),
					"err":    err,
				}).Error("method failed")
			}
		}
	}()

	// Declaring an empty struct of crs.
	result1 := new(CrsResp)
	// Fetching crs (call req status) by REL_ATTR,
	// 	- crs' REL_ATTR refers to code attribute, you can always check w/ SDM schema using bop_sinfo,
	// 		eg: `bop_sinfo -a crs` // will show all attributes of crs
	//		eg: `bop_sinfo -f crs` // will show all factory-related stuff of crs
	//  - OP is a code for Open status
	//  - [rest.SdmReturnAllFields] - fetching all fields
	//  - and writing them to our result struct pointer.
	if err := c.GetByRelAttr("crs", "OP", rest.SdmReturnAllFields, result1); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "Run", "GetByRelAttr"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method":  casdmwsgo.Concat(">", "Rest", "Run", "GetByRelAttr"),
		"result1": fmt.Sprintf("%#v", *result1),
	}).Println("ok")

	// [rest.ResultTypeDecoded] means that the result will be decode into the map.
	result2 := new(rest.ResultTypeDecoded)
	if err := c.GetByRelAttr("crs", "OP", rest.SdmReturnAllFields, result2); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "Run", "GetByRelAttr"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Rest", "Run", "GetByRelAttr"),
		"result": fmt.Sprintf("%#v", *result2),
	}).Println("ok")

	return nil
}

func (c *RestSample) upload() error {
	// Logout on any result.
	defer func() {
		if err := c.Logout(); err != nil {
			// Just log an error, nothig serious if we won't know that logout was unsuccessfull.
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": casdmwsgo.Concat(">", "Rest", "Upload", "Logout"),
					"err":    err,
				}).Error("method failed")
			}
		}
	}()

	// Uploading upload.png from the current folder as Sample.png with description,
	// filesystem files are binary but there is an option to upload Base64 encoded files too.
	attmntid, err := c.UploadFile("./upload.png", "Sample.png", "Some Valuable Description", rest.AttmntContentTypeBinary)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "Upload", "UploadFile"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method":    casdmwsgo.Concat(">", "Rest", "Upload", "UploadFile"),
		"attmnt ID": *attmntid,
	}).Println("ok")

	// Check [rest.SdmRestData] for more details about how data should be defined.
	data := &rest.SdmRestData{
		"lrel_attachments_requests": &rest.SdmRestData{
			"attmnt": &rest.SdmRestData{
				"@REL_ATTR": *attmntid,
			},
			"cr": &rest.SdmRestData{
				"@REL_ATTR": "cr:410206", // replace with <cr:id> from ur test env
			},
		},
	}

	// [rest.ResultTypeDecoded] means that result will be decode into the map
	result := new(rest.ResultTypeDecoded)
	// Creating a new entry that will store relation between newly uploaded attachment
	// and existing request in the system.
	if err := c.Create("lrel_attachments_requests", nil, data, result); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "Upload", "Create"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Rest", "Upload", "Create"),
		"result": *result,
	}).Println("ok")
	return nil
}

func (c *RestSample) uploadAndDownload() error {
	// Logout on any result.
	defer func() {
		if err := c.Logout(); err != nil {
			// Just log an error, nothig serious if we won't know that logout was unsuccessfull.
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "Logout"),
					"err":    err,
				}).Error("method failed")
			}
		}
	}()

	// Uploading string wrapped into a byte array as secret.txt with a description,
	// if a attachment is encoded in Base64 it can upload as-is with [rest.AttmntContentTypeBase64] flag.
	attmntid, err := c.Upload([]byte("we are super secret"), "secret.txt", "text/plain", "Some Valuable Description", rest.AttmntContentTypeBinary)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "UploadFile"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method":    casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "UploadFile"),
		"attmnt ID": *attmntid,
	}).Println("ok")

	// Check [rest.SdmRestData] for more details about how data should be defined,
	// in case of CA SDM Security logic, only linked to an object attachment can be
	// downloaded.
	data := &rest.SdmRestData{
		"lrel_attachments_requests": &rest.SdmRestData{
			"attmnt": &rest.SdmRestData{
				"@REL_ATTR": *attmntid,
			},
			"cr": &rest.SdmRestData{
				"@REL_ATTR": "cr:410206",
			},
		},
	}
	// [rest.ResultTypeDecoded] means that result will be decode into the map
	result := new(rest.ResultTypeDecoded)
	// Creating a new entry that will store relation between newly uploaded attachment
	// and existing request in the system.
	// In case of CA SDM Security logic, only linked to an object attachment can be
	// downloaded.
	if err := c.Create("lrel_attachments_requests", nil, data, result); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "Create"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "Create"),
		"result": *result,
	}).Println("ok")
	// Downloading an attachment to secret_file.txt in the current folder.
	if err := c.DownloadFile(*attmntid, "./secret_file.txt"); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "UploadAndDownload", "DownloadFile"),
			"err":    err,
		}).Error("method failed")
		return err
	}
	return nil
}
