package mid

import (
	"context"
	"fmt"
	"runtime/debug"

	"example.com/service/app/api/metrics"
)

func Panics(ctx context.Context, handler Handler) (err error) {

	defer func() {
		if rec := recover(); rec != nil {
			trace := debug.Stack()
			err = fmt.Errorf("Panic [%v] Trace [%s]", rec, string(trace))
			metrics.AddPanics(ctx)
		}
	}()
	return handler(ctx)
}
