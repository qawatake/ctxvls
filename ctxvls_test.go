package ctxvls_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qawatake/ctxvls"
)

func TestValuesFrom(t *testing.T) {
	t.Run("with no value", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues[string, any](ctx, "a")
		got := ctxvls.ValuesFrom[string, any](ctx, "a")
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("called twice with the same key", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues[string, any](ctx, "a", "a", 1, 2)
		ctx = ctxvls.WithValues[string, any](ctx, "a", "b", 2, 3)
		got := ctxvls.ValuesFrom[string, any](ctx, "a")
		want := []any{"a", 1, 2, "b", 2, 3}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("called twice with different keys", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, "a", 1, 2)
		ctx = ctxvls.WithValues(ctx, "b", 2, 3)
		got := ctxvls.ValuesFrom[string, int](ctx, "a")
		want := []int{1, 2}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent contexts", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := ctxvls.WithValues(ctx, "a", 1, 2)
		ctx2 := ctxvls.WithValues(ctx, "a", 2, 3)
		got1 := ctxvls.ValuesFrom[string, int](ctx1, "a")
		got2 := ctxvls.ValuesFrom[string, int](ctx2, "a")
		want1 := []int{1, 2}
		want2 := []int{2, 3}
		if diff := cmp.Diff(want1, got1); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(want2, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent contexts 2", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := ctxvls.WithValues(ctx, "a", 1, 2)
		ctx2 := ctxvls.WithValues(ctx1, "a", 2, 3)
		got1 := ctxvls.ValuesFrom[string, int](ctx1, "a")
		got2 := ctxvls.ValuesFrom[string, int](ctx2, "a")
		want1 := []int{1, 2}
		want2 := []int{1, 2, 2, 3}
		if diff := cmp.Diff(want1, got1); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(want2, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent 3", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, "a", 1, 2)
		got1 := ctxvls.ValuesFrom[string, int](ctx, "a")
		got1[0] = 3
		got2 := ctxvls.ValuesFrom[string, int](ctx, "a")
		want := []int{1, 2}
		if diff := cmp.Diff(want, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Parallel", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, "a", 1, 2)
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			i := i
			wg.Add(1)
			go func() {
				ctxvls.WithValues(ctx, "a", i)
				wg.Done()
			}()
		}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				ctxvls.ValuesFrom[string, int](ctx, "a")
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
