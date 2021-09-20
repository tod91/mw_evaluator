// A Module for handling incoming REST requests from the client(in our case curl)

package handler

import (
	"encoding/json"
	"log"
	"mw_evaluator/errtracker"
	"mw_evaluator/math"
	"mw_evaluator/models"
	"mw_evaluator/parser"
	"mw_evaluator/validator"
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
		panic(err.Error())
	}

	rdyForParsing := parser.PreProcessExp(input["expression"])
	tokens := parser.Parse(rdyForParsing)
	ok, _ := validator.IsOk(tokens, rdyForParsing, "/evaluate")

	if !ok {
		panic("expression didn't pass validation")
	}

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

func validate(w http.ResponseWriter, r *http.Request) {
	var input map[string]string

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		panic(err.Error())
	}

	resp := models.ValidateResp{}

	rdyForParsing := parser.PreProcessExp(input["expression"])
	tokens := parser.Parse(rdyForParsing)
	resp.Valid, err = validator.IsOk(tokens, rdyForParsing, "/validate")

	if err != nil {
		resp.Reason = err.Error()
	}

	jData, err := json.Marshal(&resp)

	if err != nil {
		panic("cannot marshal resp")
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jData)

}

func errors(w http.ResponseWriter, _ *http.Request) {
	resp, err := errtracker.Tracker.GetAll()
	if err != nil {
		panic("cannot get asd")
	}

	jData, err := json.Marshal(resp)

	if err != nil {
		panic("cannot marshal resp")
	}
	w.Header().Add("Content-Type", "application/json")

	w.Write(jData)
}
