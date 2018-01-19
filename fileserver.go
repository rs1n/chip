package chip

import (
	"net/http"

	"github.com/go-chi/chi"
)

const filePattern = "/*"

func ServeRoot(router chi.Router, root string) {
	fs := fileServerForRoot(root)
	router.Get(filePattern, fs.ServeHTTP)
}

func fileServerForRoot(root string) http.Handler {
	return http.FileServer(http.Dir(root))
}
