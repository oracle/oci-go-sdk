package common

import (
	"net/http"
	"crypto/sha256"
	"io/ioutil"
	"encoding/base64"
	"crypto/rsa"
	"crypto"
	"crypto/rand"
	"fmt"
	"strings"
)

type RequestSigner interface {
	Sign(r *http.Request) error
}

type OCIRequestSigner struct {
	PrivateKey *rsa.PrivateKey
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
	var unencodedSig []byte
	unencodedSig, e := rsa.SignPKCS1v15(rand.Reader, signer.PrivateKey, crypto.SHA256, hashed)

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

func (signer *OCIRequestSigner) Sign(request *http.Request) (err error) {
	signer.computeSignature(request)
	return
}
