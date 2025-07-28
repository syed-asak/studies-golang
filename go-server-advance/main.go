package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Middleware for logging each request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// formHandler handles POST form submissions with validation.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	if name == "" || address == "" {
		http.Error(w, "Name and Address fields are required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Post request successful\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler serves a simple greeting for GET requests.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello! Welcome to the advanced Go server.")
}

// apiHandler returns a simple JSON response
func apiHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Time    string `json:"time"`
	}

	res := Response{
		Status:  "success",
		Message: "This is a JSON response",
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	mux := http.NewServeMux()

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	// Handlers
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/api", apiHandler)

	// Wrap with logging middleware
	loggedMux := loggingMiddleware(mux)

	// Server setup
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	// Graceful shutdown
	idleConnsClosed := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("🚀 Server is running at http://localhost:8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped")
}
