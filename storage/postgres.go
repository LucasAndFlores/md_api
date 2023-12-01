package storage

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func NewPostgresConnection(ctx context.Context) (*pgxpool.Pool, error) { 
     
    var connectionString = os.Getenv("POSTGRES_CONNECTION_STRING")

    if connectionString == "" {
        return nil, errors.New("The env variable POSTGRES_CONNECTION_STRING is not defined")
    }

    conn, err := pgxpool.New(ctx, connectionString)

    if err != nil {
        log.Println("Something happen with db connection: ", err)
        return nil, err
    }

    return conn, nil  
}
