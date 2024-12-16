postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_user

dropdb: 
	docker exec -it postgres17 dropdb e_commerce_user

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_user?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_user?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

server:
	go run main.go

evans:
	evans --host localhost --port 9090 -r

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb  --grpc-gateway_opt paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migrateup new_migration sqlc server evans proto