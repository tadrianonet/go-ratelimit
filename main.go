package main

import (
	"fmt"
	"go-ratelimit/middlewares"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	http.Handle("/hello", middlewares.RateLimitMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("Server is running on http://localhost:8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
