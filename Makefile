##! Compose
up_local:
	docker-compose -f docker-compose-local.yaml up -d
up_galaxy:
	docker-compose -f docker-compose-galaxy.yaml up -d
down_local:
	docker-compose -f docker-compose-local.yaml down
down_galaxy:
	docker-compose -f docker-compose-galaxy.yaml down
down: down_local down_galaxy

##! Build docker images
_build_core_image:
	docker build --file=Dockerfile.core --tag=go-graphql_galaxy-core:latest .
_build_base_image: _build_core_image
	docker build --file=Dockerfile.base --tag=go-graphql_galaxy-base:latest .
build_app_image: _build_base_image
	docker build --file=Dockerfile.app --tag=go-graphql_galaxy-app:latest .

##! Local app build & run
_bin_exists:
	mkdir -p bin
local_build: _bin_exists
	go build -o bin/server cmd/server/main.go

local_run: local_build
	APP_CONFIG_PATH=config/local.yaml ./bin/server

local_run_galaxy: local_build
	APP_CONFIG_PATH=config/local-galaxy.yaml ./bin/server

##! Local gql go generator
_build_generate: bin_exists
	go build -o bin/generate cmd/generate/main.go
generate: _build_generate
	./bin/generate
# go get github.com/99designs/gqlgen
# go run github.com/99designs/gqlgen generate

##! Local integration tests
genqlient_generate:
	go get github.com/Khan/genqlient
	go run github.com/Khan/genqlient

integration_test:
	APP_CONFIG_PATH=../config/local.yaml go test tests/integration_test.go

integration_test_verbose:
	APP_CONFIG_PATH=../config/local.yaml go test tests/integration_test.go -v

##! sh to images
sh_base:
	docker run --rm -it --entrypoint sh go-graphql_galaxy-base
sh_app:
	docker run --rm -it --entrypoint sh go-graphql_galaxy-app

##! DB Operations
local_conn_string = "postgres://postgres:postgres@localhost:5432/postgres"
galaxy_conn_string = "postgres://arthur:xvQqwww2Kczb7cuJ2dvfPy15abC@dontpanic.k42.app/galaxy"

dump_galaxy:
	@pg_dump $(galaxy_conn_string) --exclude-table-data='*_id_seq' -f sql/db_dump_latest.sql
	@echo "Created sql/db_dump_latest.sql"
	@echo -e "\033[0;31m!!! You should (remove/change to postgres) all \"cloudsuperuser\" and \"arthur\" occurrences !!!\033[0m"

migrate_db:
	@psql $(local_conn_string) -f ./sql/db_dump.sql > /dev/null
	@echo -e "\033[0;32mMigrated.\033[0m"

clean_db:
	docker volume rm go-graphql_galaxy_pgdata

psql_local:
	@psql $(local_conn_string)

psql_galaxy:
	@psql $(galaxy_conn_string)
