package ctxvls2

import (
	"context"
)

type valueStore[V any] struct {
	parent valueser[V]
	value  V
}

type valuesStore[V any] struct {
	parent valueser[V]
	values []V
}

type valueser[V any] interface {
	Values() []V
	AnyValues() []any
}

type partitionKey[T comparable] struct {
	key T
}

func WithKeyValue[T comparable, V any](ctx context.Context, key T, value V) context.Context {
	pkey := partitionKey[T]{key: key}
	parent := storeFromContext[T, V](ctx, pkey)
	newStore := &valueStore[V]{
		parent: parent,
		value:  value,
	}
	return context.WithValue(ctx, pkey, newStore)
}

func storeFromContext[T comparable, V any](ctx context.Context, pkey partitionKey[T]) valueser[V] {
	if store, ok := ctx.Value(pkey).(valueser[V]); ok {
		return store
	}
	return (*valueStore[V])(nil)
}

func WithKeyValues[T comparable, V any](ctx context.Context, key T, values ...V) context.Context {
	if len(values) == 1 {
		return WithKeyValue(ctx, key, values[0])
	}
	pkey := partitionKey[T]{key: key}
	parent := storeFromContext[T, V](ctx, pkey)
	newStore := &valuesStore[V]{
		parent: parent,
		values: values,
	}
	return context.WithValue(ctx, pkey, newStore)
}

func ValuesFrom[T comparable, V any](ctx context.Context, key T) []V {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	return store.Values()
}

func AnyValuesFrom[T comparable, V any](ctx context.Context, key T) []any {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	return store.AnyValues()
}

func (s *valueStore[V]) AnyValues() []any {
	if s == nil {
		return nil
	}
	values := s.parent.AnyValues()
	values = append(values, s.value)
	return values
}

func (s *valueStore[V]) Values() []V {
	if s == nil {
		return nil
	}
	values := s.parent.Values()
	values = append(values, s.value)
	return values
}

func (s *valuesStore[V]) AnyValues() []any {
	if s == nil {
		return nil
	}
	values := s.parent.AnyValues()
	for _, v := range s.values {
		values = append(values, v)
	}
	return values
}

func (s *valuesStore[V]) Values() []V {
	if s == nil {
		return nil
	}
	values := s.parent.Values()
	values = append(values, s.values...)
	return values
}
