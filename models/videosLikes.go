package models

import (
	"github.com/jinzhu/gorm"
)

// VideosLikes Struct
type VideosLikes struct {
	gorm.Model
	UserID	int `json:"idUser"`
	VideoID	int `json:"idVideo"`
}