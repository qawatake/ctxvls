package ctxvls

import (
	"context"
)

func WithKeyValues[T comparable, V any](ctx context.Context, key T, values ...V) context.Context {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	newStore := &valueStore[V]{
		parent: store,
		values: values,
	}
	return context.WithValue(ctx, pkey, newStore)
}

func ValuesFrom[T comparable, V any](ctx context.Context, key T) []V {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	return valuesFrom(store)
}

func AnyValuesFrom[T comparable, V any](ctx context.Context, key T) []any {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	return anyValuesFrom(store)
}

func valuesFrom[V any](s *valueStore[V]) []V {
	if s == nil {
		return nil
	}
	numValues := 0
	for ss := s; ss != nil; ss = ss.parent {
		numValues += len(ss.values)
	}
	values := make([]V, numValues)
	i := numValues - 1
	for ss := s; ss != nil; ss = ss.parent {
		for j := len(ss.values) - 1; j >= 0; j-- {
			values[i] = ss.values[j]
			i--
		}
	}
	return values
}

func anyValuesFrom[V any](s *valueStore[V]) []any {
	if s == nil {
		return nil
	}
	numValues := 0
	for ss := s; ss != nil; ss = ss.parent {
		numValues += len(ss.values)
	}
	values := make([]any, numValues)
	i := numValues - 1
	for ss := s; ss != nil; ss = ss.parent {
		for j := len(ss.values) - 1; j >= 0; j-- {
			values[i] = ss.values[j]
			i--
		}
	}
	return values
}

type valueStore[V any] struct {
	parent *valueStore[V]
	values []V
}

type partitionKey[T comparable] struct {
	key T
}

func storeFromContext[T comparable, V any](ctx context.Context, pkey partitionKey[T]) *valueStore[V] {
	if store, ok := ctx.Value(pkey).(*valueStore[V]); ok {
		return store
	}
	return nil
}
