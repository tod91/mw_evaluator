// A Module for handling incoming REST requests from the client(in our case curl)

package handler

import (
	"encoding/json"
	"log"
	"mw_evaluator/math"
	"mw_evaluator/parser"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/evaluate", evaluate)
	http.HandleFunc("/validate", validate)
	http.HandleFunc("/errors", errors)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func evaluate(w http.ResponseWriter, r *http.Request) {
	var input map[string]string

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		{
			panic(err.Error())
		}
	}

	tokens := parser.Parse(input["expression"])
	result := math.Eval(tokens)

	resp := map[string]int{
		"result": result,
	}

	jData, err := json.Marshal(resp)
	if err != nil {
		panic("cannot marshal resp")
	}
	w.Header().Add("Content-Type", "application/json")

	w.Write(jData)
}

func validate(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func errors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}
