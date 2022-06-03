#!/bin/bash

#protoc -I=./protobuf --go_out=internal/ports/adapters/driving/pb --go_opt=paths=source_relative ./protobuf/*.proto

# Path to this plugin
PROTOC_GEN_TS_PATH="frontend/node_modules/.bin/protoc-gen-ts"

protoc \
    -I=./protobuf \
    --go_opt=paths=source_relative \
    --go_out=internal/ports/adapters/driving/pb \
    --go-grpc_out=internal/ports/adapters/driving/pb \
    --go-grpc_opt=paths=source_relative \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:frontend/src/pb" \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:"frontend/src/pb" \
    ./protobuf/baseService.proto
#    --ts_out="frontend/src/pb" \
