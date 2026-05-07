package main

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "testing"
)

func resetDB() {
  cars = make(map[string]Car)
  nextID = 1
}

func TestCreateCar(t *testing.T) {
  resetDB()
  body := []byte(`{"brand":"Toyota", "model":"Camry", "year":2020, "price":25000}`)
  req := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewBuffer(body))
  w := httptest.NewRecorder()

  createCar(w, req)

  if w.Code != http.StatusCreated {
    t.Errorf("waiting for 201, got %d", w.Code)
  }
}

func TestCreateCar_ValidationError(t *testing.T) {
  resetDB()
  body := []byte(`{"brand":"Toyota", "model":"Camry", "year":1800, "price":25000}`)
  req := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewBuffer(body))
  w := httptest.NewRecorder()

  createCar(w, req)

  if w.Code != http.StatusBadRequest {
    t.Errorf("waiting for 400, got %d", w.Code)
  }
}

func TestDeleteCar(t *testing.T) {
  resetDB()
  cars["1"] = Car{ID: "1", Brand: "Mazda", Model: "3", Year: 2015, Price: 10000}

  req := httptest.NewRequest(http.MethodDelete, "/cars/1", nil)
  req.SetPathValue("id", "1")
  w := httptest.NewRecorder()

  deleteCar(w, req)

  if w.Code != http.StatusNoContent {
    t.Errorf("waiting for 204, got %d", w.Code)
  }
}