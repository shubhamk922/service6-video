package mid

import (
	"context"
	"fmt"
	"runtime/debug"
)

func Panics(ctx context.Context, handler Handler) (err error) {

	defer func() {
		if rec := recover(); rec != nil {
			trace := debug.Stack()
			err = fmt.Errorf("Panic [%v] Trace [%s]", rec, string(trace))
		}
	}()
	return handler(ctx)
}
