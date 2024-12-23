proto:
    rm -f pb/*.ts
    rm -f pb/google/protobuf/*.ts
	protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=pb \
    --ts_proto_opt=nestJs=true \
    --proto_path=src/proto/product-service \
    --proto_path=/usr/local/include \
    src/proto/product-service/*.proto

.PHONY: proto
