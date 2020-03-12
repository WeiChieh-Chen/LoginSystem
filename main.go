package main

import (
	"log"
	"net/http"
)

func main() {
	m := new(AuthMux)
	if err := http.ListenAndServe(":8001", m); err != nil {
		log.Fatal(err)
	}
}
