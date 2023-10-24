package main

import (
	"fmt"
	"groupie-tracker/functions"
	"log"
	"net/http"
)

func main() {
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout"))))
	http.HandleFunc("/", functions.HandleError)
	fmt.Println("Starting server: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
