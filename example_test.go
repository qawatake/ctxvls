package ctxvls_test

import (
	"context"
	"fmt"

	"github.com/qawatake/ctxvls"
)

func Example() {
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
