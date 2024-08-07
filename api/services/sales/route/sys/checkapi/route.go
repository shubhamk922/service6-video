package checkapi

import (
	"example.com/service/foundation/web"
)

func Routes(mux *web.App) {
	mux.HandleFunc("GET /liveness", liveness)
	mux.HandleFunc("GET /readiness", readiness)
	mux.HandleFunc("GET /testerror", testError)

}
