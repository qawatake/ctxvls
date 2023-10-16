package example_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qawatake/ctxvls/internal/example"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = example.AddToContext(ctx, example.KV{"a", 1}, example.KV{"b", 2})
	ctx = example.AddToContext(ctx, example.KV{"b", 2}, example.KV{"c", 3})
	got := example.FromContext(ctx)
	want := []example.KV{{"a", 1}, {"b", 2}, {"b", 2}, {"c", 3}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

func ExampleFromContext() {
	ctx := context.Background()
	ctx = example.AddToContext(ctx, example.KV{"a", 1}, example.KV{"b", 2})
	ctx = example.AddToContext(ctx, example.KV{"b", 2}, example.KV{"c", 3})
	kvs := example.FromContext(ctx)
	fmt.Println(kvs)
	// Output:
	// [{a 1} {b 2} {b 2} {c 3}]
}
