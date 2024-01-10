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
	})
}
