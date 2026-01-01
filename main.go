package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "%s", "Hello World")
	})

	fmt.Println("startig http server on port 8080")
	http.ListenAndServe(":8080", nil)
	fmt.Println("stopped http server")
}
