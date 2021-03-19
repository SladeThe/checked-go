package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint32ToInt8(t *testing.T) {
	require.Equal(t, int8(0), Uint32ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Uint32ToInt8(math.MaxInt8))
	require.Panics(t, func() { Uint32ToInt8(math.MaxInt8 + 1) })
}

func TestUint32ToInt16(t *testing.T) {
	require.Equal(t, int16(0), Uint32ToInt16(0))
	require.Equal(t, int16(math.MaxInt16), Uint32ToInt16(math.MaxInt16))
	require.Panics(t, func() { Uint32ToInt16(math.MaxInt16 + 1) })
}

func TestUint32ToInt32(t *testing.T) {
	require.Equal(t, int32(0), Uint32ToInt32(0))
	require.Equal(t, int32(math.MaxInt32), Uint32ToInt32(math.MaxInt32))
	require.Panics(t, func() { Uint32ToInt32(math.MaxInt32 + 1) })
}

func TestUint32ToInt64(t *testing.T) {
	require.Equal(t, int64(0), Uint32ToInt64(0))
	require.Equal(t, int64(math.MaxUint32), Uint32ToInt64(math.MaxUint32))
}

func TestUint32ToUint8(t *testing.T) {
	require.Equal(t, uint8(0), Uint32ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Uint32ToUint8(math.MaxUint8))
	require.Panics(t, func() { Uint32ToUint8(math.MaxUint8 + 1) })
}

func TestUint32ToUint16(t *testing.T) {
	require.Equal(t, uint16(0), Uint32ToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), Uint32ToUint16(math.MaxUint16))
	require.Panics(t, func() { Uint32ToUint16(math.MaxUint16 + 1) })
}

func TestUint32ToUint64(t *testing.T) {
	require.Equal(t, uint64(0), Uint32ToUint64(0))
	require.Equal(t, uint64(math.MaxUint32), Uint32ToUint64(math.MaxUint32))
}
