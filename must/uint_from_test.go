package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt8ToUint(t *testing.T) {
	require.Panics(t, func() { Int8ToUint(-1) })
	require.Equal(t, uint(0), Int8ToUint(0))
	require.Equal(t, uint(math.MaxInt8), Int8ToUint(math.MaxInt8))
}

func TestInt16ToUint(t *testing.T) {
	require.Panics(t, func() { Int16ToUint(-1) })
	require.Equal(t, uint(0), Int16ToUint(0))
	require.Equal(t, uint(math.MaxInt16), Int16ToUint(math.MaxInt16))
}

func TestInt32ToUint(t *testing.T) {
	require.Panics(t, func() { Int32ToUint(-1) })
	require.Equal(t, uint(0), Int32ToUint(0))
	require.Equal(t, uint(math.MaxInt32), Int32ToUint(math.MaxInt32))
}

func TestUint8ToUint(t *testing.T) {
	require.Equal(t, uint(0), Uint8ToUint(0))
	require.Equal(t, uint(math.MaxUint8), Uint8ToUint(math.MaxUint8))
}

func TestUint16ToUint(t *testing.T) {
	require.Equal(t, uint(0), Uint16ToUint(0))
	require.Equal(t, uint(math.MaxUint16), Uint16ToUint(math.MaxUint16))
}

func TestUint32ToUint(t *testing.T) {
	require.Equal(t, uint(0), Uint32ToUint(0))
	require.Equal(t, uint(math.MaxUint32), Uint32ToUint(math.MaxUint32))
}
