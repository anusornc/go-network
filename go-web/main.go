package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ex : http://localhost:80/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// print : Server is running on port 80
	fmt.Println("Server is running on port 80")
	http.ListenAndServe(":80", nil)
}
