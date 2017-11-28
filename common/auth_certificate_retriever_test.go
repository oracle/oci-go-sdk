package common

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	expectedCert = `-----BEGIN CERTIFICATE-----
MIIHFTCCBP2gAwIBAgIRAOPfQJtFNAzYnsIFJIL0VUcwDQYJKoZIhvcNAQELBQAw
gZ8xczBxBgNVBAsTam9wYy1kZXZpY2U6OTQ6YTI6MzQ6YTk6ZTI6ZTk6NDk6Mzc6
NzY6M2I6ZDE6Mzc6NWE6MDM6NmQ6OGE6ZjM6M2I6MWY6N2Q6NDk6ZmI6OWM6ZTU6
Mzg6Yzk6NDM6Njg6Mzk6NDM6ZmQ6YWExKDAmBgNVBAMTH1BLSVNWQyBJZGVudGl0
eSBJbnRlcm1lZGlhdGUgcjIwHhcNMTcxMTIwMTk1ODM2WhcNMTcxMTIwMjA1ODM2
WjCCAbAxggFOMBwGA1UECxMVb3BjLWNlcnR0eXBlOmluc3RhbmNlMGcGA1UECxNg
b3BjLWluc3RhbmNlOm9jaWQxLmluc3RhbmNlLm9jMS5waHguYWJ5aHFsanJ6Y29i
dmJ1dDdrbmlxZHk2c2Y0ZnA1NjU0dmI3ZGdiZ3g2ZHEzaDM3b3luamk0dTVuMmxx
MGoGA1UECxNjb3BjLWNvbXBhcnRtZW50Om9jaWQxLmNvbXBhcnRtZW50Lm9jMS4u
YWFhYWFhYWEyM2tubjJxaTIzdXJxcXZ2b3picm8zdTVneDJ1YmRhYnl5Z3ZlNng1
NmNpZWh1MnpibHZxMFkGA1UECxNSb3BjLXRlbmFudDpvY2lkdjE6dGVuYW5jeTpv
YzE6cGh4OjE0NjA0MDY1OTI2NjA6YWFhYWFhYWFiNGZhb2ZyZmt4ZWNvaGhqdWl2
anEyNjJwdTFcMFoGA1UEAxNTb2NpZDEuaW5zdGFuY2Uub2MxLnBoeC5hYnlocWxq
cnpjb2J2YnV0N2tuaXFkeTZzZjRmcDU2NTR2YjdkZ2JneDZkcTNoMzdveW5qaTR1
NW4ybHEwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDiSb9X4Rl0QtZS
uRuxoio2NQL2S9ueZiBnSX+6+ohPrNLty4P2MzeUC6a1Fc4gkJj5DVcLNlMatnOM
pIWhq10NnOvuxTg3Fnu3e9QPGWdIFysojSh6fugWkOhPNl7t9MtQbjOx4FNkyNRT
WKP6kkFUFpLjJYsHRNu0IJ/Lk5DdYlsJT8nUhywIO7nl/iN6VUl84nyIAuqiBbK5
CD7r6/7QJ2E7pyZQUvRGFjMosWTb8UJpgTrrvdlSzaU9Fxu8W0Ev7CLhEuRICOKn
IHY/xIgMUqIskxGGS/Vmrsk9pHw4dCs5Zx+oEWpoz9iDGdYnUi7KSL3fdrOfIg3l
a/SMwJ+MQCnHAinoxL5VW4HMCdCYIUXWavUlVPURk21/zPfsfKU8Ml7aCBpHQ+7q
iORbrZ0pa/tO5cYjWZ/cKs9cpfhv54mQlrdBy+FoZcYUZavXoRkYt90MGqCq33qb
HEUgimySQ+s8IC4R4/fJNlB7//+R/r/qrlQV2JGakiGCvvvN9XCFV+nzeuiu8kJr
T3r0RqJNyWdCix+84bpUxOUVGNr+7Z+TCX+70tBoTL1Zv4efFkCbqHvmvXb8hgFH
5rQO+19WnmneA0oKv2S7iDi5qO8loLg7FEhJ1FKy4StRUKpwhVBPuu8F9xQknIpK
CJzyJDdfLRtHEpmVej+rwrlZOmdSVwIDAQABozgwNjATBgNVHSUEDDAKBggrBgEF
BQcDAjAfBgNVHSMEGDAWgBQKKSDx5jhmP8SLpYrtwMQsrYdKfDANBgkqhkiG9w0B
AQsFAAOCAgEAxnifneny4EOqbIz4acFAjZmUE0rt++Wkyj6TJuXzib/vkEzGo5p1
PUw32d4DNSbAUtHiNJvbtIdsnLcO+D7j3NB/FlnHrSQ76JeABxAqYKYg9mntNFLj
R04oJmHNOUYk2svAa/XnfEAkuxItSRiqYOADYer8TDiw3+9KpaPjqH1pjIw3Ue7D
TfBlA2DVu62or5sg+1atzyYfC/GlRBnWH3SRp/cRdI3l6kqgn1sJgVAbTum1Kl/t
Vj3v3bOwCzlWuKX2NePviVoyweKQ+jILqyli/PTsLOVJ5U1UmYzaFVsSRgV5igbj
82f/ufI8OfjTfuVwY359Drva1CcTvRuwDFhcb3urTz684hBXnAEQd/vLP/lrbuFI
Ln49+wlsWaUSlUgCI3BmlBH3Dd04KBxSvI83S08e3xQCOOKpnqjFpsb9te6PMwoi
fnHMYZTTaL9kno57ptgOQfN4QCVftfUiG3WEFVxfqpnOtUf7IzYWlW6h2DTPT+y9
fuT5cUA93NF8bhKIbz8F4bscExld1sLAKa5xdCkI/oy6fN+rG6aqlm4Iy2+RfPss
uDr1HvBHkhVJmzCbhH3qyFRnrLnkij6pzmNV9zUEQ7EqPMUISAf7H7sAyBm+W89z
mVe2I48cw2cH3kyZApR6NG2EMloOvmNQ0FK9gkj64cub0ndimjEOqwE=
-----END CERTIFICATE-----
`
	expectedPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
)

