package models

import (
	"github.com/jinzhu/gorm"
)

// Comment Struct
type Comment struct {
	gorm.Model
	UserID			int 	`json:"idUser"`
	VideoID			int 	`json:"videoId"`
	NameUser		string 	`json:"nameUser"`
	ProfilePic		string 	`json:"profilePic"`
	Text			string	`json:"text"`
	Day 			int		`json:"day"`
	Month 			string	`json:"month"`
	Year 			int		`json:"year"`
	Likes 			*int		`json:"likes"`
	Dislikes 		*int		`json:"dislikes"`
	Replies			[]Reply	`json:"replies"`
}