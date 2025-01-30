ARG BASE_IMAGE=go-graphql_galaxy-base
FROM ${BASE_IMAGE} as builder

# FROM scratch
# ! DEBUG
FROM alpine:latest

# ! DEBUG
RUN apk add --no-cache bash

WORKDIR /bin
COPY --from=builder /bin/app /bin/app

# Local env (local db)
COPY --from=builder /src/config/compose-local.yaml /config/compose-local.yaml
# Local env (galaxy db)
COPY --from=builder /src/config/compose-galaxy.yaml /config/compose-galaxy.yaml

# Cloud kubernetes
COPY --from=builder /src/config/cloud-galaxy.yaml /config/cloud-galaxy.yaml

EXPOSE 8080

ENTRYPOINT ["/bin/app"]