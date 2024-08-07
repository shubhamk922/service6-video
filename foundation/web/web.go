package web

import (
	"context"
	"errors"
	"net/http"
	"os"
	"syscall"
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

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
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
			// This will be the end of food chain , from here it goes back to mux , hence any sort of error need to be handle error and return nil
			if validateError(err) {
				// Server got restart
				a.SignalShutdown()
				return
			}
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)

}

func (a *App) HandleFuncNoMiddleware(pattern string, handle Handler, mw ...MidHandler) {

	

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
			// This will be the end of food chain , from here it goes back to mux , hence any sort of error need to be handle error and return nil
			if validateError(err) {
				// Server got restart
				a.SignalShutdown()
				return
			}
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)

}

func validateError(err error) bool {

	// Ignore syscall.EPIPE and syscall.ECONNRESET errors which occurs
	// when a write operation happens on the http.ResponseWriter that
	// has simultaneously been disconnected by the client (TCP
	// connections is broken). For instance, when large amounts of
	// data is being written or streamed to the client.
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://gosamples.dev/broken-pipe/
	// https://gosamples.dev/connection-reset-by-peer/

	switch {
	case errors.Is(err, syscall.EPIPE):

		// Usually, you get the broken pipe error when you write to the connection after the
		// RST (TCP RST Flag) is sent.
		// The broken pipe is a TCP/IP error occurring when you write to a stream where the
		// other end (the peer) has closed the underlying connection. The first write to the
		// closed connection causes the peer to reply with an RST packet indicating that the
		// connection should be terminated immediately. The second write to the socket that
		// has already received the RST causes the broken pipe error.
		return false

	case errors.Is(err, syscall.ECONNRESET):

		// Usually, you get connection reset by peer error when you read from the
		// connection after the RST (TCP RST Flag) is sent.
		// The connection reset by peer is a TCP/IP error that occurs when the other end (peer)
		// has unexpectedly closed the connection. It happens when you send a packet from your
		// end, but the other end crashes and forcibly closes the connection with the RST
		// packet instead of the TCP FIN, which is used to close a connection under normal
		// circumstances.
		return false
	}

	return true
}
