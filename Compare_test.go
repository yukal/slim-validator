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

// go test -v -run TestCompareMin .

func TestCompareMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "min"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value exceeds the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must be at least 4")
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length exceeds the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf("love")

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf("love")

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf("love")

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length exceeds the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length exceeds the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "val1",
				2: "val2",
				3: "val3",
				4: "val4",
			}

			g.It("success when the length exceeds the min threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the min threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than the min threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := compare("min", proto, value)
				g.Assert(result).Equal("must contain at least 8 items")
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(10)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It("failure when given an invalid value", func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareMax .

func TestCompareMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "max"`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the value reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value exceeds the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must be up to 4")
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf("love")

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf("love")

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf("love")

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "val1",
				2: "val2",
				3: "val3",
				4: "val4",
			}

			g.It("success when the length is less than the max threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("success when the length reaches the max threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length exceeds the max threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("max", proto, value)
				g.Assert(result).Equal("must contain up to 2 items")
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(10)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It("failure when given an invalid value", func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareEq .

func TestCompareEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eq" (equal)`, func() {
		g.Describe("numeric", func() {
			g.It("success when the value equals a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(4)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the value is less than a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(-4)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})

			g.It("failure when the value exceeds a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(8)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must be exactly 4")
			})
		})

		// ...

		g.Describe("string", func() {
			g.It("success when the length equals a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf("love")

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf("")

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 4 characters")
			})

			g.It("failure when the length exceeds a threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf("love")

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length equals a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(arrFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than a threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(arrFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(arrFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length equals a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(sliceFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than a threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(sliceFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(sliceFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "val1",
				2: "val2",
				3: "val3",
				4: "val4",
			}

			g.It("success when the length equals a threshold", func() {
				proto := reflect.ValueOf(4)
				value := reflect.ValueOf(mapFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is less than a threshold", func() {
				proto := reflect.ValueOf(8)
				value := reflect.ValueOf(mapFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				proto := reflect.ValueOf(2)
				value := reflect.ValueOf(mapFilled)

				result := compare("eq", proto, value)
				g.Assert(result).Equal("must contain exactly 2 items")
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(10)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It("failure when given an invalid value", func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareMatch .

func TestCompareMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "match"`, func() {
		g.It("success when value match the mask", func() {
			proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
			value := reflect.ValueOf("b0fb0c19711bcf3b73f41c909f66bded")

			g.Assert(compare("match", proto, value)).Equal("")
		})

		g.It("failure when given an empty mask", func() {
			proto := reflect.ValueOf(``)
			value := reflect.ValueOf("abra")

			g.Assert(compare("match", proto, value)).Equal("")
		})

		g.It("failure when a value does not match the mask", func() {
			proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
			value := reflect.ValueOf("Z0fz0c19711bcf3b73f41c909f66bded")

			g.Assert(compare("match", proto, value)).Equal(MsgNotValid)
		})

		g.It("failure when given invalid mask", func() {
			proto := reflect.ValueOf(nil)
			value := reflect.ValueOf("cadabra")

			g.Assert(compare("match", proto, value)).Equal(MsgInvalidRule)
		})

		g.It(`failure when given invalid value`, func() {
			proto := reflect.ValueOf(``)
			value := reflect.ValueOf(nil)

			g.Assert(compare("match", proto, value)).Equal(MsgInvalidValue)
		})
	})
}

// go test -v -run TestCompareEachMatch .

func TestCompareEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachMatch"`, func() {
		g.Describe(`array`, func() {
			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([2]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"f41c909f66bdedb0fb0c19711bcf3b73",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([0]string{})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([2]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"zzz",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgNotValid)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf([1]string{"str"})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf([1]string{"str"})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"f41c909f66bdedb0fb0c19711bcf3b73",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([]string{})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"zzz",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgNotValid)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf([]string{"str"})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf([]string{"b0fb0c19711bcf3b73f41c909f66bded"})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidValue)
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(map[int]string{
					1: "b0fb0c19711bcf3b73f41c909f66bded",
					2: "f41c909f66bdedb0fb0c19711bcf3b73",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(map[int]string{})

				g.Assert(compare("eachMatch", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(map[int]string{
					1: "b0fb0c19711bcf3b73f41c909f66bded",
					2: "zzz",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgNotValid)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf(map[int]string{
					1: "str",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(map[int]string{
					1: "b0fb0c19711bcf3b73f41c909f66bded",
				})

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("eachMatch", proto, value)).Equal(MsgInvalidValue)
			})
		})
	})
}

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
