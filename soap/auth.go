package soap

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"os"
	"strconv"

	"golang.org/x/crypto/pkcs12"
)

// LoginIntfc is an interface for unified login structs
type LoginIntfc interface {
	// Login returns sid (session id) in case of successfull login.
	// SID is the first argument in the most of WSDL methods.
	Login(SoapWebService) (*int, error)
}

type loginUserPassword struct {
	username string
	password string
	policy   string
}

// NewLoginUserPassword is a built-in function to create an instance of
// SoapLogin compatible struct. With authorization by CA SDM username, password, and policy.
func NewLoginUserPassword(username, password, policy string) *loginUserPassword {
	return &loginUserPassword{
		username: username,
		password: password,
		policy:   policy,
	}
}

func (l *loginUserPassword) Login(ws SoapWebService) (*int, error) {
	if l.policy == "" {
		resp, err := ws.Login(&Login{
			Username: &l.username,
			Password: &l.password,
		})
		if err != nil {
			return nil, err
		}
		return resp.LoginReturn, nil
	}
	resp, err := ws.LoginService(&LoginService{
		Username: &l.username,
		Password: &l.password,
		Policy:   &l.policy,
	})
	if err != nil {
		return nil, err
	}
	return resp.LoginServiceReturn, nil
}

type loginManaged struct {
	policy string
	cert   string
}

// NewLoginManaged is a built-in function to create an instance of
// SoapLogin compatible struct. With authorization by CA SDM policy and certificate.
//
// Pre-requirement steps:
//  1. Create new Policy under Admin Tab > SOAP > Policies;
//  2. Specify Proxy Contact
//  3. Run `pdm_pki -p <POLICY_NAME> -f`
//  4. Distribute newly created `<POLICY_NAME>.p12` file
func NewLoginManaged(policy, cert string) *loginManaged {
	return &loginManaged{
		policy: policy,
		cert:   cert,
	}
}

func (l *loginManaged) Login(ws SoapWebService) (*int, error) {
	sign, err := l.readCert()
	if err != nil {
		return nil, err
	}
	resp, err := ws.LoginServiceManaged(&LoginServiceManaged{
		Policy:           &l.policy,
		Encrypted_policy: sign,
	})
	if err != nil {
		return nil, err
	}
	sid, err := strconv.Atoi(*resp.LoginServiceManagedReturn)
	if err != nil {
		return nil, err
	}
	return &sid, nil
}

func (l *loginManaged) readCert() (*string, error) {
	data, err := os.ReadFile(l.cert)
	if err != nil {
		return nil, err
	}

	pk, _, err := pkcs12.Decode(data, l.policy)
	if err != nil {
		return nil, err
	}
	rsaPK := pk.(*rsa.PrivateKey)

	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(l.policy))
	hashed := h.Sum(nil)
	var signature []byte
	signature, err = rsa.SignPKCS1v15(rand.Reader, rsaPK, crypto.SHA1, hashed)
	if err != nil {
		return nil, err
	}
	return StringPtr(base64.StdEncoding.EncodeToString(signature)), nil
}

type loginArtifact struct {
	userid   string
	artifact string
}

// NewLoginArtifact is a built-in function to create an instance of
// SoapLogin compatible struct. With authorization by CA SDM userid and EEM artifact.
//
// *NOT IMPLEMENTED*
func NewLoginArtifact(userid, artifact string) *loginArtifact {
	return &loginArtifact{
		userid:   userid,
		artifact: artifact,
	}
}

// Not Implemented
// I dont have EEM instance, so... pull requests are very appreciated u know
// Login with EEM artifact
func (l *loginArtifact) Login(ws SoapWebService) (*int, error) {
	if l.artifact == "" {
		art, err := ws.GetArtifact(&GetArtifact{
			// I dont have EEM instance so this is not implemented yet
		})
		if err != nil {
			return nil, err
		}
		l.artifact = *art.GetArtifactReturn
	}
	resp, err := ws.LoginWithArtifact(&LoginWithArtifact{
		Userid:   &l.userid,
		Artifact: &l.artifact,
	})
	if err != nil {
		return nil, err
	}
	return resp.LoginWithArtifactReturn, nil
}
