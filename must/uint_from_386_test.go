package must

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/SladeThe/checked-go/consts"
)

func TestInt64ToUint(t *testing.T) {
	require.Panics(t, func() { Int64ToUint(-1) })
	require.Equal(t, uint(0), Int64ToUint(0))
	require.Equal(t, uint(consts.MaxUint), Int64ToUint(consts.MaxUint))
	require.Panics(t, func() { Int64ToUint(consts.MaxUint + 1) })
}

func TestUint64ToUint(t *testing.T) {
	require.Equal(t, uint(0), Uint64ToUint(0))
	require.Equal(t, uint(consts.MaxUint), Uint64ToUint(consts.MaxUint))
	require.Panics(t, func() { Uint64ToUint(consts.MaxUint + 1) })
}
