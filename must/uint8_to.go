package must

import (
	"fmt"
	"math"
)

// Uint8ToInt8 converts uint8 to int8 or panics.
func Uint8ToInt8(v uint8) int8 {
	if v > math.MaxInt8 {
		panic(fmt.Sprintf("Uint8ToInt8(%d)", v))
	}

	return int8(v)
}

// Uint8ToInt16 converts uint8 to int16 or panics.
func Uint8ToInt16(v uint8) int16 {
	return int16(v)
}

// Uint8ToInt32 converts uint8 to int32 or panics.
func Uint8ToInt32(v uint8) int32 {
	return int32(v)
}

// Uint8ToInt64 converts uint8 to int64 or panics.
func Uint8ToInt64(v uint8) int64 {
	return int64(v)
}

// Uint8ToUint16 converts uint8 to uint16 or panics.
func Uint8ToUint16(v uint8) uint16 {
	return uint16(v)
}

// Uint8ToUint32 converts uint8 to uint32 or panics.
func Uint8ToUint32(v uint8) uint32 {
	return uint32(v)
}

// Uint8ToUint64 converts uint8 to uint64 or panics.
func Uint8ToUint64(v uint8) uint64 {
	return uint64(v)
}
