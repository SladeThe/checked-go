package must

import (
	"fmt"
	"math"
)

// UintToInt64 converts uint to int64 or panics.
func UintToInt64(v uint) int64 {
	if v > math.MaxInt64 {
		panic(fmt.Sprintf("UintToInt64(%d)", v))
	}

	return int64(v)
}

// UintToUint32 converts uint to uint32 or panics.
func UintToUint32(v uint) uint32 {
	if v > math.MaxUint32 {
		panic(fmt.Sprintf("UintToUint32(%d)", v))
	}

	return uint32(v)
}
