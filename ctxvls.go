package ctxvls

import (
	"context"

	ctxvls "github.com/qawatake/ctxvls/internal/ctxvls"
)

// WithKeyValues returns a copy of parent that stores values for key.
func WithKeyValues[T comparable](ctx context.Context, key T, values ...any) context.Context {
	return ctxvls.WithKeyValues(ctx, key, values...)
}

// ValuesFromByKey returns all the values stored for key by WithKeyValues.
func ValuesFromByKey[T comparable](ctx context.Context, key T) []any {
	return ctxvls.AnyValuesFrom[T, any](ctx, key)
}

// WithValues returns a copy of parent that stores values.
func WithValues[T any](ctx context.Context, values ...T) context.Context {
	return ctxvls.WithKeyValues(ctx, partitionKey[T]{}, values...)
}

// ValuesFrom returns all the values stored for type T by WithValues.
// Note that ValuesFrom does not return values stored by WithKeyValues even if value types are the same.
func ValuesFrom[T any](ctx context.Context) []T {
	return ctxvls.ValuesFrom[partitionKey[T], T](ctx, partitionKey[T]{})
}

type partitionKey[T any] struct{}