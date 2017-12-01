package auth

import (
	"crypto/rsa"
	"github.com/oracle/oci-go-sdk/common"
)

type instancePrincipalConfigurationProvider struct {
	keyProvider *instancePrincipalKeyProvider
	region      *common.Region
}

func InstancePrincipalConfigurationProvider() (provider common.ConfigurationProvider, err error) {
	keyProvider, err := NewInstancePrincipalKeyProvider()
	if err != nil {
		return
	}
	provider = instancePrincipalConfigurationProvider{keyProvider: keyProvider, region: nil}
	return
}

func InstancePrincipalConfigurationProviderForRegion(region common.Region) (provider common.ConfigurationProvider, err error) {
	keyProvider, err := NewInstancePrincipalKeyProvider()
	if err != nil {
		return
	}
	provider = instancePrincipalConfigurationProvider{keyProvider: keyProvider, region: &region}
	return
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
