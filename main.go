package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("POST /cars", createCar)
	http.HandleFunc("GET /cars", getCars)
	http.HandleFunc("GET /cars/{id}", getCar)
	http.HandleFunc("PUT /cars/{id}", updateCar)
	http.HandleFunc("PATCH /cars/{id}", patchCar)
	http.HandleFunc("DELETE /cars/{id}", deleteCar)


	fmt.Println("Server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}