package models

import (
	"github.com/jinzhu/gorm"
)

// Font Struct
type Font struct {
	gorm.Model
	Src				string 	`json:"src"`
	Type			string 	`json:"type"`
	FontType		string 	`json:"font"`
}