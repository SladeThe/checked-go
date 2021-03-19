package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt64ToInt8(t *testing.T) {
	require.Panics(t, func() { Int64ToInt8(math.MinInt8 - 1) })
	require.Equal(t, int8(math.MinInt8), Int64ToInt8(math.MinInt8))
	require.Equal(t, int8(0), Int64ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Int64ToInt8(math.MaxInt8))
	require.Panics(t, func() { Int64ToInt8(math.MaxInt8 + 1) })
}

func TestInt64ToInt16(t *testing.T) {
	require.Panics(t, func() { Int64ToInt16(math.MinInt16 - 1) })
	require.Equal(t, int16(math.MinInt16), Int64ToInt16(math.MinInt16))
	require.Equal(t, int16(0), Int64ToInt16(0))
	require.Equal(t, int16(math.MaxInt16), Int64ToInt16(math.MaxInt16))
	require.Panics(t, func() { Int64ToInt16(math.MaxInt16 + 1) })
}

func TestInt64ToInt32(t *testing.T) {
	require.Panics(t, func() { Int64ToInt32(math.MinInt32 - 1) })
	require.Equal(t, int32(math.MinInt32), Int64ToInt32(math.MinInt32))
	require.Equal(t, int32(0), Int64ToInt32(0))
	require.Equal(t, int32(math.MaxInt32), Int64ToInt32(math.MaxInt32))
	require.Panics(t, func() { Int64ToInt32(math.MaxInt32 + 1) })
}

func TestInt64ToUint8(t *testing.T) {
	require.Panics(t, func() { Int64ToUint8(-1) })
	require.Equal(t, uint8(0), Int64ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Int64ToUint8(math.MaxUint8))
	require.Panics(t, func() { Int64ToUint8(math.MaxUint8 + 1) })
}

func TestInt64ToUint16(t *testing.T) {
	require.Panics(t, func() { Int64ToUint16(-1) })
	require.Equal(t, uint16(0), Int64ToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), Int64ToUint16(math.MaxUint16))
	require.Panics(t, func() { Int64ToUint16(math.MaxUint16 + 1) })
}

func TestInt64ToUint32(t *testing.T) {
	require.Panics(t, func() { Int64ToUint32(-1) })
	require.Equal(t, uint32(0), Int64ToUint32(0))
	require.Equal(t, uint32(math.MaxUint32), Int64ToUint32(math.MaxUint32))
	require.Panics(t, func() { Int64ToUint32(math.MaxUint32 + 1) })
}

func TestInt64ToUint64(t *testing.T) {
	require.Panics(t, func() { Int64ToUint64(-1) })
	require.Equal(t, uint64(0), Int64ToUint64(0))
	require.Equal(t, uint64(math.MaxInt64), Int64ToUint64(math.MaxInt64))
}
