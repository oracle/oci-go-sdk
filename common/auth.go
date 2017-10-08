package common

import (
	"net/http"
	"crypto"
	"crypto/sha256"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"encoding/base64"
	"fmt"
	"strings"
	"os"
)

//RequestSigner the interface to sign a request
type RequestSigner interface {
	Sign(r *http.Request) error
}

type OCIRequestSigner struct {
	KeyProvider KeyProvider
}

func getSigningHeaders(method string) []string {
	result := []string{
		"date",
		"(request-target)",
	}

	if method == http.MethodPost || method == http.MethodPut {
		result = append(result, "content-length", "content-type", "x-content-sha256")
	}

	return result
}

func getSigningString(request *http.Request) string {
	signingHeaders := getSigningHeaders(request.Method)
	signingString := ""
	for _, header := range signingHeaders {
		if signingString != "" {
			signingString += "\n"
		}

		if header == "(request-target)" {
			signingString += fmt.Sprintf("%s: %s", header, getRequestTarget(request))
		} else {
			signingString += fmt.Sprintf("%s: %s", header, request.Header.Get(header))
		}
	}

	return signingString

}

func getRequestTarget(request *http.Request) string {
	lowercaseMethod := strings.ToLower(request.Method)
	return fmt.Sprintf("%s %s", lowercaseMethod, request.URL.RequestURI())
}

func (signer OCIRequestSigner) computeSignature(request *http.Request) (signature string, err error) {
	signingString := getSigningString(request)
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	hashed := hasher.Sum(nil)

	privateKey, err := signer.KeyProvider.PrivateRSAKey()
	if err != nil {
		return
	}

	var unencodedSig []byte
	unencodedSig, e := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if e != nil {
		err = fmt.Errorf("Can not compute signature while signing the request %s: ", e.Error())
		return
	}

	signature = base64.StdEncoding.EncodeToString(unencodedSig)
	return
}

func GetBodyHash(request http.Request) (hashString string, err error) {
	body, e := ioutil.ReadAll(request.Body)
	if e != nil {
		err = fmt.Errorf("Can not read body of request while calculating body hash: %s", e.Error())
		return
	}

	hash := sha256.Sum256(body)
	hashString = base64.StdEncoding.EncodeToString(hash[:])
	return
}

func (signer OCIRequestSigner) Sign(request *http.Request) (err error) {
	signer.computeSignature(request)
	return
}

type KeyProvider interface {
	PrivateRSAKey() (*rsa.PrivateKey, error)
}
// ConfigurationProvider returns information about the account owner
type ConfigurationProvider interface {
	KeyProvider
	TenancyOCID() (string, error)
	UserOCID() (string, error)
	KeyFingerPrint() (string, error)
}

type EnvironmentConfigurationProvider struct {
	PrivateKeyPassword string
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
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "key_file")
	var ok bool
	var value string
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read PrivateKey from env variable %s", environmentVariable)
		return
	}

	pemFileContent, err := ioutil.ReadFile(value)
	if err != nil {
		Debugln("Can not read PrivateKey location from environment variable: " + environmentVariable)
		return
	}

	key, err = privateKeyFromBytes(pemFileContent, &p.PrivateKeyPassword)
	return
}

func (p EnvironmentConfigurationProvider) TenancyOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "tenancy")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from env variable %s", environmentVariable)
	}
	return
}

func (p EnvironmentConfigurationProvider) UserOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "user")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from env variable %s", environmentVariable)
	}
	return
}

func (p EnvironmentConfigurationProvider) KeyFingerPrint() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "fingerprint")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from env variable %s", environmentVariable)
	}
	return
}


