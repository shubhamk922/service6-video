package mid

import (
	"context"

	"example.com/service/app/api/errs"
	"example.com/service/foundation/logger"
)

// Errors handles errors coming out of the call chain.
func Errors(ctx context.Context, log *logger.Logger, next Handler) error {
	err := next(ctx)

	if err == nil {
		return nil
	}

	log.Error(ctx, "message", "ERROR", err.Error())

	if errs.IsError(err) {
		return errs.GetError(err)
	}
	return errs.Newf(errs.Unknown, errs.Unknown.String())

}
