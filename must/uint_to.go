package must

import (
	"fmt"
	"math"

	"github.com/SladeThe/checked-go/consts"
)

// UintToInt converts uint to int or panics.
func UintToInt(v uint) int {
	if v > consts.MaxInt {
		panic(fmt.Sprintf("UintToInt(%d)", v))
	}

	return int(v)
}

// UintToInt8 converts uint to int8 or panics.
func UintToInt8(v uint) int8 {
	if v > math.MaxInt8 {
		panic(fmt.Sprintf("UintToInt8(%d)", v))
	}

	return int8(v)
}

// UintToInt16 converts uint to int16 or panics.
func UintToInt16(v uint) int16 {
	if v > math.MaxInt16 {
		panic(fmt.Sprintf("UintToInt16(%d)", v))
	}

	return int16(v)
}

// UintToInt32 converts uint to int32 or panics.
func UintToInt32(v uint) int32 {
	if v > math.MaxInt32 {
		panic(fmt.Sprintf("UintToInt32(%d)", v))
	}

	return int32(v)
}

// UintToUint8 converts uint to uint8 or panics.
func UintToUint8(v uint) uint8 {
	if v > math.MaxUint8 {
		panic(fmt.Sprintf("UintToUint8(%d)", v))
	}

	return uint8(v)
}

// UintToUint16 converts uint to uint16 or panics.
func UintToUint16(v uint) uint16 {
	if v > math.MaxUint16 {
		panic(fmt.Sprintf("UintToUint16(%d)", v))
	}

	return uint16(v)
}

// UintToUint64 converts uint to uint64 or panics.
func UintToUint64(v uint) uint64 {
	return uint64(v)
}
