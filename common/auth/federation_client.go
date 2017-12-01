package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
	"strings"
	"sync"
)

// federationClient is a client to retrieve the security token for an instance principal necessary to sign a request.
// It also provides the private key whose corresponding public key is used to retrieve the security token.
type federationClient interface {
	PrivateKey() (*rsa.PrivateKey, error)
	SecurityToken() (string, error)
}

// x509FederationClient retrieves a security token from Auth service.
type x509FederationClient struct {
	tenancyId                         string
	sessionKeySupplier                sessionKeySupplier
	leafCertificateRetriever          x509CertificateRetriever
	intermediateCertificateRetrievers []x509CertificateRetriever
	securityToken                     securityToken
	authClient                        *common.BaseClient
	mux                               sync.Mutex
}

func newX509FederationClient(region common.Region, tenancyId string, leafCertificateRetriever x509CertificateRetriever, intermediateCertificateRetrievers []x509CertificateRetriever) federationClient {
	client := &x509FederationClient{
		tenancyId:                         tenancyId,
		leafCertificateRetriever:          leafCertificateRetriever,
		intermediateCertificateRetrievers: intermediateCertificateRetrievers,
	}
	client.sessionKeySupplier = newSessionKeySupplier()
	client.authClient = newAuthClient(region, client)
	return client
}

func newAuthClient(region common.Region, provider common.KeyProvider) *common.BaseClient {
	dispatcher := common.DefaultHttpDispatcher()
	signer := common.OciRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: []string{"date", "(request-target)"}, // "host" is not needed for the federation endpoint.  Don't ask me why.
		BodyHeaders:    common.DefaultBodyHeaders,
	}
	client := common.NewBaseClient(signer, &dispatcher, "") // Why BaseClient requires "region"?
	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "auth", string(region))
	client.BasePath = "v1/x509"
	return &client
}

// For authClient to sign requests to X509 Federation Endpoint
func (c *x509FederationClient) KeyID() (string, error) {
	tenancy := c.tenancyId
	fingerPrint := fingerPrint(c.leafCertificateRetriever.Certificate())
	return fmt.Sprintf("%s/fed-x509/%s", tenancy, fingerPrint), nil
}

// For authClient to sign requests to X509 Federation Endpoint
func (c *x509FederationClient) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return c.leafCertificateRetriever.PrivateKey(), nil
}

func (c *x509FederationClient) PrivateKey() (privateKey *rsa.PrivateKey, err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err = c.renewSecurityTokenIfNotValid(); err != nil {
		return
	}
	privateKey = c.sessionKeySupplier.PrivateKey()
	return
}

func (c *x509FederationClient) SecurityToken() (token string, err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if err = c.renewSecurityTokenIfNotValid(); err != nil {
		return
	}
	token = c.securityToken.String()
	return
}

func (c *x509FederationClient) renewSecurityTokenIfNotValid() (err error) {
	if c.securityToken == nil || !c.securityToken.Valid() {
		err = c.renewSecurityToken()
	}
	return
}

func (c *x509FederationClient) renewSecurityToken() (err error) {
	if err = c.sessionKeySupplier.Refresh(); err != nil {
		return
	}

	if err = c.leafCertificateRetriever.Refresh(); err != nil {
		return
	}

	updatedTenancyId := extractTenancyIdFromCertificate(c.leafCertificateRetriever.Certificate())
	if c.tenancyId != updatedTenancyId {
		err = fmt.Errorf("Unexpected update of tenancy OCID in the leaf certificate. Previous tenancy: %s, Updated: %s", c.tenancyId, updatedTenancyId)
		return
	}

	for _, retriever := range c.intermediateCertificateRetrievers {
		if err = retriever.Refresh(); err != nil {
			return
		}
	}

	c.securityToken, err = c.getSecurityToken()
	return
}

func (c *x509FederationClient) getSecurityToken() (token securityToken, err error) {
	request := c.makeX509FederationRequest()
	var httpRequest http.Request
	if httpRequest, err = common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "", request); err != nil {
		common.Logln(err)
		return
	}

	var httpResponse *http.Response
	defer common.CloseBodyIfValid(httpResponse)
	if httpResponse, err = c.authClient.Call(context.Background(), &httpRequest); err != nil {
		common.Logln(err)
		return
	}

	response := x509FederationResponse{}
	if err = common.UnmarshalResponse(httpResponse, &response); err != nil {
		common.Logln(err)
		return
	}

	return newInstancePrincipalToken(response.Token.Token)
}

type x509FederationRequest struct {
	X509FederationDetails `contributesTo:"body"`
}

