package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"unicode/utf8"
)

const (
	NON_ZERO = "NonZero"

	MsgMinStrLen      = "must contain at least %d characters"
	MsgMaxStrLen      = "must contain up to %d characters"
	MsgMinSetLen      = "must contain at least %d items"
	MsgMaxSetLen      = "must contain up to %d items"
	MsgMin            = "must be at least %d"
	MsgMax            = "must be up to %d"
	MsgNotValid       = "is not valid"
	MsgEmpty          = "is empty"
	MsgInvalidValue   = "has invalid value"
	MsgInvalidRule    = "has invalid rule"
	MsgInvalidBodyVal = "invalid body value"
)

type Group []any
type Rule [2]any

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
	case "validator.Group":

		for n := 0; n < rules.Len(); n++ {
			item := reflect.Indirect(reflect.ValueOf(
				rules.Index(n).Interface(),
			))

			if hint := checkField(item, value); hint != "" {
				return hint
			}
		}

	case "validator.Rule":
		action := rules.Index(0).Elem().String()
		proto := rules.Index(1).Elem()

		return compare(action, proto, value)

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

	if !proto.IsValid() {
		return MsgInvalidRule
	}

	switch action {
	case "min":
		return filterMin(proto, value)

	case "max":
		return filterMax(proto, value)

	case "match":
		if !IsMatch(proto, value) {
			return MsgNotValid
		}

	case "eachMatch":
		if (proto.Kind() != reflect.String) || (proto.Len() == 0) {
			return MsgInvalidRule
		}

		if !IsEachMatch(proto.String(), value) {
			return MsgNotValid
		}
	}

	return ""
}

func filterMin(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgMinStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgMinSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgMin, proto.Interface())
	}

	if !IsMin(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func filterMax(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgMaxStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgMaxSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgMax, proto.Interface())
	}

	if !IsMax(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func IsMatch(reg, value reflect.Value) (flag bool) {
	if reg.Kind() == reflect.String && value.Kind() == reflect.String {
		flag, _ = regexp.MatchString(reg.String(), value.String())
	}

	return
}

func IsEachMatch(reg string, value reflect.Value) bool {
	isValid := false

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		if isValid = value.Type().Elem().Kind() == reflect.String; !isValid {
			return false
		}

		for n := 0; n < value.Len(); n++ {
			matched, _ := regexp.MatchString(reg, value.Index(n).String())
			isValid = isValid && matched
		}

	case reflect.Map:
		if isValid = value.Type().Elem().Kind() == reflect.String; !isValid {
			return false
		}

		iter := value.MapRange()

		for iter.Next() {
			// k := iter.Key()
			// v := iter.Value()

			matched, _ := regexp.MatchString(reg, iter.Value().String())
			isValid = isValid && matched
		}
	}

	return isValid
}
