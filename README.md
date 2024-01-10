# slim-validator
A simple validation package for Go

## Usage

```go
type Article struct {
  Id    uint   `json:"id"`
  Phone string `json:"phone"`
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
}

article := Article{
  Id:    12,
  Phone: "+380001234567",
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
