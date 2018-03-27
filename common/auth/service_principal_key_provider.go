// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
)

type servicePrincipalKeyProvider struct {
	federationClient federationClient
}

func newServicePrincipalKeyProvider(tenancyID, region, cert, key string, intermediates []string, passphrase *string) (provider *servicePrincipalKeyProvider, err error) {
	leafCertificateRetriever := newStaticX509CertificateRetriever([]byte(cert), []byte(key), passphrase)

	intermediateCertificateRetrievers := []x509CertificateRetriever{}
	for _, intermediate := range intermediates {
		intermediateCertificateRetrievers =
			append(intermediateCertificateRetrievers, newStaticX509CertificateRetriever([]byte(intermediate), []byte(key), passphrase))
	}

	federationClient := newX509FederationClient(
		common.Region(region), tenancyID, leafCertificateRetriever, intermediateCertificateRetrievers, false)

	provider = &servicePrincipalKeyProvider{federationClient: federationClient}
	return
}

func (p *servicePrincipalKeyProvider) PrivateRSAKey() (privateKey *rsa.PrivateKey, err error) {
	if privateKey, err = p.federationClient.PrivateKey(); err != nil {
		err = fmt.Errorf("failed to get private key: %s", err.Error())
		return nil, err
	}
	return privateKey, nil
}

func (p *servicePrincipalKeyProvider) KeyID() (string, error) {
	var securityToken string
	var err error
	if securityToken, err = p.federationClient.SecurityToken(); err != nil {
		return "", fmt.Errorf("failed to get security token: %s", err.Error())
	}

	return fmt.Sprintf("ST$%s", securityToken), nil
}
