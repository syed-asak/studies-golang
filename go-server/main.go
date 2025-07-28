package main

import (
	"fmt"
	"log"
	"net/http"
)

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
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	// refering dir static where static files are
	// http.FileServer(...) Returns an http.Handler that serves HTTP requests with the contents of the file system.

	http.Handle("/", fileServer)
	// http.Handle("/", fileServer):Tells the HTTP server to use the fileServer handler for all requests to the root path /

	http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/form", formHandler)When a user accesses http://localhost:8080/form, the formHandler function will handle the request.Typically used to serve or process a form (e.g., HTML form submission).

	http.HandleFunc("/hello", helloHandler)
	// When a user accesses http://localhost:8080/hello, the helloHandler function will respond. Often used for returning a greeting or demo response.

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { //is a common pattern in Go for starting an HTTP server and handling any startup errors.
		log.Fatal(err) //http.ListenAndServe(":8080", nil) starts a web server listening on port 8080, second argument nil means:
	}
}
