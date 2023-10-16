package ctxvls_test

import (
	"context"
	"testing"

	"github.com/qawatake/ctxvls"
)

type key struct{}

func BenchmarkWithKeyValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, key{}, 10)
	}
}

func BenchmarkWithKeyValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, key{}, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkValuesFromByKey(b *testing.B) {
	ctx := context.Background()
	for j := 0; j < 10; j++ {
		ctx = ctxvls.WithValues(ctx, key{}, j)
	}
	for i := 0; i < b.N; i++ {
		ctxvls.ValuesFrom[key, int](ctx, key{})
	}
}

func BenchmarkWithKeyValue_multiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		for j := 0; j < 1000; j++ {
			ctx = ctxvls.WithValues(ctx, key{}, j)
		}
	}
}

func BenchmarkWithKeyValues_multiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		for j := 0; j < 1000; j++ {
			ctx = ctxvls.WithValues(ctx, key{}, j+1, j+2, j+3, j+4, j+5, j+6, j+7, j+8, j+9, j+10)
		}
	}
}
