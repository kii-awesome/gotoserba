package entities

type Category struct {
	ID   uint `json:"id"`
	Name string `json:"name"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   uint `json:"id"`
	Name string `json:"name"`
}

type CategoryWebResponse struct {
	Code   int    `json:"id"`
	Status string `json:"status"`
	Data   []Category  `json:"data"`
}
