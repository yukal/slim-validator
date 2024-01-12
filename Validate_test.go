package validator

import (
	"math"
	"strings"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go clean -testcache
// go test -v -cover .
// go test -v -cover -run TestValidate .

// go test -v -run TestValidateMin .

func TestValidateMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "min"`, func() {
		type Article struct {
			Title   string         `json:"title"`
			Age     uint8          `json:"age"`
			Images  []string       `json:"images"`
			Phones  [4]string      `json:"phones"`
			Options map[int]string `json:"options"`
		}

		g.Describe("numeric", func() {
			g.It("success when the value exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 18},
					},
				}

				hints := filter.Validate(Article{Age: 21})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the value reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 21},
					},
				}

				hints := filter.Validate(Article{Age: 21})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the value is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"min", 18},
					},
				}

				hints := filter.Validate(Article{Age: 16})
				g.Assert(len(hints)).Equal(1)
				g.Assert(hints[0]).Equal("age must be at least 18")
			})
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 10},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 20},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"min", 30},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain at least 30 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 2},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 4},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"min", 8},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain at least 8 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 2},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 4},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"min", 8},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain at least 8 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

			g.It("success when the length exceeds the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 2},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 4},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than the min threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"min", 8},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain at least 8 items")
			})
		})

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"min", 0},
						},
					}

					hints := filter.Validate(Article{})
					g.Assert(len(hints)).Equal(0, hints)
				}
			})

			g.It("failure when missing rule value", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"min"},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})
		})
	})
}

// go test -v -run TestValidateMax .

func TestValidateMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "max"`, func() {
		type Article struct {
			Title   string         `json:"title"`
			Age     uint8          `json:"age"`
			Images  []string       `json:"images"`
			Phones  [4]string      `json:"phones"`
			Options map[int]string `json:"options"`
		}

		g.Describe("numeric", func() {
			g.It("success when the value is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 21},
					},
				}

				hints := filter.Validate(Article{Age: 18})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the value reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 21},
					},
				}

				hints := filter.Validate(Article{Age: 21})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the value exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"max", 12},
					},
				}

				hints := filter.Validate(Article{Age: 21})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be up to 12")
			})
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 30},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 20},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"max", 10},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain up to 10 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 8},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 4},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"max", 2},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain up to 2 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 8},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 4},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"max", 2},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain up to 2 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

			g.It("success when the length is less than the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 8},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("success when the length reaches the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 4},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length exceeds the max threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"max", 2},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain up to 2 items")
			})
		})

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"max", 0},
						},
					}

					hints := filter.Validate(Article{})
					g.Assert(len(hints)).Equal(0, hints)
				}
			})

			g.It("failure when missing rule proto", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"max"},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})
		})
	})
}

// go test -v -run TestValidateEq .

func TestValidateEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "eq" (equal)`, func() {
		type Article struct {
			Title   string         `json:"title"`
			Age     uint8          `json:"age"`
			Images  []string       `json:"images"`
			Phones  [4]string      `json:"phones"`
			Options map[int]string `json:"options"`
		}

		g.Describe("numeric", func() {
			g.It("success when the value equals a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 18},
					},
				}

				hints := filter.Validate(Article{Age: 18})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the value is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 21},
					},
				}

				hints := filter.Validate(Article{Age: 18})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be exactly 21")
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Rule{"eq", 18},
					},
				}

				hints := filter.Validate(Article{Age: 21})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be exactly 18")
			})
		})

		// ...

		g.Describe("string", func() {
			strFilled := "all you need is love"

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 20},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 30},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain exactly 30 characters")
			})

			g.It("failure when the length exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Rule{"eq", 10},
					},
				}

				hints := filter.Validate(Article{Title: strFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain exactly 10 characters")
			})
		})

		// ...

		g.Describe("array", func() {
			arrFilled := [4]string{"c", "o", "d", "e"}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 4},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 8},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Rule{"eq", 2},
					},
				}

				hints := filter.Validate(Article{Phones: arrFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain exactly 2 items")
			})
		})

		// ...

		g.Describe("slice", func() {
			sliceFilled := []string{"t", "e", "s", "t"}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 4},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 8},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Rule{"eq", 2},
					},
				}

				hints := filter.Validate(Article{Images: sliceFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain exactly 2 items")
			})
		})

		// ...

		g.Describe("map", func() {
			mapFilled := map[int]string{
				1: "We all live in a yellow submarine",
				2: "While My Guitar Gently Weeps",
				3: "All you need is love",
				4: "Let it be",
			}

			g.It("success when the length equals a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 4},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is less than a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 8},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain exactly 8 items")
			})

			g.It("failure when the value exceeds a threshold", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Rule{"eq", 2},
					},
				}

				hints := filter.Validate(Article{Options: mapFilled})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain exactly 2 items")
			})
		})

		// ...

		g.Describe(`emptiness`, func() {
			fieldsToCheck := []string{"Age", "Title", "Images", "Options"}
			article := Article{
				Age:    21,
				Title:  "All you need is love",
				Phones: [4]string{"0001234567"},
				Images: []string{"img1", "img2"},
				Options: map[int]string{
					1: "one",
					2: "two",
				},
			}

			g.It("success when given a zero proto and empty value", func() {
				for _, fieldName := range fieldsToCheck {
					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"eq", 0},
						},
					}

					hints := filter.Validate(Article{})
					g.Assert(len(hints)).Equal(0, hints)
				}
			})

			g.It("failure when missing rule value", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{"eq"},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})

			g.It("failure when given an empty rule", func() {
				for _, fieldName := range fieldsToCheck {
					expectMsg := strings.ToLower(fieldName) + " " + MsgInvalidRule

					filter := Filter{
						{
							Field: fieldName,
							Check: Rule{},
						},
					}

					hints := filter.Validate(article)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal(expectMsg)
				}
			})
		})
	})
}

// go test -v -run TestValidateRange .

func TestValidateRange(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "range"`, func() {
		type Article struct {
			Title   string            `json:"title"`
			Age     uint8             `json:"age"`
			Images  []string          `json:"images"`
			Phones  [4]string         `json:"phones"`
			Options map[string]string `json:"options"`
		}

		g.Describe(`numeric`, func() {
			g.It("success when the value matches the range", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				hints := filter.Validate(Article{Age: 18})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when given below-range value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				hints := filter.Validate(Article{Age: 16})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be in the range 18..21")
			})

			g.It("failure when given above-range value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				hints := filter.Validate(Article{Age: 31})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be in the range 18..21")
			})
		})

		// ...

		g.Describe(`string`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{4, 20},
					},
				}

				hints := filter.Validate(Article{Title: "all you need is love"})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{25, 45},
					},
				}

				hints := filter.Validate(Article{Title: "all you need is love"})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain 25..45 characters")
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Title",
						Check: Range{3, 18},
					},
				}

				hints := filter.Validate(Article{Title: "all you need is love"})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("title must contain 3..18 characters")
			})
		})

		// ...

		g.Describe(`array`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Range{1, 4},
					},
				}

				hints := filter.Validate(Article{
					Phones: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Range{10, 80},
					},
				}

				hints := filter.Validate(Article{
					Phones: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Phones",
						Check: Range{1, 3},
					},
				}

				hints := filter.Validate(Article{
					Phones: [4]string{"t", "e", "s", "t"},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("phones must contain 1..3 items")
			})
		})

		// ...

		g.Describe(`slice`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{1, 4},
					},
				}

				hints := filter.Validate(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{10, 80},
					},
				}

				hints := filter.Validate(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Images",
						Check: Range{1, 3},
					},
				}

				hints := filter.Validate(Article{
					Images: []string{"jpeg", "jpg", "png", "gif"},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("images must contain 1..3 items")
			})
		})

		// ...

		g.Describe(`map`, func() {
			g.It("success when the length matches the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{1, 4},
					},
				}

				hints := filter.Validate(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when the length is below the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{10, 80},
					},
				}

				hints := filter.Validate(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain 10..80 items")
			})

			g.It("failure when the length is above the range", func() {
				filter := Filter{
					{
						Field: "Options",
						Check: Range{1, 3},
					},
				}

				hints := filter.Validate(Article{
					Options: map[string]string{
						"jpeg": "image/jpeg",
						"jpg":  "image/jpeg",
						"png":  "image/png",
						"gif":  "image/gif",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("options must contain 1..3 items")
			})
		})

		// ...

		g.Describe("emptiness", func() {
			g.It("failure if at least 1 element of the range is empty", func() {
				protos := []Range{
					{},
					{nil, 15},
					{15, nil},
				}

				for _, item := range protos {
					filter := Filter{
						{
							Field: "Age",
							Check: Range{item[0], item[1]},
						},
					}

					hints := filter.Validate(Article{Age: 31})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("age " + MsgInvalidRule)
				}
			})

			g.It("failure when given an empty value", func() {
				filter := Filter{
					{
						Field: "Age",
						Check: Range{18, 21},
					},
				}

				hints := filter.Validate(Article{})
				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("age must be in the range 18..21")
			})
		})
	})
}

