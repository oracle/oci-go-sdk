package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
)

type instancePrincipalConfigurationProvider struct {
	keyProvider *instancePrincipalKeyProvider
	region      *common.Region
}

func InstancePrincipalConfigurationProvider() (common.ConfigurationProvider, error) {
	var err error
	var keyProvider *instancePrincipalKeyProvider
	if keyProvider, err = newInstancePrincipalKeyProvider(); err != nil {
		return nil, fmt.Errorf("failed to create a new key provider for instance principal: %s", err.Error())
	}
	return instancePrincipalConfigurationProvider{keyProvider: keyProvider, region: nil}, nil
}

func InstancePrincipalConfigurationProviderForRegion(region common.Region) (common.ConfigurationProvider, error) {
	var err error
	var keyProvider *instancePrincipalKeyProvider
	if keyProvider, err = newInstancePrincipalKeyProvider(); err != nil {
		return nil, fmt.Errorf("failed to create a new key provider for instance principal: %s", err.Error())
	}
	return instancePrincipalConfigurationProvider{keyProvider: keyProvider, region: &region}, nil
}

func (p instancePrincipalConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.keyProvider.PrivateRSAKey()
}

func (p instancePrincipalConfigurationProvider) KeyID() (string, error) {
	return p.keyProvider.KeyID()
}

func (p instancePrincipalConfigurationProvider) TenancyOCID() (string, error) {
	return "", nil
}

func (p instancePrincipalConfigurationProvider) UserOCID() (string, error) {
	return "", nil
}

func (p instancePrincipalConfigurationProvider) KeyFingerPrint() (string, error) {
	return "", nil
}

func (p instancePrincipalConfigurationProvider) Region() (string, error) {
	if p.region == nil {
		return string(p.keyProvider.RegionForFederationClient()), nil
	}
	return string(*p.region), nil
}
