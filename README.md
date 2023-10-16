# ctxvls

[![Go Reference](https://pkg.go.dev/badge/github.com/qawatake/ctxvls.svg)](https://pkg.go.dev/github.com/qawatake/ctxvls)
[![test](https://github.com/qawatake/ctxvls/actions/workflows/test.yaml/badge.svg)](https://github.com/qawatake/ctxvls/actions/workflows/test.yaml)

Package `ctxvls` provide APIs for storing and retrieving collections of values in context, and it is designed to be safe for concurrent use.

```go
package main

import (
  "context"
  "fmt"

  "github.com/qawatake/ctxvls"
)

type MyKey struct{}

func main() {
	ctx := context.Background()
	ctx = ctxvls.WithValues(ctx, "a", 1, 2, 3)
	ctx = ctxvls.WithValues(ctx, "a", 2, 3, 4)
	ctx = ctxvls.WithValues(ctx, "b", 3, 4, 5)
	valuesForA := ctxvls.ValuesFrom[string, int](ctx, "a")
	valuesForB := ctxvls.ValuesFrom[string, int](ctx, "b")
	fmt.Println(valuesForA)
	fmt.Println(valuesForB)
	// Output:
	// [1 2 3 2 3 4]
	// [3 4 5]
}
```
