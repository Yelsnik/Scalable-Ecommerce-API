proto:
	protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=pb \
    --ts_proto_opt=nestJs=true \
    --proto_path=src/proto/product-service \
    --proto_path=/usr/local/include \
    src/proto/product-service/*.proto

proto-cart:
    protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=pb \
    --ts_proto_opt=nestJs=true \
    --proto_path=src/proto/cart-service \
    --proto_path=/usr/local/include \
    src/proto/cart-service/*.proto

removeProto:
    rm -f pb/*.ts
    rm -f pb/google/protobuf/*.ts

.PHONY: proto
