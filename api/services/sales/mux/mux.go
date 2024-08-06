package mux

import (
	"os"

	"example.com/service/api/services/sales/route/sys/checkapi"
	"example.com/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.

func WebAPI(shutdown chan os.Signal) *web.App {

	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
