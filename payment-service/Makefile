postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root e_commerce_payment

dropdb: 
	docker exec -it postgres17 dropdb e_commerce_payment

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_payment?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/e_commerce_payment?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

mock:
	mockgen -package mockdb -destination db/mock/store.go payment-service/db/sqlc Store
	mockgen -package clientdb -destination client/mock/client.go payment-service/client ClientInterface
	mockgen -package helpmck -destination gapi/helpers/mock/helper_mock.go payment-service/gapi/helpers Helper
	mockgen -package stripemck -destination stripe/mock/stripe_mock.go payment-service/stripe Stripe


sqlc:
	sqlc generate

server:
	go run main.go

evans:
	evans --host localhost --port 6060 -r

proto:
	rm -f payment/*.go
	protoc --proto_path=proto --go_out=payment --go_opt=paths=source_relative \
    --go-grpc_out=payment --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: postgres createdb dropdb migrateup new_migration sqlc server evans proto