package main

import (
	"fmt"
	"log"
	"net/http"
)

func RecoveryFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered:%v", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h(w, r)
	}
}

func getHelloMessage() (string, error) {
	return "Hello Backend", nil
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Query().Get("panic") == "1" {
		panic("simulated panic in HelloHandle")
	}

	message, err := getHelloMessage()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	_, fmtErr := fmt.Fprintln(w, message)
	if fmtErr != nil {
		log.Printf("yanıt yazılamadı: %v", fmtErr)
	}
}

func main() {
	http.HandleFunc("/hello", RecoveryFunc(HelloHandle))
	fmt.Println("server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
