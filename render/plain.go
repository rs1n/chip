package render

import "net/http"

func Plain(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Status(w, status)

	if _, err := w.Write([]byte(data)); err != nil {
		panic(err)
	}
}
