package models

import (
	"github.com/jinzhu/gorm"
)

// Video2D Struct
type Video2D struct {
	gorm.Model
	UserID			int 	`json:"userId"`
	Src				string 	`json:"src"`
	Type			string 	`json:"type"`
	Thumbnail 		string	`json:"thumbnail"`
	Ratio 			float64	`json:"ratio"`
}