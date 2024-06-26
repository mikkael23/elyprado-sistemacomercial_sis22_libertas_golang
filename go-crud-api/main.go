package main

import (
	"log"
	"net/http"

	"go-crud-api/routers"
)

func main() {
	router := routers.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
