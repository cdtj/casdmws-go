package rest

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	casdmwsgo "github.com/cdtj/casdmws-go"
	"github.com/sirupsen/logrus"
)

// client is the http.client with fields used by CA Service Desk REST API.
type client struct {
	*http.Client
	restURI    string
	repoID     string
	serverName string
	schema     SdmRestSchema
	auth       Login
}

// newClient is inetrnal method to spawn a client struct
func newClient(restURI, repoID, serverName string, schema SdmRestSchema, auth Login, httpCli *http.Client) *client {
	if httpCli == nil {
		httpCli = http.DefaultClient
	}
	return &client{
		Client:     httpCli,
		restURI:    strings.TrimSuffix(restURI, "/"),
		repoID:     repoID,
		serverName: serverName,
		schema:     schema,
		auth:       auth,
	}
}

// GetFilter is enum to switch between supported filters
//
// Deprecated
type GetFilter uint8

// Deprecated
const (
	GetFilterAll GetFilter = iota
	GetFilterWC
	GetFilterID
	GetFilterCommonName
	GetFilterRelAttr
	GetFilterCollection
	GetFilterUndefined
)

// get is internal method to perform all supported fetches via GET request
func (c *client) get(factory, filter, param string, returnFlds []string, result ResultType) error {
	url := factory
	if filter != "" {
		url = casdmwsgo.Concat("/", url, filter)
	}
	if param != "" {
		url = casdmwsgo.Concat("?", url, param)
	}
	r, err := c.reqPrepare(http.MethodGet, url, returnFlds, nil)
	if err != nil {
		return err
	}
	return c.reqDo(r, result)
}

// post is an internal method to perform a POST request
func (c *client) post(factory string, fields []string, data *SdmRestData, result ResultType) error {
	r, err := c.reqPrepare(http.MethodPost, factory, fields, data)
	if err != nil {
		return err
	}
	if err := c.reqNullifyFields(r, data); err != nil {
		return err
	}
	return c.reqDo(r, result)
}

// put is internal method to perform a PUT request
func (c *client) put(factory, id string, returnFlds []string, data *SdmRestData, result ResultType) error {
	r, err := c.reqPrepare(http.MethodPost, factory, returnFlds, data)
	if err != nil {
		return err
	}
	if err := c.reqNullifyFields(r, data); err != nil {
		return err
	}
	return c.reqDo(r, result)
}

// put is internal method to perform a PUT request
func (c *client) delete(factory, id string, result ResultType) error {
	r, err := c.reqPrepare(http.MethodDelete, strings.Join([]string{factory, id}, "/"), nil, nil)
	if err != nil {
		return err
	}
	return c.reqDo(r, result)
}

func (c *client) login() error {
	return c.auth.Login(c)
}

func (c *client) logout() error {
	return c.auth.Logout(c)
}

const attachmentBoundary = "*****MessageBoundary*****"

// upload uploads []byte as Service Desk attachment with options specififed.
func (c *client) upload(data []byte, fileName, mime, description string, content AttmntContentType) (*string, error) {
	logrus.WithFields(logrus.Fields{
		"method":   casdmwsgo.Concat(">", "Rest", "Upload"),
		"fileName": fileName,
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"method":   casdmwsgo.Concat(">", "Rest", "Upload"),
			"fileName": fileName,
			"duration": time.Since(dt),
		}).Trace("end")
	}()
	url := casdmwsgo.Concat("?",
		"/attmnt/",
		casdmwsgo.Concat("&",
			casdmwsgo.Concat("=", "repositoryId", c.repoID),
			casdmwsgo.Concat("=", "serverName", c.serverName),
			casdmwsgo.Concat("=", "mimeType", mime),
			casdmwsgo.Concat("=", "description", fileName),
		))
	buf, err := c.reqPrepareUpload(fileName, description, content, data)
	if err != nil {
		return nil, err
	}
	r, err := c.reqPrepare(http.MethodPost, url, nil, buf)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "multipart/form-data; BOUNDARY="+attachmentBoundary)

	result := new(ResultTypeAttmnt)
	if err := c.reqDo(r, result); err != nil {
		return nil, err
	}
	str := string(*result)
	return &str, nil
}

// download downloads an attachment by attachment id.
func (c *client) download(attmntID string) ([]byte, error) {
	url := casdmwsgo.Concat("/", "attmnt", attmntID, "file-resource")
	result := new(ResultTypeRaw)
	if err := c.get(url, "", "", nil, result); err != nil {
		return nil, err
	}
	return *result, nil
}

