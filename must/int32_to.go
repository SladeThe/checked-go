package must

import (
	"fmt"
	"math"
)

// Int32ToInt8 converts int32 to int8 or panics.
func Int32ToInt8(v int32) int8 {
	if v < math.MinInt8 || v > math.MaxInt8 {
		panic(fmt.Sprintf("Int32ToInt8(%d)", v))
	}

	return int8(v)
}

// Int32ToInt16 converts int32 to int16 or panics.
func Int32ToInt16(v int32) int16 {
	if v < math.MinInt16 || v > math.MaxInt16 {
		panic(fmt.Sprintf("Int32ToInt16(%d)", v))
	}

	return int16(v)
}

// Int32ToInt64 converts int32 to int64 or panics.
func Int32ToInt64(v int32) int64 {
	return int64(v)
}

// Int32ToUint8 converts int32 to uint8 or panics.
func Int32ToUint8(v int32) uint8 {
	if v < 0 || v > math.MaxUint8 {
		panic(fmt.Sprintf("Int32ToUint8(%d)", v))
	}

	return uint8(v)
}

// Int32ToUint16 converts int32 to uint16 or panics.
func Int32ToUint16(v int32) uint16 {
	if v < 0 || v > math.MaxUint16 {
		panic(fmt.Sprintf("Int32ToUint16(%d)", v))
	}

	return uint16(v)
}

// Int32ToUint32 converts int32 to uint32 or panics.
func Int32ToUint32(v int32) uint32 {
	if v < 0 {
		panic(fmt.Sprintf("Int32ToUint32(%d)", v))
	}

	return uint32(v)
}

// Int32ToUint64 converts int32 to uint64 or panics.
func Int32ToUint64(v int32) uint64 {
	if v < 0 {
		panic(fmt.Sprintf("Int32ToUint64(%d)", v))
	}

	return uint64(v)
}
