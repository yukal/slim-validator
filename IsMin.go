package validator

import "reflect"

// https://go.dev/ref/spec#Numeric_types
func IsMin(proto, value any) bool {
	types := reflect.ValueOf(value).Kind().String() + ":" + reflect.ValueOf(proto).Kind().String()

	switch types {

	// int

	case "int8:int8":
		return value.(int8) >= proto.(int8)

	case "int8:int16":
		return int16(value.(int8)) >= proto.(int16)

	case "int8:int32", "int8:rune":
		return int32(value.(int8)) >= proto.(int32)

	case "int8:int64":
		return int64(value.(int8)) >= proto.(int64)

	case "int8:int":
		return int(value.(int8)) >= proto.(int)

	case "int16:int8":
		return value.(int16) >= int16(proto.(int8))

	case "int16:int16":
		return value.(int16) >= proto.(int16)

	case "int16:int32", "int16:rune":
		return int32(value.(int16)) >= proto.(int32)

	case "int16:int64":
		return int64(value.(int16)) >= proto.(int64)

	case "int16:int":
		return int(value.(int16)) >= proto.(int)

	case "int32:int8", "rune:int8":
		return value.(int32) >= int32(proto.(int8))

	case "int32:int16", "rune:int16":
		return value.(int32) >= int32(proto.(int16))

	case "int32:int32", "int32:rune", "rune:int32", "rune:rune":
		return value.(int32) >= proto.(int32)

	case "int32:int64", "rune:int64":
		return int64(value.(int32)) >= proto.(int64)

	case "int32:int", "rune:int":
		return int(value.(int32)) >= proto.(int)

	case "int64:int8":
		return value.(int64) >= int64(proto.(int8))

	case "int64:int16":
		return value.(int64) >= int64(proto.(int16))

	case "int64:int32", "int64:rune":
		return value.(int64) >= int64(proto.(int32))

	case "int64:int64":
		return value.(int64) >= proto.(int64)

	case "int64:int":
		return value.(int64) >= int64(proto.(int))

	case "int:int8":
		return value.(int) >= int(proto.(int8))

	case "int:int16":
		return value.(int) >= int(proto.(int16))

	case "int:int32", "int:rune":
		return value.(int) >= int(proto.(int32))

	case "int:int64":
		return int64(value.(int)) >= proto.(int64)

	case "int:int":
		return value.(int) >= proto.(int)

	// ............................................................
	// int & uint

	case "int8:uint8", "int8:byte":
		if value.(int8) >= 0 {
			return uint8(value.(int8)) >= proto.(uint8)
		}

	case "int8:uint16":
		if value.(int8) >= 0 {
			return uint16(value.(int8)) >= proto.(uint16)
		}

	case "int8:uint32":
		if value.(int8) >= 0 {
			return uint32(value.(int8)) >= proto.(uint32)
		}

	case "int8:uint64":
		if value.(int8) >= 0 {
			return uint64(value.(int8)) >= proto.(uint64)
		}

	case "int8:uint":
		if value.(int8) >= 0 {
			return uint(value.(int8)) >= proto.(uint)
		}

	// ...

	case "int16:uint8", "int16:byte":
		return value.(int16) >= int16(proto.(uint8))

	case "int16:uint16":
		if value.(int16) >= 0 {
			return uint16(value.(int16)) >= proto.(uint16)
		}

	case "int16:uint32":
		if value.(int16) >= 0 {
			return uint32(value.(int16)) >= proto.(uint32)
		}

	case "int16:uint64":
		if value.(int16) >= 0 {
			return uint64(value.(int16)) >= proto.(uint64)
		}

	case "int16:uint":
		if value.(int16) >= 0 {
			return uint(value.(int16)) >= proto.(uint)
		}

	// ...

	case "int32:uint8", "int32:byte", "rune:uint8", "rune:byte":
		return value.(int32) >= int32(proto.(uint8))

	case "int32:uint16", "rune:uint16":
		return value.(int32) >= int32(proto.(uint16))

	case "int32:uint32", "rune:uint32":
		if value.(int32) >= 0 {
			return uint32(value.(int32)) >= proto.(uint32)
		}

	case "int32:uint64", "rune:uint64":
		if value.(int32) >= 0 {
			return uint64(value.(int32)) >= proto.(uint64)
		}

	case "int32:uint", "rune:uint":
		if value.(int32) >= 0 {
			return uint(value.(int32)) >= proto.(uint)
		}

	// ...

	case "int64:uint8", "int64:byte":
		return value.(int64) >= int64(proto.(uint8))

	case "int64:uint16":
		return value.(int64) >= int64(proto.(uint16))

	case "int64:uint32":
		return value.(int64) >= int64(proto.(uint32))

	case "int64:uint64":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) >= uint64(proto.(uint64))
		}

	case "int64:uint":
		if value.(int64) >= 0 {
			return uint64(value.(int64)) >= uint64(proto.(uint))
		}

	// ...

	case "int:uint8", "int:byte":
		return value.(int) >= int(proto.(uint8))

	case "int:uint16":
		return value.(int) >= int(proto.(uint16))

	case "int:uint32":
		if value.(int) >= 0 {
			return uint64(value.(int)) >= uint64(proto.(uint32))
		}

	case "int:uint64":
		if value.(int) >= 0 {
			return uint64(value.(int)) >= proto.(uint64)
		}

	case "int:uint":
		if value.(int) >= 0 {
			return uint(value.(int)) >= proto.(uint)
		}

	// ............................................................
	// uint

	case "uint8:uint8", "uint8:byte", "byte:uint8", "byte:byte":
		return value.(uint8) >= proto.(uint8)

	case "uint8:uint16", "byte:uint16":
		return uint16(value.(uint8)) >= proto.(uint16)

	case "uint8:uint32", "byte:uint32":
		return uint32(value.(uint8)) >= proto.(uint32)

	case "uint8:uint64", "byte:uint64":
		return uint64(value.(uint8)) >= proto.(uint64)

	case "uint8:uint", "byte:uint":
		return uint(value.(uint8)) >= proto.(uint)

	case "uint16:uint8", "uint16:byte":
		return value.(uint16) >= uint16(proto.(uint8))

	case "uint16:uint16":
		return value.(uint16) >= proto.(uint16)
	case "uint16:uint32":
		return uint32(value.(uint16)) >= proto.(uint32)
	case "uint16:uint64":
		return uint64(value.(uint16)) >= proto.(uint64)
	case "uint16:uint":
		return uint(value.(uint16)) >= proto.(uint)

	case "uint32:uint8", "uint32:byte":
		return value.(uint32) >= uint32(proto.(uint8))
	case "uint32:uint16":
		return value.(uint32) >= uint32(proto.(uint16))
	case "uint32:uint32":
		return value.(uint32) >= proto.(uint32)
	case "uint32:uint64":
		return uint64(value.(uint32)) >= proto.(uint64)
	case "uint32:uint":
		return uint(value.(uint32)) >= proto.(uint)

	case "uint64:uint8", "uint64:byte":
		return value.(uint64) >= uint64(proto.(uint8))
	case "uint64:uint16":
		return value.(uint64) >= uint64(proto.(uint16))
	case "uint64:uint32":
		return value.(uint64) >= uint64(proto.(uint32))
	case "uint64:uint64":
		return value.(uint64) >= proto.(uint64)
	case "uint64:uint":
		return value.(uint64) >= uint64(proto.(uint))

	case "uint:uint8", "uint:byte":
		return value.(uint) >= uint(proto.(uint8))
	case "uint:uint16":
		return value.(uint) >= uint(proto.(uint16))
	case "uint:uint32":
		return value.(uint) >= uint(proto.(uint32))
	case "uint:uint64":
		return uint64(value.(uint)) >= proto.(uint64)
	case "uint:uint":
		return value.(uint) >= proto.(uint)

	// ............................................................
	// uint & int

	case "uint8:int8", "byte:int8":
		if proto.(int8) >= 0 {
			return value.(uint8) >= uint8(proto.(int8))
		}
		return true

	case "uint8:int16", "byte:int16":
		return int16(value.(uint8)) >= proto.(int16)

	case "uint8:int32", "uint8:rune", "byte:int32", "byte:rune":
		return int32(value.(uint8)) >= proto.(int32)

	case "uint8:int64", "byte:int64":
		return int64(value.(uint8)) >= proto.(int64)

	case "uint8:int", "byte:int":
		return int(value.(uint8)) >= proto.(int)

	// ...

	case "uint16:int8":
		if proto.(int8) >= 0 {
			return value.(uint16) >= uint16(proto.(int8))
		}
		return true

	case "uint16:int16":
		if proto.(int16) >= 0 {
			return value.(uint16) >= uint16(proto.(int16))
		}
		return true

	case "uint16:int32", "uint16:rune":
		return int32(value.(uint16)) >= proto.(int32)

	case "uint16:int64":
		return int64(value.(uint16)) >= proto.(int64)

	case "uint16:int":
		return int(value.(uint16)) >= proto.(int)

	// ...

	case "uint32:int8":
		if proto.(int8) >= 0 {
			return value.(uint32) >= uint32(proto.(int8))
		}
		return true

	case "uint32:int16":
		if proto.(int16) >= 0 {
			return value.(uint32) >= uint32(proto.(int16))
		}
		return true

	case "uint32:int32", "uint32:rune":
		if proto.(int32) >= 0 {
			return value.(uint32) >= uint32(proto.(int32))
		}
		return true

	case "uint32:int64":
		return int64(value.(uint32)) >= proto.(int64)

	case "uint32:int":
		return int64(value.(uint32)) >= int64(proto.(int))

	// ...

	case "uint64:int8":
		if proto.(int8) >= 0 {
			return value.(uint64) >= uint64(proto.(int8))
		}
		return true

	case "uint64:int16":
		if proto.(int16) >= 0 {
			return value.(uint64) >= uint64(proto.(int16))
		}
		return true

	case "uint64:int32", "uint64:rune":
		if proto.(int32) >= 0 {
			return value.(uint64) >= uint64(proto.(int32))
		}
		return true

	case "uint64:int64":
		if proto.(int64) >= 0 {
			return value.(uint64) >= uint64(proto.(int64))
		}
		return true

	case "uint64:int":
		if proto.(int) >= 0 {
			return value.(uint64) >= uint64(proto.(int))
		}
		return true

	// ...

	case "uint:int8":
		if proto.(int8) >= 0 {
			return value.(uint) >= uint(proto.(int8))
		}
		return true

	case "uint:int16":
		if proto.(int16) >= 0 {
			return value.(uint) >= uint(proto.(int16))
		}
		return true

	case "uint:int32":
		if proto.(int32) >= 0 {
			return value.(uint) >= uint(proto.(int32))
		}
		return true

	case "uint:int64":
		if proto.(int64) >= 0 {
			return uint64(value.(uint)) >= uint64(proto.(int64))
		}
		return true

	case "uint:int":
		if proto.(int) >= 0 {
			return value.(uint) >= uint(proto.(int))
		}
		return true

	}

	return false
}
