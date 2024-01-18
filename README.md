# slim-validator
A simple validation package for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/yukal/slim-validator.svg)](https://pkg.go.dev/github.com/yukal/slim-validator)

## Install
```bash
# get latest version
go get github.com/yukal/slim-validator

# get a specific version
go get github.com/yukal/slim-validator@v0.12.0
```

Import then:

```bash
# import using alias
import validator "github.com/yukal/slim-validator"

# import using parentheses
import (
  validator "github.com/yukal/slim-validator"
)

# import using a context
# In that case, you will be able to use the validator directly without any aliasing
import (
  . "github.com/yukal/slim-validator"
)
```

## Usage

The two methods you will need for validation are the `IsValid()` and the `Validate()` after creating a validation filter. See a detailed example below:

```go
type Article struct {
  Id     uint      `json:"id"`
  Sex    uint8     `json:"sex"`
  Title  string    `json:"title"`
  Phone  string    `json:"phone"`
  Images []string  `json:"images"`
  Date   time.Time `json:"date"`
}

filter := validator.Filter{
  {
    Field: "Id",
    // id must be at least 1
    Check: validator.Rule{"min", 1},
  },
  {
    Field: "Sex",
    // sex must be in the range 1..2
    Check: validator.Range{1, 2},
  },
  {
    Field: "Title",
    // title must contain up to 64 characters
    Check: validator.Rule{"max", 64},
  },
  {
    Field: "Phone",

    // phone must match the mask of the regular expression
    Check: validator.Rule{"match", `^\+38\d{10}$`},

    // will not validate if the value is not passed
    Optional: true,
  },
  {
    Field: "Images",

    // images must match the mask of the regular expression
    Check: validator.Rule{"each:match", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
  },
  {
    Field: "Date",

    // date must exactly contain the 2024 year
    Check: validator.Rule{"year", 2024},
  },
}

article := Article{
  Id:    12,
  Sex:   1,
  Title: "We all live in a yellow submarine",
  Phone: "+380001234567",
  Images: []string{
    "https://img.it/5e8aa4647a6fd1545346e4375fedf14b.jpeg",
    "https://img.it/fe14b5e8aa46475346e4375a6fd15df4.jpg",
  },
  Date: time.Now(),
}

if !filter.IsValid(article) {
  fmt.Println("article is not valid!")
}

hints := filter.Validate(article)

for _, hint := range hints {
  fmt.Println(hint)
}
```

## Validation Rules
### NON_ZERO

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

### Match

Checks if the passed value matches the regular expression.
This rule only works with **string**

```go
{
  Field: "Phone",
  Check: validator.Rule{"match", `(?i)^[0-9a-f]{32}$`},
}
```

### Min

Compares the compliance between the prototype and value, the value must correspond to the specified prototype within the minimum threshold. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
proto := 1

// id must be at least 1
{
  Field: "Id",
  Check: validator.Rule{"min", proto},
}
```

When working with kinds of **array**, **slice**, and **map**, the minimum number (according to the specified prototype) of elements inside will check

```go
// images must contain at least 1 items
{
  Field: "Images",
  Check: validator.Rule{"min", 1},
}
```

When working with **string** values, the minimum length (according to the specified prototype) of the string will be checked using the [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
// title must contain at least 16 characters
{
  Field: "Title",
  Check: validator.Rule{"min", 16},
}
```

### Max

Compares the compliance between the prototype and value, the value must correspond to the specified prototype within the maximum threshold. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
proto := 255

// id must be up to 255
{
  Field: "Id",
  Check: validator.Rule{"max", proto},
}
```

When working with kinds of **array**, **slice**, and **map**, the maximum number (according to the specified prototype) of elements inside will check

```go
// images must contain up to 3 items
{
  Field: "Images",
  Check: validator.Rule{"max", 3},
}
```

When working with **string** values, the maximum length (according to the specified prototype) of the string will be checked using the [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
// title must contain up to 15 characters
{
  Field: "Title",
  Check: validator.Rule{"max", 15},
}
```

### Equal

Compares the compliance between the prototype and value, the value must exactly equal the specified prototype. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
// sex must be exactly 1
{
  Field: "Sex",
  Check: validator.Rule{"eq", 1},
}
```

When working with kinds of **array**, **slice**, and **map**, the validator will check if the capacity is equal to the specified number

```go
// images must contain exactly 2 items
{
  Field: "Images",
  Check: validator.Rule{"eq", 2},
}
```

When working with **string** values, the validator will check if the length is equal to the specified number. It uses [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
// title must contain exactly 15 characters
{
  Field: "Title",
  Check: validator.Rule{"eq", 15},
}
```

### Range

Compares the compliance between the prototype and the value, the value must match the specified range between the minimum and maximum threshold. The types that this rule works with are:
**int8**, **int16**, **int32**, **int64**, **int**, **uint8**, **uint16**, **uint32**, **uint64**, **uint**, **string**, **array**, **slice**, **map**

```go
// sex must be in the range 1..2
{
  Field: "Sex",
  Check: validator.Range{1, 2},
}
```

When working with kinds of **array**, **slice**, and **map**, the validator will check whether the collection capacity matches the specified range

```go
// images must contain 5..8 items
{
  Field: "Images",
  Check: validator.Range{5, 8},
}
```

When working with **string** values, the validator will check whether the length matches the specified range. It uses [utf8.RuneCountInString](https://pkg.go.dev/unicode/utf8#RuneCountInString)

```go
// title must contain 1..200 characters
{
  Field: "Title",
  Check: validator.Range{1, 200},
}
```

### Year

Compares the compliance between the prototype and the value, the value must match the specified year. This rule only works with **[time.Time](https://pkg.go.dev/time)**

```go
// date must be exactly 2024
{
  Field: "Date",
  Check: validator.Rule{"year", 2024},
}
```

## Validation Modifiers

The modifiers works in conjunction with the validation rules mentioned above in this document

### Each

The "each" modifier checks the correspondence between the prototype and each element of the **array**, **slice**, or **map**. Thus, each element will be checked according to the defined threshold rule

```go
validator.Rule{"each:min", 16},
validator.Rule{"each:max", 128},
validator.Rule{"each:eq", 0},
validator.Rule{"each:range", []int{5, 10}},
validator.Rule{"each:match", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
```

### Date

The "date" modifier checks the correspondence between the prototype and the struct value with type [time.Time](https://pkg.go.dev/time#Time). In the context of this validator, the "date" modifier was intended to work with simple time values, without comparing milli, micro, and nanoseconds. A prototype can be specified in [RFC3339](https://pkg.go.dev/time#pkg-constants) string, [int64](https://pkg.go.dev/time#example-Unix), and [time](https://pkg.go.dev/time).

```go
// passing date as int64
// following the rule "min", the date must be at least "now"
validator.Rule{"date:min", time.Now().Unix()},

// passing date as time.Time{}
// following the rule "max", the date must be up to "now + 24 hours"
validator.Rule{"date:max", time.Now().Add(time.Hour * 24)},

// passing date as string
// following the rule "eq", the date must be exactly "2024-01-01T15:04:05Z"
validator.Rule{"date:eq", "2024-01-01T15:04:05Z"},

// using timezone
validator.Rule{"date:eq", "2024-01-01T15:04:05+02:00"},
```

### Time

The "time" modifier checks the correspondence between the prototype and the struct value with type [time.Time](https://pkg.go.dev/time#Time). In the context of this validator, the "time" modifier is intended to work with more precise time values, including comparisons of [nanoseconds](https://pkg.go.dev/time#Time.UnixNano). The prototype can be specified as a 64-bit string in nanoseconds, [int64](https://pkg.go.dev/time#Time.UnixNano) with nanoseconds, and [time](https://pkg.go.dev/time).

```go
// passing time as int64
// following the rule "min", the time must be at least "now"
validator.Rule{"time:min", time.Now().Unix()},

// passing time as time.Time{}
// following the rule "max", the time must be up to "now + 24 hours"
validator.Rule{"time:max", time.Now().Add(time.Hour * 24)},

// passing time as string
// following the rule "eq", the time must be exactly "1705337534239050689"
validator.Rule{"time:eq", "1705337534239050689"},
```

### Fields

The "fields" modifier checks the equality between the threshold rule and the number of successfully validated fields. This modifier must be placed last of the specified validator rules. Otherwise, an incorrect result will be returned. Please note, that this modifier should be placed without the "Optional" and the "Field" parameters, as it runs after validating the struct fields

```go
{
  // Must contain at least one valid field after validation, 
  // otherwise the validator will return "invalid body value"
  Check: validator.Rule{"fields:min", 1},
}
```
