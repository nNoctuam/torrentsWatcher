#!/bin/bash

protoc -I=./protobuf --go_out=internal/pb  --js_out=import_style=commonjs,binary:frontend/src/pb ./protobuf/*.proto
