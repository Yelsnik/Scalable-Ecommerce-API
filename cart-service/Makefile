postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_cart

dropdb: 
	docker exec -it postgres17 dropdb e_commerce_cart

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_cart?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_cart?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Yelsnik/simplebank/db/sqlc Store

proto:
	rm -f cart/*.go
	protoc --proto_path=proto --go_out=cart --go_opt=paths=source_relative \
    --go-grpc_out=cart --go-grpc_opt=paths=source_relative \
    proto/*.proto

proto-product:
	rm -f product/*.go
	protoc --proto_path=proto --go_out=product --go_opt=paths=source_relative \
    --go-grpc_out=product --go-grpc_opt=paths=source_relative \
    proto/product/*.proto

evans:
	evans --host localhost --port 7070 -r


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server new_migration mock proto evans