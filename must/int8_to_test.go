package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt8ToInt16(t *testing.T) {
	require.Equal(t, int16(math.MinInt8), Int8ToInt16(math.MinInt8))
	require.Equal(t, int16(0), Int8ToInt16(0))
	require.Equal(t, int16(math.MaxInt8), Int8ToInt16(math.MaxInt8))
}

func TestInt8ToInt32(t *testing.T) {
	require.Equal(t, int32(math.MinInt8), Int8ToInt32(math.MinInt8))
	require.Equal(t, int32(0), Int8ToInt32(0))
	require.Equal(t, int32(math.MaxInt8), Int8ToInt32(math.MaxInt8))
}

func TestInt8ToInt64(t *testing.T) {
	require.Equal(t, int64(math.MinInt8), Int8ToInt64(math.MinInt8))
	require.Equal(t, int64(0), Int8ToInt64(0))
	require.Equal(t, int64(math.MaxInt8), Int8ToInt64(math.MaxInt8))
}

func TestInt8ToUint8(t *testing.T) {
	require.Panics(t, func() { Int8ToUint8(-1) })
	require.Equal(t, uint8(0), Int8ToUint8(0))
	require.Equal(t, uint8(math.MaxInt8), Int8ToUint8(math.MaxInt8))
}

func TestInt8ToUint16(t *testing.T) {
	require.Panics(t, func() { Int8ToUint16(-1) })
	require.Equal(t, uint16(0), Int8ToUint16(0))
	require.Equal(t, uint16(math.MaxInt8), Int8ToUint16(math.MaxInt8))
}

func TestInt8ToUint32(t *testing.T) {
	require.Panics(t, func() { Int8ToUint32(-1) })
	require.Equal(t, uint32(0), Int8ToUint32(0))
	require.Equal(t, uint32(math.MaxInt8), Int8ToUint32(math.MaxInt8))
}

func TestInt8ToUint64(t *testing.T) {
	require.Panics(t, func() { Int8ToUint64(-1) })
	require.Equal(t, uint64(0), Int8ToUint64(0))
	require.Equal(t, uint64(math.MaxInt8), Int8ToUint64(math.MaxInt8))
}
