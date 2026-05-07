package main

import (
	"errors"
	"sync"
)

type Car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}

func (c *Car) Validate() error {
	if c.Brand == "" || c.Model == "" { 
        return errors.New("brand and model cannot be empty") 
    }
	if c.Year < 1885 || c.Year > 2026 { 
        return errors.New("incorrect year") 
    }
	if c.Price <= 0 { 
        return errors.New("price must be greater than zero") 
    }
	return nil
}

var (
	cars  = make(map[string]Car)
	nextID = 1
	mu     sync.Mutex
)