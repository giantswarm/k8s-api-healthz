package main

import (
	"context"
	"fmt"
	"os"

	"github.com/giantswarm/k8s-api-healthz/pkg/project"

	"github.com/giantswarm/microerror"
	flag "github.com/spf13/pflag"
)

type Flag struct {
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

func main() {
	err := mainError()
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
}

func mainError() error {
	var err error

	var f Flag
	flag.StringVar(&f.ApiEndpoint, "api-endpoint", "https://127.0.0.1:443/healthz", "K8s API url that will be used for the API connection. Only a secure connection is supported.")
	flag.StringVar(&f.ApiCertPath, "api-cert", "/etc/kubernetes/ssl/apiserver-crt.pem", "Path to the client certificate file to authenticate with against the API.")
	flag.StringVar(&f.ApiCACertPath, "api-ca-cert", "/etc/kubernetes/ssl/apiserver-ca.pem", "Path to the CA file containing the issuer for the client certificate passed via --api-cert.")
	flag.StringVar(&f.ApiKeyPath, "api-key", "/etc/kubernetes/ssl/apiserver-key.pem", "Path to the private key file to authenticate with against the API.")
	flag.StringVar(&f.EtcdEndpoint, "etcd-endpoint", "https://127.0.0.1:2379/health", "URL that will be used for the etcd connection. Only a secure connection is supported.")
	flag.StringVar(&f.EtcdCertPath, "etcd-cert", "/etc/kubernetes/ssl/etcd/client-crt.pem", "Path to the client certificate file to authenticate with against etcd.")
	flag.StringVar(&f.EtcdCACertPath, "etcd-ca-cert", "/etc/kubernetes/ssl/etcd/client-ca.pem", "Path to the CA file containing the issuer of the client certificate passed via --etcd-cert.")
	flag.StringVar(&f.EtcdKeyPath, "etcd-key", "/etc/kubernetes/ssl/etcd/client-key.pem", "Path to the private key file to authenticate with against etcd.")
	flag.IntVar(&f.Port, "port", 8089, "TCP port on which the HTTP server will be listening.")

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("%s:%s - %s", project.Name(), project.Version(), project.GitSHA())
		return nil
	}
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		flag.Usage()
		return nil
	}
	flag.Parse()

	var healthz *Healthz
	{
		healthzConfig := HealthzConfig{
			ApiEndpoint:    f.ApiEndpoint,
			ApiCertPath:    f.ApiCertPath,
			ApiCACertPath:  f.ApiCACertPath,
			ApiKeyPath:     f.ApiKeyPath,
			EtcdEndpoint:   f.EtcdEndpoint,
			EtcdCertPath:   f.EtcdCertPath,
			EtcdCACertPath: f.EtcdCACertPath,
			EtcdKeyPath:    f.EtcdKeyPath,
			Port:           f.Port,
		}

		healthz, err = NewHealthz(healthzConfig)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	ctx := context.Background()
	err = healthz.Boot(ctx)
	if err != nil {
		return microerror.Mask(err)
	}
	return nil
}
