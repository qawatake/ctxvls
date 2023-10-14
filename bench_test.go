package ctxvls_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qawatake/ctxvls"
)

type key struct{}

func BenchmarkWithKeyValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		for j := 0; j < 1000; j++ {
			ctx = ctxvls.WithKeyValues(ctx, key{}, j)
		}
	}
}

func BenchmarkWithKeyValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		for j := 0; j < 1000; j++ {
			ctx = ctxvls.WithKeyValues(ctx, key{}, j+1, j+2, j+3, j+4, j+5, j+6, j+7, j+8, j+9, j+10)
		}
	}
}

func BenchmarkValuesFromByKey(b *testing.B) {
	ctx := context.Background()
	for j := 0; j < 1000; j++ {
		ctx = ctxvls.WithKeyValues(ctx, key{}, j)
	}
	for i := 0; i < b.N; i++ {
		ctxvls.ValuesFromByKey(ctx, key{})
	}
}

type KV struct {
	Key   string
	Value string
}

func TestValues(t *testing.T) {
	ctx := context.Background()
	ctx = ctxvls.WithValues[KV](ctx,
		KV{
			Key:   "a",
			Value: "1",
		},
		KV{
			Key:   "a",
			Value: "2",
		},
	)
	got := ctxvls.ValuesFrom[KV](ctx)
	want := []KV{
		{Key: "a", Value: "1"},
		{Key: "a", Value: "2"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}