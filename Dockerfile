FROM node:13 AS frontendBuilder
WORKDIR /var/torrentsWatcherFrontend
RUN npm i -g @vue/cli && npm ci
ADD frontend .
RUN npm ci && npm run build

FROM golang:1.16 AS builder
WORKDIR /var/torrentsWatcher
ADD . .
COPY --from=frontendBuilder /var/torrentsWatcherFrontend/dist /var/torrentsWatcher/cmd/torrentsWatcher/dist
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags '-static' -s -w" -o /go/bin/torrentsWatcher ./cmd/torrentsWatcher/main.go

FROM scratch
# configurations
WORKDIR /
# the main program:
COPY --from=builder /go/bin/torrentsWatcher /torrentsWatcher
CMD ["/torrentsWatcher"]
