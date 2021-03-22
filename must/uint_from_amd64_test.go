package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64ToUint(t *testing.T) {
	require.Panics(t, func() { Int64ToUint(-1) })
	require.Equal(t, uint(0), Int64ToUint(0))
	require.Equal(t, uint(math.MaxInt64), Int64ToUint(math.MaxInt64))
}

func TestUint64ToUint(t *testing.T) {
	require.Equal(t, uint(0), Uint64ToUint(0))
	require.Equal(t, uint(math.MaxUint64), Uint64ToUint(math.MaxUint64))
}
