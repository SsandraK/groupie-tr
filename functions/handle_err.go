package functions

import (
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		HomePage(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Error 404")
		return
	}
}
