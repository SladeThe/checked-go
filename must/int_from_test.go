package must

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/SladeThe/checked-go/consts"
)

func TestInt8ToInt(t *testing.T) {
	require.Equal(t, math.MinInt8, Int8ToInt(math.MinInt8))
	require.Equal(t, 0, Int8ToInt(0))
	require.Equal(t, math.MaxInt8, Int8ToInt(math.MaxInt8))
}

func TestInt16ToInt(t *testing.T) {
	require.Equal(t, math.MinInt16, Int16ToInt(math.MinInt16))
	require.Equal(t, 0, Int16ToInt(0))
	require.Equal(t, math.MaxInt16, Int16ToInt(math.MaxInt16))
}

func TestInt32ToInt(t *testing.T) {
	require.Equal(t, math.MinInt32, Int32ToInt(math.MinInt32))
	require.Equal(t, 0, Int32ToInt(0))
	require.Equal(t, math.MaxInt32, Int32ToInt(math.MaxInt32))
}

func TestUint8ToInt(t *testing.T) {
	require.Equal(t, 0, Uint8ToInt(0))
	require.Equal(t, math.MaxUint8, Uint8ToInt(math.MaxUint8))
}

func TestUint16ToInt(t *testing.T) {
	require.Equal(t, 0, Uint16ToInt(0))
	require.Equal(t, math.MaxUint16, Uint16ToInt(math.MaxUint16))
}

func TestUint64ToInt(t *testing.T) {
	require.Equal(t, 0, Uint64ToInt(0))
	require.Equal(t, consts.MaxInt, Uint64ToInt(consts.MaxInt))
	require.Panics(t, func() { Uint64ToInt(consts.MaxInt + 1) })
}
