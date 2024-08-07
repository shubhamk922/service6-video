package mid

import (
	"context"
	"net/http"

	"example.com/service/app/api/errs"
	"example.com/service/app/api/mid"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
)

var codeStatus [17]int

// init maps out the error codes to http status codes.
func init() {
	codeStatus[errs.OK.Value()] = http.StatusOK
	codeStatus[errs.Canceled.Value()] = http.StatusGatewayTimeout
	codeStatus[errs.Unknown.Value()] = http.StatusInternalServerError
	codeStatus[errs.InvalidArgument.Value()] = http.StatusBadRequest
	codeStatus[errs.DeadlineExceeded.Value()] = http.StatusGatewayTimeout
	codeStatus[errs.NotFound.Value()] = http.StatusNotFound
	codeStatus[errs.AlreadyExists.Value()] = http.StatusConflict
	codeStatus[errs.PermissionDenied.Value()] = http.StatusForbidden
	codeStatus[errs.ResourceExhausted.Value()] = http.StatusTooManyRequests
	codeStatus[errs.FailedPrecondition.Value()] = http.StatusBadRequest
	codeStatus[errs.Aborted.Value()] = http.StatusConflict
	codeStatus[errs.OutOfRange.Value()] = http.StatusBadRequest
	codeStatus[errs.Unimplemented.Value()] = http.StatusNotImplemented
	codeStatus[errs.Internal.Value()] = http.StatusInternalServerError
	codeStatus[errs.Unavailable.Value()] = http.StatusServiceUnavailable
	codeStatus[errs.DataLoss.Value()] = http.StatusInternalServerError
	codeStatus[errs.Unauthenticated.Value()] = http.StatusUnauthorized
}

// This is a function to create a Middleware Handler type which executes errors middleware functionality
func Errors(log *logger.Logger) web.MidHandler {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			if err := mid.Errors(ctx, log, hdl); err != nil {
				errs := err.(errs.Error)

				if err := web.Respond(ctx, w, errs, codeStatus[errs.Code.Value()]); err != nil {
					return err
				}

				// THink about err type as signals , so that when we receive err as shutdown signal , we will handle to shutdown service
				if web.IsShutdown(err) {
					return err
				}
			}

			return mid.Errors(ctx, log, hdl)
		}

		return h
	}
	return m
}
