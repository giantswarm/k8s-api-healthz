package main

import (
	"context"
	"fmt"
	"github.com/giantswarm/k8s-api-healthz/pkg/project"
	"os"

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
	flag.StringVar(&f.ApiEndpoint, "api-endpoint", "https://127.0.0.1:443/healthz", "K8s API url that will be used for the api connection. Only secure connection is supported.")
	flag.StringVar(&f.ApiCertPath, "api-cert", "/etc/kubernetes/ssl/apiserver-crt.pem", "Path to the cert file to authenticate against api.")
	flag.StringVar(&f.ApiCACertPath, "api-ca-cert", "/etc/kubernetes/ssl/apiserver-ca.pem", "Path to the cacert file to authenticate against api.")
	flag.StringVar(&f.ApiKeyPath, "api-key", "/etc/kubernetes/ssl/apiserver-key.pem", "Path to the key file to authenticate against api.")
	flag.StringVar(&f.EtcdEndpoint, "etcd-endpoint", "https://127.0.0.1:2379/health", "ETCD url that will be used for the etcd connection. Only secure connection is supported.")
	flag.StringVar(&f.EtcdCertPath, "etcd-cert", "/etc/kubernetes/ssl/apiserver-crt.pem", "Path to the cert file to authenticate against etcd.")
	flag.StringVar(&f.EtcdCACertPath, "etcd-ca-cert", "/etc/kubernetes/ssl/apiserver-ca.pem", "Path to the cacert file to authenticate against etcd.")
	flag.StringVar(&f.EtcdKeyPath, "etcd-key", "/etc/kubernetes/ssl/apiserver-key.pem", "Path to the key file to authenticate against etcd.")
	flag.IntVar(&f.Port, "port", 8089, "Define a port on which the http server will be running.")

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
			ApiEndpoint:    f.EtcdEndpoint,
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
