package must

import "fmt"

// Int64ToUint converts int64 to uint or panics.
func Int64ToUint(v int64) uint {
	if v < 0 {
		panic(fmt.Sprintf("Int64ToUint(%d)", v))
	}

	return uint(v)
}

// Uint64ToUint converts uint64 to uint or panics.
func Uint64ToUint(v uint64) uint {
	return uint(v)
}
