package functions

import (
	"fmt"
	"net/http"
)

// Handle different URL paths within the web application.
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
