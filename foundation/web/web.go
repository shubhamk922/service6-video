package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// this is a type of Handler that we want to use

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// Creating a wrapper over standard library mux

type App struct {
	*http.ServeMux // My App is everything a mux is
	shutdown       chan os.Signal
}

func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
	}
}

func (a *App) HandleFunc(pattern string, handle Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {
		// we can write any code that we want

		if err := handle(r.Context(), w, r); err != nil {
			// error handling
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)

}
