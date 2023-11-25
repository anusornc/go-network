package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func main() {
	// Handle request with gorilla/mux
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // get all params
		title := vars["title"]	// get title param
		page := vars["page"]	// get page param

		// create the variable contains the HTML
		b := []byte("<h1>Book: " + title + "</h1><br><h2>Page: " + page + "</h2>")


		//fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		w.Write(b) // write data to response
	})

	// Handle request for calcuate like /cal/10/plus/20 return total = 30
	r.HandleFunc("/cal/{num1}/{operator}/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // get all params
		num1 := vars["num1"]	// get num1 param
		num2 := vars["num2"]	// get num2 param
		operator := vars["operator"]	// get operator param

		total := 0

		// convert num1 and num2 to integers
		num1Int, _ := strconv.Atoi(num1)
		num2Int, _ := strconv.Atoi(num2)

		// calculate
		if operator == "plus" {
			total = num1Int + num2Int
		} else if operator == "minus" {
			total = num1Int - num2Int
		} else if operator == "multiply" {
			total = num1Int * num2Int
		} else if operator == "divide" {
			total = num1Int / num2Int
		}

		// create the variable contains the HTML
		b := []byte("<h1>Calculate: " + num1 + " " + operator + " " + num2 + "</h1><br><h2>Total: " + strconv.Itoa(total) + "</h2>")
		w.Write(b) // write data to response

	})


	// "/" is the route path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!") // write data to response
	})

	// /about is the route path
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My name is Anusorn")
	})

	// static/ is the route path
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// listen to port 8080
	http.ListenAndServe(":8080", r)

}
