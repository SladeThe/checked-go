package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt32ToInt8(t *testing.T) {
	require.Panics(t, func() { Int32ToInt8(math.MinInt8 - 1) })
	require.Equal(t, int8(math.MinInt8), Int32ToInt8(math.MinInt8))
	require.Equal(t, int8(0), Int32ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Int32ToInt8(math.MaxInt8))
	require.Panics(t, func() { Int32ToInt8(math.MaxInt8 + 1) })
}

func TestInt32ToInt16(t *testing.T) {
	require.Panics(t, func() { Int32ToInt16(math.MinInt16 - 1) })
	require.Equal(t, int16(math.MinInt16), Int32ToInt16(math.MinInt16))
	require.Equal(t, int16(0), Int32ToInt16(0))
	require.Equal(t, int16(math.MaxInt16), Int32ToInt16(math.MaxInt16))
	require.Panics(t, func() { Int32ToInt16(math.MaxInt16 + 1) })
}

func TestInt32ToInt64(t *testing.T) {
	require.Equal(t, int64(math.MinInt32), Int32ToInt64(math.MinInt32))
	require.Equal(t, int64(0), Int32ToInt64(0))
	require.Equal(t, int64(math.MaxInt32), Int32ToInt64(math.MaxInt32))
}

func TestInt32ToUint8(t *testing.T) {
	require.Panics(t, func() { Int32ToUint8(-1) })
	require.Equal(t, uint8(0), Int32ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Int32ToUint8(math.MaxUint8))
	require.Panics(t, func() { Int32ToUint8(math.MaxUint8 + 1) })
}

func TestInt32ToUint16(t *testing.T) {
	require.Panics(t, func() { Int32ToUint16(-1) })
	require.Equal(t, uint16(0), Int32ToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), Int32ToUint16(math.MaxUint16))
	require.Panics(t, func() { Int32ToUint16(math.MaxUint16 + 1) })
}

func TestInt32ToUint32(t *testing.T) {
	require.Panics(t, func() { Int32ToUint32(-1) })
	require.Equal(t, uint32(0), Int32ToUint32(0))
	require.Equal(t, uint32(math.MaxInt32), Int32ToUint32(math.MaxInt32))
}

func TestInt32ToUint64(t *testing.T) {
	require.Panics(t, func() { Int32ToUint64(-1) })
	require.Equal(t, uint64(0), Int32ToUint64(0))
	require.Equal(t, uint64(math.MaxInt32), Int32ToUint64(math.MaxInt32))
}
