postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root synapsis_challenge

dropdb:
	docker exec -it postgres16 dropdb synapsis_challenge

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/synapsis_challenge?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/synapsis_challenge?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/alifdwt/synapsis-backend-challenge/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown test server mock