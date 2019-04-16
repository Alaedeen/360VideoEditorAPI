package models

import (
	"github.com/jinzhu/gorm"
)

// AddedShapes Struct
type AddedShapes struct {
	gorm.Model
	ProjectID		int 	`json:"projectId"`
	Image			string 	`json:"image"`
	Type	 		string	`json:"type"`
	IDElement 		string	`json:"id"`
}