package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUintToInt(t *testing.T) {
	require.Equal(t, 0, UintToInt(0))
	require.Equal(t, math.MaxInt, UintToInt(math.MaxInt))
	require.Panics(t, func() { UintToInt(math.MaxInt + 1) })
}

func TestUintToInt8(t *testing.T) {
	require.Equal(t, int8(0), UintToInt8(0))
	require.Equal(t, int8(math.MaxInt8), UintToInt8(math.MaxInt8))
	require.Panics(t, func() { UintToInt8(math.MaxInt8 + 1) })
}

func TestUintToInt16(t *testing.T) {
	require.Equal(t, int16(0), UintToInt16(0))
	require.Equal(t, int16(math.MaxInt16), UintToInt16(math.MaxInt16))
	require.Panics(t, func() { UintToInt16(math.MaxInt16 + 1) })
}

func TestUintToInt32(t *testing.T) {
	require.Equal(t, int32(0), UintToInt32(0))
	require.Equal(t, int32(math.MaxInt32), UintToInt32(math.MaxInt32))
	require.Panics(t, func() { UintToInt32(math.MaxInt32 + 1) })
}

func TestUintToUint8(t *testing.T) {
	require.Equal(t, uint8(0), UintToUint8(0))
	require.Equal(t, uint8(math.MaxUint8), UintToUint8(math.MaxUint8))
	require.Panics(t, func() { UintToUint8(math.MaxUint8 + 1) })
}

func TestUintToUint16(t *testing.T) {
	require.Equal(t, uint16(0), UintToUint16(0))
	require.Equal(t, uint16(math.MaxUint16), UintToUint16(math.MaxUint16))
	require.Panics(t, func() { UintToUint16(math.MaxUint16 + 1) })
}

func TestUintToUint64(t *testing.T) {
	require.Equal(t, uint64(0), UintToUint64(0))
	require.Equal(t, uint64(math.MaxUint), UintToUint64(math.MaxUint))
}
