FROM alpine:3.12
RUN apk add --no-cache ca-certificates bash
COPY serverscom-cloud-controller-manager /bin/serverscom-cloud-controller-manager
ENTRYPOINT ["/bin/serverscom-cloud-controller-manager"]
