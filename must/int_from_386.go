package must

import (
	"fmt"
	"math"
)

// Int64ToInt converts int64 to int or panics.
func Int64ToInt(v int64) int {
	if v < math.MinInt || v > math.MaxInt {
		panic(fmt.Sprintf("Int64ToInt(%d)", v))
	}

	return int(v)
}

// Uint32ToInt converts uint32 to int or panics.
func Uint32ToInt(v uint32) int {
	if v > math.MaxInt {
		panic(fmt.Sprintf("Uint32ToInt(%d)", v))
	}

	return int(v)
}
