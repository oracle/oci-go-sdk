package common

import (
	"crypto/rsa"
	"fmt"
)

const (
	regionUrl                     = `http://169.254.169.254/opc/v1/instance/region`
	leafCertificateUrl            = `http://169.254.169.254/opc/v1/identity/cert.pem`
	leafCertificateKeyUrl         = `http://169.254.169.254/opc/v1/identity/key.pem`
	intermediateCertificateUrl    = `http://169.254.169.254/opc/v1/identity/intermediate.pem`
	intermediateCertificateKeyUrl = ``
)

// instancePrincipalKeyProvider implements KeyProvider to provide a key ID and its corresponding private key
// for an instance principal by getting a security token via x509FederationClient.
//
// The region name of the endpoint for x509FederationClient is obtained from the metadata service on the compute
// instance.
type instancePrincipalKeyProvider struct {
	regionForFederationClient Region
	federationClient          federationClient
}

// newInstancePrincipalKeyProvider creates and returns an instancePrincipalKeyProvider instance based on
// x509FederationClient.
//
// NOTE: There is a race condition between PrivateRSAKey() and KeyID().  These two pieces are tightly coupled; KeyID
// includes a security token obtained from Auth service by giving a public key which is paired with PrivateRSAKey.
// The x509FederationClient caches the security token in memory until it is expired.  Thus, even if a client obtains a
// KeyID that is not expired at the moment, the PrivateRSAKey that the client acquires at a next moment could be
// invalid because the KeyID could be already expired.
func newInstancePrincipalKeyProvider() (provider *instancePrincipalKeyProvider, err error) {
	region, err := getRegionForFederationClient(regionUrl)
	if err != nil {
		return
	}

	leafCertificateRetriever := newUrlBasedX509CertificateRetriever(leafCertificateUrl, leafCertificateKeyUrl)
	intermediateCertificateRetrievers := []x509CertificateRetriever{
		newUrlBasedX509CertificateRetriever(intermediateCertificateUrl, intermediateCertificateKeyUrl),
	}

	if err = leafCertificateRetriever.Refresh(); err != nil {
		return
	}
	tenancyId := extractTenancyIdFromCertificate(leafCertificateRetriever.Certificate())

	federationClient := newX509FederationClient(
		region, tenancyId, leafCertificateRetriever, intermediateCertificateRetrievers)

	provider = &instancePrincipalKeyProvider{regionForFederationClient: region, federationClient: federationClient}
	return
}

func getRegionForFederationClient(url string) (r Region, err error) {
	body, err := httpGet(url)
	if err != nil {
		Logln(err)
		return
	}
	return StringToRegion(body.String())
}

func (p *instancePrincipalKeyProvider) RegionForFederationClient() Region {
	return p.regionForFederationClient
}

func (p *instancePrincipalKeyProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.federationClient.PrivateKey()
}

func (p *instancePrincipalKeyProvider) KeyID() (keyID string, err error) {
	var securityToken string
	if securityToken, err = p.federationClient.SecurityToken(); err != nil {
		return
	}
	keyID = fmt.Sprintf("ST$%s", securityToken)
	return
}
