postgres:
	docker run --name=simple_bank_db -p 5436:5432 -e POSTGRES_USER='root' -e PGUSER='root' -e POSTGRES_PASSWORD='secret' -d postgres:latest

createdb:
	winpty docker exec -it simple_bank_db createdb --username=root --owner=root simple_bank_db

dropdb:
	winpty docker exec -it simple_bank_db dropdb simple_bank_db

migrateup:
	migrate -path schema -database "postgresql://root:secret@localhost:5436/simple_bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path schema -database "postgresql://root:secret@localhost:5436/simple_bank_db?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown