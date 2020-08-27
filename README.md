[![CircleCI](https://circleci.com/gh/giantswarm/k8s-api-healthz.svg?style=shield)](https://circleci.com/gh/giantswarm/k8s-api-healthz) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/k8s-api-healthz/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/k8s-api-healthz)

# k8s-api-healthz

This is simple service thats suppose to be running on the master node in order to properly and securely check health of the API via http.
Health check is done by accessing the `/healthz` endpoint on the k8s API and etcd endpoint with use of certificates in order to do proper auth over https.

The reasoning behind creating yet another health service is that in cloud providers such as AWS or Azure its very hard to securely access the `/healthz` endpoint on the api as with disabled both `anonymous-access` and `insecure port`.  And in order to access `/healthz` endpoint on such https  you need to provide credentials/certs.

So instead of accessing directly k8s api or etcd api you access this service via  http.
