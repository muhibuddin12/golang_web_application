package main

import (
	"fmt"
	"net/http"

	"github.com/muhibuddin12/golang_web_application/handler"
)

func main() {
	fmt.Println("Running server")

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler.GetServeMux(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
