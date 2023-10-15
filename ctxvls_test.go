package ctxvls_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qawatake/ctxvls"
)

func TestValuesFromByKey(t *testing.T) {
	t.Run("with no value", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a")
		got := ctxvls.ValuesFromByKey(ctx, "a")
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("called twice with the same key", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a", "a", 1, 2)
		ctx = ctxvls.WithKeyValues(ctx, "a", "b", 2, 3)
		got := ctxvls.ValuesFromByKey(ctx, "a")
		want := []any{"a", 1, 2, "b", 2, 3}
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

	t.Run("Independent 3", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a", 1, 2)
		got1 := ctxvls.ValuesFromByKey(ctx, "a")
		got1[0] = 3
		got2 := ctxvls.ValuesFromByKey(ctx, "a")
		want := []any{1, 2}
		if diff := cmp.Diff(want, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Parallel", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		ctx = ctxvls.WithKeyValues(ctx, "a", 1, 2)
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			i := i
			wg.Add(1)
			go func() {
				ctxvls.WithKeyValues(ctx, "a", i)
				wg.Done()
			}()
		}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				ctxvls.ValuesFromByKey(ctx, "a")
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
func TestValuesFrom(t *testing.T) {
	t.Run("with no value", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues[any](ctx)
		got := ctxvls.ValuesFrom[string](ctx)
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("called twice", func(t *testing.T) {
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, 1, 2)
		ctx = ctxvls.WithValues(ctx, 2, 3)
		got := ctxvls.ValuesFrom[int](ctx)
		want := []int{1, 2, 2, 3}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Independent contexts", func(t *testing.T) {
		ctx := context.Background()
		ctx1 := ctxvls.WithValues(ctx, 1, 2)
		ctx2 := ctxvls.WithValues(ctx, 2, 3)
		got1 := ctxvls.ValuesFrom[int](ctx1)
		got2 := ctxvls.ValuesFrom[int](ctx2)
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
		ctx1 := ctxvls.WithValues(ctx, 1, 2)
		ctx2 := ctxvls.WithValues(ctx1, 2, 3)
		got1 := ctxvls.ValuesFrom[int](ctx1)
		got2 := ctxvls.ValuesFrom[int](ctx2)
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
		ctx = ctxvls.WithValues(ctx, 1, 2)
		got1 := ctxvls.ValuesFrom[int](ctx)
		got1[0] = 3
		got2 := ctxvls.ValuesFrom[int](ctx)
		want := []int{1, 2}
		if diff := cmp.Diff(want, got2); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Parallel", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		ctx = ctxvls.WithValues(ctx, 1, 2)

		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			i := i
			wg.Add(1)
			go func() {
				ctxvls.WithValues(ctx, i, i+1, i+2)
				wg.Done()
			}()
		}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				ctxvls.ValuesFrom[int](ctx)
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
