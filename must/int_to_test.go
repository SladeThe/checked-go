package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntToInt8(t *testing.T) {
	require.Panics(t, func() { IntToInt8(math.MinInt8 - 1) })
	require.Equal(t, int8(math.MinInt8), IntToInt8(math.MinInt8))
	require.Equal(t, int8(0), IntToInt8(0))
	require.Equal(t, int8(math.MaxInt8), IntToInt8(math.MaxInt8))
	require.Panics(t, func() { IntToInt8(math.MaxInt8 + 1) })
}

func TestIntToInt16(t *testing.T) {
	require.Panics(t, func() { IntToInt16(math.MinInt16 - 1) })
	require.Equal(t, int16(math.MinInt16), IntToInt16(math.MinInt16))
	require.Equal(t, int16(0), IntToInt16(0))
	require.Equal(t, int16(math.MaxInt16), IntToInt16(math.MaxInt16))
	require.Panics(t, func() { IntToInt16(math.MaxInt16 + 1) })
}

func TestIntToInt64(t *testing.T) {
	require.Equal(t, int64(math.MinInt), IntToInt64(math.MinInt))
	require.Equal(t, int64(0), IntToInt64(0))
	require.Equal(t, int64(math.MaxInt), IntToInt64(math.MaxInt))
}

func TestIntToUint(t *testing.T) {
	require.Panics(t, func() { IntToUint(-1) })
	require.Equal(t, uint(0), IntToUint(0))
	require.Equal(t, uint(math.MaxInt), IntToUint(math.MaxInt))
}

func TestIntToUint8(t *testing.T) {
	require.Panics(t, func() { IntToUint8(-1) })
	require.Equal(t, uint8(0), IntToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), IntToUint8(math.MaxUint8))
	require.Panics(t, func() { IntToUint8(math.MaxUint8 + 1) })
}

func TestIntToUint16(t *testing.T) {
	require.Panics(t, func() { IntToUint16(-1) })
	require.Equal(t, uint16(0), IntToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), IntToUint16(math.MaxUint16))
	require.Panics(t, func() { IntToUint16(math.MaxUint16 + 1) })
}

func TestIntToUint64(t *testing.T) {
	require.Panics(t, func() { IntToUint64(-1) })
	require.Equal(t, uint64(0), IntToUint64(0))
	require.Equal(t, uint64(math.MaxInt), IntToUint64(math.MaxInt))
}
