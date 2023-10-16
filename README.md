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
	key := MyKey{}
	ctx = ctxvls.WithKeyValues(ctx, key, 1, "b", 'x')
	ctx = ctxvls.WithKeyValues(ctx, key, 2, "c", 'y')
	values := ctxvls.ValuesFromByKey(ctx, key)
	fmt.Println(values)
	// Output:
	// [1 b 120 2 c 121]
}
```
