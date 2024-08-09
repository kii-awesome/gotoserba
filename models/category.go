package models

type Category struct {
	ID   uint `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Name string `gorm:"name" json:"name"`
	Description string `gorm:"description" json:"description"`
}

type CategoryWebResponse struct {
	Code   int    `json:"id"`
	Status string `json:"status"`
	Data   []Category  `json:"data"`
}
