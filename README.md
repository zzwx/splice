[![https://github.com/zzwx/splice](doc/splice-gobadge.svg)](https://pkg.go.dev/github.com/zzwx/splice)

# Splice

Package splice is an generics implementation of JavaScript's `array.splice` function for `[]T` where `T` is constrained as `any`.

# Example

```go
import "github.com/zzwx/splice/v2"

var months = []string{"Jan", "March", "April", "June"}
splice.Splice(&months, 1, 0, "Feb") // inserts at index 1
fmt.Println(months)
deleted := splice.Splice(&months, 4, 1, "May") // replaces 1 element at index 4
fmt.Println(months)
fmt.Println(deleted)
// Output:
// [Jan Feb March April June]
// [Jan Feb March April May]
// [June]
```