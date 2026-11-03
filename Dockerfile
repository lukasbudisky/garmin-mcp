# Builder
FROM golang:1.26.4-alpine AS builder

WORKDIR /opt

COPY . /opt

RUN apk add --no-cache ca-certificates \
    && go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o "garmin-mcp"


# Garmin MCP
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /opt/garmin-mcp /garmin-mcp

EXPOSE 8080

ENTRYPOINT ["/garmin-mcp", "http"]

CMD ["-a", "0.0.0.0", "-p", "8080"]
