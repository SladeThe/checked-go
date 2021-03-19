package must

import (
	"fmt"
	"math"
)

// Uint64ToInt8 converts uint64 to int8 or panics.
func Uint64ToInt8(v uint64) int8 {
	if v > math.MaxInt8 {
		panic(fmt.Sprintf("Uint64ToInt8(%d)", v))
	}

	return int8(v)
}

// Uint64ToInt16 converts uint64 to int16 or panics.
func Uint64ToInt16(v uint64) int16 {
	if v > math.MaxInt16 {
		panic(fmt.Sprintf("Uint64ToInt16(%d)", v))
	}

	return int16(v)
}

// Uint64ToInt32 converts uint64 to int32 or panics.
func Uint64ToInt32(v uint64) int32 {
	if v > math.MaxInt32 {
		panic(fmt.Sprintf("Uint64ToInt32(%d)", v))
	}

	return int32(v)
}

// Uint64ToInt64 converts uint64 to int64 or panics.
func Uint64ToInt64(v uint64) int64 {
	if v > math.MaxInt64 {
		panic(fmt.Sprintf("Uint64ToInt64(%d)", v))
	}

	return int64(v)
}

// Uint64ToUint8 converts uint64 to uint8 or panics.
func Uint64ToUint8(v uint64) uint8 {
	if v > math.MaxUint8 {
		panic(fmt.Sprintf("Uint64ToUint8(%d)", v))
	}

	return uint8(v)
}

// Uint64ToUint16 converts uint64 to uint16 or panics.
func Uint64ToUint16(v uint64) uint16 {
	if v > math.MaxUint16 {
		panic(fmt.Sprintf("Uint64ToUint16(%d)", v))
	}

	return uint16(v)
}

// Uint64ToUint32 converts uint64 to uint32 or panics.
func Uint64ToUint32(v uint64) uint32 {
	if v > math.MaxUint32 {
		panic(fmt.Sprintf("Uint64ToUint32(%d)", v))
	}

	return uint32(v)
}
