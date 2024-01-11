# slim-validator
A simple validation package for Go

## Usage

```go
type Article struct {
  Id     uint     `json:"id"`
  Title  string   `json:"title"`
  Phone  string   `json:"phone"`
  Images []string `json:"images"`
}

filter := validator.Filter{
  {
    Field: "Id",
    Check: validator.NON_ZERO,
  },
  {
    Field: "Title",
    // title must contain up to 15 characters
    Check: validator.Rule{"max", 15},
  },
  {
    Field: "Phone",
    Check: validator.Rule{"match", `^\+38\d{10}$`},
  },
  {
    Field: "Images",
    Check: validator.Group{
      // images must contain at least 1 items
      validator.Rule{"min", 1},

      // images must contain up to 3 items
      validator.Rule{"max", 3},

      validator.Rule{"eachMatch", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
    },
  },
}

article := Article{
  Id:    12,
  Phone: "+380001234567",
  Images: []string{
    "https://img.it/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
    "https://img.it/fe14b5e8aa46475346e4375a6fd15df4.jpg",
  },
}

hints := filter.Validate(article)

for _, hint := range hints {
  fmt.Println(hint)
}
```

## Validation Rules
#### NON_ZERO

Checks the passed value for [non-zero](https://go.dev/ref/spec#The_zero_value) [ [1](https://pkg.go.dev/reflect#Value.IsZero) ] [ [2](https://golangbyexample.com/go-default-zero-value-all-types/) ].
The types that this rule works with are:
**bool**, **string**, **array**, **slice**, **map**, **chan**, **struct**, **func**, **interface**,
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**,
**float32**, **float64**, **complex64**, **complex128**

```go
{
  Field: "Id",
  Check: validator.NON_ZERO,
}
```

#### Match

Checks if the passed value matches the regular expression.
This rule only works with **string**

```go
{
  Field: "Phone",
  Check: validator.Rule{"match", `(?i)^[0-9a-f]{32}$`},
}
```

#### Each Match

Check whether any element of a collection matches a regular expression. This rule works with: **array**, **slice**, and **map** that contain string values

```go
{
  Field: "Images",
  Check: validator.Rule{"eachMatch", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
}
```

#### Min

Compares the compliance between the prototype and value, the value must correspond to the specified prototype within the minimum threshold. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
proto := 1

{
  Field: "Images",
  Check: validator.Rule{"min", proto},
}
```

When working with kinds of **array**, **slice**, and **map**, the minimum number (according to the specified prototype) of elements inside will check

```go
article := Article{
  Images: []string{},
}

filter := validator.Filter{
  {
    Field: "Images",
    Check: validator.Rule{"min", 1},
  },
}

// []string{"images must contain at least 1 items"}
hints := filter.Validate(article)
```

When working with **string** values, the minimum length (according to the specified prototype) of the string will be checked using the [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
Article{
  Title: string{"somebody to love"},
}

filter := validator.Filter{
  {
    Field: "Title",
    Check: validator.Rule{"min", 1},
  },
}
```

#### Max

Compares the compliance between the prototype and value, the value must correspond to the specified prototype within the maximum threshold. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
proto := 10

{
  Field: "Images",
  Check: validator.Rule{"max", proto},
}
```

When working with kinds of **array**, **slice**, and **map**, the maximum number (according to the specified prototype) of elements inside will check

```go
article := Article{
  Images: []string{"img1", "img2", "img3", "img4"},
}

filter := validator.Filter{
  {
    Field: "Images",
    // images must contain up to 3 items
    Check: validator.Rule{"max", 3},
  },
}
```

When working with **string** values, the maximum length (according to the specified prototype) of the string will be checked using the [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
Article{
  Title: string{"Somebody to love"},
}

filter := validator.Filter{
  {
    Field: "Title",
    // title must contain up to 15 characters
    Check: validator.Rule{"max", 15},
  },
}
```
