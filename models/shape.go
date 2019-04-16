package models

import (
	"github.com/jinzhu/gorm"
)

// Shape Struct
type Shape struct {
	gorm.Model
	Src				string 	`json:"src"`
	Type			string 	`json:"type"`
}