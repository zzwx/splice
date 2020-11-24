[![github.com/zzwx/splice](doc/splice-gobadge.svg)](https://pkg.go.dev/github.com/zzwx/splice)

# Splice

Package splice is a simple implementation of JavaScript's array.splice function for []string in go (golang).

# Usage

```go
import "github.com/zzwx/splice"

// Delete one element starting with index #0, and squeeze in ["0","1"]
source := []string{`a`}
deleted := splice.Strings(&source, 0, 1, "0", "1")
fmt.Println("source:", source)
fmt.Println("deleted:", deleted)
// source: [0 1]
// deleted: [a]

// Emulate one argument source.splice(1)
source = []string{"a", "b", "c"}
deleted = splice.Strings(&source, 1, len(source))
fmt.Println("source:", source)
fmt.Println("deleted:", deleted)
// source: [a]
// deleted: [b c]
```