package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt16ToInt8(t *testing.T) {
	require.Panics(t, func() { Int16ToInt8(math.MinInt8 - 1) })
	require.Equal(t, int8(math.MinInt8), Int16ToInt8(math.MinInt8))
	require.Equal(t, int8(0), Int16ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Int16ToInt8(math.MaxInt8))
	require.Panics(t, func() { Int16ToInt8(math.MaxInt8 + 1) })
}

func TestInt16ToInt32(t *testing.T) {
	require.Equal(t, int32(math.MinInt16), Int16ToInt32(math.MinInt16))
	require.Equal(t, int32(0), Int16ToInt32(0))
	require.Equal(t, int32(math.MaxInt16), Int16ToInt32(math.MaxInt16))
}

func TestInt16ToInt64(t *testing.T) {
	require.Equal(t, int64(math.MinInt16), Int16ToInt64(math.MinInt16))
	require.Equal(t, int64(0), Int16ToInt64(0))
	require.Equal(t, int64(math.MaxInt16), Int16ToInt64(math.MaxInt16))
}

func TestInt16ToUint8(t *testing.T) {
	require.Panics(t, func() { Int16ToUint8(-1) })
	require.Equal(t, uint8(0), Int16ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Int16ToUint8(math.MaxUint8))
	require.Panics(t, func() { Int16ToUint8(math.MaxUint8 + 1) })
}

func TestInt16ToUint16(t *testing.T) {
	require.Panics(t, func() { Int16ToUint16(-1) })
	require.Equal(t, uint16(0), Int16ToUint16(0))
	require.Equal(t, uint16(math.MaxInt16), Int16ToUint16(math.MaxInt16))
}

func TestInt16ToUint32(t *testing.T) {
	require.Panics(t, func() { Int16ToUint32(-1) })
	require.Equal(t, uint32(0), Int16ToUint32(0))
	require.Equal(t, uint32(math.MaxInt16), Int16ToUint32(math.MaxInt16))
}

func TestInt16ToUint64(t *testing.T) {
	require.Panics(t, func() { Int16ToUint64(-1) })
	require.Equal(t, uint64(0), Int16ToUint64(0))
	require.Equal(t, uint64(math.MaxInt16), Int16ToUint64(math.MaxInt16))
}
