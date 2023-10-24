package functions

import (
	"html/template"
	"net/http"
)

// handle incoming HTTP requests and serve an HTML template
func HomePage(w http.ResponseWriter, r *http.Request) {
	//The CORS header to make requests to this server from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	get_artists := GetArtists()

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error 01", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, get_artists); err != nil {
		http.Error(w, "Internal Server Error 02", http.StatusInternalServerError)
		return
	}
}
