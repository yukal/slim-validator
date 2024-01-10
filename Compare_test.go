package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestCompare .

// go test -v -run TestCompareNonZero .

func TestCompareNonZero(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule NON_ZERO`, func() {
		nonEmptyValues := []any{
			int8(-1),
			int16(-1),
			int32(-1),
			int64(-1),
			int(-1),
			uint8(1),
			uint16(1),
			uint32(1),
			uint64(1),
			uint(1),
			float32(1),
			float64(-1),
			complex64(1),
			complex128(-1),
			true,
			"ok",
		}

		for _, val := range nonEmptyValues {
			val := val

			g.It(fmt.Sprintf("success if given a non-zero %T", val), func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val)

				result := compare(NON_ZERO, proto, value)
				g.Assert(result).Equal("")
			})
		}

		var nonEmptyInterface interface{} = 1
		nonEmptyStructs := [][]any{
			{"array", [1]int{100}},
			{"slice", []int{100}},
			{"map", map[int]string{1: "ok"}},
			{"chan", make(chan int)},
			{"struct", time.Now()},
			{"func", func() string { return "ok" }},
			{"interface", nonEmptyInterface},
		}

		for _, val := range nonEmptyStructs {
			val := val

			g.It("success if given a non-zero "+val[0].(string), func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val[1])

				result := compare(NON_ZERO, proto, value)
				g.Assert(result).Equal("")
			})
		}

		emptyValues := []any{
			*new(int8),       // 0
			*new(int16),      // 0
			*new(int32),      // 0
			*new(int64),      // 0
			*new(int),        // 0
			*new(uint8),      // 0
			*new(uint16),     // 0
			*new(uint32),     // 0
			*new(uint64),     // 0
			*new(uint),       // 0
			*new(float32),    // 0.0
			*new(float64),    // 0.0
			*new(complex64),  // 0.0
			*new(complex128), // 0.0
			*new(bool),       // false
			*new(string),     // ""
		}

		for _, val := range emptyValues {
			val := val

			g.It(fmt.Sprintf("failure if given a zero %T", val), func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val)

				result := compare(NON_ZERO, proto, value)
				g.Assert(result).Equal(MsgEmpty)
			})
		}

		emptyStructs := [][]any{
			{"array", *new([0]string)},
			{"slice", *new([]string)},
			{"map", *new(map[int]string)},
			{"chan", *new(chan int)},
			{"struct", *new(struct{})},
			{"func", *new(func())},
		}

		for _, val := range emptyStructs {
			val := val

			g.It("failure if given a zero "+val[0].(string), func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(val[1])

				result := compare(NON_ZERO, proto, value)
				g.Assert(result).Equal(MsgEmpty)
			})
		}

		g.It("failure if given a zero interface", func() {
			proto := reflect.ValueOf(*new(interface{}))
			value := reflect.ValueOf(*new(interface{}))

			result := compare(NON_ZERO, proto, value)
			g.Assert(result).Equal(MsgInvalidValue)
		})
	})
}