// Attmnt is a struct used as attachment payload during upload.
type Attmnt struct {
	XMLName     xml.Name    `xml:"attmnt" json:"-"`
	Repo        *AttmntRepo `json:"repository" xml:"repository"`
	FileName    string      `json:"orig_file_name" xml:"orig_file_name"`
	AttmntName  string      `json:"attmnt_name" xml:"attmnt_name"`
	Description string      `json:"description" xml:"description"`
}

// Attmnt is a struct used as attachment payload during upload.
type AttmntRepo struct {
	Id string `json:"@id" xml:"id,attr"`
}

// AttmntReq is a struct used as attachment payload during upload.
type AttmntReq struct {
	XMLName xml.Name `xml:"attmnt" json:"-"`
	Attmnt  *Attmnt  `json:"attmnt"`
}

// AttmntContentType is a flag to define file encdoing.
type AttmntContentType uint8

const (
	AttmntContentTypeBinary AttmntContentType = iota // for most cases
	AttmntContentTypeBase64                          // when data is encoded in Base64
)

func (c *client) reqPrepareUpload(fileName, description string, content AttmntContentType, data []byte) (*bytes.Buffer, error) {
	attmntPayload := &AttmntReq{
		Attmnt: &Attmnt{
			Repo: &AttmntRepo{
				Id: c.repoID,
			},
			FileName:    fileName,
			AttmntName:  fileName,
			Description: description,
		},
	}
	body := new(bytes.Buffer)
	charSet := "UTF-8"
	transferEnc := "base64"
	if content != AttmntContentTypeBase64 {
		transferEnc = "binary"
	}
	CR := "\n\r"

	// this is probably a bug but when using JSON schema,
	// SDM throws HTTP 415 (Content Type Unsupported)
	xmlSchema := SdmRestSchemaXML
	xmlPayload, err := xmlSchema.Marshal(attmntPayload.Attmnt)
	if err != nil {
		return nil, err
	}

	body.Write([]byte("--" + attachmentBoundary + CR))
	body.Write([]byte("Content-Disposition: form-data; name=\"payload\"" + CR))
	body.Write([]byte("Content-Type: " + string(xmlSchema) + "; CHARACTERSET=" + charSet + CR))
	body.Write([]byte(CR))
	body.Write(xmlPayload)
	// Depending w3.org there is no need in 2 CRs before boundary,
	// but they were in a Java SDM Sample so I'm keeping them.
	// body.Write([]byte(CR))
	// body.Write([]byte(CR))
	body.Write([]byte("--" + attachmentBoundary + CR))
	body.Write([]byte(
		"Content-Disposition: form-data; name=\"" +
			fileName + "\"; filename=\"" +
			fileName + "\"" + CR))
	body.Write([]byte("Content-Type: application/octet-stream" + CR))
	body.Write([]byte("Content-Transfer-Encoding: " + transferEnc + CR))
	// IMPORTANT: using CR here (as requested mentioned by docs) causes empty line which crashes bin files,
	// so \r is omited.
	body.Write([]byte("\n"))
	if content == AttmntContentTypeBase64 {
		body.Write(data[strings.IndexByte(string(data), ',')+1:])
	} else {
		body.Write(data)
	}
	// Java SDK sample says: CRLF is important! It indicates end of binary boundary.
	// but I think it's not.
	// body.Write([]byte(CR))
	// body.Write([]byte(CR))
	// End of multipart/form-data.
	body.Write([]byte("--" + attachmentBoundary + "--"))
	// body.Write([]byte(CR))

	return body, nil
}

// reqPrepare is a base method to create a new [http.Request] with needed preparations
func (c *client) reqPrepare(method, url string, returnFlds []string, data interface{}) (*http.Request, error) {
	var buf *bytes.Buffer
	switch v := data.(type) {
	case *bytes.Buffer:
		buf = v
	default:
		bodyData, err := c.schema.Marshal(data)
		if err != nil {
			return nil, err
		}
		logrus.WithFields(logrus.Fields{
			"method":   casdmwsgo.Concat(">", "Rest", "reqPrepare"),
			"bodyData": string(bodyData),
		}).Trace("preparing request")
		buf = bytes.NewBuffer(bodyData)
	}
	// url must start with "/"
	if !strings.HasPrefix(url, "/") {
		url = casdmwsgo.Concat("", "/", url)
	}

	r, err := http.NewRequest(method,
		casdmwsgo.Concat("", c.restURI, url),
		buf,
	)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=utf-8", c.schema))
	r.Header.Add("Accept", string(c.schema))
	r.Header.Add("Content-Length", strconv.Itoa(buf.Len()))
	r.Header.Add("Host", r.Host)
	if err := c.reqReturnFields(r, returnFlds); err != nil {
		return nil, err
	}
	// Requestes are the same for all methods but rest_access doesn't need
	// an authorization so this comparsion is the easiest way to implement
	// everything in a signle function.
	if url != "/rest_access" || method != http.MethodPost {
		if err := c.reqAuth(r, url); err != nil {
			return nil, err
		}
	}
	logrus.WithFields(logrus.Fields{
		"method":         casdmwsgo.Concat(">", "Rest", "reqPrepare"),
		"HTTP":           r.Method,
		"Content-Length": r.Header.Get("Content-Length"),
		"Auth":           r.Header.Get("Authorization"),
		"XKey":           r.Header.Get("X-AccessKey"),
		"URL":            url,
	}).Trace("headers")
	return r, nil
}

