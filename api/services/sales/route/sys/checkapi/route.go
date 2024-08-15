package checkapi

import (
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

func Routes(build string, mux *web.App, log *logger.Logger, db *sqlx.DB) {

	api := NewApp(build, log, db)

	mux.HandleFuncNoMiddleware("GET /liveness", api.liveness)
	mux.HandleFuncNoMiddleware("GET /readiness", api.readiness)
	mux.HandleFunc("GET /testerror", api.testError)
	mux.HandleFunc("GET /testpanic", api.testPanic)

}
