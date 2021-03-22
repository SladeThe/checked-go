package must

import (
	"fmt"

	"github.com/SladeThe/checked-go/consts"
)

// Int64ToUint converts int64 to uint or panics.
func Int64ToUint(v int64) uint {
	if v < 0 || v > consts.MaxUint {
		panic(fmt.Sprintf("Int64ToUint(%d)", v))
	}

	return uint(v)
}

// Uint64ToUint converts uint64 to uint or panics.
func Uint64ToUint(v uint64) uint {
	if v > consts.MaxUint {
		panic(fmt.Sprintf("Uint64ToUint(%d)", v))
	}

	return uint(v)
}
