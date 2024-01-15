package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

const (
	NON_ZERO = "NonZero"

	MsgMinStrLen      = "must contain at least %v characters"
	MsgMaxStrLen      = "must contain up to %v characters"
	MsgEqStrLen       = "must contain exactly %v characters"
	MsgRangeStrLen    = "must contain %v..%v characters"
	MsgMinSetLen      = "must contain at least %v items"
	MsgMaxSetLen      = "must contain up to %v items"
	MsgEqSetLen       = "must contain exactly %v items"
	MsgRangeSetLen    = "must contain %v..%v items"
	MsgMin            = "must be at least %v"
	MsgMax            = "must be up to %v"
	MsgEq             = "must be exactly %v"
	MsgRange          = "must be in the range %v..%v"
	MsgNotValid       = "is not valid"
	MsgEmpty          = "is empty"
	MsgUnsupportType  = "has unsupported type to validate"
	MsgInvalidValue   = "has invalid value"
	MsgInvalidRule    = "has invalid rule"
	MsgInvalidBodyVal = "invalid body value"
)

var (
	refNil = reflect.ValueOf(nil)
)

type Group []any
type Range [2]any
type Rule [2]any

type FilterItem struct {
	Field    string
	Check    any
	Optional bool
}

type Filter []FilterItem

// Checks the fields of the structure according to the specified rules.
// Returns false or true, respectively
func (filter Filter) IsValid(data any) bool {
	return len(filter.Validate(data)) == 0
}

// Checks the fields of the structure according to the specified rules.
// Returns a slice with error hints if at least one field is not valid,
// otherwise, it will return an empty slice
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

		if hint := checkOthers(rules, successFields); hint != "" {
			hints = append(hints, hint)
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

	case "validator.Range":
		return compare("range", rules, value)

	case "validator.Rule":
		action := rules.Index(0).Elem().String()
		proto := rules.Index(1).Elem()

		return compare(action, proto, value)

	case "string":
		action := rules.String()
		proto := reflect.ValueOf(nil)

		return compare(action, proto, value)

	case "interface {}":
		unpackedRules := reflect.Indirect(reflect.ValueOf(rules.Interface()))
		if hint := checkField(unpackedRules, value); hint != "" {
			return hint
		}
	}

	return ""
}

func checkOthers(rules reflect.Value, successFields int) string {
	var (
		action = ""
		value  = reflect.ValueOf(nil)
		proto  reflect.Value
	)

	switch rules.Type().String() {
	case "validator.Rule":
		action = rules.Index(0).Elem().String()
		proto = rules.Index(1).Elem()

		if action[0:6] == "fields" {
			action = action[7:]
			value = reflect.ValueOf(successFields)
		}

		if hint := compare(action, proto, value); hint != "" {
			return MsgInvalidBodyVal
		}

	default:
		return MsgInvalidRule
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
	case "range":
		return filterRange(proto, value)

	case "min":
		return filterMin(proto, value)

	case "max":
		return filterMax(proto, value)

	case "eq":
		return filterEq(proto, value)

	case "match":
		return filterMatch(proto, value)

	// modifiers
	case "each:range", "each:min", "each:max", "each:eq", "each:match":
		return filterEach(action[5:], proto, value)

	case "date:min", "date:max", "date:eq":
		return filterDate(action[5:], proto, value)

	case "year":
		if !IsYearEqual(proto.Interface(), value.Interface()) {
			return fmt.Sprintf(MsgEq, proto.Interface())
		}

	default:
		return MsgInvalidRule
	}

	return ""
}

func filterRange(proto, value reflect.Value) string {
	hint := ""

	valMin := proto.Index(0)
	valMax := proto.Index(1)

	if valMin.Equal(refNil) || valMax.Equal(refNil) {
		return MsgInvalidRule
	}

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgRangeStrLen, valMin.Interface(), valMax.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgRangeSetLen, valMin.Interface(), valMax.Interface())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		hint = fmt.Sprintf(MsgRange, valMin.Interface(), valMax.Interface())

	default:
		hint = MsgUnsupportType
	}

	if !IsMin(valMin.Interface(), value.Interface()) {
		return hint
	}
	if !IsMax(valMax.Interface(), value.Interface()) {
		return hint
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

func filterEq(proto, value reflect.Value) string {
	hint := ""

	switch value.Kind() {
	case reflect.String:
		value = reflect.ValueOf(utf8.RuneCountInString(value.String()))
		hint = fmt.Sprintf(MsgEqStrLen, proto.Interface())

	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		value = reflect.ValueOf(value.Len())
		hint = fmt.Sprintf(MsgEqSetLen, proto.Interface())

	default:
		hint = fmt.Sprintf(MsgEq, proto.Interface())
	}

	if !IsEqual(proto.Interface(), value.Interface()) {
		return hint
	}

	return ""
}

func filterEach(action string, proto, value reflect.Value) string {
	switch action {
	case "match":
		if (proto.Kind() != reflect.String) || (proto.Len() == 0) {
			return MsgInvalidRule
		}
	}

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for n := 0; n < value.Len(); n++ {
			if hint := compare(action, proto, value.Index(n)); hint != "" {
				return fmt.Sprintf("item[%v] "+hint, n)
			}
		}

		return ""

	case reflect.Map:
		iter := value.MapRange()

		for iter.Next() {
			if hint := compare(action, proto, iter.Value()); hint != "" {
				return fmt.Sprintf("item[%v] "+hint, iter.Key())
			}
		}

		return ""
	}

	return MsgUnsupportType
}

func filterDate(action string, proto, value reflect.Value) string {
	var tmProto, tmValue int64

	if proto.Equal(refNil) {
		return MsgInvalidRule
	}

	if value.Equal(refNil) {
		return MsgInvalidValue
	}

	switch proto.Type().String() + ":" + value.Type().String() {
	case "int64:time.Time":
		tmProto = proto.Int()
		tmValue = value.Interface().(time.Time).Unix()

	case "time.Time:time.Time":
		tmProto = proto.Interface().(time.Time).Unix()
		tmValue = value.Interface().(time.Time).Unix()

	case "string:time.Time":
		t, err := time.Parse(time.RFC3339, proto.String())
		if err != nil {
			return MsgInvalidRule
		}

		tmProto = t.Unix()
		tmValue = value.Interface().(time.Time).Unix()

	default:
		return MsgUnsupportType
	}

	switch action {
	case "min":
		if tmValue < tmProto {
			return fmt.Sprintf(MsgMin, time.Unix(tmProto, 0).UTC().Format(time.RFC3339))
		}

	case "max":
		if tmValue > tmProto {
			return fmt.Sprintf(MsgMax, time.Unix(tmProto, 0).UTC().Format(time.RFC3339))
		}

	case "eq":
		if tmValue != tmProto {
			return fmt.Sprintf(MsgEq, time.Unix(tmProto, 0).UTC().Format(time.RFC3339))
		}

	default:
		return MsgInvalidRule
	}

	return ""
}

func filterMatch(reg, value reflect.Value) string {
	match, err := regexp.MatchString(reg.String(), value.String())

	switch {
	case err != nil:
		return MsgInvalidRule

	case !match:
		return MsgNotValid
	}

	return ""
}
