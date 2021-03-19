package must

import (
	"fmt"
	"math"
)

// Uint16ToInt8 converts uint16 to int8 or panics.
func Uint16ToInt8(v uint16) int8 {
	if v > math.MaxInt8 {
		panic(fmt.Sprintf("Uint8ToInt8(%d)", v))
	}

	return int8(v)
}

// Uint16ToInt16 converts uint16 to int16 or panics.
func Uint16ToInt16(v uint16) int16 {
	if v > math.MaxInt16 {
		panic(fmt.Sprintf("Uint16ToInt16(%d)", v))
	}

	return int16(v)
}

// Uint16ToInt32 converts uint16 to int32 or panics.
func Uint16ToInt32(v uint16) int32 {
	return int32(v)
}

// Uint16ToInt64 converts uint16 to int64 or panics.
func Uint16ToInt64(v uint16) int64 {
	return int64(v)
}

// Uint16ToUint8 converts uint16 to uint8 or panics.
func Uint16ToUint8(v uint16) uint8 {
	if v > math.MaxUint8 {
		panic(fmt.Sprintf("Uint16ToUint8(%d)", v))
	}

	return uint8(v)
}

// Uint16ToUint32 converts uint16 to uint32 or panics.
func Uint16ToUint32(v uint16) uint32 {
	return uint32(v)
}

// Uint16ToUint64 converts uint16 to uint64 or panics.
func Uint16ToUint64(v uint16) uint64 {
	return uint64(v)
}
