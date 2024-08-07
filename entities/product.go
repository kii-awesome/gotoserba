package entities

import (
	"time"
)

type Product struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Image string `json:"image"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Image string `json:"image"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Image string `json:"image"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductWebResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data   []Product  `json:"data"`
}