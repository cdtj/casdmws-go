package rest

import (
	"os"

	casdmwsgo "github.com/cdtj/casdmws-go"
)

// SdmRest is a struct that implements Client for CA Service Desk REST API
type SdmRest struct {
	SdmRestInterface
}

// Create creates a new entry in the factory specified with writing result back in result structure with fields requested
//
//	func Create(c *RestClient) error {
//		data := &rest.SdmRestData{
//			"lrel_attachments_requests": &rest.SdmRestData{
//				"attmnt": &rest.SdmRestData{
//					"@REL_ATTR": *attmntid,
//				},
//				"cr": &rest.SdmRestData{
//					"@REL_ATTR": "cr:410206",
//				},
//			},
//		}
//		result := new(rest.ResultTypeDecoded)
//		err := c.Create("lrel_attachments_requests", nil, data, result); err != nil {
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print creation result
//		return nil
//	}
func (c *SdmRest) Create(factory string, fields []string, data *SdmRestData, result ResultType) error {
	return c.post(factory, fields, data, result)
}

// GetAll fetches all entries of selected factory with requested fields
//
//	func GetAll(c *RestClient) error {
//		shift := &ResultShift{Start: 1, Size: 10}
//		result := new(ResultDecoded)
//		err := c.GetAll("cr", shift, []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a some map that contains first 10 result of the query
//		return nil
//	}
func (c *SdmRest) GetAll(factory string, shift *ResultShift, returnFlds []string, result ResultType) error {
	return c.get(factory, "", shift.String(), returnFlds, result)
}

// GetByID fetches an entry of selected factory matched by id with requested fields
//
//	func GetByID(c *RestClient) error {
//		result := new(ResultDecoded)
//		err := c.GetByID("cr", NewSdmID(42), []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains ref_num and status of request with id 42
//		return nil
//	}
func (c *SdmRest) GetByID(factory string, id SdmIDType, returnFlds []string, result ResultType) error {
	return c.get(factory, id.ToID(), "", returnFlds, result)
}

// GetByPersid fetches an entry with id and factory extracted from persid with requested fields
//
//	func GetByPersid(c *RestClient) error {
//		result := new(ResultDecoded)
//		err := c.GetByPersid("cr:42", []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains ref_num and status of request with id 42
//		return nil
//	}
func (c *SdmRest) GetByPersid(persid string, returnFlds []string, result ResultType) error {
	p, err := NewSdmPersid(persid)
	if err != nil {
		return err
	}
	return c.get(p.factory, p.ToID(), "", returnFlds, result)
}

// GetByCommonName fetches entries of selected factory matched by their COMMON_NAMEs with requested fields.
// COMMON_NAME is not unique for most factories so shift might be helpuf.
//
//	func GetByCommonName(c *RestClient) error {
//		result := new(ResultDecoded)
//		err := c.GetByCommonName("cr", "R42", nil, []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains ref_num and status of request with ref_num R42
//		return nil
//	}
func (c *SdmRest) GetByCommonName(factory, name string, shift *ResultShift, returnFlds []string, result ResultType) error {
	return c.get(factory, casdmwsgo.Concat("-", "COMMON_NAME", name), shift.String(), returnFlds, result)
}

// GetByRelAttr fetches an entry of selected factory matched by REL_ATTR with requested fields.
// REL_ATTR is an external key for the factory. Can be obtained by `bop_sinfo -f <factory_name>`
//
//	func GetByRelAttr(c *RestClient) error {
//		result := new(ResultDecoded)
//		err := c.GetByRelAttr("cr", "42", []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains ref_num and status of request with id 42
//		return nil
//	}
func (c *SdmRest) GetByRelAttr(factory, val string, returnFlds []string, result ResultType) error {
	return c.get(factory, casdmwsgo.Concat("-", "REL_ATTR", val), "", returnFlds, result)
}

// GetByWhereClause fetches entries of selected factory matched by SQL-like search criteria with requested fields
//
//	func GetByWhereClause(c *RestClient) error {
//		result := new(ResultDecoded)
//		shift := &ResultShift{Start: 1, Size: 10}
//		err := c.GetByWhereClause("cr", "active = 1 OR id = 42", shift, []string{"ref_num", "status"}, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains ref_num and status of first 10 requests that are active (active = 1) OR with id = 42
//		return nil
//	}
func (c *SdmRest) GetByWhereClause(factory, wc string, shift *ResultShift, returnFlds []string, result ResultType) error {
	query := casdmwsgo.Concat("=", "WC", wc)
	if shift.String() != "" {
		query = casdmwsgo.Concat("&", query, shift.String())
	}
	return c.get(factory, "", query, returnFlds, result)
}

