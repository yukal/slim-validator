package validator

import (
	"math"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -v -cover .
// go test -v -cover -run TestValidate .

// go test -v -run TestValidateMatch .

func TestValidateMatch(t *testing.T) {
	type Article struct {
		Hash string `json:"hash"`
	}

	g := Goblin(t)

	g.Describe(`Rule "match"`, func() {
		msgInvalidRule := "hash " + MsgInvalidRule
		msgInvalidValue := "hash " + MsgNotValid

		g.It("success when the value matches the mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			hints := filter.Validate(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(len(hints)).Equal(0, hints)
		})

		g.It("success when given an empty mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", ``},
				},
			}

			hints := filter.Validate(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(len(hints)).Equal(0, hints)
		})

		g.It("failure when the value does not match the mask", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			hints := filter.Validate(Article{
				Hash: "Z0zZ0z19711zZz3z73z41z909z66zZzZ",
			})

			g.Assert(len(hints)).Equal(1, hints)
			g.Assert(hints[0]).Equal(msgInvalidValue)
		})

		g.It("failure when missing rule value", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match"},
				},
			}

			hints := filter.Validate(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(len(hints)).Equal(1, hints)
			g.Assert(hints[0]).Equal(msgInvalidRule)
		})

		g.It("failure when given an empty rule", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{},
				},
			}

			hints := filter.Validate(Article{
				Hash: "b0fb0c19711bcf3b73f41c909f66bded",
			})

			g.Assert(len(hints)).Equal(1, hints)
			g.Assert(hints[0]).Equal(msgInvalidRule)
		})

		g.It("failure when given an empty value", func() {
			filter := Filter{
				{
					Field: "Hash",
					Check: Rule{"match", `(?i)^[0-9a-f]{32}$`},
				},
			}

			hints := filter.Validate(Article{})
			g.Assert(len(hints)).Equal(1, hints)
			g.Assert(hints[0]).Equal(msgInvalidValue)
		})
	})
}

// go test -v -run TestValidateEachMatch .

func TestValidateEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eachMatch"`, func() {
		msgInvalidRule := "hash " + MsgInvalidRule
		msgInvalidValue := "hash " + MsgNotValid

		g.Describe(`array`, func() {
			type Array struct {
				Hash  [2]string `json:"hash"`
				Empty [0]string `json:"empty"`
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Empty",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Array{
					Empty: [0]string{},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				hints := filter.Validate(Array{
					Hash: [2]string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Hash []string `json:"hash"`
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				hints := filter.Validate(Slice{
					Hash: []string{
						"b0fb0c19711bcf3b73f41c909f66bded",
						"37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Hash map[int]string `json:"hash"`
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when given an empty list", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when at least 1 value does not match", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "Z0zZ0z19711zZz3z73z41z909z66zZzZ",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch"},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"eachMatch", ``},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})

			g.It("failure when given an empty rule", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "b0fb0c19711bcf3b73f41c909f66bded",
						2: "37763f73e30e7b0bfbfffb9643c1cbc8",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidRule)
			})
		})

	})
}

// go test -v -run TestValidateNonZero .

func TestValidateNonZero(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule NON_ZERO`, func() {
		type Article struct {
			// numeric

			Int8       int8
			Int16      int16
			Int32      int32
			Int64      int64
			Int        int
			Uint8      uint8
			Uint16     uint16
			Uint32     uint32
			Uint64     uint64
			Uint       uint
			Float32    float32
			Float64    float64
			Complex64  complex64
			Complex128 complex128

			// flag

			Bool bool

			// lengthable

			String string
			Array  [1]int
			Slice  []int
			Map    map[int]string
			Chan   chan int

			// struct

			Struct    time.Time
			Func      func() string
			Interface interface{}
		}

		type TestItem struct {
			Field string
			Val   Article
		}

		nonEmptyValues := []TestItem{
			{Field: "Int8", Val: Article{Int8: math.MinInt8}},
			{Field: "Int16", Val: Article{Int16: math.MinInt16}},
			{Field: "Int32", Val: Article{Int32: math.MinInt32}},
			{Field: "Int64", Val: Article{Int64: math.MinInt64}},
			{Field: "Int", Val: Article{Int: math.MinInt}},
			{Field: "Uint8", Val: Article{Uint8: uint8(1)}},
			{Field: "Uint16", Val: Article{Uint16: uint16(1)}},
			{Field: "Uint32", Val: Article{Uint32: uint32(1)}},
			{Field: "Uint64", Val: Article{Uint64: uint64(1)}},
			{Field: "Uint", Val: Article{Uint: uint(1)}},
			{Field: "Float32", Val: Article{Float32: float32(math.MaxFloat32)}},
			{Field: "Float64", Val: Article{Float64: float64(math.MaxFloat64)}},
			{Field: "Complex64", Val: Article{Complex64: complex64(1)}},
			{Field: "Complex128", Val: Article{Complex128: complex128(1)}},
			{Field: "Bool", Val: Article{Bool: true}},
			{Field: "String", Val: Article{String: "ok"}},
			{Field: "Array", Val: Article{Array: [1]int{100}}},
			{Field: "Slice", Val: Article{Slice: []int{100}}},
			{Field: "Map", Val: Article{Map: map[int]string{1: "ok"}}},
			{Field: "Chan", Val: Article{Chan: make(chan int)}},
			{Field: "Struct", Val: Article{Struct: time.Now()}},
			{Field: "Func", Val: Article{Func: func() string { return "ok" }}},
			{Field: "Interface", Val: Article{Interface: "ok"}},
		}

		for _, item := range nonEmptyValues {
			item := item // (!) save the context

			g.It("success if given a non-zero "+item.Field, func() {
				filter := Filter{
					{
						Field: item.Field,
						Check: NON_ZERO,
					},
				}

				hints := filter.Validate(item.Val)
				g.Assert(len(hints)).Equal(0, hints)
			})
		}

		// failure

		emptyValues := []TestItem{
			{Field: "Int8", Val: Article{Int8: *new(int8)}},
			{Field: "Int16", Val: Article{Int16: *new(int16)}},
			{Field: "Int32", Val: Article{Int32: *new(int32)}},
			{Field: "Int64", Val: Article{Int64: *new(int64)}},
			{Field: "Int", Val: Article{Int: *new(int)}},
			{Field: "Uint8", Val: Article{Uint8: *new(uint8)}},
			{Field: "Uint16", Val: Article{Uint16: *new(uint16)}},
			{Field: "Uint32", Val: Article{Uint32: *new(uint32)}},
			{Field: "Uint64", Val: Article{Uint64: *new(uint64)}},
			{Field: "Uint", Val: Article{Uint: *new(uint)}},
			{Field: "Float32", Val: Article{Float32: *new(float32)}},
			{Field: "Float64", Val: Article{Float64: *new(float64)}},
			{Field: "Complex64", Val: Article{Complex64: *new(complex64)}},
			{Field: "Complex128", Val: Article{Complex128: *new(complex128)}},
			{Field: "Bool", Val: Article{Bool: *new(bool)}},
			{Field: "String", Val: Article{String: *new(string)}},
			{Field: "Array", Val: Article{Array: *new([1]int)}},
			{Field: "Slice", Val: Article{Slice: *new([]int)}},
			{Field: "Map", Val: Article{Map: *new(map[int]string)}},
			{Field: "Chan", Val: Article{Chan: *new(chan int)}},
			{Field: "Struct", Val: Article{Struct: *new(time.Time)}},
			{Field: "Func", Val: Article{Func: *new(func() string)}},
			{Field: "Interface", Val: Article{Interface: *new(interface{})}},
		}

		for _, item := range emptyValues {
			item := item // (!) save the context

			g.It("failure if given a zero "+item.Field, func() {
				filter := Filter{
					{
						Field: item.Field,
						Check: NON_ZERO,
					},
				}

				hints := filter.Validate(item.Val)
				expectMsg := item.Field + " " + MsgEmpty

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(expectMsg)
			})
		}
	})
}