// go test -v -run TestValidateYear .

func TestValidateYear(t *testing.T) {
	type Article struct {
		Date time.Time `json:"date"`
	}

	g := Goblin(t)

	g.Describe(`Rule "year"`, func() {
		g.It("success when the value matches a specific year", func() {
			tm, err := time.Parse(time.RFC3339, "2024-12-25T16:04:05Z")
			g.Assert(err).IsNil(err)

			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
				},
			}

			hints := filter.Validate(Article{Date: tm})
			g.Assert(len(hints)).Equal(0, hints)
		})

		g.It("failure when the value is not match", func() {
			filter := Filter{
				{
					Field: "Date",
					Check: Rule{"year", 2024},
				},
			}

			hints := filter.Validate(Article{
				Date: *new(time.Time),
			})

			g.Assert(len(hints)).Equal(1, hints)
			g.Assert(hints[0]).Equal("date must be exactly 2024")
		})
	})
}

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

// go test -v -run TestValidateEachMatchDeprecated .

func TestValidateEachMatchDeprecated(t *testing.T) {
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

// go test -v -run TestValidateEachMin .

func TestValidateEachMin(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:min"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int            `json:"pages"`
				Bands   [2]string         `json:"bands"`
				Artists [2][2]string      `json:"artists"`
				Songs   [2][]string       `json:"songs"`
				Albums  [2]map[int]string `json:"albums"`
			}

			type Emptyness struct {
				Pages   [0]int            `json:"pages"`
				Bands   [0]string         `json:"bands"`
				Artists [0][0]string      `json:"artists"`
				Songs   [0][]string       `json:"songs"`
				Albums  [0]map[int]string `json:"albums"`
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 15},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{5, 15},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be at least 10")
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 7},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 9},
						},
					}

					hints := filter.Validate(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 12},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain at least 12 characters")
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain at least 4 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int            `json:"pages"`
				Bands   []string         `json:"bands"`
				Artists [][2]string      `json:"artists"`
				Artist  [][0]string      `json:"artist"`
				Songs   [][]string       `json:"songs"`
				Albums  []map[int]string `json:"albums"`
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 15},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{5, 15},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be at least 10")
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 7},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 12},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain at least 12 characters")
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Slice{
						Artist: [][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain at least 4 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int            `json:"pages"`
				Bands   map[int]string            `json:"bands"`
				Artists map[string][2]string      `json:"artists"`
				Artist  map[string][0]string      `json:"artist"`
				Songs   map[string][]string       `json:"songs"`
				Albums  map[string]map[int]string `json:"albums"`
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 15},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:min", 10},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[Title] must be at least 10")
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 7},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:min", 12},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[2] must contain at least 12 characters")
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[Nirvana] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
								"High Hopes",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
								"Heart-Shaped Box",
								"Come As You Are",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Metallica] must contain at least 4 items")
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length exceeds the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 1},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 2},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 3},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length is less than the min threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:min", 4},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
								1975: "Toys in the Attic",
								1976: "Rocks",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Scorpions] must contain at least 4 items")
				})
			})
		})

	})
}

