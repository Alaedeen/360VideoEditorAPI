package models

import (
	"github.com/jinzhu/gorm"
)

// RepliesLikes Struct
type RepliesLikes struct {
	gorm.Model
	UserID		int `json:"idUser"`
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
	ReplyID 	int `json:"idReply"`
}