package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// UserRepository ...
type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id uint) (models.User, error)
	GetUserBy(keys []string, values []interface{}) (models.User, error)
	CreateUser( u models.User) (models.User, error)
	DeleteUser(id uint)(error)
	UpdateUser(u models.User,id uint)(error)
	GetUserVideos(u models.User) ([]models.Video, error)
}

// UserRepo ...
type UserRepo struct {
	Db *gorm.DB
}



// GetUsers ...
func (r *UserRepo) GetUsers() ([]models.User, error){
	 var Users []models.User

	err:=r.Db.Model(&models.User{}).Find(&Users).Error
	
	return Users, err
}

// GetUser ...
func (r *UserRepo) GetUser( id uint) (models.User, error){
	var User models.User
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

// GetUserBy ...
func (r *UserRepo) GetUserBy(keys []string, values []interface{}) (models.User, error){
	var User models.User
	var m map[string]interface{}
	m = make(map[string]interface{})
	for index,value := range keys{
		m[value] = values[index]
	}
	err := r.Db.Where(m).Find(&User).Error
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
		err :=r.Db.Delete(&user).Error
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
func (r *UserRepo) GetUserVideos(u models.User)([]models.Video, error){
	user := u
	videos := []models.Video{}
	err:=r.Db.Model(&user).Related(&videos).Error
	return videos,err
}