proto-product:
	rm -f $(name)-service/product/*.go
	protoc --proto_path=proto --go_out=$(name)-service/product --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/product --go-grpc_opt=paths=source_relative \
    proto/product-service/*.proto

proto-user:
	rm -f $(name)-service/user/user-service/*.go
	protoc --proto_path=proto --go_out=$(name)-service/user --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/user --go-grpc_opt=paths=source_relative \
    proto/user-service/*.proto

proto-cart:
	rm -f $(name)-service/cart/*.go
	protoc --proto_path=proto --go_out=$(name)-service/cart --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/cart --go-grpc_opt=paths=source_relative \
    proto/cart-service/*.proto

proto-payment:
	rm -f $(name)-service/payment/*.go
	protoc --proto_path=proto --go_out=$(name)-service/payment --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/payment --go-grpc_opt=paths=source_relative \
    proto/payment-service/*.proto