package chip

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	xmiddleware "github.com/sknv/chip/middleware"
)

// BootstrapRouter plugs standard middleware.
func BootstrapRouter(router chi.Router) {
	router.Use(middleware.Logger)
	router.Use(xmiddleware.Recoverer)
}
