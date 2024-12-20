package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type Connection struct {
	Conn *pgx.Conn
}

func NewConnection(ctx context.Context) *Connection {
	config := NewConfig()

	conn, err := pgx.Connect(ctx, config.CreateDbUrl())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к БД: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Не удалось подключиться к постгре: %v\n", err)
	}

	return &Connection{Conn: conn}
}
