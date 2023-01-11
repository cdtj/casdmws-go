package rest

import (
	"net/http"
)

type SdmRestInterface interface {
	get(factory, filter, param string, returnFlds []string, result ResultType) error
	post(factory string, returnFlds []string, data *SdmRestData, result ResultType) error
	put(factory, id string, returnFlds []string, data *SdmRestData, result ResultType) error
	delete(factory, id string, result ResultType) error
	// attachments related methods
	upload(data []byte, fileName, mime, description string, content AttmntContentType) (*string, error)
	download(attmntID string) ([]byte, error)
	// it feels like a bad idea to wrap those over interface
	// since login and logout are implemented as POST and DELETE.
	login() error
	logout() error
}

// NewSdmRest is a quick definition of a proper CA Service Desk REST WebServices Client
// with default implementation.
func NewSdmRest(restURI, repoID, serverName string, schema SdmRestSchema, auth Login, httpCli *http.Client) *SdmRest {
	return NewSdmRestWithCli(newClient(restURI, repoID, serverName, schema, auth, httpCli))
}

// NewSdmRestWithCli is a quick definition of a proper CA Service Desk REST WebServices Client
// with custom implementation.
func NewSdmRestWithCli(cli SdmRestInterface) *SdmRest {
	return &SdmRest{cli}
}
