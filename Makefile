postgres:
	docker run --name service_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it service_postgres createdb --username=root --owner=root station_service

dropdb:
	docker exec -it service_postgres dropdb station_service

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/station_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/station_service?sslmode=disable" -verbose down

test: 
	go test -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, test