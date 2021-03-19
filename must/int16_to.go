package must

import (
	"fmt"
	"math"
)

// Int16ToInt8 converts int16 to int8 or panics.
func Int16ToInt8(v int16) int8 {
	if v < math.MinInt8 || v > math.MaxInt8 {
		panic(fmt.Sprintf("Int16ToInt8(%d)", v))
	}

	return int8(v)
}

// Int16ToInt32 converts int16 to int32 or panics.
func Int16ToInt32(v int16) int32 {
	return int32(v)
}

// Int16ToInt64 converts int16 to int64 or panics.
func Int16ToInt64(v int16) int64 {
	return int64(v)
}

// Int16ToUint8 converts int16 to uint8 or panics.
func Int16ToUint8(v int16) uint8 {
	if v < 0 || v > math.MaxUint8 {
		panic(fmt.Sprintf("Int16ToUint8(%d)", v))
	}

	return uint8(v)
}

// Int16ToUint16 converts int16 to uint16 or panics.
func Int16ToUint16(v int16) uint16 {
	if v < 0 {
		panic(fmt.Sprintf("Int16ToUint16(%d)", v))
	}

	return uint16(v)
}

// Int16ToUint32 converts int16 to uint32 or panics.
func Int16ToUint32(v int16) uint32 {
	if v < 0 {
		panic(fmt.Sprintf("Int16ToUint32(%d)", v))
	}

	return uint32(v)
}

// Int16ToUint64 converts int16 to uint64 or panics.
func Int16ToUint64(v int16) uint64 {
	if v < 0 {
		panic(fmt.Sprintf("Int16ToUint64(%d)", v))
	}

	return uint64(v)
}
