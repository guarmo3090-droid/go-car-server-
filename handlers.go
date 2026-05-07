package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func sendErr(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var c Car
	json.NewDecoder(r.Body).Decode(&c)
	if err := c.Validate(); err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	mu.Lock()
	c.ID = strconv.Itoa(nextID)
	nextID++
	cars[c.ID] = c
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func getCars(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	list := []Car{}
	for _, c := range cars {
		list = append(list, c)
	}
	json.NewEncoder(w).Encode(list)
}

