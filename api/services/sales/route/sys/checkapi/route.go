package checkapi

import (
	"example.com/service/foundation/web"
)

func Routes(mux *web.App) {
	mux.HandleFuncNoMiddleware("GET /liveness", liveness)
	mux.HandleFuncNoMiddleware("GET /readiness", readiness)
	mux.HandleFunc("GET /testerror", testError)
	mux.HandleFunc("GET /testpanic", testPanic)

}
