package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/songlinm/basiccalc/calc"
)

// CalcReq represents the data payload in an evaluation request
type CalcReq struct {
	Notion     string `json:"notion"`
	Expression string `json:"expression"`
}

// Reply represents an expression's evaluation result
type Reply struct {
	Ret    float64 `json:"result"`
	ErrMsg string  `json:"error,omitempty"`
}

// Print usage on the default page
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `POST raw expression to /calculator e.g.
curl -H 'Content-Type: application/json' -X POST -d '{"notion": "prefix", "expression": "/ 8 3"}' <base-url>/calculator
curl -H 'Content-Type: application/json' -X POST -d '{"notion": "infix", "expression": "( 1 + 9 ) / 2 * ( 3 + 8 )"}' <base-url>/calculator
`)
}

// serves the calculator endpoint
func calculator(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" { // sending data to the server should be a POST
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// RESTful: expects json
	if contentType := req.Header.Get("Content-type"); contentType != "application/json" {
		w.Header().Set("Accept", "application/json")
		http.Error(w, "Expects Content-type: application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Decode the request
	var (
		rawData []byte
		calcReq = &CalcReq{}
		rep     = &Reply{}
		err     error
	)

	// Notice here we read everything in the request payload in one go.
	// Notice the calc library, in contrast, can handle streaming data, i.e.
	// calc.Prefix(r io.Reader) the parameter r does not require everything to
	// be read in one go. Therefore, if there is future need to handle extremely long
	// expressions, we can have another endpoint (or an v2 endpoint) that takes
	// content type of application/octet-stream, which can be directly passed into
	// the calc.Prefix or calc.Infix APIs
	if rawData, err = ioutil.ReadAll(req.Body); err != nil {
		http.Error(w, ``, http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(rawData, calcReq); err != nil {
		http.Error(w, fmt.Sprintf(`Invalid JSON: %v`, err), http.StatusBadRequest)
		return
	}

	w.Header().Set(`Content-Type`, `application/json`)

	switch calcReq.Notion {
	case `prefix`:
		rep.Ret, err = calc.Prefix(strings.NewReader(calcReq.Expression))
	case `infix`:
		rep.Ret, err = calc.Infix(strings.NewReader(calcReq.Expression))
	default:
		err = fmt.Errorf(`Only prefix and infix notions are supported`)
	}

	if err != nil {
		rep.ErrMsg = err.Error()
	}

	encoded, encErr := json.Marshal(rep)
	if encErr != nil {
		http.Error(w, ``, http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, `%s`, string(encoded))
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/calculator", calculator)

	log.Println(`Starting BasicCalc server...`)
	if err := http.ListenAndServe(`:8080`, nil); err != nil {
		log.Panicf(`Failed to serve HTTP: %v`, err)
	}
}
