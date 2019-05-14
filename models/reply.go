package models

import (
	"github.com/jinzhu/gorm"
)

// Reply Struct
type Reply struct {
	gorm.Model
	UserID			int 	`json:"idUser"`
	CommentID		int 	`json:"commentId"`
	NameUser		string 	`json:"nameUser"`
	ProfilePic		string 	`json:"profilePic"`
	Text			string	`json:"text"`
	Day 			int		`json:"day"`
	Month 			string	`json:"month"`
	Year 			int		`json:"year"`
	Likes 			*int		`json:"likes"`
	Dislikes 		*int		`json:"dislikes"`
}