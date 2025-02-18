package services

import (
	database "backend/database/gen"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	ClientDB *database.Queries
)

// TODO: switch to connection pool to be safe under goroutinces
func CreateConnection(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic("Database not found xdd")
	}
	ClientDB = database.New(conn)
	return conn
}
