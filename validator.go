package validator

import (
	"reflect"
)

const (
	NON_ZERO = "NonZero"

	MsgNotValid       = "is not valid"
	MsgEmpty          = "is empty"
	MsgInvalidValue   = "has invalid value"
	MsgInvalidRule    = "has invalid rule"
	MsgInvalidBodyVal = "invalid body value"
)

type FilterItem struct {
	Field    string
	Check    any
	Optional bool
}

type Filter []FilterItem

// data should be a type of struct{ ... }
func (filter Filter) Validate(data any) []string {
	refValData := reflect.Indirect(reflect.ValueOf(data))
	refTypData := refValData.Type()

	hints := make([]string, 0, refTypData.NumField())
	successFields := 0

	for _, filterStruct := range filter {
		// field := refTypData.Field(i)
		// field, fieldExist := refTypData.FieldByName(filterStruct.Field)
		rules := reflect.Indirect(reflect.ValueOf(filterStruct.Check))

		if field, exist := refTypData.FieldByName(filterStruct.Field); exist {
			tagName := filterStruct.Field
			// tagName = field.Tag.Get("json")

			if tagNameJson, exist := field.Tag.Lookup("json"); exist {
				tagName = tagNameJson
			}

			value := refValData.FieldByName(filterStruct.Field)

			if filterStruct.Optional && value.IsZero() {
				continue
			}

			if hint := checkField(rules, value); hint != "" {
				hints = append(hints, tagName+" "+hint)
			} else {
				successFields++
			}

			continue
		}
	}

	return hints
}

func checkField(rules, value reflect.Value) string {
	switch rules.Type().String() {

	case "string":
		action := rules.String()
		proto := reflect.ValueOf(nil)

		return compare(action, proto, value)
	}

	return ""
}

func compare(action string, proto, value reflect.Value) string {
	if !value.IsValid() {
		return MsgInvalidValue
	}

	switch action {
	case NON_ZERO:
		if value.IsZero() {
			return MsgEmpty
		}
		return ""
	}

	return ""
}