// reqDo is a base method to proccess with preparted [http.Request]
func (c *client) reqDo(r *http.Request, result ResultType) error {
	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	logrus.WithFields(logrus.Fields{
		"method":     casdmwsgo.Concat(">", "Rest", "reqDo"),
		"statusCode": resp.StatusCode,
		"URL":        r.URL,
	}).Trace("response")

	if resp.StatusCode >= 400 {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"method": casdmwsgo.Concat(">", "Rest", "reqDo"),
				"data":   string(data),
				"err":    err,
			}).Error("readAll failed")
			return &RestError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "reqDo"),
			"data":   string(data),
			"resp":   resp.Status,
			"code":   resp.StatusCode,
		}).Error("bad response")
		switch resp.StatusCode {
		case http.StatusBadRequest:
			return &RestError{
				Code:    http.StatusBadRequest,
				Message: string(data),
			}
		case http.StatusUnauthorized:
			return &RestError{
				Code:    http.StatusUnauthorized,
				Message: string(data),
			}
		case http.StatusNotFound:
			return &RestError{
				Code:    http.StatusNotFound,
				Message: string(data),
			}
		case http.StatusConflict:
			return &RestError{
				Code:    http.StatusConflict,
				Message: string(data),
			}
		case http.StatusInternalServerError:
			return &RestError{
				Code:    http.StatusInternalServerError,
				Message: string(data),
			}
		case http.StatusUnsupportedMediaType:
			return &RestError{
				Code:    http.StatusUnsupportedMediaType,
				Message: string(data),
			}
		}
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		switch v := result.(type) {
		case *ResultTypeAttmnt:
			loc, err := resp.Location()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": casdmwsgo.Concat(">", "Rest", "reqDo"),
					"err":    err,
				}).Error("respone location error")
				return err
			}
			locArr := strings.Split(loc.Path, "/")
			persid := locArr[len(locArr)-1]
			if len(locArr) == 0 {
				return ErrRestUploadFailed
			}
			*v = ResultTypeAttmnt(persid)
		case *ResultTypeRaw:
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"method": casdmwsgo.Concat(">", "Rest", "reqDo"),
					"err":    err,
				}).Error("read raw body error")
				return err
			}
			*v = ResultTypeRaw(data)
		default:
			if err = c.schema.Decode(resp.Body, result); err != nil {
				return err
			}
		}
	}
	return nil
}

// reqNullifyFields moves NULL values from input data to a specific array in the header,
// this is how nullifying in CA SDM REST API works.
func (c *client) reqNullifyFields(r *http.Request, data *SdmRestData) error {
	nulls := make([]string, 0)
	for k, v := range *data {
		switch v.(type) {
		case *sdmNull:
			delete(*data, k)
			nulls = append(nulls, k)
		}
	}
	if len(nulls) > 0 {
		r.Header.Add("X-AttrsToNull", strings.Join(nulls, ","))
	}
	return nil
}

// reqReturnFields sets an array of fields that need to be returned in response.
func (c *client) reqReturnFields(r *http.Request, fields []string) error {
	if len(fields) > 0 {
		r.Header.Set("X-Obj-Attrs", strings.Join(fields, ","))
	}
	return nil
}

// reqAuth authenticates the request by adding authentication token to the header,
// this method does all needed calls to obtain the correct auth token.
func (c *client) reqAuth(r *http.Request, url string) error {
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Rest", "reqAuth"),
		"login":  fmt.Sprintf("%v", c.auth),
		"url":    url,
	}).Trace("starting > preparing token")
	err := c.auth.Token(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": casdmwsgo.Concat(">", "Rest", "reqAuth"),
			"err":    err,
		}).Trace("error > preparing token")
		if errors.Is(err, ErrRestNotLoggedIn) || errors.Is(err, ErrRestExpired) {
			if err := c.auth.Login(c); err != nil {
				return err
			}
			return c.reqAuth(r, url)
		} else {
			return err
		}
	}
	logrus.WithFields(logrus.Fields{
		"method": casdmwsgo.Concat(">", "Rest", "reqAuth"),
		"login":  fmt.Sprintf("%v", c.auth),
	}).Trace("ending > preparing token")
	return nil
}
