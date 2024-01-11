package validator

import (
	"math"
	"strings"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

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
