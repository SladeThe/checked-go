package must

import (
	"fmt"
	"math"
)

// Uint32ToInt8 converts uint32 to int8 or panics.
func Uint32ToInt8(v uint32) int8 {
	if v > math.MaxInt8 {
		panic(fmt.Sprintf("Uint32ToInt8(%d)", v))
	}

	return int8(v)
}

// Uint32ToInt16 converts uint32 to int16 or panics.
func Uint32ToInt16(v uint32) int16 {
	if v > math.MaxInt16 {
		panic(fmt.Sprintf("Uint32ToInt16(%d)", v))
	}

	return int16(v)
}

// Uint32ToInt32 converts uint32 to int32 or panics.
func Uint32ToInt32(v uint32) int32 {
	if v > math.MaxInt32 {
		panic(fmt.Sprintf("Uint32ToInt32(%d)", v))
	}

	return int32(v)
}

// Uint32ToInt64 converts uint32 to int64 or panics.
func Uint32ToInt64(v uint32) int64 {
	return int64(v)
}

// Uint32ToUint8 converts uint32 to uint8 or panics.
func Uint32ToUint8(v uint32) uint8 {
	if v > math.MaxUint8 {
		panic(fmt.Sprintf("Uint32ToUint8(%d)", v))
	}

	return uint8(v)
}

// Uint32ToUint16 converts uint32 to uint16 or panics.
func Uint32ToUint16(v uint32) uint16 {
	if v > math.MaxUint16 {
		panic(fmt.Sprintf("Uint32ToUint16(%d)", v))
	}

	return uint16(v)
}

// Uint32ToUint64 converts uint32 to uint64 or panics.
func Uint32ToUint64(v uint32) uint64 {
	return uint64(v)
}
