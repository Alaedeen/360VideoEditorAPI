package models

import (
	"github.com/jinzhu/gorm"
)

// TagElements Struct
type TagElements struct {
	gorm.Model
	AddedTagsID		int 	`json:"tagId"`
	Image			string 	`json:"image"`
	Type	 		string	`json:"type"`
	IDElement 		string	`json:"id"`
}