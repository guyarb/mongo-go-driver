package conn_test

import (
	"testing"

	"github.com/10gen/mongo-go-driver/yamgo/internal/conntest"
	"github.com/10gen/mongo-go-driver/yamgo/internal/testutil"
	. "github.com/10gen/mongo-go-driver/yamgo/private/conn"
	"github.com/stretchr/testify/require"
)

func TestTracked_Inc(t *testing.T) {
	t.Parallel()

	c := &conntest.MockConnection{}
	require.True(t, c.Alive())

	tracked := Tracked(c)
	require.True(t, tracked.Alive())
	require.True(t, c.Alive())

	tracked.Inc()
	require.True(t, tracked.Alive())
	require.True(t, c.Alive())

	testutil.RequireNoErrorOnClose(t, tracked)
	require.True(t, tracked.Alive())
	require.True(t, c.Alive())

	testutil.RequireNoErrorOnClose(t, tracked)
	require.False(t, tracked.Alive())
	require.False(t, c.Alive())
}
