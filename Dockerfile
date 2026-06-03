FROM alpine:3.23

ARG TARGETARCH

RUN apk add --no-cache ca-certificates

COPY ./k8s-api-healthz-linux-${TARGETARCH}  /k8s-api-healthz 

ENTRYPOINT ["/k8s-api-healthz"]