type X509FederationDetails struct {
	Certificate              string   `mandatory:"true" json:"certificate,omitempty"`
	PublicKey                string   `mandatory:"true" json:"publicKey,omitempty"`
	IntermediateCertificates []string `mandatory:"false" json:"intermediateCertificates,omitempty"`
}

type x509FederationResponse struct {
	Token `presentIn:"body"`
}

type Token struct {
	Token string `mandatory:"true" json:"token,omitempty"`
}

func (c *x509FederationClient) makeX509FederationRequest() *x509FederationRequest {
	certificate := c.sanitizeCertificateString(string(c.leafCertificateRetriever.CertificatePemRaw()))
	publicKey := c.sanitizeCertificateString(string(c.sessionKeySupplier.PublicKeyPemRaw()))
	var intermediateCertificates []string
	for _, retriever := range c.intermediateCertificateRetrievers {
		intermediateCertificates = append(intermediateCertificates, c.sanitizeCertificateString(string(retriever.CertificatePemRaw())))
	}

	details := X509FederationDetails{
		Certificate:              certificate,
		PublicKey:                publicKey,
		IntermediateCertificates: intermediateCertificates,
	}
	return &x509FederationRequest{details}
}

func (c *x509FederationClient) sanitizeCertificateString(certString string) string {
	certString = strings.Replace(certString, "-----BEGIN CERTIFICATE-----", "", -1)
	certString = strings.Replace(certString, "-----END CERTIFICATE-----", "", -1)
	certString = strings.Replace(certString, "-----BEGIN PUBLIC KEY-----", "", -1)
	certString = strings.Replace(certString, "-----END PUBLIC KEY-----", "", -1)
	certString = strings.Replace(certString, "\n", "", -1)
	return certString
}

// sessionKeySupplier provides an RSA keypair which can be re-generated by calling Refresh().
type sessionKeySupplier interface {
	Refresh() error
	PrivateKey() *rsa.PrivateKey
	PublicKeyPemRaw() []byte
}

// inMemorySessionKeySupplier implements sessionKeySupplier to vend an RSA keypair.
// Refresh() generates a new RSA keypair with a random source, and keeps it in memory.
//
// inMemorySessionKeySupplier is not thread-safe.
type inMemorySessionKeySupplier struct {
	keySize         int
	privateKey      *rsa.PrivateKey
	publicKeyPemRaw []byte
}

// newSessionKeySupplier creates and returns a sessionKeySupplier instance which generates key pairs of size 2048.
func newSessionKeySupplier() sessionKeySupplier {
	return &inMemorySessionKeySupplier{keySize: 2048}
}

// Refresh() is failure atomic, i.e., PrivateKey() and PublicKeyPemRaw() would return their previous values
// if Refresh() fails.
func (s *inMemorySessionKeySupplier) Refresh() (err error) {
	common.Debugln("Refreshing session key")

	var privateKey *rsa.PrivateKey
	privateKey, err = rsa.GenerateKey(rand.Reader, s.keySize)
	if err != nil {
		common.Logln(err)
		return
	}

	var publicKeyAsnBytes []byte
	if publicKeyAsnBytes, err = x509.MarshalPKIXPublicKey(privateKey.Public()); err != nil {
		common.Logln(err)
		return
	}
	publicKeyPemRaw := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyAsnBytes,
	})

	s.privateKey = privateKey
	s.publicKeyPemRaw = publicKeyPemRaw
	return
}

func (s *inMemorySessionKeySupplier) PrivateKey() *rsa.PrivateKey {
	if s.privateKey == nil {
		return nil
	}

	c := *s.privateKey
	return &c
}

func (s *inMemorySessionKeySupplier) PublicKeyPemRaw() []byte {
	if s.publicKeyPemRaw == nil {
		return nil
	}

	c := make([]byte, len(s.publicKeyPemRaw))
	copy(c, s.publicKeyPemRaw)
	return c
}

type securityToken interface {
	fmt.Stringer
	Valid() bool
}

type instancePrincipalToken struct {
	tokenString string
	jwtToken    *jwtToken
}

func newInstancePrincipalToken(tokenString string) (newToken securityToken, err error) {
	jwtToken, err := parseJwt(tokenString)
	if err != nil {
		common.Logln(err)
		return
	}
	newToken = &instancePrincipalToken{tokenString, jwtToken}
	return newToken, nil
}

func (t *instancePrincipalToken) String() string {
	return t.tokenString
}

func (t *instancePrincipalToken) Valid() bool {
	return !t.jwtToken.expired()
}
