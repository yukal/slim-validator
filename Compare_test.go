package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go clean -testcache
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

// go test -v -run TestCompareRange .

func TestCompareRange(t *testing.T) {
	var (
		strFilled   = "love"
		arrFilled   = [4]string{"c", "o", "d", "e"}
		sliceFilled = []string{"t", "e", "s", "t"}
		mapFilled   = map[string]string{
			"i": "val1",
			"t": "val2",
			"e": "val3",
			"m": "val4",
		}
	)

	g := Goblin(t)

	g.Describe(`Rule "range"`, func() {
		g.Describe(`numeric`, func() {
			g.It("success when the value matches the range", func() {
				proto := reflect.ValueOf(Range{1, 25})
				value := reflect.ValueOf(15)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when given below-range value", func() {
				proto := reflect.ValueOf(Range{15, 25})
				value := reflect.ValueOf(5)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must be in the range 15..25")
			})

			g.It("failure when given above-range value", func() {
				proto := reflect.ValueOf(Range{15, 25})
				value := reflect.ValueOf(55)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must be in the range 15..25")
			})
		})

		// ...

		g.Describe(`array`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(arrFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(sliceFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(mapFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 items")
			})
		})

		// ...

		g.Describe(`string`, func() {
			g.It("success when the length matches the range", func() {
				proto := reflect.ValueOf(Range{1, 4})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("")
			})

			g.It("failure when the length is below the range", func() {
				proto := reflect.ValueOf(Range{10, 80})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 10..80 characters")
			})

			g.It("failure when the length is above the range", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal("must contain 1..3 characters")
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure if at least 1 element of the range is invalid", func() {
				value := reflect.ValueOf(arrFilled)
				protos := []Range{
					{},
					{nil, 4},
					{4, nil},
				}

				for _, item := range protos {
					proto := reflect.ValueOf(item)
					result := compare("range", proto, value)

					g.Assert(result).Equal(MsgInvalidRule)
				}
			})

			g.It("failure when given invalid range", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(strFilled)

				result := compare("range", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid value", func() {
				proto := reflect.ValueOf(Range{1, 3})
				value := reflect.ValueOf(nil)

				result := compare("range", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareYear .

func TestCompareYear(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value matches a specific year", func() {
			tm, err := time.Parse(time.RFC3339, "2023-12-25T16:04:05Z")
			g.Assert(err).IsNil(err)

			proto := reflect.ValueOf(2023)
			value := reflect.ValueOf(tm)

			result := compare("year", proto, value)
			g.Assert(result).Equal("")
		})

		g.It("failure when the value is not match", func() {
			proto := reflect.ValueOf(2024)
			value := reflect.ValueOf(*new(time.Time))

			result := compare("year", proto, value)
			g.Assert(result).Equal("must be exactly 2024")
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

// go test -v -run TestCompareEachMatchDeprecated .

func TestCompareEachMatchDeprecated(t *testing.T) {
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

// go test -v -run TestCompareEachMin .

func TestCompareEachMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:min"`, func() {
		g.Describe(`array`, func() {
			g.Describe(`array:numeric`, func() {
				g.It("success when the element value exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([2]int{15, 25})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the min threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]int{15, 25})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([0]int{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([2]int{5, 15})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must be at least 10")
				})
			})

			// ...

			g.Describe(`array:string`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([2]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(20)
					value := reflect.ValueOf([2]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([0]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(18)
					value := reflect.ValueOf([2]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must contain at least 18 characters")
				})
			})

			// ...

			g.Describe(`array:array`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([0][0]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([2][3]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`array:slice`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([0][]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Lady Madonna",
							"Let it be",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[1] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`array:map`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([0]map[int]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Lady Madonna",
							4: "Let it be",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[1] must contain at least 4 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.Describe(`slice:numeric`, func() {
				g.It("success when the element value exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{15, 25})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the min threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]int{15, 25})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{5, 15})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must be at least 10")
				})
			})

			// ...

			g.Describe(`slice:string`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(20)
					value := reflect.ValueOf([]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(18)
					value := reflect.ValueOf([]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must contain at least 18 characters")
				})
			})

			// ...

			g.Describe(`slice:array`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([][2]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`slice:slice`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([][]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Lady Madonna",
							"Let it be",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[1] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`slice:map`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([]map[int]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Lady Madonna",
							4: "Let it be",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[1] must contain at least 4 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.Describe(`map:numeric`, func() {
				g.It("success when the element value exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 25,
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the min threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 25,
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{
						"first":  5,
						"second": 15,
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[first] must be at least 10")
				})
			})

			// ...

			g.Describe(`map:string`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]string{
						"first":  "We all live in a yellow submarine",
						"second": "All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(20)
					value := reflect.ValueOf(map[string]string{
						"first":  "We all live in a yellow submarine",
						"second": "All you need is love",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(18)
					value := reflect.ValueOf(map[string]string{
						"first":  "No pain no gain",
						"second": "Fight fire with fire",
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[first] must contain at least 18 characters")
				})
			})

			// ...

			g.Describe(`map:array`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string][2]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`map:slice`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string][]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
							"Lady Madonna",
							"Let it be",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[Scorpions] must contain at least 4 items")
				})
			})

			// ...

			g.Describe(`map:map`, func() {
				g.It("success when the element length exceeds the min threshold", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the min threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string]map[int]string{})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Lady Madonna",
							4: "Let it be",
						},
						"Scorpions": {
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:min", proto, value)
					g.Assert(result).Equal("item[Scorpions] must contain at least 4 items")
				})
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf("Here In My Heart")

				result := compare("each:min", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It(`failure when given an invalid value`, func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("each:min", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareEachMax .

func TestCompareEachMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:max"`, func() {
		g.Describe(`array`, func() {
			g.Describe(`array:numeric`, func() {
				g.It("success when the element value is less than the max threshold", func() {
					proto := reflect.ValueOf(30)
					value := reflect.ValueOf([2]int{15, 25})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the max threshold", func() {
					proto := reflect.ValueOf(25)
					value := reflect.ValueOf([2]int{15, 25})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([0]int{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([2]int{5, 15})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[1] must be up to 10")
				})
			})

			// ...

			g.Describe(`array:string`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(40)
					value := reflect.ValueOf([2]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf([2]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([0]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[1] must contain up to 15 characters")
				})
			})

			// ...

			g.Describe(`array:array`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0][0]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][3]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`array:slice`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0][]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`array:map`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0]map[int]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Let it be",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.Describe(`slice:numeric`, func() {
				g.It("success when the element value is less than the max threshold", func() {
					proto := reflect.ValueOf(30)
					value := reflect.ValueOf([]int{15, 25})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the max threshold", func() {
					proto := reflect.ValueOf(25)
					value := reflect.ValueOf([]int{15, 25})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{5, 15})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[1] must be up to 10")
				})
			})

			// ...

			g.Describe(`slice:string`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(40)
					value := reflect.ValueOf([]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf([]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[1] must contain up to 15 characters")
				})
			})

			// ...

			g.Describe(`slice:array`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][2]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][3]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`slice:slice`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
						{
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`slice:map`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Let it be",
						},
						{
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[0] must contain up to 2 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.Describe(`map:numeric`, func() {
				g.It("success when the element value is less than the max threshold", func() {
					proto := reflect.ValueOf(30)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 25,
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element value reaches the max threshold", func() {
					proto := reflect.ValueOf(25)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 25,
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{
						"first":  5,
						"second": 15,
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[second] must be up to 10")
				})
			})

			// ...

			g.Describe(`map:string`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(40)
					value := reflect.ValueOf(map[string]string{
						"first":  "We all live in a yellow submarine",
						"second": "All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf(map[string]string{
						"first":  "We all live in a yellow submarine",
						"second": "All you need is love",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[string]string{
						"first":  "No pain no gain",
						"second": "Fight fire with fire",
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[second] must contain up to 15 characters")
				})
			})

			// ...

			g.Describe(`map:array`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][2]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][3]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`map:slice`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
							"Let it be",
						},
						"Scorpions": {
							"No pain no gain",
							"Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain up to 2 items")
				})
			})

			// ...

			g.Describe(`map:map`, func() {
				g.It("success when the element length is less than the max threshold", func() {
					proto := reflect.ValueOf(3)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when the element length reaches the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
							3: "Let it be",
						},
						"Scorpions": {
							1: "No pain no gain",
							2: "Here in my heart",
						},
					})

					result := compare("each:max", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain up to 2 items")
				})
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf("Here In My Heart")

				result := compare("each:max", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It(`failure when given an invalid value`, func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("each:max", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareEachEq .

func TestCompareEachEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:eq"`, func() {
		g.Describe(`array`, func() {
			g.Describe(`array:numeric`, func() {
				g.It("success when the element value equals the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]int{15, 15})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([0]int{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element value is less than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]int{5, 10})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must be exactly 15")
				})

				g.It("failure when the element value is greater than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]int{25, 30})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must be exactly 15")
				})
			})

			// ...

			g.Describe(`array:string`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([2]string{
						"test",
						"code",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([0]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf([2]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 33 characters")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([2]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 15 characters")
				})
			})

			// ...

			g.Describe(`array:array`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0][0]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][1]string{
						{
							"We all live in a yellow submarine",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe(`array:slice`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0][]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe(`array:map`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([0]map[int]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([2]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.Describe(`slice:numeric`, func() {
				g.It("success when the element value equals the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]int{15, 15})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]int{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element value is less than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]int{5, 10})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must be exactly 15")
				})

				g.It("failure when the element value is greater than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]int{25, 30})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must be exactly 15")
				})
			})

			// ...

			g.Describe(`slice:string`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf([]string{
						"test",
						"code",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf([]string{
						"We all live in a yellow submarine",
						"All you need is love",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 33 characters")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf([]string{
						"No pain no gain",
						"Fight fire with fire",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 15 characters")
				})
			})

			// ...

			g.Describe(`slice:array`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][0]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				// g.It("failure when the element length is less than expected", func() {
				// 	proto := reflect.ValueOf(2)
				// 	value := reflect.ValueOf([][2]string{
				// 		{
				// 			"We all live in a yellow submarine",
				// 			"All you need is love",
				// 		},
				// 		{
				// 			"No pain no gain",
				// 		},
				// 	})

				// 	result := compare("each:eq", proto, value)
				// 	g.Assert(result).Equal("item[1] must contain exactly 2 items")
				// })

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([][2]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe(`slice:slice`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([][]string{
						{
							"We all live in a yellow submarine",
							"All you need is love",
						},
						{
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe(`slice:map`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
							2: "Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[1] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf([]map[int]string{
						{
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						{
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[0] must contain exactly 1 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.Describe(`map:numeric`, func() {
				g.It("success when the element value equals the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 15,
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf(map[string]int{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element value is less than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[string]int{
						"first":  5,
						"second": 15,
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[first] must be exactly 15")
				})

				g.It("failure when the element value is greater than the expected number", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[string]int{
						"first":  15,
						"second": 30,
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[second] must be exactly 15")
				})
			})

			// ...

			g.Describe(`map:string`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(4)
					value := reflect.ValueOf(map[int]string{
						1: "test",
						2: "code",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(10)
					value := reflect.ValueOf([]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(33)
					value := reflect.ValueOf(map[int]string{
						1: "We all live in a yellow submarine",
						2: "All you need is love",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[2] must contain exactly 33 characters")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(15)
					value := reflect.ValueOf(map[int]string{
						1: "No pain no gain",
						2: "Fight fire with fire",
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[2] must contain exactly 15 characters")
				})
			})

			// ...

			g.Describe(`map:array`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][2]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][0]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][1]string{
						"Beatles": {
							"We all live in a yellow submarine",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][3]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"While My Guitar Gently Weeps",
							"All you need is love",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain exactly 2 items")
				})
			})

			// ...

			g.Describe(`map:slice`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
							"Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Scorpions] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf(map[string][]string{
						"Beatles": {
							"We all live in a yellow submarine",
							"All you need is love",
						},
						"Scorpions": {
							"No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe(`map:map`, func() {
				g.It("success when the element length match", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
							2: "Send Me an Angel",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("success when given an empty data list", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("")
				})

				g.It("failure when the element length is less than expected", func() {
					proto := reflect.ValueOf(2)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Scorpions] must contain exactly 2 items")
				})

				g.It("failure when the element value is greater than expected", func() {
					proto := reflect.ValueOf(1)
					value := reflect.ValueOf(map[string]map[int]string{
						"Beatles": {
							1: "We all live in a yellow submarine",
							2: "All you need is love",
						},
						"Scorpions": {
							1: "No pain no gain",
						},
					})

					result := compare("each:eq", proto, value)
					g.Assert(result).Equal("item[Beatles] must contain exactly 1 items")
				})
			})
		})

		// ...

		g.Describe("invalidity", func() {
			g.It("failure when given an invalid threshold", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf("Here In My Heart")

				result := compare("each:eq", proto, value)
				g.Assert(result).Equal(MsgInvalidRule)
			})

			g.It(`failure when given an invalid value`, func() {
				proto := reflect.ValueOf(10)
				value := reflect.ValueOf(nil)

				result := compare("each:eq", proto, value)
				g.Assert(result).Equal(MsgInvalidValue)
			})
		})
	})
}

// go test -v -run TestCompareEachMatch .

func TestCompareEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:match"`, func() {
		g.Describe(`array`, func() {
			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([2]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"f41c909f66bdedb0fb0c19711bcf3b73",
				})

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([0]string{})

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([2]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"zzz",
				})

				errorMsg := "item[1] " + MsgNotValid
				g.Assert(compare("each:match", proto, value)).Equal(errorMsg)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf([1]string{"str"})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf([1]string{"str"})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidValue)
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

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([]string{})

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"zzz",
				})

				errorMsg := "item[1] " + MsgNotValid
				g.Assert(compare("each:match", proto, value)).Equal(errorMsg)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf([]string{"str"})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf([]string{"b0fb0c19711bcf3b73f41c909f66bded"})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidValue)
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

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("success when given an empty slice", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(map[int]string{})

				g.Assert(compare("each:match", proto, value)).Equal("")
			})

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(map[int]string{
					1: "b0fb0c19711bcf3b73f41c909f66bded",
					2: "zzz",
				})

				errorMsg := "item[2] " + MsgNotValid
				g.Assert(compare("each:match", proto, value)).Equal(errorMsg)
			})

			g.It("failure when given an empty mask", func() {
				proto := reflect.ValueOf(``)
				value := reflect.ValueOf(map[int]string{
					1: "str",
				})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It("failure when given invalid mask", func() {
				proto := reflect.ValueOf(nil)
				value := reflect.ValueOf(map[int]string{
					1: "b0fb0c19711bcf3b73f41c909f66bded",
				})

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidRule)
			})

			g.It(`failure when given invalid value`, func() {
				proto := reflect.ValueOf(`(?i)^[0-9a-f]{32}$`)
				value := reflect.ValueOf(nil)

				g.Assert(compare("each:match", proto, value)).Equal(MsgInvalidValue)
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
