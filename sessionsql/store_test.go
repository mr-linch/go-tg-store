package sessionsql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mr-linch/go-tg-store/sessionsql"
	"github.com/stretchr/testify/assert"
)

func TestStoreQueries(t *testing.T) {
	tableName := "session"
	for _, test := range []struct {
		Name    string
		Queries sessionsql.Queries
	}{
		{"SQLite3", sessionsql.SQLite3},
		{"PostgreSQL", sessionsql.PostgreSQL},
		{"MySQL", sessionsql.MySQL},
	} {
		t.Run(test.Name, func(t *testing.T) {
			t.Run("Setup", func(t *testing.T) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				assert.NoError(t, err)
				defer db.Close()

				store := sessionsql.New(db, tableName, test.Queries)

				mock.
					ExpectExec(fmt.Sprintf(test.Queries.Setup, tableName)).
					WillReturnResult(sqlmock.NewResult(0, 0))

				err = store.Setup(context.Background())
				assert.NoError(t, err)

				assert.NoError(t, mock.ExpectationsWereMet())
			})

			t.Run("Set", func(t *testing.T) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				assert.NoError(t, err)
				defer db.Close()

				store := sessionsql.New(db, tableName, test.Queries)

				mock.
					ExpectExec(fmt.Sprintf(test.Queries.Set, tableName)).
					WithArgs("key", []byte("value")).
					WillReturnResult(sqlmock.NewResult(0, 0))

				err = store.Set(context.Background(), "key", []byte("value"))
				assert.NoError(t, err)

				assert.NoError(t, mock.ExpectationsWereMet())
			})

			t.Run("Get", func(t *testing.T) {
				t.Run("Found", func(t *testing.T) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					assert.NoError(t, err)
					defer db.Close()

					store := sessionsql.New(db, tableName, test.Queries)

					mock.
						ExpectQuery(fmt.Sprintf(test.Queries.Get, tableName)).
						WithArgs("key").
						WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow([]byte("value")))

					v, err := store.Get(context.Background(), "key")
					assert.NoError(t, err)
					assert.Equal(t, []byte("value"), v)

					assert.NoError(t, mock.ExpectationsWereMet())
				})

				t.Run("NotFound", func(t *testing.T) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					assert.NoError(t, err)
					defer db.Close()

					store := sessionsql.New(db, tableName, test.Queries)

					mock.
						ExpectQuery(fmt.Sprintf(test.Queries.Get, tableName)).
						WithArgs("key").
						WillReturnError(sql.ErrNoRows)

					v, err := store.Get(context.Background(), "key")
					assert.NoError(t, err)
					assert.Nil(t, v)

					assert.NoError(t, mock.ExpectationsWereMet())
				})

				t.Run("OtherError", func(t *testing.T) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					assert.NoError(t, err)
					defer db.Close()

					store := sessionsql.New(db, tableName, test.Queries)

					mock.
						ExpectQuery(fmt.Sprintf(test.Queries.Get, tableName)).
						WithArgs("key").
						WillReturnError(sql.ErrConnDone)

					v, err := store.Get(context.Background(), "key")
					assert.ErrorIs(t, err, sql.ErrConnDone)
					assert.Nil(t, v)

					assert.NoError(t, mock.ExpectationsWereMet())
				})
			})

			t.Run("Del", func(t *testing.T) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				assert.NoError(t, err)
				defer db.Close()

				store := sessionsql.New(db, tableName, test.Queries)

				mock.
					ExpectExec(fmt.Sprintf(test.Queries.Del, tableName)).
					WithArgs("key").
					WillReturnResult(sqlmock.NewResult(0, 1))

				err = store.Del(context.Background(), "key")
				assert.NoError(t, err)

				assert.NoError(t, mock.ExpectationsWereMet())
			})
		})
	}
}
