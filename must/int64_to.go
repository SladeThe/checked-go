package must

import (
	"fmt"
	"math"
)

// Int64ToInt8 converts int64 to int8 or panics.
func Int64ToInt8(v int64) int8 {
	if v < math.MinInt8 || v > math.MaxInt8 {
		panic(fmt.Sprintf("Int64ToInt8(%d)", v))
	}

	return int8(v)
}

// Int64ToInt16 converts int64 to int16 or panics.
func Int64ToInt16(v int64) int16 {
	if v < math.MinInt16 || v > math.MaxInt16 {
		panic(fmt.Sprintf("Int64ToInt16(%d)", v))
	}

	return int16(v)
}

// Int64ToInt32 converts int64 to int32 or panics.
func Int64ToInt32(v int64) int32 {
	if v < math.MinInt32 || v > math.MaxInt32 {
		panic(fmt.Sprintf("Int64ToInt32(%d)", v))
	}

	return int32(v)
}

// Int64ToUint8 converts int64 to uint8 or panics.
func Int64ToUint8(v int64) uint8 {
	if v < 0 || v > math.MaxUint8 {
		panic(fmt.Sprintf("Int64ToUint8(%d)", v))
	}

	return uint8(v)
}

// Int64ToUint16 converts int64 to uint16 or panics.
func Int64ToUint16(v int64) uint16 {
	if v < 0 || v > math.MaxUint16 {
		panic(fmt.Sprintf("Int64ToUint16(%d)", v))
	}

	return uint16(v)
}

// Int64ToUint32 converts int64 to uint32 or panics.
func Int64ToUint32(v int64) uint32 {
	if v < 0 || v > math.MaxUint32 {
		panic(fmt.Sprintf("Int64ToUint32(%d)", v))
	}

	return uint32(v)
}

// Int64ToUint64 converts int64 to uint64 or panics.
func Int64ToUint64(v int64) uint64 {
	if v < 0 {
		panic(fmt.Sprintf("Int64ToUint64(%d)", v))
	}

	return uint64(v)
}
