package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64ToInt(t *testing.T) {
	require.Equal(t, math.MinInt64, Int64ToInt(math.MinInt64))
	require.Equal(t, 0, Int64ToInt(0))
	require.Equal(t, math.MaxInt64, Int64ToInt(math.MaxInt64))
}

func TestUint32ToInt(t *testing.T) {
	require.Equal(t, 0, Uint32ToInt(0))
	require.Equal(t, math.MaxUint32, Uint32ToInt(math.MaxUint32))
}
