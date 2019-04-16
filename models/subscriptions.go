package models

import (
	"github.com/jinzhu/gorm"
)

// Subscriptions Struct
type Subscriptions struct {
	gorm.Model
	UserID	int `json:"idSubscriber"`
	IDSubscribed	int `json:"idSubscribed"`
}