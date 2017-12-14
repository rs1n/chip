package render

import (
	"encoding/json"
	"net/http"
)

// Json renders a Json response.
func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	Status(w, status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// BindJson binds a Json payload to a destination object.
func BindJson(req *http.Request, v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}
