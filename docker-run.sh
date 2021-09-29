echo "{\
        \"PORT_GRPC\": \"${PORT_GRPC}\"\
      }" > /app/dist/config.json
/app/torrentsWatcher &
envoy -c /etc/envoy/envoy.yaml
