package models

import (
	"github.com/jinzhu/gorm"
)

// Picture Struct
type Picture struct {
	gorm.Model
	UserID			int 	`json:"userId"`
	Src				string 	`json:"src"`
	Type			string 	`json:"type"`
	Ratio 			float64	`json:"ratio"`
}