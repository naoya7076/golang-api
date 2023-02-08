package common

import (
	"context"
	"net/http"
)

// traceIDKey 型を空の構造体 struct{}にしているのは、無駄な情報を持たせないようにすること でメモリの節約をする効果があるから
type traceIDKey struct{}

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

type userNameKey struct{}

func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})
	if usernameStr, ok := id.(string); ok {
		return usernameStr
	}
	return ""
}

func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)

	return req
}
