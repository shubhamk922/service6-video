package mux

import (
	"net/http"

	"example.com/service/api/services/sales/route/sys/checkapi"
)

// WebAPI constructs a http.Handler with all application routes bound.

func WebAPI() *http.ServeMux {

	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
