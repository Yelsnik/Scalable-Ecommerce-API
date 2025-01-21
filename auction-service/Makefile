postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_auction

dropdb: 
	docker exec -it postgres17 dropdb e_commerce_auction

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_auction?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_auction?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

server:
	go run main.go

evans:
	evans --host localhost --port 3030 -r

.PHONY: postgres createdb dropdb migrateup migratedown new_migration sqlc server evans