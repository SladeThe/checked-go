package must

import (
	"fmt"
	"math"
)

// IntToInt8 converts int to int8 or panics.
func IntToInt8(v int) int8 {
	if v < math.MinInt8 || v > math.MaxInt8 {
		panic(fmt.Sprintf("IntToInt8(%d)", v))
	}

	return int8(v)
}

// IntToInt16 converts int to int16 or panics.
func IntToInt16(v int) int16 {
	if v < math.MinInt16 || v > math.MaxInt16 {
		panic(fmt.Sprintf("IntToInt16(%d)", v))
	}

	return int16(v)
}

// IntToInt64 converts int to int64 or panics.
func IntToInt64(v int) int64 {
	return int64(v)
}

// IntToUint converts int to uint or panics.
func IntToUint(v int) uint {
	if v < 0 {
		panic(fmt.Sprintf("IntToUint(%d)", v))
	}

	return uint(v)
}

// IntToUint8 converts int to uint8 or panics.
func IntToUint8(v int) uint8 {
	if v < 0 || v > math.MaxUint8 {
		panic(fmt.Sprintf("IntToUint8(%d)", v))
	}

	return uint8(v)
}

// IntToUint16 converts int to uint16 or panics.
func IntToUint16(v int) uint16 {
	if v < 0 || v > math.MaxUint16 {
		panic(fmt.Sprintf("IntToUint16(%d)", v))
	}

	return uint16(v)
}

// IntToUint64 converts int to uint64 or panics.
func IntToUint64(v int) uint64 {
	if v < 0 {
		panic(fmt.Sprintf("IntToUint64(%d)", v))
	}

	return uint64(v)
}
