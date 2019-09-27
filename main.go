package main

import (
	"flag"
	"fmt"

	"github.com/giantswarm/microerror"
)

func main() {
	err := mainError()
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
}

func mainError() error {
	var err error

	apiEndpoint := flag.String("api-endpoint", "https://127.0.0.1:443/healthz", "K8s API url that will be used for the api connection. Only secure connection is supported.")
	apiCertPath := flag.String("api-cert", "/etc/kubernetes/ssl/apiserver-crt.pem", "Path to the cert file to authenticate against api.")
	apiCACertPath := flag.String("api-ca-cert", "/etc/kubernetes/ssl/apiserver-ca.pem", "Path to the cacert file to authenticate against api.")
	apiKeyPath := flag.String("api-key", "/etc/kubernetes/ssl/apiserver-key.pem", "Path to the key file to authenticate against api.")
	etcdEndpoint := flag.String("etcd-endpoint", "https://127.0.0.1:2379/health", "ETCD url that will be used for the etcd connection. Only secure connection is supported.")
	etcdCertPath := flag.String("etcd-cert", "/etc/kubernetes/ssl/apiserver-crt.pem", "Path to the cert file to authenticate against etcd.")
	etcdCACertPath := flag.String("etcd-ca-cert", "/etc/kubernetes/ssl/apiserver-ca.pem", "Path to the cacert file to authenticate against etcd.")
	etcdKeyPath := flag.String("etcd-key", "/etc/kubernetes/ssl/apiserver-key.pem", "Path to the key file to authenticate against etcd.")
	port := flag.Int("port", 8089, "Define a port on which the http server will be running.")
	flag.Parse()

	var healthz *Healthz
	{
		healthzConfig := HealtzConfig{
			ApiEndpoint:    *apiEndpoint,
			ApiCertPath:    *apiCertPath,
			ApiCACertPath:  *apiCACertPath,
			ApiKeyPath:     *apiKeyPath,
			EtcdEndpoint:   *etcdEndpoint,
			EtcdCertPath:   *etcdCertPath,
			EtcdCACertPath: *etcdCACertPath,
			EtcdKeyPath:    *etcdKeyPath,
			Port:           *port,
		}

		healthz, err = NewHealtz(healthzConfig)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = healthz.BootHealthServer()
	if err != nil {
		return microerror.Mask(err)
	}
	return nil
}
