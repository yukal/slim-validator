// For Coverage
//
// These tests are intended to improve the indicators of the test coverage, which has a minor role
// at the time of writing, given the predicted state of the behavior of the main code, since this file
// tests behavior that should never be executed but is only an insurance against a change in the future.

package validator

import (
	"reflect"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -v -run TestCoverageFilter .

func TestCoverageFilter(t *testing.T) {
	g := Goblin(t)

	g.Describe(`For Coverage: filter`, func() {
		g.It("failure filterMin when given invalid value", func() {
			hint := filterMin(
				reflect.ValueOf(NON_ZERO),
				reflect.ValueOf(nil),
			)

			g.Assert(hint).Equal(MsgInvalidValue, hint)
		})

		g.It("failure filterMax when given invalid value", func() {
			hint := filterMax(
				reflect.ValueOf(NON_ZERO),
				reflect.ValueOf(nil),
			)

			g.Assert(hint).Equal(MsgInvalidValue, hint)
		})

		g.It("failure filterEq when given invalid value", func() {
			hint := filterEq(
				reflect.ValueOf(NON_ZERO),
				reflect.ValueOf(nil),
			)

			g.Assert(hint).Equal(MsgInvalidValue, hint)
		})

		g.It("failure filterRange when given invalid value", func() {
			hint := filterRange(
				reflect.ValueOf(Range{1, 2}),
				reflect.ValueOf(nil),
			)

			g.Assert(hint).Equal(MsgInvalidValue, hint)
		})

		g.It("failure filterDate when given invalid rule", func() {
			now := time.Now()

			hint := filterDate(
				"invalid-rule",
				reflect.ValueOf(now),
				reflect.ValueOf(now),
			)

			g.Assert(hint).Equal(MsgInvalidRule, hint)
		})

		g.It("failure filterTime when given invalid rule", func() {
			now := time.Now()

			hint := filterTime(
				"invalid-rule",
				reflect.ValueOf(now),
				reflect.ValueOf(now),
			)

			g.Assert(hint).Equal(MsgInvalidRule, hint)
		})
	})
}
