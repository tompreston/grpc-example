// Example using http.ServeMux, not gRPC. ChatGPT generated this.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware to log requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

// Middleware to add a custom header
func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "MiddlewareHeader")
		next.ServeHTTP(w, r) // Call the next handler
	})
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// About page handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page!")
}

func main() {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// Wrap the ServeMux with middleware
	combinedMiddleware := loggingMiddleware(headerMiddleware(mux))

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: combinedMiddleware,
	}

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
