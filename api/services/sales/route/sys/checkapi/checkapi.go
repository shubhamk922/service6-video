package checkapi

import (
	"context"
	"math/rand"
	"net/http"

	"example.com/service/app/api/errs"
	"example.com/service/foundation/web"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)

}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK) // what is we dont need to use json then , we nedd code to prepare respoonse
}

func testError(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {

	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this is trusted")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK) // what is we dont need to use json then , we nedd code to prepare respoonse
}
