ports:
* frontend: `env PORT_HTTP` => `inner 10000` (go app)
* gRPC: `env PORT_GRPC` => `inner 10001` envoy => `inner 10002` (go app)


