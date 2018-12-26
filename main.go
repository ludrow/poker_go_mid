package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world")
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Print(err)
	}
}
