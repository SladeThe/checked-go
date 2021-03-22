package must

import "fmt"

// Int8ToUint converts int8 to uint or panics.
func Int8ToUint(v int8) uint {
	if v < 0 {
		panic(fmt.Sprintf("Int8ToUint(%d)", v))
	}

	return uint(v)
}

// Int16ToUint converts int16 to uint or panics.
func Int16ToUint(v int16) uint {
	if v < 0 {
		panic(fmt.Sprintf("Int16ToUint(%d)", v))
	}

	return uint(v)
}

// Int32ToUint converts int32 to uint or panics.
func Int32ToUint(v int32) uint {
	if v < 0 {
		panic(fmt.Sprintf("Int32ToUint(%d)", v))
	}

	return uint(v)
}

// Uint8ToUint converts uint8 to uint or panics.
func Uint8ToUint(v uint8) uint {
	return uint(v)
}

// Uint16ToUint converts uint16 to uint or panics.
func Uint16ToUint(v uint16) uint {
	return uint(v)
}

// Uint32ToUint converts uint32 to uint or panics.
func Uint32ToUint(v uint32) uint {
	return uint(v)
}
