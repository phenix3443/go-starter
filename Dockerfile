ARG GO_VERSION=1.24.0-bookworm

FROM golang:${GO_VERSION} AS builder

WORKDIR /workspace
COPY . .
RUN make build

FROM golang:${GO_VERSION}
RUN apt-get update && apt-get install -y jq mariadb-client
COPY --from=builder /workspace/build/server /usr/bin/server
CMD ["server"]
