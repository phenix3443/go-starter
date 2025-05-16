ARG GO_VERSION=1.24.3-bookworm

FROM golang:${GO_VERSION} AS builder

WORKDIR /workspace
COPY . .
RUN make build

FROM golang:${GO_VERSION}
RUN apt-get update && apt-get install -y jq mariadb-client
COPY --from=builder /workspace/build/main /usr/bin/main
CMD ["main"]
