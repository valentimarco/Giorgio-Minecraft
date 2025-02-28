FROM cosmtrek/air:v1.61.7

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest