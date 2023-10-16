package ctxvls_test

import (
	"context"
	"testing"

	"github.com/qawatake/ctxvls"
)

type key struct{}

func BenchmarkWithValues_one(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, key{}, 10)
	}
}

func BenchmarkWithValues_multiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, key{}, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkValuesFrom(b *testing.B) {
	ctx := context.Background()
	for j := 0; j < 10; j++ {
		ctx = ctxvls.WithValues(ctx, key{}, j)
	}
	for i := 0; i < b.N; i++ {
		ctxvls.ValuesFrom[key, int](ctx, key{})
	}
}
