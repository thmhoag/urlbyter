# Start from goreleaser so we can pull it's entrypoint
# and binary
FROM goreleaser/goreleaser as goreleaser

FROM golang:1.14.14-alpine as build-sdk

RUN apk add --no-cache \
    bash \
    curl \
    git

COPY --from=goreleaser /entrypoint.sh /entrypoint.sh
COPY --from=goreleaser /bin/goreleaser /bin/goreleaser

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "-h" ]