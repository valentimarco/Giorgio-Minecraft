package database

import (
	database "backend/database/gen"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct{
	*pgx.Conn
	*database.Queries
}

// TODO: switch to connection pool to be safe under goroutinces
func CreateConnection(ctx context.Context) *DB {
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic("Database not found xdd")
	}
	
	return &DB{
		conn,
		database.New(conn),
	}
	
}