// go test -v -run TestValidateEachMax .

func TestValidateEachMax(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:max"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int            `json:"pages"`
				Bands   [2]string         `json:"bands"`
				Artists [2][2]string      `json:"artists"`
				Songs   [2][]string       `json:"songs"`
				Albums  [2]map[int]string `json:"albums"`
			}

			type Emptyness struct {
				Pages   [0]int            `json:"pages"`
				Bands   [0]string         `json:"bands"`
				Artists [0][0]string      `json:"artists"`
				Songs   [0][]string       `json:"songs"`
				Albums  [0]map[int]string `json:"albums"`
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 30},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 25},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 15},
						},
					}

					hints := filter.Validate(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{5, 15},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[1] must be up to 10")
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 18},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain up to 10 characters")
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
								"Lady Madonna",
								"Let It Be",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
								1975: "Toys in the Attic",
								1976: "Rocks",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain up to 1 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int            `json:"pages"`
				Bands   []string         `json:"bands"`
				Artists [][2]string      `json:"artists"`
				Artist  [][0]string      `json:"artist"`
				Songs   [][]string       `json:"songs"`
				Albums  []map[int]string `json:"albums"`
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 30},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 25},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{15, 25},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 15},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{5, 15},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[1] must be up to 10")
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 18},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain up to 10 characters")
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Artist: [][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Here in My Heart",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"Lithium",
								"Smells Like Teen Spirit",
								"Heart-Shaped Box",
								"Come As You Are",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
								1975: "Toys in the Attic",
								1976: "Rocks",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain up to 1 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int            `json:"pages"`
				Bands   map[int]string            `json:"bands"`
				Artists map[string][2]string      `json:"artists"`
				Artist  map[string][0]string      `json:"artist"`
				Songs   map[string][]string       `json:"songs"`
				Albums  map[string]map[int]string `json:"albums"`
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 30},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element value reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 25},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 25,
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 15},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 value exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[Prologue] must be up to 10")
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 18},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:max", 10},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain up to 10 characters")
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[Nirvana] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
								"High Hopes",
							},
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
								"Heart-Shaped Box",
								"Come As You Are",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Nirvana] must contain up to 1 items")
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length is less than the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 3},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when the element length reaches the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1996: "Pure Instinct",
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 2},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when at least 1 element length exceeds the max threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:max", 1},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
								1975: "Toys in the Attic",
								1976: "Rocks",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Aerosmith] must contain up to 1 items")
				})
			})
		})

	})
}

// go test -v -run TestValidateEachEq .

