package mux

import (
	"os"

	"example.com/service/api/services/api/mid"
	"example.com/service/api/services/sales/route/sys/checkapi"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// WebAPI constructs a http.Handler with all application routes bound.

func WebAPI(build string, log *logger.Logger, db *sqlx.DB, shutdown chan os.Signal) *web.App {

	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	checkapi.Routes(build, mux, log, db)

	return mux
}
