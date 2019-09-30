package main

import (
	"crypto/x509"
	"github.com/giantswarm/microerror"
	"io/ioutil"
)

// CertPoolFromFile returns an x509.CertPool containing the certificates
// in the given PEM-encoded file.
// Returns an error if the file could not be read, a certificate could not
// be parsed, or if the file does not contain any certificates
func CertPoolFromFile(filename string) (*x509.CertPool, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, microerror.Maskf(err, "can't read CA file: %v", filename)
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, microerror.Maskf(err, "failed to append certificates from file: %s", filename)
	}
	return cp, nil
}
