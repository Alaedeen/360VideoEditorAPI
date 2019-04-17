package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// UserRepository ...
type UserRepository interface {
	GetUsers(u []models.User) ([]models.User, error)
	GetUser(u models.User ,id uint) (models.User, error)
	CreateUser( u models.User) (models.User, error)
	DeleteUser(id uint)(error)
	UpdateUser(u models.User,id uint)(error)
	GetUserVideos(u models.User, v []models.Video ) ([]models.Video, error)
}

// UserRepo ...
type UserRepo struct {
	Db *gorm.DB
}



// GetUsers ...
func (r *UserRepo) GetUsers(u []models.User) ([]models.User, error){
	 Users := u

	err:=r.Db.Find(&Users).Error
	
	return Users, err
}

// GetUser ...
func (r *UserRepo) GetUser(u models.User , id uint) (models.User, error){
	User :=u
	subscriptions := []models.Subscriptions{}
	videosLikes := []models.VideosLikes{}
	videosDislikes := []models.VideosDislikes{}
	commentsLikes := []models.CommentsLikes{}
	commentsDislikes := []models.CommentsDislikes{}
	repliesLikes := []models.RepliesLikes{}
	repliesDislikes := []models.RepliesDislikes{}
	err := r.Db.First(&User,id).Error
	r.Db.Model(&User).Related(&subscriptions)
	r.Db.Model(&User).Related(&videosLikes)
	r.Db.Model(&User).Related(&videosDislikes)
	r.Db.Model(&User).Related(&commentsLikes)
	r.Db.Model(&User).Related(&commentsDislikes)
	r.Db.Model(&User).Related(&repliesLikes)
	r.Db.Model(&User).Related(&repliesDislikes)
	User.Subscriptions=subscriptions
	User.VideosLikes=videosLikes
	User.VideosDislikes=videosDislikes
	User.CommentsLikes=commentsLikes
	User.CommentsDislikes=commentsDislikes
	User.RepliesLikes=repliesLikes
	User.RepliesDislikes=repliesDislikes

	return User,err
}

// CreateUser ...
func (r *UserRepo) CreateUser(u models.User) (models.User, error){
	User :=u
	err :=r.Db.Create(&User).Error
	return User, err
}

// DeleteUser ...
func (r *UserRepo) DeleteUser(id uint)(error){
	user := models.User{}
	err := r.Db.First(&user,id).Error
	if err != nil {
		return err
	}else{
		user.ID=id
		err :=r.Db.Unscoped().Delete(&user).Error
		return err
	}
	
}

// UpdateUser ...
func (r *UserRepo) UpdateUser(u models.User,id uint)(error){
	user := models.User{}
	err := r.Db.First(&user,id).Error
	if err != nil {
		return err
	}
	u.ID=id
	err1 :=r.Db.Model(&user).Updates(&u).Error
	return err1

}

// GetUserVideos ...
func (r *UserRepo) GetUserVideos(u models.User, v []models.Video )([]models.Video, error){
	user := u
	video := v
	err:=r.Db.Model(&user).Related(&video).Error
	return video,err
}