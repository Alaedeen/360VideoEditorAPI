package models

import (
	"github.com/jinzhu/gorm"
)

// UploadRequest Struct
type UploadRequest struct {
	gorm.Model
	UserID			int 		`json:"userId"`
	Title			string 		`json:"title"`
	UploadDay 		int			`json:"uploadDay"`
	UploadMonth 	string		`json:"uploadMonth"`
	UploadYear 		int			`json:"uploadYear"`
	Thumbnail 		string		`json:"thumbnail"`
	Src 			string		`json:"src"`
	AFrame 			string		`json:"aFrame"`
}