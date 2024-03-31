.PHONY: postgresinit
postgresinit:
	docker run --name chatapp -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

.PHONY: postgres
postgres:
	docker exec -it chatapp psql

.PHONY: createdb
createdb:
	docker exec -it chatapp createdb --username=root --owner=root go-chat

.PHONY: dropdb
dropdb:
	docker exec -it chatapp dropdb go-chat

.PHONY: mygrateup
mygrateup:
	migrate -path=db/migrations -database="postgresql://root:password@localhost:5433/go-chat?sslmode=disable" up

.PHONY: mygratedown
mygratedown:
	migrate -path=db/migrations -database="postgresql://root:password@localhost:5433/go-chat?sslmode=disable" down     