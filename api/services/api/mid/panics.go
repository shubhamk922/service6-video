package mid

import (
	"context"
	"net/http"

	"example.com/service/app/api/mid"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
)

func Panics(log *logger.Logger) web.MidHandler {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Panics(ctx, hdl)

		}
		return h
	}

	return m

}
