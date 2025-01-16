rabbitmq-docker:
	docker run -d --hostname rabbitmq --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:3-management

rabbitmq:
	docker start rabbitmq

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_notification

dropdb: 
	docker exec -it postgres17 dropdb e_commerce_notification

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_notification?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_notification?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

server:
	go run main.go

