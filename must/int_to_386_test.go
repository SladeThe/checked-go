package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntToInt32(t *testing.T) {
	require.Equal(t, int32(math.MinInt32), IntToInt32(math.MinInt32))
	require.Equal(t, int32(0), IntToInt32(0))
	require.Equal(t, int32(math.MaxInt32), IntToInt32(math.MaxInt32))
}

func TestIntToUint32(t *testing.T) {
	require.Panics(t, func() { IntToUint32(-1) })
	require.Equal(t, uint32(0), IntToUint32(0))
	require.Equal(t, uint32(math.MaxInt32), IntToUint32(math.MaxInt32))
}
