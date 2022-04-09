
CONTAINER_NAME := goreddit-db
DB_NAME := goreddit


postgres:
	docker run --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

startdb:
	docker start $(CONTAINER_NAME)

stopdb:
	docker stop $(CONTAINER_NAME)

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb --username=root --owner=root $(DB_NAME)

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down


.PHONY: postgres startdb stopdb createdb dropdb migrateup migratedown