package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(w, "Hello, World once again :)")
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
