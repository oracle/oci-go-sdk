package common

import (
	"net/http"
	"crypto"
	"crypto/sha256"
	"crypto/rsa"
	"crypto/rand"
	"io/ioutil"
	"encoding/base64"
	"fmt"
	"strings"
)

//RequestSigner the interface to sign a request
type RequestSigner interface {
	Sign(r *http.Request) error
}

//KeyProvider interface that wraps information about the key's account owner
type KeyProvider interface {
	PrivateRSAKey() (*rsa.PrivateKey, error)
	KeyID() (string, error)
}

var signerVersion = "1"
type OCIRequestSigner struct {
	KeyProvider KeyProvider
}

func getSigningHeaders(method string) []string {
	result := []string{
		"date",
		"(request-target)",
		"host",
	}

	if method == http.MethodPost || method == http.MethodPut {
		result = append(result, "content-length", "content-type", "x-content-sha256")
	}

	return result
}

func getSigningString(request *http.Request) string {
	signingHeaders := getSigningHeaders(request.Method)
	signingParts := make([]string, len(signingHeaders))
	for i, part := range signingHeaders {
		var value string
		switch part {
		case "(request-target)":
			value = getRequestTarget(request)
		case "host":
			value = request.URL.Host
		default:
			value = request.Header.Get(part)
		}
		signingParts[i] = fmt.Sprintf("%s: %s", part, value)
	}

	signingString := strings.Join(signingParts, "\n")
	return signingString

}

func getRequestTarget(request *http.Request) string {
	lowercaseMethod := strings.ToLower(request.Method)
	return fmt.Sprintf("%s %s", lowercaseMethod, request.URL.RequestURI())
}

func (signer OCIRequestSigner) computeSignature(request *http.Request) (signature string, err error) {
	signingString := getSigningString(request)
	Debugf("Signing string is: %s", signingString)
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
	var signature string
	if signature,  err = signer.computeSignature(request); err != nil {
		return
	}

	signigHeaders := strings.Join(getSigningHeaders(request.Method), " ")


	var keyID string
	if keyID, err = signer.KeyProvider.KeyID();  err != nil {
		return
	}

	authValue := fmt.Sprintf("Signature version=\"%s\",headers=\"%s\",keyId=\"%s\",algorithm=\"rsa-sha256\",signature=\"%s\"",
		signerVersion, signigHeaders, keyID, signature)

	request.Header.Add("Authorization", authValue)

	return
}

