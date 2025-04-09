package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primary_key` //;AUTO_INCREMENT"
	Name string `json:"name"`
}
