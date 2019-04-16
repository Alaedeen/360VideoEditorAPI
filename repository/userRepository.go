package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// UserRepository ...
type UserRepository interface {
	GetUsers(u []models.User) ([]models.User, error)
	GetUser(id int, u models.User) (models.User, error)
	CreateUser( u models.User) (models.User, error)
	DeleteUser()
	UpdateUser()
	GetUserVideos(u models.User, v []models.Video ) ([]models.Video, error)
}

// UserRepo ...
type UserRepo struct {
	Db *gorm.DB
}

// Result Struct)
type Result struct {
	ID 			uint
	Isbn		string `json:"isbn"`
	Title		string `json:"title"`
}

// GetUsers ...
func (r *UserRepo) GetUsers(u []models.User) ([]models.User, error){
	 Users := u

	r.Db.Find(&Users)
	
	return Users, nil
}

// GetUser ...
func (r *UserRepo) GetUser(id int, b models.User) (models.User, error){
	User :=b

	err := r.Db.Where("id = ?", uint(id)).First(&User).Error

	return User,err
}

// CreateUser ...
func (r *UserRepo) CreateUser(b models.User) (models.User, error){
	User :=b
	err :=r.Db.Create(&User).Error
	return User, err
}

// DeleteUser ...
func (r *UserRepo) DeleteUser(){

}

// UpdateUser ...
func (r *UserRepo) UpdateUser(){

}

// GetUserVideos ...
func (r *UserRepo) GetUserVideos(u models.User, v []models.Video )([]models.Video, error){
	user := u
	video := v
	err:=r.Db.Model(&user).Related(&video).Error
	return video,err
}