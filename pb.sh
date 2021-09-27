#!/bin/bash

#protoc -I=./protobuf --go_out=internal/pb --go_opt=paths=source_relative ./protobuf/*.proto

# Path to this plugin
PROTOC_GEN_TS_PATH="frontend/node_modules/.bin/protoc-gen-ts"

protoc \
    -I=./protobuf \
    --go_out=internal/pb \
    --go_opt=paths=source_relative \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:frontend/src/pb" \
    --ts_out="frontend/src/pb" \
    ./protobuf/*.proto