func TestValidateEachEq(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:eq" (equal)`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int            `json:"pages"`
				Bands   [2]string         `json:"bands"`
				Artists [2][2]string      `json:"artists"`
				Songs   [2][]string       `json:"songs"`
				Albums  [2]map[int]string `json:"albums"`
			}

			type Emptyness struct {
				Pages   [0]int            `json:"pages"`
				Bands   [0]string         `json:"bands"`
				Artists [0][0]string      `json:"artists"`
				Songs   [0][]string       `json:"songs"`
				Albums  [0]map[int]string `json:"albums"`
			}

			g.Describe("array:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{15, 15},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Emptyness{
						Pages: [0]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{5, 10},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be exactly 15")
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{25, 30},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be exactly 15")
				})
			})

			// ...

			g.Describe("array:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Emptyness{
						Bands: [0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain exactly 9 characters")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Array{
						Bands: [2]string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain exactly 10 characters")
				})
			})

			// ...

			g.Describe("array:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Emptyness{
						Artists: [0][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain exactly 4 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Array{
						Artists: [2][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("array:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Emptyness{
						Songs: [0][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[1] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Array{
						Songs: [2][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("array:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Emptyness{
						Albums: [0]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[1] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Array{
						Albums: [2]map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain exactly 1 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int            `json:"pages"`
				Bands   []string         `json:"bands"`
				Artists [][2]string      `json:"artists"`
				Artist  [][0]string      `json:"artist"`
				Songs   [][]string       `json:"songs"`
				Albums  []map[int]string `json:"albums"`
			}

			g.Describe("slice:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{15, 15},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{5, 10},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be exactly 15")
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Slice{
						Pages: []int{25, 30},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be exactly 15")
				})
			})

			// ...

			g.Describe("slice:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Aerosmith",
							"Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Metallica",
							"Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain exactly 9 characters")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Slice{
						Bands: []string{
							"Led Zeppelin",
							"Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain exactly 10 characters")
				})
			})

			// ...

			g.Describe("slice:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							{
								"Thomas William Hamilton",
								"Joseph Michael Kramer",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Slice{
						Artist: [][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain exactly 4 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][2]string{
							{
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
							{
								"Krist Anthony Novoselic",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("slice:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
								"Send Me an Angel",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[1] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{
							{
								"We all live in a yellow submarine",
								"All you need is love",
							},
							{
								"No pain no gain",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("slice:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[1] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{
							{
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							{
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain exactly 1 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int            `json:"pages"`
				Bands   map[int]string            `json:"bands"`
				Artists map[string][2]string      `json:"artists"`
				Artist  map[string][0]string      `json:"artist"`
				Songs   map[string][]string       `json:"songs"`
				Albums  map[string]map[int]string `json:"albums"`
			}

			g.Describe("map:numeric", func() {
				g.It("success when the element value equals a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 15,
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the value is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    5,
							"Prologue": 15,
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[Title] must be exactly 15")
				})

				g.It("failure when the value exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:eq", 15},
						},
					}

					hints := filter.Validate(Map{
						Pages: map[string]int{
							"Title":    15,
							"Prologue": 30,
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[Prologue] must be exactly 15")
				})
			})

			// ...

			g.Describe("map:string", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Aerosmith",
							2: "Scorpions",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 9},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Metallica",
							2: "Nirvana",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[2] must contain exactly 9 characters")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Bands",
							Check: Rule{"each:eq", 10},
						},
					}

					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Led Zeppelin",
							2: "Pink Floyd",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain exactly 10 characters")
				})
			})

			// ...

			g.Describe("map:array", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Aerosmith": {
								"Steven Victor Tallarico",
								"Anthony Joseph Perry",
							},
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artist",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Map{
						Artist: map[string][0]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Scorpions": {
								"Klaus Meine",
								"Rudolf Schenker",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[Scorpions] must contain exactly 4 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][2]string{
							"Nirvana": {
								"Kurt Donald Cobain",
								"David Eric Grohl",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[Nirvana] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("map:slice", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Aerosmith": {
								"Walk This Way",
								"Dream On",
							},
							"Scorpions": {
								"Rock You Like a Hurricane",
								"Still Loving You",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Led Zeppelin": {
								"Kashmir",
								"Stairway To Heaven",
							},
							"Pink Floyd": {
								"Wish You Were Here",
							},
						},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Pink Floyd] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{
							"Nirvana": {
								"Lithium",
								"Smells Like Teen Spirit",
							},
							"Metallica": {
								"One",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Nirvana] must contain exactly 1 items")
				})
			})

			// ...

			g.Describe("map:map", func() {
				g.It("success when the element length equals a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
								2007: "Humanity",
							},
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 4},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is less than a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 2},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Scorpions] must contain exactly 2 items")
				})

				g.It("failure when the element length exceeds a threshold", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:eq", 1},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{
							"Aerosmith": {
								1973: "Aerosmith",
								1974: "Get Your Wings",
							},
							"Scorpions": {
								1999: "Eye II Eye",
							},
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Aerosmith] must contain exactly 1 items")
				})
			})
		})
	})
}

