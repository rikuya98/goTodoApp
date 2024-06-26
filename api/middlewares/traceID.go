package middlewares

import (
	"context"
	"sync"
)

var (
	LogNo int = 1
	mu    sync.Mutex
)

type traceIDKey struct{}

func newTraceID() int {
	var no int

	mu.Lock()
	no = LogNo
	LogNo += 1
	mu.Unlock()

	return no
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	traceID, ok := ctx.Value(traceIDKey{}).(int)
	if !ok {
		return 0
	}
	return traceID
}
