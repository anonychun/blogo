package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresClient(t *testing.T) {
	c, err := NewPostgresClient()
	require.NoError(t, err)
	require.NotNil(t, c)

	t.Run("get conn", func(t *testing.T) {
		assert.NotNil(t, c.Conn())
	})

	t.Run("close", func(t *testing.T) {
		assert.NoError(t, c.Close())
	})
}
