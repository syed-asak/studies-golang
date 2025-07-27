package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles form submissions.
// It parses form data and displays the name and address fields.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler returns a simple greeting.
// It only handles GET requests to the "/hello" path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "Get" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	// Create a file server for serving static files from the "static" directory.
	fileServer := http.FileServer(http.Dir("./static"))

	// Route for serving static files (e.g., index.html, styles.css).
	http.Handle("/", fileServer)

	// Route for handling form submissions at /form.
	http.HandleFunc("/form", formHandler)

	// Route for responding with a simple greeting at /hello.
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080\n")

	// Start the server and log any error if it fails to start.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
