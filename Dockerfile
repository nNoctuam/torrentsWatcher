FROM node:13 AS frontendBuilder
WORKDIR /var/torrentsWatcherFrontend
#RUN npm i -g @vue/cli
ADD frontend/package.json .
ADD frontend/package-lock.json .
RUN npm ci
ADD frontend .
RUN npm run build-prod

FROM golang:1.17 AS builder
RUN apt-get update && apt-get install -y ca-certificates openssl
WORKDIR /var/torrentsWatcher
COPY go.mod go.sum ./
COPY vendor ./vendor/
COPY config ./config/
COPY cmd ./cmd/
COPY internal ./internal/
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags '-static' -s -w" -o /go/bin/torrentsWatcher ./cmd/torrentsWatcher/main.go

FROM envoyproxy/envoy:v1.17.0
# configurations
WORKDIR /
# the main program:
COPY docker-run.sh /
COPY --from=frontendBuilder /var/torrentsWatcherFrontend/dist /dist
COPY --from=builder /go/bin/torrentsWatcher /torrentsWatcher
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY envoy.yaml /etc/envoy/
RUN touch /dist/config.json && chown envoy /dist/config.json
CMD ["/docker-run.sh"]
