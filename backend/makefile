# Apply .env file to makefile
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

docker_cmd = docker compose -f ../compose.dev.yml

NAME ?= create_name_table

up:
	$(docker_cmd) up -d
logs:
	$(docker_cmd) logs -f
down:
	$(docker_cmd) down
build:
	$(docker_cmd) build
generate-binds:
	$(docker_cmd) exec backend sqlc generate

create-entity:
	$(docker_cmd) exec backend migrate create -ext sql -dir database/migrations -seq ${NAME}
	
apply:
	$(docker_cmd) exec backend migrate -database ${DB_URL} -path database/migrations up
	$(docker_cmd) exec backend sqlc generate

