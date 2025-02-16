# install golang-migrate cli to manage sql migrations
install-migrate-cli:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz migrate
	sudo mv migrate /usr/bin

NAME ?= create_name_table

create-migration:
	migrate create -ext sql -dir database/migrations -seq ${NAME}



