package models

type Dessert struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
