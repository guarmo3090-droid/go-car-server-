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

func getCar(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  defer mu.Unlock()
  c, exists := cars[r.PathValue("id")]
  if !exists {
    sendErr(w, http.StatusNotFound, "car not found")
    return
  }
  json.NewEncoder(w).Encode(c)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  var c Car
  json.NewDecoder(r.Body).Decode(&c)
  if err := c.Validate(); err != nil {
    sendErr(w, http.StatusBadRequest, err.Error())
    return
  }

  mu.Lock()
  defer mu.Unlock()
  if _, exists := cars[id]; !exists {
    sendErr(w, http.StatusNotFound, "avtomibile ne znaishlyosha")
    return
  }
  c.ID = id
  cars[id] = c
}

func patchCar(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  var updates map[string]interface{}
  json.NewDecoder(r.Body).Decode(&updates)

  mu.Lock()
  defer mu.Unlock()
  c, exists := cars[id]
  if !exists {
    sendErr(w, http.StatusNotFound, "automobile ne bulo znaideno")
    return
  }

  if brand, ok := updates["brand"].(string); ok { c.Brand = brand }
  if model, ok := updates["model"].(string); ok { c.Model = model }
  if year, ok := updates["year"].(float64); ok { c.Year = int(year) }
  if price, ok := updates["price"].(float64); ok { c.Price = int(price) }
  
  cars[id] = c
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  delete(cars, r.PathValue("id"))
  mu.Unlock()
  w.WriteHeader(http.StatusNoContent)
}



