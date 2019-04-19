package models

import (
	"github.com/jinzhu/gorm"
)

// Test Struct
type Test struct {
	gorm.Model
	UserID			Date
}