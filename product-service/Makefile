proto:
    rm -f pb/*.ts
	protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=pb \
    --proto_path=src/proto \
    src/proto/*.proto

.PHONY: proto