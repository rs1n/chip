package chip

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	xmiddleware "github.com/rs1n/chip/middleware"
)

// BootstrapRouter plugs standard middleware.
func BootstrapRouter(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(xmiddleware.Recoverer)
}
