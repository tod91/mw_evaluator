// A Module for handling incoming REST requests from the client(in our case curl)

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"mw_evaluator/math"
	"mw_evaluator/parser"
	"net/http"
)

func evaluate(w http.ResponseWriter, r *http.Request) {
	var input map[string]string

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		panic(err.Error())
	}

	result := parser.Parse(input["expression"])
	resp := math.Eval(result)

	fmt.Print(resp)
	response := make(map[string]int)
	response["result"] = 42
	//marshal responce and return
	w.Write([]byte(input["expression"] + "\n"))

}

func validate(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func errors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func StartServer() {
	http.HandleFunc("/evaluate", evaluate)
	http.HandleFunc("/validate", validate)
	http.HandleFunc("/errors", errors)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
