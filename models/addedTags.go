package models

import (
	"github.com/jinzhu/gorm"
)

// AddedTags Struct
type AddedTags struct {
	gorm.Model
	ProjectID		int 			`json:"projectId"`
	IDTag	 		string			`json:"id"`
	Shapes			[]TagElements 	`json:"shapes"`
}