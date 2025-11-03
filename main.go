package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("hello backend")
}

func main() {
	http.HandleFunc("/hello", HelloHandle)
	fmt.Println("server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
