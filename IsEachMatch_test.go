package validator

import (
	"reflect"
	"testing"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestIsEachMatch .

func TestIsEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`IsEachMatch`, func() {
		g.Describe("slice", func() {
			reg := `(?i)^[0-9a-f]{32}$`

			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"f41c909f66bdedb0fb0c19711bcf3b73",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty slice", func() {
				result := IsEachMatch(reg, reflect.ValueOf([]string{}))
				g.Assert(result).IsTrue()
			})

			g.It(`success when given an empty mask`, func() {
				result := IsEachMatch(``, reflect.ValueOf([]string{"str"}))
				g.Assert(result).IsTrue()
			})

			// g.It("failure when given nil instead of mask", func() {
			// 	result := IsEachMatch(nil, reflect.ValueOf([]string{}))
			// 	g.Assert(result).IsFalse()
			// })

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"zzz",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsFalse()
			})

		})

		g.Describe("array", func() {
			reg := `^38[0-9]{10}$`

			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf([2]string{
					"380001234567",
					"380007654321",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty array", func() {
				proto := reflect.ValueOf([0]string{})
				result := IsEachMatch(reg, proto)
				g.Assert(result).IsTrue()
			})

			g.It(`success when given an empty mask`, func() {
				proto := reflect.ValueOf([1]string{"str"})
				result := IsEachMatch(``, proto)
				g.Assert(result).IsTrue()
			})

			// g.It("failure when given nil instead of mask", func() {
			// 	proto := reflect.ValueOf([0]string{})
			// 	result := IsEachMatch(nil, proto)
			// 	g.Assert(result).IsFalse()
			// })

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf([]string{
					"380001234567",
					"0001234567",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsFalse()
			})

		})

		g.Describe("map", func() {
			reg := `^https\://img\.domain\.com/[0-9A-Fa-f]{32}\.(?:pn|jpe?)g$`

			g.It("success when values match the mask", func() {
				proto := reflect.ValueOf(map[string]string{
					"img1": "https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
					"img2": "https://img.domain.com/4792592a98f8b9143de71d1db403d163.jpg",
					"img3": "https://img.domain.com/92f2b876b8ea94f711d2173539e73802.png",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty map", func() {
				proto := reflect.ValueOf(map[string]string{})
				result := IsEachMatch(reg, proto)
				g.Assert(result).IsTrue()
			})

			g.It(`success when given an empty mask`, func() {
				proto := reflect.ValueOf(map[string]string{"k": "v"})
				result := IsEachMatch(``, proto)
				g.Assert(result).IsTrue()
			})

			// g.It("failure when given nil instead of mask", func() {
			// 	proto := reflect.ValueOf(map[string]string{})
			// 	result := IsEachMatch(nil, proto)
			// 	g.Assert(result).IsFalse()
			// })

			g.It("failure when at least 1 value does not match", func() {
				proto := reflect.ValueOf(map[string]string{
					"img1": "https://img.domain.com/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
					"img2": "https://hack.it/animation.gif",
				})

				result := IsEachMatch(reg, proto)
				g.Assert(result).IsFalse()
			})
		})

		g.Describe("emptiness", func() {
			g.It(`success when given an empty mask`, func() {
				proto := reflect.ValueOf([]string{
					"b0fb0c19711bcf3b73f41c909f66bded",
					"f41c909f66bdedb0fb0c19711bcf3b73",
				})

				result := IsEachMatch(``, proto)
				g.Assert(result).IsTrue()
			})

			g.It("success when given an empty value", func() {
				proto := reflect.ValueOf([]string{})
				result := IsEachMatch(`(?i)^[0-9a-f]{32}$`, proto)
				g.Assert(result).IsTrue()
			})
		})
	})
}
