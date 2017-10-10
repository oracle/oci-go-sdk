package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

// ConfigurationProvider wraps information about the account owner
type ConfigurationProvider interface {
	KeyProvider
	TenancyOCID() (string, error)
	UserOCID() (string, error)
	CompartmentOCID() (string, error)
	KeyFingerPrint() (string, error)
}

// ConfigurationProvider reads configuration from environment variables
type EnvironmentConfigurationProvider struct {
	PrivateKeyPassword        string
	EnvironmentVariablePrefix string
}

// PrivateKeyFromBytes is a helper function that will produce a RSA private
// key from bytes.
func privateKeyFromBytes(pemData []byte, password *string) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {
		decrypted := pemBlock.Bytes
		if x509.IsEncryptedPEMBlock(pemBlock) {
			if password == nil {
				e = fmt.Errorf("private_key_password is required for encrypted private keys")
				return
			}
			if decrypted, e = x509.DecryptPEMBlock(pemBlock, []byte(*password)); e != nil {
				return
			}
		}

		key, e = x509.ParsePKCS1PrivateKey(decrypted)

	} else {
		e = fmt.Errorf("PEM data was not found in buffer")
		return
	}
	return
}

func (p EnvironmentConfigurationProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "private_key_path")
	var ok bool
	var value string
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		return nil, fmt.Errorf("Can not read PrivateKey from env variable: %s", environmentVariable)
	}

	pemFileContent, err := ioutil.ReadFile(value)
	if err != nil {
		Debugln("Can not read PrivateKey location from environment variable: " + environmentVariable)
		return
	}

	key, err = privateKeyFromBytes(pemFileContent, &p.PrivateKeyPassword)
	return
}

func (p EnvironmentConfigurationProvider) KeyID() (keyID string, err error) {
	ocid, err := p.TenancyOCID()
	if err != nil {
		return
	}

	userocid, err := p.UserOCID()
	if err != nil {
		return
	}

	fingerPrint, err := p.KeyFingerPrint()
	if err != nil {
		return
	}

	return fmt.Sprintf("%s/%s/%s", ocid, userocid, fingerPrint), nil
}

func (p EnvironmentConfigurationProvider) TenancyOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "tenancy_ocid")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from environment variable %s", environmentVariable)
	}
	return
}

func (p EnvironmentConfigurationProvider) CompartmentOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "compartment_ocid")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read comparment from environment variable %s", environmentVariable)
	}
	return
}

func (p EnvironmentConfigurationProvider) UserOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "user_ocid")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from environment variable %s", environmentVariable)
	}
	return
}

func (p EnvironmentConfigurationProvider) KeyFingerPrint() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "fingerprint")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from environment variable %s", environmentVariable)
	}
	return
}