func TestUrlBasedX509CertificateRetriever_RefreshWithoutPrivateKeyUrl(t *testing.T) {
	certServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedCert)
	}))
	defer certServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, "")
	err := retriever.Refresh()

	assert.NoError(t, err)

	assert.Equal(t, []byte(expectedCert), retriever.CertificatePemRaw())
	actualCert := retriever.Certificate()
	actualCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: actualCert.Raw})
	assert.Equal(t, []byte(expectedCert), actualCertPem)

	assert.Nil(t, retriever.PrivateKeyPemRaw())
	assert.Nil(t, retriever.PrivateKey())
}

func TestUrlBasedX509CertificateRetriever_RefreshWithPrivateKeyUrl(t *testing.T) {
	certServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedCert)
	}))
	defer certServer.Close()
	privateKeyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedPrivateKey)
	}))
	defer privateKeyServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, privateKeyServer.URL)
	err := retriever.Refresh()

	assert.NoError(t, err)

	assert.Equal(t, []byte(expectedCert), retriever.CertificatePemRaw())
	actualCert := retriever.Certificate()
	actualCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: actualCert.Raw})
	assert.Equal(t, []byte(expectedCert), actualCertPem)

	assert.Equal(t, []byte(expectedPrivateKey), retriever.PrivateKeyPemRaw())
	actualPrivateKey := retriever.PrivateKey()
	actualPrivateKeyPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(actualPrivateKey)})
	assert.Equal(t, []byte(expectedPrivateKey), actualPrivateKeyPem)
}

func TestUrlBasedX509CertificateRetriever_RefreshCertNotFound(t *testing.T) {
	certServer := httptest.NewServer(http.NotFoundHandler())
	defer certServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, "")
	err := retriever.Refresh()

	assert.Error(t, err)
}

func TestUrlBasedX509CertificateRetriever_RefreshPrivateKeyNotFound(t *testing.T) {
	certServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedCert)
	}))
	defer certServer.Close()
	privateKeyServer := httptest.NewServer(http.NotFoundHandler())
	defer privateKeyServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, privateKeyServer.URL)
	err := retriever.Refresh()

	assert.Error(t, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func TestUrlBasedX509CertificateRetriever_RefreshCertInternalServerError(t *testing.T) {
	certServer := httptest.NewServer(http.HandlerFunc(internalServerError))
	defer certServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, "")
	err := retriever.Refresh()

	assert.Error(t, err)
}

func TestUrlBasedX509CertificateRetriever_RefreshPrivateKeyInternalServerError(t *testing.T) {
	certServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedCert)
	}))
	defer certServer.Close()
	privateKeyServer := httptest.NewServer(http.HandlerFunc(internalServerError))
	defer privateKeyServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, privateKeyServer.URL)
	err := retriever.Refresh()

	assert.Error(t, err)
}
