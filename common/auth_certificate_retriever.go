package common

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"sync"
)

// x509CertificateRetriever provides an X509 certificate with the RSA private key
type x509CertificateRetriever interface {
	Refresh() error
	CertificatePemRaw() []byte
	Certificate() *x509.Certificate
	PrivateKeyPemRaw() []byte
	PrivateKey() *rsa.PrivateKey
}

// urlBasedX509CertificateRetriever retrieves PEM-encoded X509 certificates from the given URLs.
type urlBasedX509CertificateRetriever struct {
	certUrl           string
	privateKeyUrl     string
	passphrase        string
	certificatePemRaw []byte
	certificate       *x509.Certificate
	privateKeyPemRaw  []byte
	privateKey        *rsa.PrivateKey
	mux               sync.Mutex
}

func newUrlBasedX509CertificateRetriever(certUrl string, privateKeyUrl string) x509CertificateRetriever {
	retriever := &urlBasedX509CertificateRetriever{
		certUrl:       certUrl,
		privateKeyUrl: privateKeyUrl,
		passphrase:    "", // No passphrase for the private key for Compute instances
		mux:           sync.Mutex{},
	}
	return retriever
}

// Refresh() is failure atomic, i.e., CertificatePemRaw(), Certificate(), PrivateKeyPemRaw(), and PrivateKey() would
// return their previous values if Refresh() fails.
func (r *urlBasedX509CertificateRetriever) Refresh() (err error) {
	Debugln("Refreshing certificate")

	r.mux.Lock()
	defer r.mux.Unlock()

	var certificatePemRaw []byte
	var certificate *x509.Certificate
	if certificatePemRaw, certificate, err = r.renewCertificate(r.certUrl); err != nil {
		return
	}

	var privateKeyPemRaw []byte
	var privateKey *rsa.PrivateKey
	if r.privateKeyUrl != "" {
		if privateKeyPemRaw, privateKey, err = r.renewPrivateKey(r.privateKeyUrl, r.passphrase); err != nil {
			return
		}
	}

	r.certificatePemRaw = certificatePemRaw
	r.certificate = certificate
	r.privateKeyPemRaw = privateKeyPemRaw
	r.privateKey = privateKey
	return
}

func (r *urlBasedX509CertificateRetriever) renewCertificate(url string) (certificatePemRaw []byte, certificate *x509.Certificate, err error) {
	var body bytes.Buffer
	body, err = httpGet(url)
	if err != nil {
		Logln(err)
		return
	}

	certificatePemRaw = body.Bytes()
	var block *pem.Block
	block, _ = pem.Decode(certificatePemRaw)
	if certificate, err = x509.ParseCertificate(block.Bytes); err != nil {
		Logln(err)
		return nil, nil, err
	}

	return
}

func (r *urlBasedX509CertificateRetriever) renewPrivateKey(url string, passphrase string) (privateKeyPemRaw []byte, privateKey *rsa.PrivateKey, err error) {
	var body bytes.Buffer
	body, err = httpGet(url)
	if err != nil {
		Logln(err)
		return
	}

	privateKeyPemRaw = body.Bytes()
	if privateKey, err = privateKeyFromBytes(privateKeyPemRaw, &passphrase); err != nil {
		Logln(err)
		return nil, nil, err
	}

	return
}

func (r *urlBasedX509CertificateRetriever) CertificatePemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.certificatePemRaw == nil {
		return nil
	}

	c := make([]byte, len(r.certificatePemRaw))
	copy(c, r.certificatePemRaw)
	return c
}

func (r *urlBasedX509CertificateRetriever) Certificate() *x509.Certificate {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.certificate == nil {
		return nil
	}

	c := *r.certificate
	return &c
}

func (r *urlBasedX509CertificateRetriever) PrivateKeyPemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.privateKeyPemRaw == nil {
		return nil
	}

	c := make([]byte, len(r.privateKeyPemRaw))
	copy(c, r.privateKeyPemRaw)
	return c
}

func (r *urlBasedX509CertificateRetriever) PrivateKey() *rsa.PrivateKey {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.privateKey == nil {
		return nil
	}

	c := *r.privateKey
	return &c
}
