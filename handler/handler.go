package handler

import (
	"log"
	"net/http"
)

func evaluate(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func validate(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func errors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func HandleRequest() {
	http.HandleFunc("/evaluate", evaluate)
	http.HandleFunc("/validate", validate)
	http.HandleFunc("/errors", errors)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
