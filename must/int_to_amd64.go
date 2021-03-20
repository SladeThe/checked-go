package must

import (
	"fmt"
	"math"
)

// IntToInt32 converts int to int32 or panics.
func IntToInt32(v int) int32 {
	if v < math.MinInt32 || v > math.MaxInt32 {
		panic(fmt.Sprintf("IntToInt32(%d)", v))
	}

	return int32(v)
}

// IntToUint32 converts int to uint32 or panics.
func IntToUint32(v int) uint32 {
	if v < 0 || v > math.MaxUint32 {
		panic(fmt.Sprintf("IntToUint32(%d)", v))
	}

	return uint32(v)
}
