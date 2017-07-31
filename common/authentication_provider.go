// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

import "fmt"

// AuthenticationProvider provides required information to sign requests for Oracle Cloud Infrastructure
type AuthenticationProvider interface {
	//KeyId used to sign the requests
	KeyID() string
	//Bytes representing the private key
	PrivateKey() []byte
	//Passphrase for the private key
	PrivateKeyPassphrase() string
}


type AuthenticationDetailsProvider struct {
	Fingerprint string
	TenantID string
	UserID string
	PemKey []byte
}

/// TODO expand these methods
func (provider AuthenticationDetailsProvider) KeyID() string {
	return fmt.Sprintf("%s/%s/%s", provider.TenantID, provider.UserID, provider.Fingerprint)
}

func (provider AuthenticationDetailsProvider) PrivateKey() []byte {
	return provider.PemKey
}

func (provider AuthenticationDetailsProvider) PrivateKeyPassphrase() string {
	return ""
}




