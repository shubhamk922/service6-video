package mux

import (
	"os"

	"example.com/service/api/services/sales/route/sys/checkapi"
	"example.com/service/app/api/mid"
	"example.com/service/foundation/logger"
	"example.com/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.

func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {

	mux := web.NewApp(shutdown, mid.Logger(log))

	checkapi.Routes(mux)

	return mux
}
