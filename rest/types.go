package rest

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"

	casdmwsgo "github.com/cdtj/casdmws-go"
	"github.com/sirupsen/logrus"
)

// SdmRestData is a common interface for passing any attributes to SDM API,
// following rules must be met:
//   - factory name must be always on top,
//   - SDM objects must have @REL_ATTR with Related Attr as value,
//     for example:
//     -- use persistent_id for cr,
//     -- UUID for cnt,
//     -- id for attmnt,
//     -- and code for status,
//   - plain objects can just key-value pair.
//
// Example:
//
//	data := &SdmRestData{
//		"grpmem": &SdmRestData{
//			"member": &SdmRestData{
//				"@REL_ATTR": NewSdmUUID("<UUID_HERE>"),
//			},
//			"group": &SdmRestData{
//				"@REL_ATTR": NewSdmUUID("<UUID_HERE>"),
//			},
//			"manager": false,
//		},
//	}
type SdmRestData map[string]interface{}

// SdmRestSchema defenies content type of request/response.
type SdmRestSchema string

var (
	SdmRestSchemaJSON SdmRestSchema = "application/json" // to send and retrive data as JSON
	SdmRestSchemaXML  SdmRestSchema = "application/xml"  // to send and retrive data as XML
)

// NewSchema parses schema from string.
func NewSchema(s string) SdmRestSchema {
	switch strings.ToUpper(s) {
	case strings.ToUpper(string(SdmRestSchemaJSON)), "JSON":
		return SdmRestSchemaJSON
	case strings.ToUpper(string(SdmRestSchemaXML)), "XML":
		return SdmRestSchemaXML
	}
	return ""
}

// Marshal marshalises struct into []byte with method depending on schema.
func (s SdmRestSchema) Marshal(v interface{}) ([]byte, error) {
	switch s {
	case SdmRestSchemaJSON:
		return json.Marshal(v)
	case SdmRestSchemaXML:
		return xml.Marshal(v)
	default:
		return nil, errors.New("schema not implemented:" + string(s))
	}
}

// Decode decodes content into struct with method depending on schema.
func (s SdmRestSchema) Decode(r io.Reader, v interface{}) error {
	switch s {
	case SdmRestSchemaJSON:
		dec := json.NewDecoder(r)
		if err := dec.Decode(&v); err != nil {
			logrus.WithFields(logrus.Fields{
				"method": "reqDo",
				"err":    err,
			}).Error("json decode failed")
			return err
		}
		logrus.Debugf("v: %v\n", v)
		return nil
	case SdmRestSchemaXML:
		dec := xml.NewDecoder(r)
		if err := dec.Decode(&v); err != nil {
			logrus.WithFields(logrus.Fields{
				"method": "reqDo",
				"err":    err,
			}).Error("xml decode failed")
			return err
		}
		logrus.Debugf("v: %v\n", v)
		return nil
	default:
		return errors.New("schema not implemented:" + string(s))
	}
}

// SdmIDType is a data type that implements logic of CA Service Desk API acceptable IDs.
type SdmIDType interface {
	ToID() string
}

// NewSdmIDType is a function that sould convert incoming interface into a proper SDM ID depending on it's param.
func NewSdmIDType(intfc interface{}) SdmIDType {
	switch v := intfc.(type) {
	case SdmUUID:
		return v
	case string:
		if len(v) == 32 {
			return NewSdmUUID(v)
		}
		return NewSdmID(v)
	default:
		return NewSdmID(v)
	}
}

// SdmUUID is a CA Service Desk UUID type.
type SdmUUID [16]byte

// String returns SdmUUID as string.
func (id SdmUUID) String() string {
	return strings.ToUpper(hex.EncodeToString(id[:]))
}

// ToID returns SdmUUID as SDM ID compatible string.
func (id SdmUUID) ToID() string {
	return fmt.Sprintf("U'%s'", id.String())
}

// NewSdmUUID create a new SdmUUID from string.
func NewSdmUUID(s string) (id SdmUUID) {
	log.Println("s:", s)
	if len(s) == 32 {
		hex, err := hex.DecodeString(s)
		if err != nil {
			return
		}
		for i := 0; i < len(hex) || i < 16; i++ {
			id[i] = hex[i]
		}
	}
	return
}

// SdmID is a basic SDM positive number id.
type SdmID uint

// NewSdmID is a function that sould convert incoming interface into a proper SDM ID depending on it's param.
func NewSdmID(intfc interface{}) SdmID {
	switch v := intfc.(type) {
	case string:
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil || i > math.MaxUint {
			return 0
		}
		return SdmID(i)
	case int, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return NewSdmID(fmt.Sprintf("%d", v))
	default:
		return 0
	}
}

// ToID returns SdmID as SDM ID compatible string.
func (id SdmID) ToID() string {
	return strconv.FormatUint(uint64(id), 10)
}

// SdmPersid is a CA Service Desk unique identifier which represents concatenation of table (factory) and id.
type SdmPersid struct {
	factory string
	id      SdmIDType
}

// NewSdmPersid returns SdmPersid parsed from the passed string.
func NewSdmPersid(s string) (*SdmPersid, error) {
	arr := strings.Split(s, ":")
	if len(arr) == 2 {
		return &SdmPersid{
			factory: arr[0],
			id:      NewSdmIDType(arr[1]),
		}, nil
	}
	return nil, ErrInvalidPersid
}

// ToID returns SdmPersid as SDM ID compatible string.
func (id *SdmPersid) ToID() string {
	return id.id.ToID()
}

// sdmNull is used as NULL value in CA SDM data
type sdmNull struct{}

// SdmNull is used as NULL value in CA SDM data
func SdmNull() *sdmNull {
	return new(sdmNull)
}

// SdmReturnAllFields pass this as fields parametr to
// fetch all attributes of returned entity.
var SdmReturnAllFields = []string{"*"}

// ResultShift is a struct to pass param to fetches
// to perform shifting resuls.
//
// Start < 1 will be 1.
type ResultShift struct {
	Start int
	Size  int
}

// Next increases start pos with size, the method can be called in a loop to fetch the next results.
// Note: CA Service Desk REST API doesn't throw an exception when start or size reaches out of boundaries.
func (s *ResultShift) Next() *ResultShift {
	s.Start = s.Start + s.Size
	return s
}

// String converts ResultShift in URL compatible strng.
func (s *ResultShift) String() string {
	if s == nil {
		return ""
	}
	return casdmwsgo.Concat("&", casdmwsgo.Concat("=", "start", strconv.Itoa(s.Start)), casdmwsgo.Concat("=", "size", strconv.Itoa(s.Size)))
}

// ResultType is an interface to be passed to response writer
type ResultType interface{}

// ResultTypeRaw POINTER should be passed as request param to get response as-is
type ResultTypeRaw []byte

// ResultTypeAttmnt POINTER should be passed as request param to get attachment id as response
type ResultTypeAttmnt string

// ResultTypeDecoded POINTER should be passed as request param to get response decoded into map[string]interface{}
type ResultTypeDecoded map[string]interface{}
