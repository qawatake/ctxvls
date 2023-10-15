# ctxvls

`ctxvls` provides simple APIs for storing and retrieving values in a context.

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
