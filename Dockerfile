ARG GO_VERSION=1.24.3-bookworm

FROM golang:${GO_VERSION} AS builder

WORKDIR /workspace
COPY . .
RUN make build

FROM golang:${GO_VERSION}
RUN apt-get update && apt-get install -y jq mariadb-client

# Create a non-root user and switch to it
RUN useradd -m appuser
USER appuser

COPY --from=builder /workspace/build/main /usr/bin/main

# Add a HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD pgrep main || exit 1

CMD ["main"]
