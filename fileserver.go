package chip

import (
	"net/http"

	"github.com/go-chi/chi"
)

const filePattern = "/*"

func ServeRoot(r chi.Router, root string) {
	fs := fileServerForRoot(root)
	r.Get(filePattern, fs.ServeHTTP)
}

func fileServerForRoot(root string) http.Handler {
	return http.FileServer(http.Dir(root))
}
