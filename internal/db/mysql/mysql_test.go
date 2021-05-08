package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, c)

	t.Run("get conn", func(t *testing.T) {
		assert.NotNil(t, c.Conn())
	})

	t.Run("close conn", func(t *testing.T) {
		assert.NoError(t, c.Close())
	})
}
