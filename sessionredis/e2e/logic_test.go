package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type Store interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}

func logic(t *testing.T, store Store) {
	t.Helper()

	err := store.Set(context.Background(), "key", []byte("value"))
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
