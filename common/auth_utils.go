package common

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

func httpGet(url string) (body bytes.Buffer, err error) {
	response, err := http.Get(url)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		if dump, e := httputil.DumpResponse(response, true); e == nil {
			Logf("Dump Response %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	if err != nil {
		return
	}

	defer response.Body.Close()
	if _, err = body.ReadFrom(response.Body); err != nil {
		return
	}

	err = checkStatusCode(url, response)
	return
}

func checkStatusCode(url string, response *http.Response) (err error) {
	if response.StatusCode < 200 || 300 <= response.StatusCode {
		var body []byte
		if body, err = ioutil.ReadAll(response.Body); err != nil {
			Logln(err)
			return
		}
		return fmt.Errorf("HTTP Get failed: URL: %s, Status: %s, Message: %s", url, response.Status, string(body))
	}
	return nil
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

func extractTenancyIdFromCertificate(cert *x509.Certificate) string {
	for _, nameAttr := range cert.Subject.Names {
		value := nameAttr.Value.(string)
		if strings.HasPrefix(value, "opc-tenant:") {
			return value[len("opc-tenant:"):]
		}
	}
	return ""
}

func fingerPrint(certificate *x509.Certificate) string {
	fingerPrint := sha1.Sum(certificate.Raw)
	return colonSeparatedString(fingerPrint)
}

func colonSeparatedString(fingerPrint [sha1.Size]byte) string {
	spaceSeparated := fmt.Sprintf("% x", fingerPrint)
	return strings.Replace(spaceSeparated, " ", ":", -1)
}
