package must

import (
	"fmt"
)

// Int8ToInt16 converts int8 to int16 or panics.
func Int8ToInt16(v int8) int16 {
	return int16(v)
}

// Int8ToInt32 converts int8 to int32 or panics.
func Int8ToInt32(v int8) int32 {
	return int32(v)
}

// Int8ToInt64 converts int8 to int64 or panics.
func Int8ToInt64(v int8) int64 {
	return int64(v)
}

// Int8ToUint8 converts int8 to uint8 or panics.
func Int8ToUint8(v int8) uint8 {
	if v < 0 {
		panic(fmt.Sprintf("Int8ToUint8(%d)", v))
	}

	return uint8(v)
}

// Int8ToUint16 converts int8 to uint16 or panics.
func Int8ToUint16(v int8) uint16 {
	if v < 0 {
		panic(fmt.Sprintf("Int8ToUint16(%d)", v))
	}

	return uint16(v)
}

// Int8ToUint32 converts int8 to uint32 or panics.
func Int8ToUint32(v int8) uint32 {
	if v < 0 {
		panic(fmt.Sprintf("Int8ToUint32(%d)", v))
	}

	return uint32(v)
}

// Int8ToUint64 converts int8 to uint64 or panics.
func Int8ToUint64(v int8) uint64 {
	if v < 0 {
		panic(fmt.Sprintf("Int8ToUint64(%d)", v))
	}

	return uint64(v)
}
