package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
)

type servicePrincipalConfigurationProvider struct {
	keyProvider *servicePrincipalKeyProvider
	region      string
}

// NewServicePrincipalConfigurationProvider will create a new service principal configuration provider
func NewServicePrincipalConfigurationProvider(tenancyID, region, cert, key string, intermediates []string, passphrase *string) (common.ConfigurationProvider, error) {
	var err error
	var keyProvider *servicePrincipalKeyProvider
	if keyProvider, err = newServicePrincipalKeyProvider(tenancyID, region, cert, key, intermediates, passphrase); err != nil {
		return nil, fmt.Errorf("failed to create a new key provider: %s", err.Error())
	}
	return servicePrincipalConfigurationProvider{keyProvider: keyProvider, region: region}, nil
}

func (p servicePrincipalConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.keyProvider.PrivateRSAKey()
}

func (p servicePrincipalConfigurationProvider) KeyID() (string, error) {
	return p.keyProvider.KeyID()
}

func (p servicePrincipalConfigurationProvider) TenancyOCID() (string, error) {
	return "", nil
}

func (p servicePrincipalConfigurationProvider) UserOCID() (string, error) {
	return "", nil
}

func (p servicePrincipalConfigurationProvider) KeyFingerprint() (string, error) {
	return "", nil
}

func (p servicePrincipalConfigurationProvider) Region() (string, error) {
	return p.region, nil
}
