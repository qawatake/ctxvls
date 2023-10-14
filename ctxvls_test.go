package ctxvls_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qawatake/ctxvls"
)

func TestWithKeyValue(t *testing.T) {
	t.Run("with no value", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a")
		got := ctxvls.ValuesFromByKey(ctx, "a")
		var want []any = nil
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("called twice with the same key", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a", 1, 2)
		ctx = ctxvls.WithKeyValues(ctx, "a", 2, 3)
		got := ctxvls.ValuesFromByKey(ctx, "a")
		want := []any{1, 2, 2, 3}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("called twice with different keys", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a", 1, 2)
		ctx = ctxvls.WithKeyValues(ctx, "b", 2, 3)
		got := ctxvls.ValuesFromByKey(ctx, "a")
		want := []any{1, 2}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent contexts", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := ctxvls.WithKeyValues(ctx, "a", 1, 2)
		ctx2 := ctxvls.WithKeyValues(ctx, "a", 2, 3)
		got1 := ctxvls.ValuesFromByKey(ctx1, "a")
		got2 := ctxvls.ValuesFromByKey(ctx2, "a")
		want1 := []any{1, 2}
		want2 := []any{2, 3}
		if diff := cmp.Diff(want1, got1); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(want2, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent contexts 2", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := ctxvls.WithKeyValues(ctx, "a", 1, 2)
		ctx2 := ctxvls.WithKeyValues(ctx1, "a", 2, 3)
		got1 := ctxvls.ValuesFromByKey(ctx1, "a")
		got2 := ctxvls.ValuesFromByKey(ctx2, "a")
		want1 := []any{1, 2}
		want2 := []any{1, 2, 2, 3}
		if diff := cmp.Diff(want1, got1); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(want2, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Parallel", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(1000)
		for i := 0; i < 1000; i++ {
			i := i
			go func() {
				ctxvls.WithKeyValues(ctx, "a", i)
				ctxvls.ValuesFromByKey(ctx, "a")
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
