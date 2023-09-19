package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
	"wbTech/cmd/config"
)

func PostgresConn(ctx context.Context, maxAttempts int, sc config.StorageConfig) (pool *pgx.Conn, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	err = DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgx.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil

	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}
	return pool, nil
}
