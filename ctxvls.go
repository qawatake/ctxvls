// Package ctxvls provide APIs for storing and retrieving collections of values in context, and it is designed to be safe for concurrent use.
package ctxvls

import (
	"context"

	ctxvls "github.com/qawatake/ctxvls/internal/ctxvls3"
)

// WithKeyValues returns a copy of parent in which values are associated with key.
// WithKeyValues is used together with ValuesFromByKey.
func WithKeyValues[T comparable](ctx context.Context, key T, values ...any) context.Context {
	return ctxvls.WithKeyValues(ctx, key, values...)
}

// ValuesFromByKey returns all the values associated with key and stored by WithKeyValues.
func ValuesFromByKey[T comparable](ctx context.Context, key T) []any {
	return ctxvls.ValuesFrom[T, any](ctx, key)
}

// WithValues returns a copy of parent that stores values.
// WithValues[T] is used together with ValuesFrom[T].
func WithValues[T any](ctx context.Context, values ...T) context.Context {
	return ctxvls.WithKeyValues(ctx, partitionKey[T]{}, values...)
}

// ValuesFrom returns all the values of type T stored by WithValues.
// Note that ValuesFrom does not return values stored by WithKeyValues even if value types are the same.
func ValuesFrom[T any](ctx context.Context) []T {
	return ctxvls.ValuesFrom[partitionKey[T], T](ctx, partitionKey[T]{})
}

type partitionKey[T any] struct{}
