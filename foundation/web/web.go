package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// this is a type of Handler that we want to use

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// Creating a wrapper over standard library mux

type App struct {
	*http.ServeMux // My App is everything a mux is
	shutdown       chan os.Signal
	mw             []MidHandler
}

func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

func (a *App) HandleFunc(pattern string, handle Handler, mw ...MidHandler) {

	handle = wrapMiddleware(mw, handle)
	handle = wrapMiddleware(a.mw, handle)

	h := func(w http.ResponseWriter, r *http.Request) {
		// we can write any code that we want

		// we cant do logging direct here as foundation packages are not allowed to do logging , so we need mniddleware logging
		v := Values{
			Now:     time.Now(),
			TraceID: uuid.NewString(),
		}

		ctx := setValues(r.Context(), &v)

		if err := handle(ctx, w, r); err != nil {
			// error handling
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)

}
