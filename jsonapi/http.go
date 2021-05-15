package jsonapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/abelce/at"
)

var errs []JsonapiError

func errs2doc(errs []JsonapiError) (string, error) {
	doc := JsonapiDocument{
		Errors: errs,
	}
	b, e := json.Marshal(doc)
	if e != nil {
		return "", errors.New("can not encode errors object")
	}
	return string(b), nil
}

func Handle404(w http.ResponseWriter) {
	ResetHTTPErrors()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	errs = append(errs, JsonapiError{
		Detail: "Not Found",
	})
	json, e := errs2doc(errs)
	if e != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, e.Error())
	}

	fmt.Fprintln(w, json)
}

func HandleHTTPError(w http.ResponseWriter, err error) {
	_, f, l, _ := runtime.Caller(1)
	errs = append(errs, JsonapiError{
		Code:   400,
		Detail: err.Error(),
		Meta: map[string]interface{}{
			"file": f,
			"line": l,
		},
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json, e := errs2doc(errs)

	if e != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, e.Error())
	}

	log.Printf("ResponseBody:%s\n", json)

	fmt.Fprintln(w, json)
}

func ResetHTTPErrors() {
	errs = nil
}

func HandleServerError(w http.ResponseWriter, e error) {
	_, f, l, _ := runtime.Caller(1)
	errData := struct {
		Code   int                    `json:"code"`
		Detail string                 `json:"detail"`
		Meta   map[string]interface{} `json:"meta"`
	}{
		Code:   500,
		Detail: e.Error(),
		Meta: map[string]interface{}{
			"file": f,
			"line": l,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	j, err := json.Marshal(errData)

	if at.Ensure(&err) {
		w.WriteHeader(500)
		fmt.Fprintln(w, e.Error())
	}
	fmt.Fprintln(w, string(j))
}
