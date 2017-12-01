package auth

import (
	"bytes"
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

func httpGet(url string) (body bytes.Buffer, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	common.IfDebug(func() {
		if dump, e := httputil.DumpResponse(response, true); e == nil {
			common.Logf("Dump Response %v", string(dump))
		} else {
			common.Debugln(e)
		}
	})

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
			common.Logln(err)
			return
		}
		return fmt.Errorf("HTTP Get failed: URL: %s, Status: %s, Message: %s", url, response.Status, string(body))
	}
	return nil
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
