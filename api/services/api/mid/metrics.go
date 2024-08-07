package mid

import (
	"context"
	"net/http"

	"example.com/service/app/api/mid"
	"example.com/service/foundation/web"
)

func Metrics() web.MidHandler {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Metrics(ctx, hdl)

		}
		return h
	}

	return m

}
