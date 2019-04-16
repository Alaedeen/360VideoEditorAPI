package models

import (
	"github.com/jinzhu/gorm"
)

// CommentsDislikes Struct
type CommentsDislikes struct {
	gorm.Model
	UserID		int `json:"idUser"`
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
}