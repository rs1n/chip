package mng

import (
	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2"

	"github.com/sknv/chip/mng/middleware"
)

// BootstrapRouter puts a Mongo session to a request context.
func BootstrapRouter(r chi.Router, session *mgo.Session) {
	r.Use(middleware.WithMgoSession(session))
}
