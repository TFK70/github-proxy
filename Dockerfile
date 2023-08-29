FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o /app/build/github-proxy /app/cmd/github-proxy/main.go

FROM scratch

COPY --from=builder /app/build/github-proxy /github-proxy
COPY --from=builder /app/config.json /config.json
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/github-proxy"]
