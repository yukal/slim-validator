package validator

import (
	"reflect"
	"time"
)

func IsYearEqual(proto, value any) bool {
	refVal := reflect.ValueOf(value)

	if refVal.Kind() != reflect.Invalid {
		if refVal.Type().String() == "time.Time" {

			return IsEqual(proto, value.(time.Time).Year())

		}
	}

	return false
}
