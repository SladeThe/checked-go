package must

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/SladeThe/checked-go/consts"
)

func TestInt64ToInt(t *testing.T) {
	require.Panics(t, func() { Int64ToInt(consts.MinInt - 1) })
	require.Equal(t, consts.MinInt, Int64ToInt(consts.MinInt))
	require.Equal(t, 0, Int64ToInt(0))
	require.Equal(t, consts.MaxInt, Int64ToInt(consts.MaxInt))
	require.Panics(t, func() { Int64ToInt(consts.MaxInt + 1) })
}

func TestUint32ToInt(t *testing.T) {
	require.Equal(t, 0, Uint32ToInt(0))
	require.Equal(t, consts.MaxInt, Uint32ToInt(consts.MaxInt))
	require.Panics(t, func() { Uint32ToInt(consts.MaxInt + 1) })
}
