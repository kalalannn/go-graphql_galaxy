up:
	docker-compose up -d

down:
	docker-compose down

clean_db:
	docker volume rm go-graphql_galaxy_pgdata

local_conn_string = "postgres://postgres:postgres@localhost:5432/postgres"
galaxy_conn_string = "postgres://arthur:xvQqwww2Kczb7cuJ2dvfPy15abC@dontpanic.k42.app/galaxy"

psql_local:
	psql $(local_conn_string)

psql_galaxy:
	psql $(galaxy_conn_string)

dump_galaxy:
	pg_dump $(galaxy_conn_string) --exclude-table-data='*_id_seq' -f sql/db_dump.sql

build:
	mkdir -p bin
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

generate:
	go run cmd/generate/main.go
# go get github.com/99designs/gqlgen
# go run github.com/99designs/gqlgen generate