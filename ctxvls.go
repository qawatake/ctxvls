// Package ctxvls provide APIs for storing and retrieving collections of values in context, and it is designed to be safe for concurrent use.
package ctxvls

import (
	"context"

	ctxvls "github.com/qawatake/ctxvls/internal/ctxvls3"
)

// WithValues returns a copy of parent in which a collection of values are associated with a key.
func WithValues[T comparable, V any](ctx context.Context, key T, values ...V) context.Context {
	return ctxvls.WithKeyValues(ctx, key, values...)
}

// ValuesFrom returns all the values associated with key.
func ValuesFrom[T comparable, V any](ctx context.Context, key T) []V {
	return ctxvls.ValuesFrom[T, V](ctx, key)
}
