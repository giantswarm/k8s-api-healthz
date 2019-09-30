package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/giantswarm/microerror"
)

const (
	maxIdleConnection = 10
	maxTimeoutSec     = 4
)

type HealtzConfig struct {
	ApiEndpoint    string
	ApiCertPath    string
	ApiCACertPath  string
	ApiKeyPath     string
	EtcdEndpoint   string
	EtcdCertPath   string
	EtcdCACertPath string
	EtcdKeyPath    string
	Port           int
}

type Healthz struct {
	apiUrl  *url.URL
	etcdUrl *url.URL
	port    int

	apiHttpTransport  *http.Transport
	apiHttpClient     *http.Client
	etcdHttpTransport *http.Transport
	etcdHttpClient    *http.Client
}

func NewHealtz(c HealtzConfig) (*Healthz, error) {
	if c.ApiEndpoint == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ApiEndpoint cannot be empty", c)
	}
	if c.ApiCertPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ApiCertPath cannot be empty", c)
	}
	if c.ApiCACertPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ApiCACertPath cannot be empty", c)
	}
	if c.ApiKeyPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ApiKeyPath cannot be empty", c)
	}
	if c.EtcdEndpoint == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.EtcdEndpoint cannot be empty", c)
	}
	if c.EtcdCertPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.EtcdCertPath cannot be empty", c)
	}
	if c.EtcdCACertPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.EtcdCACertPath cannot be empty", c)
	}
	if c.EtcdKeyPath == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.EtcdKeyPath cannot be empty", c)
	}

	apiUrl, err := url.Parse(c.ApiEndpoint)
	if err != nil {
		return nil, microerror.Maskf(err, "cannot parse api endpoint into url.URL")
	}
	etcdUrl, err := url.Parse(c.EtcdEndpoint)
	if err != nil {
		return nil, microerror.Maskf(err, "cannot parse etcd endpoint into url.URL")
	}

	apiCertPair, err := tls.LoadX509KeyPair(c.ApiCertPath, c.ApiKeyPath)
	if err != nil {
		return nil, microerror.Maskf(err, "failed to load api cert key pair")
	}
	etcdCertPair, err := tls.LoadX509KeyPair(c.EtcdCertPath, c.EtcdKeyPath)
	if err != nil {
		return nil, microerror.Maskf(err, "failed to load etcd cert key pair")
	}
	apiCaCert, err := CertPoolFromFile(c.EtcdCACertPath)
	if err != nil {
		return nil, microerror.Maskf(err, "failed to load etcd cert key pair")
	}
	etcdCaCert, err := CertPoolFromFile(c.EtcdCACertPath)
	if err != nil {
		return nil, microerror.Maskf(err, "failed to load etcd cert key pair")
	}

	apiHttpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{apiCertPair},
			ClientCAs:          apiCaCert,
			InsecureSkipVerify: false,
		},
		MaxIdleConns: maxIdleConnection,
	}
	apiHttpClient := &http.Client{
		Transport: apiHttpTransport,
		Timeout:   maxTimeoutSec * time.Second,
	}

	etcdHttpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{etcdCertPair},
			ClientCAs:          etcdCaCert,
			InsecureSkipVerify: false,
		},
		MaxIdleConns: maxIdleConnection,
	}
	etcdHttpClient := &http.Client{
		Transport: etcdHttpTransport,
		Timeout:   maxTimeoutSec * time.Second,
	}

	h := &Healthz{
		apiUrl:  apiUrl,
		etcdUrl: etcdUrl,
		port:    c.Port,

		apiHttpClient:     apiHttpClient,
		apiHttpTransport:  apiHttpTransport,
		etcdHttpClient:    etcdHttpClient,
		etcdHttpTransport: etcdHttpTransport,
	}
	return h, nil
}

// BootHealthServer will start http server with `/healthz` endpoint
func (h *Healthz) Boot(ctx context.Context) error {
	listenOn := fmt.Sprintf(":%d", h.port)

	http.HandleFunc("/healthz", h.handleHealthCheck)
	err := http.ListenAndServe(listenOn, nil)

	return microerror.Mask(err)
}

func (h *Healthz) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if h.etcdHealthCheck() && h.apiHealthCheck() {
		fmt.Fprint(w, "OK")
	} else {
		http.Error(w, "FAILED", http.StatusInternalServerError)
	}
}

func (h *Healthz) apiHealthCheck() bool {
	// be sure to close idle connection after health check is finished
	defer h.apiHttpTransport.CloseIdleConnections()

	req, err := http.NewRequest("GET", h.apiUrl.String(), nil)
	if err != nil {
		panic(fmt.Sprintf("unable to construct health check request: %q", err))
	}
	// close connection after health check request (the TCP connection gets
	// closed by deferred s.tr.CloseIdleConnections()).
	req.Header.Add("Connection", "close")

	// send request to http endpoint
	_, err = h.apiHttpClient.Do(req)
	if err != nil {
		// check failed
		return false
	}
	// all OK
	return true
}

func (h *Healthz) etcdHealthCheck() bool {
	// be sure to close idle connection after health check is finished
	defer h.etcdHttpTransport.CloseIdleConnections()

	req, err := http.NewRequest("GET", h.etcdUrl.String(), nil)
	if err != nil {
		panic(fmt.Sprintf("unable to construct health check request: %q", err))
	}
	// close connection after health check request (the TCP connection gets
	// closed by deferred s.tr.CloseIdleConnections()).
	req.Header.Add("Connection", "close")

	// send request to http endpoint
	_, err = h.apiHttpClient.Do(req)
	if err != nil {
		// check failed
		return false
	}
	// all OK
	return true
}
