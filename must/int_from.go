package must

import (
	"fmt"
	"math"
)

// Int8ToInt converts int8 to int or panics.
func Int8ToInt(v int8) int {
	return int(v)
}

// Int16ToInt converts int16 to int or panics.
func Int16ToInt(v int16) int {
	return int(v)
}

// Int32ToInt converts int32 to int or panics.
func Int32ToInt(v int32) int {
	return int(v)
}

// Uint8ToInt converts uint8 to int or panics.
func Uint8ToInt(v uint8) int {
	return int(v)
}

// Uint16ToInt converts uint16 to int or panics.
func Uint16ToInt(v uint16) int {
	return int(v)
}

// Uint64ToInt converts uint64 to int or panics.
func Uint64ToInt(v uint64) int {
	if v > math.MaxInt {
		panic(fmt.Sprintf("Uint64ToInt(%d)", v))
	}

	return int(v)
}
