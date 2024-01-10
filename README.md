# slim-validator
A simple validation package for Go

## Usage

```go
type Article struct {
  Id     uint     `json:"id"`
  Phone  string   `json:"phone"`
  Images []string `json:"images"`
}

filter := validator.Filter{
  {
    Field: "Id",
    Check: validator.NON_ZERO,
  },
  {
    Field: "Phone",
    Check: validator.Rule{"match", `^\+38\d{10}$`},
  },
  {
    Field: "Images",
    Check: validator.Rule{"eachMatch", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
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

Checks the passed value for [non-zero](https://golangbyexample.com/go-default-zero-value-all-types/).
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
This rule only works with: **string**

```go
{
  Field: "Phone",
  Check: validator.Rule{"match", `(?i)^[0-9a-f]{32}$`},
}
```

#### Each Match

Check whether any element of a collection matches a regular expression. This rule works with **array**, **slice**, and **map** that contain string values

```go
{
  Field: "Images",
  Check: validator.Rule{"eachMatch", `(?i)^https://img.it/[0-9a-f]{32}.jpe?g$`},
}
```
