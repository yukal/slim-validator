package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestFilterDate .
// go test -v -cover -run TestFilterDateMin .
// go test -v -cover -run TestFilterDateMax .
// go test -v -cover -run TestFilterDateEq .

// go test -v -cover -run TestFilterDateMin .
func TestFilterDateMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`filter "date:min"`, func() {
		now := time.Now()

		action := "min"
		expectHint := fmt.Sprintf(MsgMin, now.UTC().Format(time.RFC3339))

		g.Describe(`compute within int64 & time.Time`, func() {
			g.It("success when the value exceeds the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within string & time.Time`, func() {
			g.It("success when the value exceeds the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when given an invalid proto", func() {
				hint := filterDate(
					action,
					reflect.ValueOf("buka-ka-ka-buku-ku"),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})
		})

		g.Describe(`compute within time.Time & time.Time`, func() {
			g.It("success when the value exceeds the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than the min threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within unsupported types`, func() {
			g.It("failure when the given proto has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf([]int{}),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given value has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf([]int{}),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given proto has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(nil),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})

			g.It("failure when the given value has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(nil),
				)

				g.Assert(hint).Equal(MsgInvalidValue, hint)
			})
		})
	})
}

// go test -v -cover -run TestFilterDateMax .
func TestFilterDateMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`filter "date:max"`, func() {
		now := time.Now()

		action := "max"
		expectHint := fmt.Sprintf(MsgMax, now.UTC().Format(time.RFC3339))

		g.Describe(`compute within int64 & time.Time`, func() {
			g.It("success when the value is less than the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value exceeds the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within string & time.Time`, func() {
			g.It("success when the value is less than the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value exceeds the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when given an invalid proto", func() {
				hint := filterDate(
					action,
					reflect.ValueOf("buka-ka-ka-buku-ku"),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})
		})

		g.Describe(`compute within time.Time & time.Time`, func() {
			g.It("success when the value is less than the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("success when the value reaches the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value exceeds the max threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within unsupported types`, func() {
			g.It("failure when the given proto has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf([]int{}),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given value has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf([]int{}),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given proto has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(nil),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})

			g.It("failure when the given value has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(nil),
				)

				g.Assert(hint).Equal(MsgInvalidValue, hint)
			})
		})
	})
}

// go test -v -cover -run TestFilterDateEq .
func TestFilterDateEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`filter "date:eq"`, func() {
		now := time.Now()

		action := "eq"
		expectHint := fmt.Sprintf(MsgEq, now.UTC().Format(time.RFC3339))

		g.Describe(`compute within int64 & time.Time`, func() {
			g.It("success when the value equals a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when the value exceeds a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within string & time.Time`, func() {
			g.It("success when the value equals a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when the value exceeds a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Format(time.RFC3339)),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when given an invalid proto", func() {
				hint := filterDate(
					action,
					reflect.ValueOf("buka-ka-ka-buku-ku"),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})
		})

		g.Describe(`compute within time.Time & time.Time`, func() {
			g.It("success when the value equals a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal("", hint)
			})

			g.It("failure when the value is less than a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(-time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})

			g.It("failure when the value exceeds a threshold", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now),
					reflect.ValueOf(now.Add(time.Second)),
				)

				g.Assert(hint).Equal(expectHint, hint)
			})
		})

		g.Describe(`compute within unsupported types`, func() {
			g.It("failure when the given proto has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf([]int{}),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given value has unsupported type", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf([]int{}),
				)

				g.Assert(hint).Equal(MsgUnsupportType, hint)
			})

			g.It("failure when the given proto has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(nil),
					reflect.ValueOf(now),
				)

				g.Assert(hint).Equal(MsgInvalidRule, hint)
			})

			g.It("failure when the given value has nil", func() {
				hint := filterDate(
					action,
					reflect.ValueOf(now.Unix()),
					reflect.ValueOf(nil),
				)

				g.Assert(hint).Equal(MsgInvalidValue, hint)
			})
		})
	})
}
