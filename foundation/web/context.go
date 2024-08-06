package web

import (
	"context"
	"time"
)

type ctxKey int

const key ctxKey = 1

type Values struct {
	TraceID    string
	Now        time.Time
	statusCode int
}

func GetValues(ctx context.Context) *Values {

	v, ok := ctx.Value(key).(*Values)

	if !ok {
		return &Values{
			TraceID: "0000000-000000-000000-000000000",
			Now:     time.Now(),
		}
	}
	return v

}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}

	return v.TraceID
}

func setStatusCode(ctx context.Context, statusCode int) {

	v, ok := ctx.Value(key).(*Values)

	if !ok {
		return
	}
	v.statusCode = statusCode

}

func setValues(ctx context.Context, v *Values) context.Context {
	return context.WithValue(ctx, key, v)
}
