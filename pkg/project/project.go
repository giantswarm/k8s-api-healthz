package project

var (
	description        = "The k8s-api-healthz provides healthcehck to k8s api and etcd and provides the result on unsecure http port for proper LB health checks."
	gitSHA             = "n/a"
	name        string = "k8s-api-healthz"
	source      string = "https://github.com/giantswarm/k8s-api-healthz"
	version            = "0.2.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
