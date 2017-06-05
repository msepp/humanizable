package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msepp/humanizable/animals"
)

// Our recognized animals
var animalList = []Humanizable{
	&animals.Dog{},
	&animals.Cat{},
}

// getHandler returns a handler for the given Humanizable
func getHandler(h Humanizable) http.HandlerFunc {
	// Function closures are a thing in Golang
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse age from request into a int64 or show error if not possible
		age, err := strconv.ParseInt(r.FormValue("age"), 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Error parsing age.. sorry! Try giving query param ?age=${number} (%s)", err)
			return
		}

		// Output animal age in human years.
		// Not the explicit conversion from int64 to int, Golang is strict about types
		fmt.Fprintf(w, "%d years in %s years: %d", int(age), h.Path(), h.AgeInHumanYears(int(age)))
	}
}

// handleRoot handles our root path.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><strong>Try one of our animals</strong><br/>")

	// Output a listing of animals.
	for _, animal := range animalList {
		fmt.Fprint(w, fmt.Sprintf("<a href='/%s/?age=1'>%s</a><br/>", animal.Path(), animal.Path()))
	}

	fmt.Fprintf(w, "</html>")
}

func main() {
	var addr string

	// Parse server address setting
	flag.StringVar(&addr, "addr", ":8080", "Server address to bind to")
	flag.Parse()

	// Create a new mux for handling request paths and set up root handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	// Setup a handler for all recognized animals
	for _, animal := range animalList {
		mux.HandleFunc(fmt.Sprintf("/%s/", animal.Path()), getHandler(animal))
	}

	// Start server
	log.Printf("Starting at address %s", addr)
	log.Fatalln(http.ListenAndServe(addr, mux))
}
