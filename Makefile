proto-product:
	rm -f $(name)-service/product/*.go
	protoc --proto_path=proto --go_out=$(name)-service/product --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/product --go-grpc_opt=paths=source_relative \
    proto/product-service/*.proto

proto-user:
	rm -f $(name)-service/user/*.go
	protoc --proto_path=proto --go_out=$(name)-service/user --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/user --go-grpc_opt=paths=source_relative \
    proto/auth_service.proto

proto-cart:
	rm -f $(name)-service/cart/*.go
	protoc --proto_path=proto --go_out=$(name)-service/cart --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/cart --go-grpc_opt=paths=source_relative \
    proto/cart-service/*.proto

proto-payment:
	rm -f $(name)-service/payment/*.go
	protoc --proto_path=proto --go_out=$(name)-service/payment --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/payment --go-grpc_opt=paths=source_relative \
    proto/payment/*.proto

proto-notification:
	rm -f $(name)-service/notification/*.go
	protoc --proto_path=proto/notification --go_out=$(name)-service/notification --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/notification --go-grpc_opt=paths=source_relative \
    proto/notification/*.proto

proto-auction:
	rm -f $(name)-service/auction/*.go
	protoc --proto_path=proto/auction --go_out=$(name)-service/auction --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/auction --go-grpc_opt=paths=source_relative \
    proto/auction/*.proto


protot-product:
	protoc --plugin=./$(service)/node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=$(service)/pb/ \
    --ts_proto_opt=nestJs=true \
    --proto_path=proto/product \
    --proto_path=/usr/local/include \
    proto/product/*.proto

protot-cart:
	protoc --plugin=./$(service)/node_modules/.bin/protoc-gen-ts_proto \
	--ts_proto_out=$(service)/pb \
	--ts_proto_opt=nestJs=true \
	--proto_path=proto/cart \
	--proto_path=/usr/local/include \
	proto/cart/*.proto

protot-auth:
	protoc --plugin=./$(service)/node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=$(service)/pb \
    --ts_proto_opt=nestJs=true \
    --proto_path=proto/pb \
    --proto_path=/usr/local/include \
    proto/pb/*.proto

rm-product:
	rm -f $(service)/pb/$(service)/*.ts

