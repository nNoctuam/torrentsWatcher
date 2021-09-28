echo PORT_HTTP = $PORT_HTTP
echo PORT_GRPC = $PORT_GRPC
whoami
ls -la /dist
echo "{\
        \"PORT_GRPC\": \"${PORT_GRPC}\"\
      }" > /dist/config.json
/torrentsWatcher &
envoy -c /etc/envoy/envoy.yaml
