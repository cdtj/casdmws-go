package rest

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"hash"
	"net/http"
	"net/url"
	"strconv"
	"time"

	casdmwsgo "github.com/cdtj/casdmws-go"
	"github.com/sirupsen/logrus"
)

// Login interface implements an authentication struct
// that attaches to the REST Client and authorizes it.
//
// This file is part of implementation of standart REST Cli implementation
// and relay on it's methods.
type Login interface {
	// Login is unified login method which reduces need to implement
	// preparation for each method manually.
	//
	// You can pass nil Login to use data from the [Client] struct
	Login(*client) error
	// Logout is unified logout method which reduces need to implement
	// preparation for each method manually.
	//
	// You can pass nil Login to use data from the [Client] struct
	Logout(*client) error
	// Token returns Token that is used as Authorization HTTP Request Header.
	//
	// The reason for calling this method manually might be ONLY if
	// you're using a custom implementation of the Login method.
	Token(*http.Request) error
	// Session returns CA Service Desk REST session stored inside the sturct,
	//
	// Can be used in save/load routintes if you're restarting the
	// app more often than session TTL.
	//
	// The reason for calling this method manually might be ONLY if
	// you're using a custom implementation of the Login method.
	Session() *Access
	// UpdSession updates CA Service Desk REST session stored inside the sturct,
	//
	// Can be used in save/load routintes if you're restarting the
	// app more often than session TTL.
	//
	// The reason for calling this method manually might be ONLY if
	// you're using a custom implementation of the Login method.
	UpdSession(*Access)
}

// Access stores CA Service Desk REST session details
type Access struct {
	id             string
	accessKey      int
	expirationDate time.Time
	secretKey      string
}

// NewAccess creates an instance of Access struct.
func NewAccess(id string, accessKey int, expirationDate int64, secretKey string) *Access {
	return &Access{
		id:             id,
		accessKey:      accessKey,
		expirationDate: time.Unix(expirationDate, 0),
		secretKey:      secretKey,
	}
}

// AccessJSON matches SDM REST rest_access response body
// when Accept-Content is specified as application/json.
type AccessJSON struct {
	AccessDataJSON struct {
		ID             string `json:"@id"`
		AccessKey      int    `json:"access_key"`
		ExpirationDate int64  `json:"expiration_date"`
		SecretKey      string `json:"secret_key"`
	} `json:"rest_access"`
}

// Access method converts JSON struct to unified Access struct.
func (b *AccessJSON) Access() *Access {
	return NewAccess(b.AccessDataJSON.ID, b.AccessDataJSON.AccessKey, b.AccessDataJSON.ExpirationDate, b.AccessDataJSON.SecretKey)
}

// AccessXML matches SDM REST rest_access response body
// when Accept-Content is specified as application/xml.
type AccessXML struct {
	ID             string `xml:"id,attr"`
	AccessKey      int    `xml:"access_key"`
	ExpirationDate int64  `xml:"expiration_date"`
	SecretKey      string `xml:"secret_key"`
}

// Access method converts XML struct to unified Access struct.
func (b *AccessXML) Access() *Access {
	return NewAccess(b.ID, b.AccessKey, b.ExpirationDate, b.SecretKey)
}

// LoginRequestBody matches SDM REST rest_access request body.
type LoginRequestBody struct {
	XMLName xml.Name `xml:"rest_access" json:"-"`
	Access  *string  `json:"rest_access"`
}

type loginUserPassword struct {
	username   string
	password   string
	secret     bool
	signFields []string
	hmacAlg    *HmacAlgorithm

	// storing session data
	data *Access
}

// NewLoginUserPassword is a built-in function to create an instance of
// Login compatible struct. With authorization using username and password.
func NewLoginUserPassword(username, password string) *loginUserPassword {
	return &loginUserPassword{
		username:   username,
		password:   password,
		secret:     false,
		signFields: nil,
		data:       nil,
		hmacAlg:    nil,
	}
}

