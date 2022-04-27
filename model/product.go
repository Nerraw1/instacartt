package model

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stars       string `json:"stars"`
	Quantity    int    `json:"quantity"`
	Location    string `json:"location"`
}
