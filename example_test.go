package ctxvls_test

import (
	"context"
	"fmt"

	"github.com/qawatake/ctxvls"
)

func ExampleValuesFromByKey() {
	ctx := context.Background()
	ctx = ctxvls.WithKeyValues(ctx, "a", 1, "b", 'x')
	ctx = ctxvls.WithKeyValues(ctx, "a", 2, "c", 'y')
	ctx = ctxvls.WithKeyValues(ctx, "b", 3, "c", 'z')
	valuesForA := ctxvls.ValuesFromByKey(ctx, "a")
	valuesForB := ctxvls.ValuesFromByKey(ctx, "b")
	fmt.Println(valuesForA)
	fmt.Println(valuesForB)
	// Output:
	// [1 b 120 2 c 121]
	// [3 c 122]
}

func ExampleValuesFrom() {
	ctx := context.Background()
	var a, b, c int = 1, 2, 3
	type MyInt int
	var x, y, z MyInt = 100, 200, 300
	ctx = ctxvls.WithValues(ctx, a, b, c)
	ctx = ctxvls.WithValues(ctx, x, y, z)
	ints := ctxvls.ValuesFrom[int](ctx)
	myints := ctxvls.ValuesFrom[MyInt](ctx)
	fmt.Println(ints)
	fmt.Println(myints)
	// Output:
	// [1 2 3]
	// [100 200 300]
}

func Example_readme() {
	ctx := context.Background()
	ctx = ctxvls.WithKeyValues(ctx, "a", 1, "b", 'x')
	ctx = ctxvls.WithKeyValues(ctx, "a", 2, "c", 'y')
	values := ctxvls.ValuesFromByKey(ctx, "a")
	fmt.Println(values)
	// Output:
	// [1 b 120 2 c 121]
}
