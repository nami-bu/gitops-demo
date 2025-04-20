package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Kubernetes with Argo CD 🚀")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I'm health now, don't worry!")
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "CPU: 80%, GPU: 10%, Memory: 20%")
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
