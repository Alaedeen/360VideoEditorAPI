package models

import (
	"github.com/jinzhu/gorm"
)

// CommentsLikes Struct
type CommentsLikes struct {
	gorm.Model
	UserID		int `json:"idUser"`
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
}