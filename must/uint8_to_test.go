package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUint8ToInt8(t *testing.T) {
	require.Equal(t, int8(0), Uint8ToInt8(0))
	require.Equal(t, int8(math.MaxInt8), Uint8ToInt8(math.MaxInt8))
	require.Panics(t, func() { Uint8ToInt8(math.MaxInt8 + 1) })
}

func TestUint8ToInt16(t *testing.T) {
	require.Equal(t, int16(0), Uint8ToInt16(0))
	require.Equal(t, int16(math.MaxUint8), Uint8ToInt16(math.MaxUint8))
}

func TestUint8ToInt32(t *testing.T) {
	require.Equal(t, int32(0), Uint8ToInt32(0))
	require.Equal(t, int32(math.MaxUint8), Uint8ToInt32(math.MaxUint8))
}

func TestUint8ToInt64(t *testing.T) {
	require.Equal(t, int64(0), Uint8ToInt64(0))
	require.Equal(t, int64(math.MaxUint8), Uint8ToInt64(math.MaxUint8))
}

func TestUint8ToUint16(t *testing.T) {
	require.Equal(t, uint16(0), Uint8ToUint16(0))
	require.Equal(t, uint16(math.MaxUint8), Uint8ToUint16(math.MaxUint8))
}

func TestUint8ToUint32(t *testing.T) {
	require.Equal(t, uint32(0), Uint8ToUint32(0))
	require.Equal(t, uint32(math.MaxUint8), Uint8ToUint32(math.MaxUint8))
}

func TestUint8ToUint64(t *testing.T) {
	require.Equal(t, uint64(0), Uint8ToUint64(0))
	require.Equal(t, uint64(math.MaxUint8), Uint8ToUint64(math.MaxUint8))
}
