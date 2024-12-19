postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_cart

dropdb: 
	docker exec -it postgres17 dropdb simple_bank

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


evans:
	evans --host localhost --port 9090 -r

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server new_migration mock proto