// go test -v -run TestValidateEachRange .

func TestValidateEachRange(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:range"`, func() {
		g.Describe(`array`, func() {
			type Array struct {
				Pages   [2]int            `json:"pages"`
				Bands   [2]string         `json:"bands"`
				Artists [1][4]string      `json:"artists"`
				Songs   [1][]string       `json:"songs"`
				Albums  [1]map[int]string `json:"albums"`
			}

			g.Describe(`array:numeric`, func() {
				filter := Filter{
					{
						Field: "Pages",
						Check: Rule{"each:range", Range{35, 45}},
					},
				}

				g.It("success when the element value matches the range", func() {
					hints := filter.Validate(Array{
						Pages: [2]int{35, 45},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list within the range", func() {
					filter := Filter{
						{
							Field: "Pages",
							Check: Rule{"each:range", Range{0, 1}},
						},
					}

					hints := filter.Validate(Array{
						Pages: [2]int{},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when given an empty data list that not match the range", func() {
					hints := filter.Validate(Array{
						Pages: [2]int{},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be in the range 35..45")
				})

				g.It("failure when the element value is below the range", func() {
					hints := filter.Validate(Array{
						Pages: [2]int{15, 25},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be in the range 35..45")
				})

				g.It("failure when the element value is above the range", func() {
					hints := filter.Validate(Array{
						Pages: [2]int{135, 145},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be in the range 35..45")
				})
			})

			// ...

			g.Describe(`array:string`, func() {
				filter := Filter{
					{
						Field: "Bands",
						Check: Rule{"each:range", Range{9, 11}},
					},
				}

				g.It("success when the element length matches the range", func() {
					hints := filter.Validate(Array{
						Bands: [2]string{
							"The Beatles",
							"Aerosmith",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when given an empty data list", func() {
					hints := filter.Validate(Array{
						Bands: [2]string{},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain 9..11 characters")
				})

				g.It("failure when the element length is below the range", func() {
					hints := filter.Validate(Array{
						Bands: [2]string{
							"Queen",
							"AC/DC",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain 9..11 characters")
				})

				g.It("failure when the element length is above the range", func() {
					hints := filter.Validate(Array{
						Bands: [2]string{
							"Rolling Stones",
							"Led Zeppelin",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain 9..11 characters")
				})
			})

			// ...

			g.Describe(`array:array`, func() {
				value := Array{
					Artists: [1][4]string{{
						"Kurt Donald Cobain",
						"David Eric Grohl",
						"Angus Young",
						"Steven Victor Tallarico",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`array:slice`, func() {
				value := Array{
					Songs: [1][]string{{
						"We all live in a yellow submarine",
						"All you need is love",
						"No pain no gain",
						"Send Me an Angel",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Array{
						Songs: [1][]string{},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain 2..4 items")
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`array:map`, func() {
				value := Array{
					Albums: [1]map[int]string{{
						1973: "Aerosmith",
						1974: "Get Your Wings",
						1999: "Eye II Eye",
						2007: "Humanity",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Array{
						Albums: [1]map[int]string{},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain 2..4 items")
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain 1..2 items")
				})
			})
		})

		// ...

		g.Describe(`slice`, func() {
			type Slice struct {
				Pages   []int            `json:"pages"`
				Bands   []string         `json:"bands"`
				Artists [][4]string      `json:"artists"`
				Artist  [][4]string      `json:"artist"`
				Songs   [][]string       `json:"songs"`
				Albums  []map[int]string `json:"albums"`
			}

			g.Describe(`slice:numeric`, func() {
				filter := Filter{
					{
						Field: "Pages",
						Check: Rule{"each:range", Range{35, 45}},
					},
				}

				g.It("success when the element value matches the range", func() {
					hints := filter.Validate(Slice{
						Pages: []int{35, 45},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					hints := filter.Validate(Slice{
						Pages: []int{},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element value is below the range", func() {
					hints := filter.Validate(Slice{
						Pages: []int{15, 25},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be in the range 35..45")
				})

				g.It("failure when the element value is above the range", func() {
					hints := filter.Validate(Slice{
						Pages: []int{135, 145},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[0] must be in the range 35..45")
				})
			})

			// ...

			g.Describe(`slice:string`, func() {
				filter := Filter{
					{
						Field: "Bands",
						Check: Rule{"each:range", Range{9, 11}},
					},
				}

				g.It("success when the element length matches the range", func() {
					hints := filter.Validate(Slice{
						Bands: []string{
							"The Beatles",
							"Aerosmith",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					hints := filter.Validate(Slice{
						Bands: []string{},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					hints := filter.Validate(Slice{
						Bands: []string{
							"Queen",
							"AC/DC",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain 9..11 characters")
				})

				g.It("failure when the element length is above the range", func() {
					hints := filter.Validate(Slice{
						Bands: []string{
							"Rolling Stones",
							"Led Zeppelin",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[0] must contain 9..11 characters")
				})
			})

			// ...

			g.Describe(`slice:array`, func() {
				value := Slice{
					Artists: [][4]string{{
						"Kurt Donald Cobain",
						"David Eric Grohl",
						"Angus Young",
						"Steven Victor Tallarico",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Slice{
						Artists: [][4]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[0] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`slice:slice`, func() {
				value := Slice{
					Songs: [][]string{{
						"We all live in a yellow submarine",
						"All you need is love",
						"No pain no gain",
						"Send Me an Angel",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Slice{
						Songs: [][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[0] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`slice:map`, func() {
				value := Slice{
					Albums: []map[int]string{{
						1973: "Aerosmith",
						1974: "Get Your Wings",
						1999: "Eye II Eye",
						2007: "Humanity",
					}},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Slice{
						Albums: []map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[0] must contain 1..2 items")
				})
			})
		})

		// ...

		g.Describe(`map`, func() {
			type Map struct {
				Pages   map[string]int            `json:"pages"`
				Bands   map[int]string            `json:"bands"`
				Artists map[string][4]string      `json:"artists"`
				Artist  map[string][4]string      `json:"artist"`
				Songs   map[string][]string       `json:"songs"`
				Albums  map[string]map[int]string `json:"albums"`
			}

			g.Describe(`map:numeric`, func() {
				filter := Filter{
					{
						Field: "Pages",
						Check: Rule{"each:range", Range{35, 45}},
					},
				}

				g.It("success when the element value matches the range", func() {
					hints := filter.Validate(Map{
						Pages: map[string]int{
							"item1": 35,
							"item2": 45,
						},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					hints := filter.Validate(Map{
						Pages: map[string]int{},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element value is below the range", func() {
					hints := filter.Validate(Map{
						Pages: map[string]int{
							"item1": 15,
							"item2": 35,
						},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[item1] must be in the range 35..45")
				})

				g.It("failure when the element value is above the range", func() {
					hints := filter.Validate(Map{
						Pages: map[string]int{
							"item1": 45,
							"item2": 55,
						},
					})
					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("pages item[item2] must be in the range 35..45")
				})
			})

			// ...

			g.Describe(`map:string`, func() {
				filter := Filter{
					{
						Field: "Bands",
						Check: Rule{"each:range", Range{9, 11}},
					},
				}

				g.It("success when the element length matches the range", func() {
					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "The Beatles",
							2: "Aerosmith",
						},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					hints := filter.Validate(Map{
						Bands: map[int]string{},
					})
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Queen",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain 9..11 characters")
				})

				g.It("failure when the element length is above the range", func() {
					hints := filter.Validate(Map{
						Bands: map[int]string{
							1: "Rolling Stones",
						},
					})

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("bands item[1] must contain 9..11 characters")
				})
			})

			// ...

			g.Describe(`map:array`, func() {
				value := Map{
					Artists: map[string][4]string{
						"signers": {
							"Kurt Donald Cobain",
							"David Eric Grohl",
							"Angus Young",
							"Steven Victor Tallarico",
						},
					},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Map{
						Artists: map[string][4]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[signers] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Artists",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("artists item[signers] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`map:slice`, func() {
				value := Map{
					Songs: map[string][]string{
						"Scorpions": {
							"Rock You Like a Hurricane",
							"When the smoke is going down",
							"Still Loving You",
							"Here in my heart",
						},
					},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Map{
						Songs: map[string][]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Scorpions] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Songs",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("songs item[Scorpions] must contain 1..2 items")
				})
			})

			// ...

			g.Describe(`map:map`, func() {
				value := Map{
					Albums: map[string]map[int]string{
						"Scorpions": {
							1975: "In Trance",
							1999: "Eye II Eye",
							2000: "Moment of Glory",
							2007: "Humanity",
						},
					},
				}

				g.It("success when the element length matches the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(value)
					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("success when given an empty data list", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{2, 4}},
						},
					}

					hints := filter.Validate(Map{
						Albums: map[string]map[int]string{},
					})

					g.Assert(len(hints)).Equal(0, hints)
				})

				g.It("failure when the element length is below the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{5, 10}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Scorpions] must contain 5..10 items")
				})

				g.It("failure when the element length is above the range", func() {
					filter := Filter{
						{
							Field: "Albums",
							Check: Rule{"each:range", Range{1, 2}},
						},
					}

					hints := filter.Validate(value)

					g.Assert(len(hints)).Equal(1, hints)
					g.Assert(hints[0]).Equal("albums item[Scorpions] must contain 1..2 items")
				})
			})
		})

		// ...

		g.Describe("invalidity", func() {
			type Slice struct {
				Bands []string `json:"bands"`
			}

			g.It("success when given an empty value", func() {
				filter := Filter{
					{
						Field: "Bands",
						Check: Rule{"each:range", Range{9, 11}},
					},
				}

				hints := filter.Validate(Slice{})
				g.Assert(len(hints)).Equal(0, hints)
			})

			g.It("failure when given an invalid rule", func() {
				filter := Filter{
					{
						Field: "Bands",
						Check: Rule{"each:range"},
					},
				}

				hints := filter.Validate(Slice{
					Bands: []string{
						"The Beatles",
						"Aerosmith",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal("bands has invalid rule")
			})
		})
	})
}

// go test -v -run TestValidateEachMatch .

func TestValidateEachMatch(t *testing.T) {
	g := Goblin(t)

	g.Describe(`Rule "each:match"`, func() {
		msgInvalidRule := "hash " + MsgInvalidRule
		msgInvalidValue := "hash item[1] " + MsgNotValid

		g.Describe(`array`, func() {
			type Array struct {
				Hash  [2]string `json:"hash"`
				Empty [0]string `json:"empty"`
			}

			g.It("success when the element value matches the mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match"},
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
						Check: Rule{"each:match", ``},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match"},
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
						Check: Rule{"each:match", ``},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
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
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "Z0zZ0z19711zZz3z73z41z909z66zZzZ",
						2: "b0fb0c19711bcf3b73f41c909f66bded",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when at least 1 value is empty", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"each:match", `(?i)^[0-9a-f]{32}$`},
					},
				}

				hints := filter.Validate(Map{
					Hash: map[int]string{
						1: "",
						2: "b0fb0c19711bcf3b73f41c909f66bded",
					},
				})

				g.Assert(len(hints)).Equal(1, hints)
				g.Assert(hints[0]).Equal(msgInvalidValue)
			})

			g.It("failure when missing mask", func() {
				filter := Filter{
					{
						Field: "Hash",
						Check: Rule{"each:match"},
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
						Check: Rule{"each:match", ``},
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
