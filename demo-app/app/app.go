package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	terminate := req.URL.Query().Get("q")
	if terminate != "" {
		fmt.Println("TERMINATING 'ACCIDENTALLY'")
		os.Exit(1)
	}
	fmt.Println(req)
	fmt.Fprintf(w, "<h1>Hello, World :)</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
