package models

import (
	"github.com/jinzhu/gorm"
)

// VideosDislikes Struct
type VideosDislikes struct {
	gorm.Model
	UserID	int `json:"idUser"`
	VideoID	int `json:"idVideo"`
}