// GetColletction fetches collection (BLREL) from the specific entry with fields writing them to result
//
//	func GetColletction(c *RestClient) error {
//		result := new(ResultDecoded)
//		shift := &ResultShift{Start: 1, Size: 10}
//		err := c.GetColletction("cr", "42", "attachments", shift, SdmReturnAllFields, result)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("%v", result) // should print a map that contains all fields of first 10 attachments attached to request with id 42
//		return nil
//	}
func (c *SdmRest) GetColletction(factory, id, collection string, shift *ResultShift, returnFlds []string, result ResultType) error {
	return c.get(factory, casdmwsgo.Concat("/", id, collection), shift.String(), returnFlds, result)
}

// Update updates an entry with data that matches factory and id provided. Result will be written back in the result interface.
//
//	func Update(c *RestClient) error {
//		result := new(ResultDecoded)
//		data := make(SdmRestData)
//		data["summary"] = "hello"
//		data["call_back_date"] = SdmNull()
//		err := c.Update("cr", "42", SdmReturnAllFields, data, result)
//		if err != nil {
//			return err
//		}
//		// should print a map that contains all fields of request with id 42
//		// where 2 fields where updated the corresponding values
//		fmt.Printf("%v", result)
//		return nil
//	}
func (c *SdmRest) Update(factory, id string, returnFlds []string, data *SdmRestData, result ResultType) error {
	return c.put(factory, id, returnFlds, data, result)
}

// Delete deletes an entry within factory and id specified. Result will be written back in the result interface.
//
//	func Delete(c *RestClient) error {
//		result := new(ResultDecoded)
//		err := c.Delete("rest_access", "42", result) // deleting rest_access session with id 42
//		if err != nil {
//			return err
//		}
//		// result will be null on success
//		fmt.Printf("%v", result)
//		return nil
//	}
func (c *SdmRest) Delete(factory, id string, result ResultType) error {
	return c.delete(factory, id, result)
}

// Login logs in
func (c *SdmRest) Login() error {
	return c.login()
}

// Logout logs out
func (c *SdmRest) Logout() error {
	return c.logout()
}

// Upload uploads []byte as attachment
//
//	func Upload(c *RestClient) error {
//		// uploading plain text as attachment
//		attmntid, err := c.Upload([]byte("this is text"), "text.txt", "plain/text", "42", AttmntContentTypeBinary)
//		if err != nil {
//			return err
//		}
//		// result will contain attachment id
//		fmt.Printf("%v", *attmntid)
//		return nil
//	}
func (c *SdmRest) Upload(data []byte, fileName, mime, description string, content AttmntContentType) (*string, error) {
	return c.upload(data, fileName, mime, description, content)
}

// UploadFile uploads file by it's path
//
//	func UploadFile(c *RestClient) error {
//		// uploading plain txt file as attachment
//		attmntid, err := c.UploadFile("./text.txt", "text.txt", "42", AttmntContentTypeBinary)
//		if err != nil {
//			return err
//		}
//		// result will contain attachment id
//		fmt.Printf("%v", *attmntid)
//		return nil
//	}
func (c *SdmRest) UploadFile(path, fileName, description string, content AttmntContentType) (*string, error) {
	f, err := os.ReadFile(path)
	// mime := mime.TypeByExtension(path)
	if err != nil {
		return nil, err
	}
	return c.upload(f, fileName, "DOC", description, content)
}

// Download downloads attachment as []byte
//
//	func Download(c *RestClient) error {
//		// downloading attachment by id
//		b, err := c.Download("42")
//		if err != nil {
//			return err
//		}
//		// result will contain file content as []byte
//		fmt.Printf("%v", b)
//		return nil
//	}
func (c *SdmRest) Download(attmntID string) ([]byte, error) {
	return c.download(attmntID)
}

// Download downloads attachment and writes it to thepath
//
//	func DownloadFile(c *RestClient) error {
//		// downloading attachment by id and writing to current dir
//		b, err := c.DownloadFile("42", "./myfile.dat")
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func (c *SdmRest) DownloadFile(attmntID string, path string) error {
	b, err := c.download(attmntID)
	if err != nil {
		return err
	}
	os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
