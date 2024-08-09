package models

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
	CreatedAt time.Time `gorm:"autoCreateTime;<-create" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

type ProductWebResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data   []Product  `json:"data"`
}