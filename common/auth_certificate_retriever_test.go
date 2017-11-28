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
	assert.Nil(t, retriever.CertificatePemRaw())
	assert.Nil(t, retriever.Certificate())
	assert.Nil(t, retriever.PrivateKeyPemRaw())
	assert.Nil(t, retriever.PrivateKey())
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
	assert.Nil(t, retriever.CertificatePemRaw())
	assert.Nil(t, retriever.Certificate())
	assert.Nil(t, retriever.PrivateKeyPemRaw())
	assert.Nil(t, retriever.PrivateKey())
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
	assert.Nil(t, retriever.CertificatePemRaw())
	assert.Nil(t, retriever.Certificate())
	assert.Nil(t, retriever.PrivateKeyPemRaw())
	assert.Nil(t, retriever.PrivateKey())
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
	assert.Nil(t, retriever.CertificatePemRaw())
	assert.Nil(t, retriever.Certificate())
	assert.Nil(t, retriever.PrivateKeyPemRaw())
	assert.Nil(t, retriever.PrivateKey())
}

const (
	anotherCert = `-----BEGIN CERTIFICATE-----
MIIHFDCCBPygAwIBAgIQRMnR4lozX+GuvxhHRBwUszANBgkqhkiG9w0BAQsFADCB
nzFzMHEGA1UECxNqb3BjLWRldmljZToxMDo5ZDpmNzplYzo2MTo4YTo3NDo5Mzpk
YzowZTpkZjpiZDo0MDpiNzo1ZjpmMToxMzo5ODoyZjo0NDpmYjpkMzowNjo4YTo5
OTo2YzplMTowNDo0MTpkZjo2YToyMTEoMCYGA1UEAxMfUEtJU1ZDIElkZW50aXR5
IEludGVybWVkaWF0ZSByMjAeFw0xNzExMjgyMDE2MjJaFw0xNzExMjgyMTE2MjJa
MIIBsDGCAU4wHAYDVQQLExVvcGMtY2VydHR5cGU6aW5zdGFuY2UwZwYDVQQLE2Bv
cGMtaW5zdGFuY2U6b2NpZDEuaW5zdGFuY2Uub2MxLnBoeC5hYnlocWxqcnpjb2J2
YnV0N2tuaXFkeTZzZjRmcDU2NTR2YjdkZ2JneDZkcTNoMzdveW5qaTR1NW4ybHEw
agYDVQQLE2NvcGMtY29tcGFydG1lbnQ6b2NpZDEuY29tcGFydG1lbnQub2MxLi5h
YWFhYWFhYTIza25uMnFpMjN1cnFxdnZvemJybzN1NWd4MnViZGFieXlndmU2eDU2
Y2llaHUyemJsdnEwWQYDVQQLE1JvcGMtdGVuYW50Om9jaWR2MTp0ZW5hbmN5Om9j
MTpwaHg6MTQ2MDQwNjU5MjY2MDphYWFhYWFhYWI0ZmFvZnJma3hlY29oaGp1aXZq
cTI2MnB1MVwwWgYDVQQDE1NvY2lkMS5pbnN0YW5jZS5vYzEucGh4LmFieWhxbGpy
emNvYnZidXQ3a25pcWR5NnNmNGZwNTY1NHZiN2RnYmd4NmRxM2gzN295bmppNHU1
bjJscTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMzUk6+x8ldxU3rS
k4p12C8FC2ZsKTzGEKL+lJJ3MgGNe4EAAF5+EkZeWKOdLDjFLzvMGhkX/x/n9evt
NgigRrfSBee10K/wPcOydDuHtS/1REC1vs6wqFQppjQ0ijtZ1enXvUa8ClJg3Xtl
jeKcSNavhsM7JyPvCVc6/sXUfz+fH689Z+WF+Wvn5/90s3N2jZAY9ym3e2N253Nc
ve68HF0hAbcKd67D/MnI+mqzZaUnY32qh6jUBKV9enKPjgDkkkPeEpStTvzZ5rn1
0fnKzr2n1gDhY0vK43tpfpjvD9EyQPduCSjWGGZrDRk7klvxzPGF7Z2eeguLndu2
+NNLJZ1SHIYR4DGT+zxATONLebOPh7lYS+urjWdO+JG/H67FMtXFE72FR+nTZOD3
EgE1MrHLyOTuxhNZcwhsLUKFwg1TOKgAOFJCGpDcw6FJwu8uMH4+Sn+UXZLGwmpl
7QhlJxGQ4IHga3VPxXZ1n2G4iNVWtM9ozFqa+eOqfBoiVbE6v+cK21pCYtmxX17z
dwUMrFJVMgVnmVkT80NDM0E4lp+MaER3Lpm2lUJ1MXT/TfzdR+89/oeJiDHv30Mt
cPbb2d6Kg5iSBIs4zmfjrvr8DUdopBpZ0eY/rV+c/HZ3kfjU2CDT56OZwh85/LHD
vEnDiybjk5HDKiJwHDD7EyUEDY4bAgMBAAGjODA2MBMGA1UdJQQMMAoGCCsGAQUF
BwMCMB8GA1UdIwQYMBaAFAopIPHmOGY/xIuliu3AxCyth0p8MA0GCSqGSIb3DQEB
CwUAA4ICAQDHsLu3tvVzKN564MsCbIAQyKavuZI7rdLXbI9VoOjsTAll1PB46AXk
sGqO8niYNxfkN0XKjgDL7PGIe+7UiTNdjoPKIBiTUYmtad82A3cDW9wiKGMEgn4X
Fj29I4z/1/uJn2rgvWsEO0YMLtZUL2IiyS+X4HpyMIYJjZ3xpg/t88IRuXlBbz3P
7q/dDl3o/T3mtklzpBu9FxyejjFLG8ExladSE4Y8dECCEh1n60/adRq63sv2Liuf
wnmXVK9e//TIrqqiQkaJOem35+Lqn0UzcZ4Y12QJoKU5k6lhFRFA8cifK0rxfWQB
SChj3xTflap75EEBCASpPiUYymvBppT3MwQbaY3CPqsGpn2IObPjJpsgCB5EMaYR
vsWsoAZbjstl2WqPpzJ6HlFXGynxTz1ogtW9SWsprsBMnOaZk9M4jp3Tws9H8hK5
UmgaMy7TUVyu6H9vT+A0E3Jr2A9SPDs1pq3NKEtWZL7hub/tDWmhCiDzG8s4ijDq
i7agiy+8AgTfgP5VzLO/pusQSs2anKGFMXSJqAS1QUvCWuMrZIxQLeiot4+oQf6K
ipFy/tzWCetPvWXjlNQ3SeOWfzeb7QVyFMkhbQxIvaTw6ux6AlPxxo6yTiOh/7Sk
brBmW9J4v5PqRE1mAQV8FAoIJ9Wta7p5kwwwwO0jjcd9uOjt8bxDdQ==
-----END CERTIFICATE-----
`
)

