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
ADD . .
COPY --from=frontendBuilder /var/torrentsWatcherFrontend/dist /var/torrentsWatcher/cmd/torrentsWatcher/dist
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags '-static' -s -w" -o /go/bin/torrentsWatcher ./cmd/torrentsWatcher/main.go

FROM scratch
# configurations
WORKDIR /
# the main program:
COPY --from=builder /go/bin/torrentsWatcher /torrentsWatcher
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/torrentsWatcher"]
