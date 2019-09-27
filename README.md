<!--

    TODO:

    - Add the project to the CircleCI:
      https://circleci.com/setup-project/gh/giantswarm/REPOSITORY_NAME

    - Import RELEASE_TOKEN variable from template repository for the builds:
      https://circleci.com/gh/giantswarm/REPOSITORY_NAME/edit#env-vars

    - Change the badge (with style=shield):
      https://circleci.com/gh/giantswarm/REPOSITORY_NAME/edit#badges
      If this is a private repository token with scope `status` will be needed.

    - Change the top level header from `# template` to `# REPOSITORY_NAME` and
      add appropriate description.

    - If the repository is public consider adding godoc badge. This should be
      the first badge separated with a single space.
      [![GoDoc](https://godoc.org/github.com/giantswarm/REPOSITORY_NAME?status.svg)](http://godoc.org/github.com/giantswarm/REPOSITORY_NAME)

-->
[![CircleCI](https://circleci.com/gh/giantswarm/template.svg?style=shield&circle-token=cbabd7d13186f190fca813db4f0c732b026f5f6c)](https://circleci.com/gh/giantswarm/template)

# k8s api healthz
This is simple service thats suppose to be running on the master node in order to properly and securely check health of the API.
That is done by accessing the `/healthz` endpoint on the k8s API and etcd endpoint with use of certificates in order to do proper auth.


The reasoning behind creating yet another health service is that in cloud providers such as AWS or Azure its very hard to securely access the `/healthz` endpoint on the api as with secure solution  you have disabled both `anonymous-access` and `insecure port` and in order to access `/healthz` endpoint on such api you need to provide credentials.

