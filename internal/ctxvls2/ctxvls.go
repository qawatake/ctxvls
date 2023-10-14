package ctxvls

import "context"

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
	return store.Values()
}

func AnyValuesFrom[T comparable, V any](ctx context.Context, key T) []any {
	pkey := partitionKey[T]{key: key}
	store := storeFromContext[T, V](ctx, pkey)
	return store.AnyValues()
}

func (s *valueStore[V]) Values() []V {
	if s == nil {
		return nil
	}
	values := s.parent.Values()
	values = append(values, s.values...)
	return values
}

func (s *valueStore[V]) AnyValues() []any {
	if s == nil {
		return nil
	}
	values := s.parent.AnyValues()
	for _, v := range s.values {
		values = append(values, v)
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
