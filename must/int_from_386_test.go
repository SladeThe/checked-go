package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64ToInt(t *testing.T) {
	require.Panics(t, func() { Int64ToInt(math.MinInt - 1) })
	require.Equal(t, math.MinInt, Int64ToInt(math.MinInt))
	require.Equal(t, 0, Int64ToInt(0))
	require.Equal(t, math.MaxInt, Int64ToInt(math.MaxInt))
	require.Panics(t, func() { Int64ToInt(math.MaxInt + 1) })
}

func TestUint32ToInt(t *testing.T) {
	require.Equal(t, 0, Uint32ToInt(0))
	require.Equal(t, math.MaxInt, Uint32ToInt(math.MaxInt))
	require.Panics(t, func() { Uint32ToInt(math.MaxInt + 1) })
}
