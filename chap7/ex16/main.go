package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bravmi/gopl/chap7/eval"
)

func parseEnv(url *url.URL) (eval.Env, error) {
	env := eval.Env{}
	for param, values := range url.Query() {
		if param == "expr" || len(values) != 1 {
			continue
		}
		val, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value for %s: %s", param, err)
		}
		env[eval.Var(param)] = val
	}
	return env, nil
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	exprStr := r.URL.Query().Get("expr")
	if exprStr == "" {
		http.Error(w, "missing 'expr' query parameter", http.StatusBadRequest)
		return
	}
	log.Printf("exprStr: %s", exprStr)

	expr, err := eval.Parse(exprStr)
	if err != nil {
		error := fmt.Sprintf("failed to parse expression: %s", err.Error())
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	log.Printf("expr: %s", expr.String())

	env, err := parseEnv(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("env: %v", env)

	fmt.Fprintln(w, expr.Eval(env))
}

func main() {
	http.HandleFunc("/calc", calcHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
