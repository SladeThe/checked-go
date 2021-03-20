package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUintToInt64(t *testing.T) {
	require.Equal(t, int64(0), UintToInt64(0))
	require.Equal(t, int64(math.MaxInt64), UintToInt64(math.MaxInt64))
	require.Panics(t, func() { UintToInt64(math.MaxInt64 + 1) })
}

func TestUintToUint32(t *testing.T) {
	require.Equal(t, uint32(0), UintToUint32(0))
	require.Equal(t, uint32(math.MaxUint32), UintToUint32(math.MaxUint32))
	require.Panics(t, func() { UintToUint32(math.MaxUint32 + 1) })
}
