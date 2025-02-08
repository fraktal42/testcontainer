package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", Ping)

	http.ListenAndServe(":5000", nil)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
