package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				if rcvr == http.ErrAbortHandler {
					// Response is already flushed.
					return
				}

				// Print a stack trace.
				log.Printf("panic: %+v", rcvr)
				debug.PrintStack()

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
