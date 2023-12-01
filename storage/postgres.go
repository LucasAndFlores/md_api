package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func NewPostgresConnection(ctx context.Context) (*pgxpool.Pool, error) { 
    connectionString := "postgresql://admin:p0stgr3s@db:5432/md_db"

    conn, err := pgxpool.New(ctx, connectionString)

    if err != nil {
        log.Println("Something happen with db connection: ", err)
        return nil, err
    }

    return conn, nil  
}