func (b *loginUserPassword) Login(c *client) error {
	logrus.WithFields(logrus.Fields{
		"loginUserPassword": b,
		"method":            casdmwsgo.Concat(">", "Rest", "Login"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginUserPassword": b,
			"method":            casdmwsgo.Concat(">", "Rest", "Login"),
			"duration":          time.Since(dt),
		}).Trace("end")
	}()
	r, err := c.reqPrepare(http.MethodPost, "rest_access", b.signFields, new(LoginRequestBody))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"method": casdmwsgo.Concat(">", "Rest", "Login", "reqPrepare"),
		}).Debug("method failed")
		return err
	}
	authToken := base64.StdEncoding.EncodeToString([]byte(b.username + ":" + b.password))
	if b.secret {
		authToken = "SDM " + authToken
	} else {
		authToken = "Basic " + authToken
	}
	r.Header.Set("Authorization", authToken)
	logrus.WithFields(logrus.Fields{
		"authorization": authToken,
		"method":        casdmwsgo.Concat(">", "Rest", "Login"),
	}).Trace("token")

	var decodedBody interface{}
	switch c.schema {
	case SdmRestSchemaJSON:
		decodedBody = new(AccessJSON)
	case SdmRestSchemaXML:
		decodedBody = new(AccessXML)
	}
	if err := c.reqDo(r, decodedBody); err != nil {
		return err
	}
	switch v := decodedBody.(type) {
	case *AccessJSON:
		logrus.WithFields(logrus.Fields{
			"restAccess": v.Access(),
			"method":     casdmwsgo.Concat(">", "Rest", "Login"),
		}).Trace("result")
		b.data = v.Access()
	case *AccessXML:
		logrus.WithFields(logrus.Fields{
			"restAccess": v.Access(),
			"method":     casdmwsgo.Concat(">", "Rest", "Login"),
		}).Trace("result")
		b.data = v.Access()
	default:
		return ErrRestUnathorized
	}

	return nil
}

// Logout deletes REST session on CA Service Desk host.
func (b *loginUserPassword) Logout(c *client) error {
	logrus.WithFields(logrus.Fields{
		"loginUserPassword": b,
		"method":            casdmwsgo.Concat(">", "Rest", "Logout"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginUserPassword": b,
			"method":            casdmwsgo.Concat(">", "Rest", "Logout"),
			"duration":          time.Since(dt),
		}).Trace("end")
	}()
	if err := c.delete("rest_access", b.data.id, nil); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"method": casdmwsgo.Concat(">", "Rest", "Logout", "delete"),
		}).Debug("method failed")
		return err
	}
	b.UpdSession(nil)
	return nil
}

func (b *loginUserPassword) Session() *Access {
	return b.data
}

func (b *loginUserPassword) UpdSession(data *Access) {
	b.data = data
}

func (b *loginUserPassword) Token(r *http.Request) error {
	logrus.WithFields(logrus.Fields{
		"loginUserPassword": b,
		"method":            casdmwsgo.Concat(">", "Rest", "Token"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginUserPassword": b,
			"method":            casdmwsgo.Concat(">", "Rest", "Token"),
			"duration":          time.Since(dt),
		}).Trace("end")
	}()
	if b.data == nil {
		return ErrRestNotLoggedIn
	}
	if b.data.expirationDate.Before(time.Now()) {
		return ErrRestExpired
	}
	// return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(b.username+":"+b.password))), nil
	r.Header.Add("X-AccessKey", strconv.Itoa(b.data.accessKey))
	return nil
}

type loginSecretKey struct {
	*loginUserPassword
}

type HmacAlgorithm uint

// Enum of HMAC Algorithms supported by CA SDM
const (
	HmacSHA1 HmacAlgorithm = iota
	HmacSHA256
	HmacSHA384
	HmacSHA512
	HmacMD5
)

// NewLoginSecretKey is a built-in function to create an instance of
// RestLogin compatible struct. With authorization by CA SDM username, password, and encrypted key.
//
// Proper CA SDM Options MUST be specified.
//  1. signFields refersh to Admin Tab > Options > WebService > string_to_sign_fields
//     If string_to_sign_fields is uninstalled the nil value should be passed as signFields.
//     Array order must be same as arguments in the option comma separated.
//  2. hmacAlg refers to Admin Tab > Options > WebService > hmac_algorithm
//
// Read more: [REST Secret Key Authentication]
// Postman Sample: [TEC10660]
//
// [REST Secret Key Authentication]: https://techdocs.broadcom.com/us/en/ca-enterprise-software/business-management/ca-service-management/17-3/reference/ca-service-desk-manager-reference-commands/technical-reference/rest-http-methods.html#concept.dita_571265def7bea5a79bc6e3bba4cd061d4d03b036_RESTSecretKeyAuthentication
// [TEC10660]: https://knowledge.broadcom.com/external/article/10660/how-to-use-secret-key-authentication-wit.html
func NewLoginSecretKey(username, password string, signFields []string, hmacAlg HmacAlgorithm) *loginSecretKey {
	return &loginSecretKey{
		loginUserPassword: &loginUserPassword{
			username:   username,
			password:   password,
			secret:     true,
			data:       nil,
			signFields: signFields,
			hmacAlg:    &hmacAlg,
		},
	}
}

