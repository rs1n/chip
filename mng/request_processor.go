package mng

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/sknv/chip/mng/middleware"
)

// GetMgoSessionForRequest returns a Mongo session from a request context.
func GetMgoSessionForRequest(r *http.Request) *mgo.Session {
	return r.Context().Value(middleware.CtxKeyMgoSession).(*mgo.Session)
}
