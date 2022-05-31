package main

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"log"
	"net/http"
)

//go:embed swagger.json
var spec []byte

func main() {
	log.SetFlags(0)
	http.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
