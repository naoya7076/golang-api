package middlewares

import (
	"context"
	"sync"
)

// traceIDKey 型を空の構造体 struct{}にしているのは、無駄な情報を持たせないようにすること でメモリの節約をする効果があるから
type traceIDKey struct{}

var (
	logNo int = 1
	mu    sync.Mutex
)

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}
