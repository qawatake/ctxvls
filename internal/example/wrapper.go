package example

import (
	"context"

	"github.com/qawatake/ctxvls"
)

type KV struct {
	Key   string
	Value any
}

func AddToContext(ctx context.Context, kvs ...KV) context.Context {
	return ctxvls.WithValues(ctx, key{}, kvs...)
}

func FromContext(ctx context.Context) []KV {
	return ctxvls.ValuesFrom[key, KV](ctx, key{})
}

type key struct{}
