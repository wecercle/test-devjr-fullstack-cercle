package connection

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	databaseQuery "github.com/wecercle/test-devjr-fullstack-cercle/core/database/postgres/query/sqlc"
)

var (
	once    sync.Once
	querier *databaseQuery.Queries
)

func Querier() *databaseQuery.Queries {
	once.Do(func() {
		db, err := sql.Open("pgx", DSN())
		if err != nil {
			panic(err)
		}
		querier = databaseQuery.New(db)
	})
	return querier
}

func DSN() string {
	host := getEnv("DATABASE_HOST", "127.0.0.1")
	// Force TCP instead of unix socket by using 127.0.0.1 if localhost is provided
	if host == "localhost" {
		host = "127.0.0.1"
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=cercle_test",
		getEnv("DATABASE_USER", "cercle_test"),
		getEnv("DATABASE_PASSWORD", "cercle_test"),
		host,
		getEnv("DATABASE_PORT", "5432"),
		getEnv("DATABASE_NAME", "cercle_test"),
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
