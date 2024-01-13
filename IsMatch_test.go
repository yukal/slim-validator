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
			hint := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf("b0fb0c19711bcf3b73f41c909f66bded"),
			)

			g.Assert(hint).Equal("")
		})

		g.It("success when given an empty mask", func() {
			hint := filterMatch(
				reflect.ValueOf(""),
				reflect.ValueOf("abra"),
			)

			g.Assert(hint).Equal("")
		})

		g.It("failure when a value does not match the mask", func() {
			hint := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf("Z0fz0c19711bcf3b73f41c909f66bded"),
			)

			g.Assert(hint).Equal(MsgNotValid)
		})

		g.It("failure when given invalid mask", func() {
			hint := filterMatch(
				reflect.ValueOf(`a(b`),
				reflect.ValueOf("a(b"),
			)

			g.Assert(hint).Equal(MsgInvalidRule)
		})

		g.It("failure when given nil mask", func() {
			hint := filterMatch(
				reflect.ValueOf(nil),
				reflect.ValueOf("cadabra"),
			)

			g.Assert(hint).Equal(MsgNotValid)
		})

		g.It("failure when given nil value", func() {
			hint := filterMatch(
				reflect.ValueOf(`(?i)^[0-9a-f]{32}$`),
				reflect.ValueOf(nil),
			)

			g.Assert(hint).Equal(MsgNotValid)
		})
	})
}
