package sessionbolt_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/mr-linch/go-tg-store/sessionbolt"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

func TestStore(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "go-tg")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "sessions.boltdb")

	db, err := bbolt.Open(path, 0666, nil)
	require.NoError(t, err)
	defer db.Close()

	store := sessionbolt.New(db, "sessions")
	require.NoError(t, err)

	err = store.Set(context.Background(), "key", []byte("value"))
	require.NoError(t, err)

	v, err := store.Get(context.Background(), "key")
	require.NoError(t, err)
	require.Equal(t, []byte("value"), v)

	err = store.Del(context.Background(), "key")
	require.NoError(t, err)

	v, err = store.Get(context.Background(), "key")
	require.NoError(t, err)
	require.Nil(t, v)
}
