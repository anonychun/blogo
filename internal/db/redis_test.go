package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedisClient(t *testing.T) {
	c, err := NewRedisClient()
	require.NoError(t, err)
	require.NotNil(t, c)

	t.Run("get conn", func(t *testing.T) {
		assert.NotNil(t, c.Conn())
	})

	t.Run("get cache", func(t *testing.T) {
		assert.NotNil(t, c.Cache())
	})

	t.Run("close", func(t *testing.T) {
		assert.NoError(t, c.Close())
	})
}
