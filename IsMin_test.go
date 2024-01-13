package validator

import (
	"fmt"
	"math"
	"testing"

	. "github.com/franela/goblin"
)

// go test .
// go test -v -run TestIsMin .

func TestIsMin(t *testing.T) {
	const (
		MIN_INT8   = int8(math.MinInt8)
		MIN_INT16  = int16(math.MinInt16)
		MIN_INT32  = int32(math.MinInt32)
		MIN_INT64  = int64(math.MinInt64)
		MIN_INT    = int(math.MinInt)
		MAX_INT8   = int8(math.MaxInt8)
		MAX_INT16  = int16(math.MaxInt16)
		MAX_INT32  = int32(math.MaxInt32)
		MAX_INT64  = int64(math.MaxInt64)
		MAX_INT    = int(math.MaxInt)
		MIN_UINT8  = uint8(0)
		MIN_UINT16 = uint16(0)
		MIN_UINT32 = uint32(0)
		MIN_UINT64 = uint64(0)
		MIN_UINT   = uint(0)
		MAX_UINT8  = uint8(math.MaxUint8)
		MAX_UINT16 = uint16(math.MaxUint16)
		MAX_UINT32 = uint32(math.MaxUint32)
		MAX_UINT64 = uint64(math.MaxUint64)
		MAX_UINT   = uint(math.MaxUint)
	)

	g := Goblin(t)

	// ................................
	// int & int

	g.Describe("int & int", func() {
		g.Describe("int8", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_INT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_INT8, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_INT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_INT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_INT32, false},
				// {"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_RUNE, true},
				// {"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_RUNE, false},
				// {"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_RUNE, false},
				// {"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_RUNE, true},
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int16", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_INT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_INT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_INT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_INT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int32", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_INT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_INT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_INT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_INT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int64", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_INT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_INT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_INT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_INT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_INT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_INT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_INT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_INT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_INT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_INT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_INT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_INT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// int & uint

	g.Describe("int & uint", func() {

		g.Describe("int8", func() {
			items := [][4]any{
				// {"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_BYTE, true},
				// {"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_BYTE, true},
				// {"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_BYTE, true},
				// {"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_BYTE, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_UINT8, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_UINT8, false},

				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_UINT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_UINT16, false},

				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_UINT32, false},

				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_UINT64, false},

				{"min-%[1]T >= min-%[2]T", MIN_INT8, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT8, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT8, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT8, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int16", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_UINT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT16, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT16, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT16, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT16, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int32", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT32, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT32, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT32, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT32, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int64", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_UINT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT64, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT64, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT64, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT64, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("int", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_UINT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_INT, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_INT, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_INT, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_INT, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// uint & uint

	g.Describe("uint & uint", func() {

		g.Describe("uint8", func() {
			items := [][4]any{
				// {"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_BYTE, true},
				// {"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_BYTE, true},
				// {"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_BYTE, true},
				// {"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_BYTE, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_UINT8, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_UINT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint16", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_UINT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint32", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_UINT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint64", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_UINT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_UINT8, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_UINT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_UINT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_UINT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_UINT16, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_UINT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_UINT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_UINT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_UINT32, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_UINT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_UINT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_UINT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_UINT64, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_UINT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_UINT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_UINT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_UINT, true},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_UINT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_UINT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_UINT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})

	// ................................
	// uint & int

	g.Describe("uint & int", func() {

		g.Describe("uint8", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_INT8, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_INT16, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_INT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT8, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT8, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT8, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT8, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				if item[3] == true {
					title = "return %[3]t  if " + title
				} else {
					title = "return %[3]t if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint16", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_INT8, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_INT32, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT16, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT16, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT16, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT16, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint32", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_INT8, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_INT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_INT64, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT32, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT32, MAX_INT, true},
				{"max-%[1]T >= min-%[2]T", MIN_UINT32, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT32, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint64", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_INT8, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_INT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_INT64, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT64, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT64, MAX_INT, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT64, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT64, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

		g.Describe("uint", func() {
			items := [][4]any{
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_INT8, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_INT8, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_INT8, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_INT8, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_INT16, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_INT16, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_INT16, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_INT16, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_INT32, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_INT32, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_INT32, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_INT32, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_INT64, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_INT64, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_INT64, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_INT64, false},
				{"min-%[1]T >= min-%[2]T", MIN_UINT, MIN_INT, false},
				{"max-%[1]T >= max-%[2]T", MAX_UINT, MAX_INT, false},
				{"max-%[1]T >= min-%[2]T", MIN_UINT, MAX_INT, true},
				{"min-%[1]T >= max-%[2]T", MAX_UINT, MIN_INT, false},
			}

			for _, item := range items {
				item := item // (!) Dont remove, it saves the context
				title := item[0].(string)

				switch item[3] {
				case false:
					title = "return %[3]t if " + title
				case true:
					title = "return %[3]t  if " + title
				}

				g.It(fmt.Sprintf(title, item[1], item[2], item[3]), func() {
					result := IsMin(item[1], item[2])
					g.Assert(result).Equal(item[3], fmt.Sprintf("Expect(%#v) Got(%#v)", item[3], result))
				})
			}
		})

	})
}
