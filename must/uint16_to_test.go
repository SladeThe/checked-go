package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint16ToInt8(t *testing.T) {
	require.Equal(t, int8(0), Uint16ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Uint16ToInt8(math.MaxInt8))
	require.Panics(t, func() { Uint16ToInt8(math.MaxInt8 + 1) })
}

func TestUint16ToInt16(t *testing.T) {
	require.Equal(t, int16(0), Uint16ToInt16(0))
	require.Equal(t, int16(math.MaxInt16), Uint16ToInt16(math.MaxInt16))
	require.Panics(t, func() { Uint16ToInt16(math.MaxInt16 + 1) })
}

func TestUint16ToInt32(t *testing.T) {
	require.Equal(t, int32(0), Uint16ToInt32(0))
	require.Equal(t, int32(math.MaxUint16), Uint16ToInt32(math.MaxUint16))
}

func TestUint16ToInt64(t *testing.T) {
	require.Equal(t, int64(0), Uint16ToInt64(0))
	require.Equal(t, int64(math.MaxUint16), Uint16ToInt64(math.MaxUint16))
}

func TestUint16ToUint8(t *testing.T) {
	require.Equal(t, uint8(0), Uint16ToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), Uint16ToUint8(math.MaxUint8))
	require.Panics(t, func() { Uint16ToUint8(math.MaxUint8 + 1) })
}

func TestUint16ToUint32(t *testing.T) {
	require.Equal(t, uint32(0), Uint16ToUint32(0))
	require.Equal(t, uint32(math.MaxUint16), Uint16ToUint32(math.MaxUint16))
}

func TestUint16ToUint64(t *testing.T) {
	require.Equal(t, uint64(0), Uint16ToUint64(0))
	require.Equal(t, uint64(math.MaxUint16), Uint16ToUint64(math.MaxUint16))
}
