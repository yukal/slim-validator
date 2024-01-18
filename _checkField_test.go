package validator

import (
	"reflect"
	"testing"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestCheckField .

func TestCheckField(t *testing.T) {
	g := Goblin(t)

	g.Describe(`checkField`, func() {
		g.Describe(`string rule`, func() {
			g.It("return empty hint", func() {
				rules := reflect.ValueOf(NON_ZERO)
				value := reflect.ValueOf("ok")
				hint := checkField(rules, value)

				g.Assert(hint).Equal("", hint)
			})

			g.It("return filled hint", func() {
				rules := reflect.ValueOf(NON_ZERO)
				value := reflect.ValueOf("")
				hint := checkField(rules, value)

				g.Assert(hint).Equal(MsgEmpty, hint)
			})
		})

		g.Describe(`validator.Rule`, func() {
			g.It("return empty hint", func() {
				rules := reflect.ValueOf(Rule{"match", `^\+38\d{10}$`})
				value := reflect.ValueOf("+380001234567")
				hint := checkField(rules, value)

				g.Assert(hint).Equal("", hint)
			})

			g.It("return filled hint", func() {
				rules := reflect.ValueOf(Rule{"match", `^\+38\d{10}$`})
				value := reflect.ValueOf("+38(000)123-45-67")
				hint := checkField(rules, value)

				g.Assert(hint).Equal(MsgNotValid, hint)
			})
		})

		g.Describe(`validator.Range`, func() {
			g.It("return empty hint", func() {
				rules := reflect.ValueOf(Range{1, 25})
				value := reflect.ValueOf(25)
				hint := checkField(rules, value)

				g.Assert(hint).Equal("", hint)
			})

			g.It("return filled hint", func() {
				rules := reflect.ValueOf(Range{1, 25})
				value := reflect.ValueOf(30)
				hint := checkField(rules, value)

				g.Assert(hint).Equal("must be in the range 1..25", hint)
			})
		})

		g.Describe(`validator.Group`, func() {
			g.It("return empty hint", func() {
				rules := reflect.ValueOf(Group{
					Rule{"min", 1},
					Rule{"each:match", `^\+38\d{10}$`},
				})

				value := reflect.ValueOf([]string{
					"+380001234567",
					"+380007654321",
				})

				hint := checkField(rules, value)
				g.Assert(hint).Equal("", hint)
			})

			g.It("return filled hint", func() {
				rules := reflect.ValueOf(Group{
					Rule{"min", 1},
				})

				value := reflect.ValueOf([]string{})
				hint := checkField(rules, value)

				g.Assert(hint).Equal("must contain at least 1 items", hint)
			})
		})
	})
}
