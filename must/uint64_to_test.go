package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint64ToInt8(t *testing.T) {
	require.Equal(t, int8(0), Uint64ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Uint64ToInt8(math.MaxInt8))
	require.Panics(t, func() { Uint64ToInt8(math.MaxInt8 + 1) })
}

func TestUint64ToInt16(t *testing.T) {
	require.Equal(t, int16(0), Uint64ToInt16(0))
	require.Equal(t, int16(math.MaxInt16), Uint64ToInt16(math.MaxInt16))
	require.Panics(t, func() { Uint64ToInt16(math.MaxInt16 + 1) })
}

func TestUint64ToInt32(t *testing.T) {
	require.Equal(t, int32(0), Uint64ToInt32(0))
	require.Equal(t, int32(math.MaxInt32), Uint64ToInt32(math.MaxInt32))
	require.Panics(t, func() { Uint64ToInt32(math.MaxInt32 + 1) })
}

func TestUint64ToInt64(t *testing.T) {
	require.Equal(t, int64(0), Uint64ToInt64(0))
	require.Equal(t, int64(math.MaxInt64), Uint64ToInt64(math.MaxInt64))
	require.Panics(t, func() { Uint64ToInt64(math.MaxInt64 + 1) })
}

func TestUint64ToUint8(t *testing.T) {
	require.Equal(t, uint8(0), Uint64ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Uint64ToUint8(math.MaxUint8))
	require.Panics(t, func() { Uint64ToUint8(math.MaxUint8 + 1) })
}

func TestUint64ToUint16(t *testing.T) {
	require.Equal(t, uint16(0), Uint64ToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), Uint64ToUint16(math.MaxUint16))
	require.Panics(t, func() { Uint64ToUint16(math.MaxUint16 + 1) })
}

func TestUint64ToUint32(t *testing.T) {
	require.Equal(t, uint32(0), Uint64ToUint32(0))
	require.Equal(t, uint32(math.MaxUint32), Uint64ToUint32(math.MaxUint32))
	require.Panics(t, func() { Uint64ToUint32(math.MaxUint32 + 1) })
}
