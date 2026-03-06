package main

import (
	"log"
	"net/http"
)

func main() {
		       mux := http.NewServeMux()
		       mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			       w.WriteHeader(http.StatusOK)
			       if _, err := w.Write([]byte("hello")); err != nil {
				       log.Printf("failed to write response: %v", err)
			       }
		       })
	       log.Println("Starting server on :8000")
	       if err := http.ListenAndServe(":8000", mux); err != nil {
		       log.Fatalf("Server failed: %v", err)
	       }
}
