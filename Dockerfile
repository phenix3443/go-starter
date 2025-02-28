FROM golang:1.22.4-bullseye AS builder

WORKDIR /workspace
COPY . .
RUN make build

FROM golang:1.22.4-bullseye
RUN apt-get update && apt-get install -y jq mariadb-client
COPY --from=builder /workspace/build/greeting /usr/bin/greeting
CMD ["greeting"]
