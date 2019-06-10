package models

import (
	"github.com/jinzhu/gorm"
)

// Video Struct
type Video struct {
	gorm.Model
	UserID			int 		`json:"userId"`
	Title			string 		`json:"title"`
	UploadDay 		int			`json:"uploadDay"`
	UploadMonth 	string		`json:"uploadMonth"`
	UploadYear 		int			`json:"uploadYear"`
	Thumbnail 		string		`json:"thumbnail"`
	Src 			string		`json:"src"`
	AFrame 			string		`json:"aFrame"`
	Likes 			*int		`json:"likes"`
	Dislikes 		*int		`json:"dislikes"`
	Views 			*int			`json:"views"`
	Comments		[]Comment	`json:"comments"`
}