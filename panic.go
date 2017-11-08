package chip

import "net/http"

func AbortHandler() {
	panic(http.ErrAbortHandler)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
