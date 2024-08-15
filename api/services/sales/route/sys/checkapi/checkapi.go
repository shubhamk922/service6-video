package checkapi

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"example.com/service/app/api/errs"
	"example.com/service/business/data/sqldb"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

type app struct {
	build string
	log   *logger.Logger
	db    *sqlx.DB
}

func NewApp(build string, log *logger.Logger, db *sqlx.DB) *app {
	return &app{
		build: build,
		db:    db,
		log:   log,
	}
}

func (api *app) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status     string `json:"status,omitempty"`
		Build      string `json:"build,omitempty"`
		Host       string `json:"host,omitempty"`
		Name       string `json:"name,omitempty"`
		PodIP      string `json:"podIP,omitempty"`
		Node       string `json:"node,omitempty"`
		Namespace  string `json:"namespace,omitempty"`
		GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
	}{
		Status:     "up",
		Build:      api.build,
		Host:       host,
		Name:       os.Getenv("KUBERNETES_NAME"),
		PodIP:      os.Getenv("KUBERNETES_POD_IP"),
		Node:       os.Getenv("KUBERNETES_NODE_NAME"),
		Namespace:  os.Getenv("KUBERNETES_NAMESPACE"),
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}

	return web.Respond(ctx, w, data, http.StatusOK)

}

func (api *app) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	// maybe this logic should be part of application lofic , hence can be part of app layer
	if err := sqldb.StatusCheck(ctx, api.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
		api.log.Info(ctx, "readiness failure", "status", status)
	}

	data := struct {
		Status string
	}{
		Status: status,
	}

	return web.Respond(ctx, w, data, statusCode) // what is we dont need to use json then , we nedd code to prepare respoonse
}

func (api *app) testError(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {

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

func (api *app) testPanic(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {

	if n := rand.Intn(100); n%2 == 0 {
		panic("We are panicking")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK) // what is we dont need to use json then , we nedd code to prepare respoonse
}
