package common

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//HttpRequestSigner the interface to sign a request
type HttpRequestSigner interface {
	Sign(r *http.Request) error
}

//KeyProvider interface that wraps information about the key's account owner
type KeyProvider interface {
	PrivateRSAKey() (*rsa.PrivateKey, error)
	KeyID() (string, error)
}

var signerVersion = "1"

//ociRequestSigner implements the http-signatures-draft spec
//as described in https://tools.ietf.org/html/draft-cavage-http-signatures-08
type ociRequestSigner struct {
	KeyProvider    KeyProvider
	GenericHeaders []string
	BodyHeaders    []string
}

var (
	defaultGenericHeaders = []string{"date", "(request-target)", "host"}
	defaultBodyHeaders    = []string{"content-length", "content-type", "x-content-sha256"}
)

func defaultRequestSigner(provider KeyProvider) HttpRequestSigner {
	return RequestSigner(provider, defaultGenericHeaders, defaultBodyHeaders)
}

func RequestSigner(provider KeyProvider, genericHeaders, bodyHeaders []string) HttpRequestSigner {
	return ociRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: genericHeaders,
		BodyHeaders:    bodyHeaders}
}

func (signer ociRequestSigner) getSigningHeaders(method string) []string {
	var result []string
	result = append(result, signer.GenericHeaders...)

	if method == http.MethodPost || method == http.MethodPut {
		result = append(result, signer.BodyHeaders...)
	}

	return result
}

func (signer ociRequestSigner) getSigningString(request *http.Request) string {
	signingHeaders := signer.getSigningHeaders(request.Method)
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

func calculateHashOfBody(request *http.Request) (err error) {
	if request.Method == http.MethodPost || request.Method == http.MethodPut {
		var hash string
		if request.ContentLength > 0 {
			hash, err = GetBodyHash(request)
			if err != nil {
				return
			}
		} else {
			hash = hashAndEncode([]byte(""))
		}
		request.Header.Set("X-Content-Sha256", hash)
	}
	return
}

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
//
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func hashAndEncode(data []byte) string {
	hashedContent := sha256.Sum256(data)
	hash := base64.StdEncoding.EncodeToString(hashedContent[:])
	return hash
}

func GetBodyHash(request *http.Request) (hashString string, err error) {
	if request.Body == nil {
		return "", fmt.Errorf("can not read body of request while calculating body hash, nil body?")
	}

	var data []byte
	bReader := request.Body
	bReader, request.Body, err = drainBody(request.Body)
	if err != nil {
		return "", fmt.Errorf("can not read body of request while calculating body hash: %s", err.Error())
	}

	data, err = ioutil.ReadAll(bReader)
	if err != nil {
		return "", fmt.Errorf("can not read body of request while calculating body hash: %s", err.Error())
	}
	hashString = hashAndEncode(data)
	return
}

func (signer ociRequestSigner) computeSignature(request *http.Request) (signature string, err error) {
	signingString := signer.getSigningString(request)
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
		err = fmt.Errorf("can not compute signature while signing the request %s: ", e.Error())
		return
	}

	signature = base64.StdEncoding.EncodeToString(unencodedSig)
	return
}

// Signs the http request, by inspecting the necessary headers. Once signed
// the request will have the proper 'Authorization' header set, otherwise
// and error is returned
func (signer ociRequestSigner) Sign(request *http.Request) (err error) {
	err = calculateHashOfBody(request)
	if err != nil {
		return
	}

	var signature string
	if signature, err = signer.computeSignature(request); err != nil {
		return
	}

	signingHeaders := strings.Join(signer.getSigningHeaders(request.Method), " ")

	var keyID string
	if keyID, err = signer.KeyProvider.KeyID(); err != nil {
		return
	}

	authValue := fmt.Sprintf("Signature version=\"%s\",headers=\"%s\",keyId=\"%s\",algorithm=\"rsa-sha256\",signature=\"%s\"",
		signerVersion, signingHeaders, keyID, signature)

	request.Header.Set("Authorization", authValue)

	return
}
