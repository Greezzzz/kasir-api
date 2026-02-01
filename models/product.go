package models

type Products struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CreatedAt  string `json:"created_at"`
	IDCategory string `json:"id_category"`
}
