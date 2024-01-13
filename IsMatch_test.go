package validator

import (
	"reflect"
	"testing"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestFilterMatch .

func TestFilterMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`filter "match"`, func() {
		g.It("success when value match the mask", func() {
			result := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf("b0fb0c19711bcf3b73f41c909f66bded"),
			)

			g.Assert(result).IsTrue()
		})

		g.It("success when given an empty mask", func() {
			result := filterMatch(
				reflect.ValueOf(""),
				reflect.ValueOf("abra"),
			)

			g.Assert(result).IsTrue()
		})

		g.It("failure when a value does not match the mask", func() {
			result := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf("Z0fz0c19711bcf3b73f41c909f66bded"),
			)

			g.Assert(result).IsFalse()
		})

		g.It("failure when given invalid mask", func() {
			result := filterMatch(
				reflect.ValueOf(nil),
				reflect.ValueOf("cadabra"),
			)

			g.Assert(result).IsFalse()
		})

		g.It("failure when given invalid value", func() {
			result := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf(nil),
			)

			g.Assert(result).IsFalse()
		})
	})
}