// Login logins in to specified Client with credentials
// stored in the struct.
// The login result will be stored as part of the struct and
// can be obtained by calling Session()
func (b *loginSecretKey) Login(c *client) error {
	logrus.WithFields(logrus.Fields{
		"loginSecretKey": b,
		"method":         casdmwsgo.Concat(">", "Rest", "Login"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginSecretKey": b,
			"method":         casdmwsgo.Concat(">", "Rest", "Login"),
			"duration":       time.Since(dt),
		}).Trace("end")
	}()
	return b.loginUserPassword.Login(c)
}

// Logout deletes REST session on CA Service Desk host.
func (b *loginSecretKey) Logout(c *client) error {
	logrus.WithFields(logrus.Fields{
		"loginSecretKey": b,
		"method":         casdmwsgo.Concat(">", "Rest", "Logout"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginSecretKey": b,
			"method":         casdmwsgo.Concat(">", "Rest", "Logout"),
			"duration":       time.Since(dt),
		}).Trace("end")
	}()
	if err := c.delete("rest_access", b.data.id, nil); err != nil {
		return err
	}
	b.UpdSession(nil)
	return nil
}

// Session returns SDM REST session stored inside the sturct.
//
// The reason for calling this method manually might be ONLY if
// you're using a custom implementation of the Login method.
func (b *loginSecretKey) Session() *Access {
	return b.data
}

// UpdSession updates session data in the current Login,
// can be used in save/load routintes if you're restarting the
// app more often than session TTL.
//
// The reason for calling this method manually might be ONLY if
// you're using a custom implementation of the Login method.
func (b *loginSecretKey) UpdSession(data *Access) {
	b.data = data
}

// Token returns Token that is used as Authorization HTTP Request Header.
//
// The reason for calling this method manually might be ONLY if
// you're using a custom implementation of the Login method.
func (b *loginSecretKey) Token(r *http.Request) error {
	logrus.WithFields(logrus.Fields{
		"loginSecretKey": b,
		"method":         casdmwsgo.Concat(">", "Rest", "Token"),
	}).Trace("token")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"loginSecretKey": b,
			"method":         casdmwsgo.Concat(">", "Rest", "Token"),
			"duration":       time.Since(dt),
		}).Trace("token")
	}()
	if b.data == nil {
		return ErrRestNotLoggedIn
	}
	if b.data.expirationDate.Before(time.Now()) {
		return ErrRestExpired
	}
	url := r.URL.Path
	query := r.URL.RawQuery
	if query != "" {
		url = casdmwsgo.Concat("?", url, query)
	}
	r.Header.Add("Authorization", b.signRequest(fmt.Sprintf("%s\n%s", r.Method, url)))
	return nil
}

func (b *loginSecretKey) signRequest(input string) string {
	logrus.WithFields(logrus.Fields{
		"secretKey": b.data.secretKey,
		"input":     input,
		"algorithm": *b.hmacAlg,
		"accessKey": b.data.accessKey,
		"method":    casdmwsgo.Concat(">", "Rest", "signRequest"),
	}).Trace("start")
	dt := time.Now()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"secretKey": b.data.secretKey,
			"input":     input,
			"algorithm": *b.hmacAlg,
			"accessKey": b.data.accessKey,
			"method":    casdmwsgo.Concat(">", "Rest", "signRequest"),
			"duration":  time.Since(dt),
		}).Trace("end")
	}()
	key := []byte(b.data.secretKey)
	h := hmac.New(b.hmacFunc(), key)
	h.Write([]byte(input))
	return fmt.Sprintf("SDM %d:%s", b.data.accessKey, url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil))))
}

func (b *loginSecretKey) hmacFunc() (h func() hash.Hash) {
	if b.hmacAlg == nil {
		return sha1.New
	}
	switch *b.hmacAlg {
	case HmacSHA1:
		return sha1.New
	case HmacSHA256:
		return sha256.New
	case HmacSHA384:
		return sha512.New384
	case HmacSHA512:
		return sha512.New
	case HmacMD5:
		return md5.New
	default:
		return sha1.New
	}
}
