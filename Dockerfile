FROM alpine:3.19

RUN apk add --no-cache ca-certificates

ADD ./k8s-api-healthz  /k8s-api-healthz 

ENTRYPOINT ["/k8s-api-healthz"]