func TestUrlBasedX509CertificateRetriever_FailureAtomicity(t *testing.T) {
	privateKeyServerFailed := false

	certServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if privateKeyServerFailed {
			fmt.Fprint(w, anotherCert)

		} else {
			fmt.Fprint(w, expectedCert)
		}
	}))
	defer certServer.Close()

	privateKeyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if privateKeyServerFailed {
			internalServerError(w, r)
		} else {
			fmt.Fprint(w, expectedPrivateKey)
		}
	}))
	defer privateKeyServer.Close()

	retriever := newUrlBasedX509CertificateRetriever(certServer.URL, privateKeyServer.URL)
	err := retriever.Refresh()

	assert.NoError(t, err)

	privateKeyServerFailed = true

	err = retriever.Refresh()

	assert.Error(t, err)
	assert.Equal(t, []byte(expectedCert), retriever.CertificatePemRaw()) // Not anotherCert but expectedCert
	actualCert := retriever.Certificate()
	actualCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: actualCert.Raw})
	assert.Equal(t, []byte(expectedCert), actualCertPem)

	assert.Equal(t, []byte(expectedPrivateKey), retriever.PrivateKeyPemRaw())
	actualPrivateKey := retriever.PrivateKey()
	actualPrivateKeyPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(actualPrivateKey)})
	assert.Equal(t, []byte(expectedPrivateKey), actualPrivateKeyPem)
}
