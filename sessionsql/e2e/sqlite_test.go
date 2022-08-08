package e2e_test

import (
	"database/sql"
	"testing"

	"github.com/mr-linch/go-tg-store/sessionsql"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestSQLite3(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	assert.NoError(t, err, "db init")

	logic(t, db, sessionsql.SQLite3)
}
