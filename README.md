# ctxvls

`ctxvls` provides simple APIs for storing and retrieving values in a context.

```go
package main

import (
  "context"
  "fmt"

  "github.com/qawatake/ctxvls"
)

func main() {
	ctx := context.Background()
	ctx = ctxvls.WithKeyValues(ctx, "a", 1, "b", 'x')
	ctx = ctxvls.WithKeyValues(ctx, "a", 2, "c", 'y')
  values := ctxvls.ValuesFromByKey(ctx, "a")
	fmt.Println(values)
	// Output:
	// [1 b 120 2 c 121]
}
```
