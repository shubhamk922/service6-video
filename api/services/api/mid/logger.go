package mid

import (
	"context"
	"net/http"

	"example.com/service/app/api/mid"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
)

// This is a type function of MidHandler defined in web

func Logger(log *logger.Logger) web.MidHandler {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, hdl)

		}
		return h
	}

	return m

}
