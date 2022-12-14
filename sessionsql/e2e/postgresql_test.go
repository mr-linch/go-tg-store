package e2e_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/mr-linch/go-tg-store/sessionsql"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestPostgreSQL(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")

	if dsn == "" {
		t.Skip("skip because POSTGRES_DSN is not set")
	}

	t.Logf("POSTGRES_DSN: %s", dsn)

	db, err := sql.Open("postgres", dsn)
	assert.NoError(t, err, "open db")
	defer db.Close()

	db.Exec("drop table if exists session")

	logic(t, db, sessionsql.PostgreSQL)
}
