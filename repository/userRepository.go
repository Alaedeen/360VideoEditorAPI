package repository

import (
	"strings"
	"errors"
	"crypto/sha1"
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// UserRepository ...
type UserRepository interface {
	GetUsers(role string, offset int,limit int) ([]models.User, error, int)
	GetUsersByName(name string, role string, offset int,limit int) ([]models.User, error, int)
	GetUser(id uint) (models.User, error)
	GetUserBy(keys []string, values []interface{}) (models.User, error)
	CreateUser( u models.User) (models.User, error)
	DeleteUser(id uint)(error)
	UpdateUser(m map[string]interface {},id uint)(error)
	GetUserVideos(u models.User, offset int,limit int) ([]models.Video, error)
	GetUserPictures(u models.User, offset int,limit int) ([]models.Picture, error)
	GetUserProjectVideos(u models.User, offset int,limit int) ([]models.Video2D, error)
	AddCommentsLikes( c models.CommentsLikes) (error)
	RemoveCommentsLikes(id int)(error)
	AddCommentsDislikes( c models.CommentsDislikes) (error)
	RemoveCommentsDislikes(id int)(error)

	AddRepliesLikes( c models.RepliesLikes) (error)
	RemoveRepliesLikes(id int)(error)
	AddRepliesDislikes( c models.RepliesDislikes) (error)
	RemoveRepliesDislikes(id int)(error)

	AddVideosLikes( c models.VideosLikes) (error)
	RemoveVideosLikes(id int)(error)
	AddVideosDislikes( c models.VideosDislikes) (error)
	RemoveVideosDislikes(id int)(error)

	AddSubscriptions( c models.Subscriptions) (error)
	RemoveSubscriptions(id int)(error)
}

// UserRepo ...
type UserRepo struct {
	Db *gorm.DB
}



// GetUsers ...
func (r *UserRepo) GetUsers(role string,offset int,limit int) ([]models.User, error, int){
	 var Users []models.User
	 var User models.User
	 var count int
	 var err error
	if role=="user" {
		err=r.Db.Where("admin = ?", false).Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ?", false).Count(&count)
	}else if role == "admin"{
		err=r.Db.Where("admin = ? AND super_admin = ?", true, false).Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ? AND super_admin = ?", true, false).Count(&count)
	}else{
		err=r.Db.Where("admin = ? AND super_admin = ?", true, true).Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ? AND super_admin = ?", true, true).Count(&count)
	}
	
	
	return Users, err, count
}

// GetUsersByName ...
func (r *UserRepo) GetUsersByName(name string, role string, offset int,limit int) ([]models.User, error, int){
	 var Users []models.User
	 var User models.User
	 var count int
	 var err error
	if role=="user" {
		err=r.Db.Where("admin = ? AND UPPER(name) LIKE ?", false, "%"+strings.ToUpper(name)+"%").Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ? AND UPPER(name) LIKE ?", false, "%"+strings.ToUpper(name)+"%").Count(&count)
	}else if role == "admin"{
		err=r.Db.Where("admin = ? AND super_admin = ? AND UPPER(name) LIKE ?", true, false, "%"+strings.ToUpper(name)+"%").Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ? AND super_admin = ? AND UPPER(name) LIKE ?", true, false, "%"+strings.ToUpper(name)+"%").Count(&count)
	}else{
		err=r.Db.Where("admin = ? AND super_admin = ? AND UPPER(name) LIKE ?", true, true, "%"+strings.ToUpper(name)+"%").Offset(offset).Limit(limit).Find(&Users).Error
		r.Db.Model(&User).Where("admin = ? AND super_admin = ? AND UPPER(name) LIKE ?", true, true, "%"+strings.ToUpper(name)+"%").Count(&count)
	}
	
	
	return Users, err, count
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
	var password string
	m = make(map[string]interface{})
	for index,value := range keys{
		if value=="password" {
			crypt := sha1.New()
			password= values[index].(string)
			crypt.Write([]byte(password))
			m[value]=crypt.Sum(nil)
		}else {
			m[value] = values[index]
		}
		
	}
	err := r.Db.Where(m).Find(&User).Error
	return User,err
}

// CreateUser ...
func (r *UserRepo) CreateUser(u models.User) (models.User, error){
	User :=u
	var user models.User
	err:= r.Db.Where(map[string]interface{}{"name": u.Name}).Find(&user).Error
	if err == nil {
		return user, errors.New("ERROR: name already used")
	}
	err= r.Db.Where(map[string]interface{}{"email": u.Email}).Find(&user).Error
	if err == nil {
		return user, errors.New("ERROR: mail already used")
	}
	err =r.Db.Create(&User).Error
	return User, err
}

// DeleteUser ...
func (r *UserRepo) DeleteUser(id uint)(error){
	user := models.User{}
	err := r.Db.First(&user,id).Error
	if err != nil {
		return err
	}
	
	user.ID=id
	err =r.Db.Delete(&user).Error
	return err
	
}

// UpdateUser ...
func (r *UserRepo) UpdateUser(m map[string]interface {},id uint)(error){
	user := models.User{}
	err:= r.Db.Where("name = ? AND id != ?",m["name"],id).Find(&user).Error
	if err == nil {
		return errors.New("ERROR: name already used")
	}
	err= r.Db.Where("email = ? AND id != ?",m["email"],id).Find(&user).Error
	if err == nil {
		return errors.New("ERROR: mail already used")
	}
	err = r.Db.First(&user,id).Error
	if err != nil {
		return err
	}
	user.ID=id
	err1 :=r.Db.Model(&user).Updates(m).Error
	return err1

}

// GetUserVideos ...
func (r *UserRepo) GetUserVideos(u models.User,offset int,limit int)([]models.Video, error){
	user := u
	videos := []models.Video{}
	err:=r.Db.Model(&user).Offset(offset).Limit(limit).Related(&videos).Error
	return videos,err
}

// GetUserPictures ...
func (r *UserRepo) GetUserPictures(u models.User, offset int,limit int) ([]models.Picture, error){
	user := u
	pictures := []models.Picture{}
	err:=r.Db.Model(&user).Offset(offset).Limit(limit).Related(&pictures).Error
	return pictures,err
}

// GetUserProjectVideos ...
func (r *UserRepo) GetUserProjectVideos(u models.User, offset int,limit int) ([]models.Video2D, error){
	user := u
	videos := []models.Video2D{}
	err:=r.Db.Model(&user).Offset(offset).Limit(limit).Related(&videos).Error
	return videos,err
}

// AddCommentsLikes ...
func (r *UserRepo)AddCommentsLikes( c models.CommentsLikes) (error){
	CommentsLike :=c
	err :=r.Db.Create(&CommentsLike).Error
	return  err
}

// RemoveCommentsLikes ...
func (r *UserRepo) RemoveCommentsLikes(id int)(error){
	like := models.CommentsLikes{}
	err := r.Db.First(&like,id).Error
	if err != nil {
		return err
	}
	like.CommentID = id
	err =r.Db.Unscoped().Delete(&like).Error
	return err
}

// AddCommentsDislikes ...
func (r *UserRepo)AddCommentsDislikes( c models.CommentsDislikes) (error){
	CommentsDislike :=c
	err :=r.Db.Create(&CommentsDislike).Error
	return  err
}

// RemoveCommentsDislikes ...
func (r *UserRepo) RemoveCommentsDislikes(id int)(error){
	dislike := models.CommentsDislikes{}
	err := r.Db.First(&dislike,id).Error
	if err != nil {
		return err
	}
	dislike.CommentID = id
	err =r.Db.Unscoped().Delete(&dislike).Error
	return err
}

// AddRepliesLikes ...
func (r *UserRepo)AddRepliesLikes( c models.RepliesLikes) (error){
	RepliesLike :=c
	err :=r.Db.Create(&RepliesLike).Error
	return  err
}

// RemoveRepliesLikes ...
func (r *UserRepo) RemoveRepliesLikes(id int)(error){
	like := models.RepliesLikes{}
	err := r.Db.First(&like,id).Error
	if err != nil {
		return err
	}
	like.ReplyID = id
	err =r.Db.Unscoped().Delete(&like).Error
	return err
}

// AddRepliesDislikes ...
func (r *UserRepo)AddRepliesDislikes( c models.RepliesDislikes) (error){
	RepliesDislike :=c
	err :=r.Db.Create(&RepliesDislike).Error
	return  err
}

// RemoveRepliesDislikes ...
func (r *UserRepo) RemoveRepliesDislikes(id int)(error){
	dislike := models.RepliesDislikes{}
	err := r.Db.First(&dislike,id).Error
	if err != nil {
		return err
	}
	dislike.ReplyID = id
	err =r.Db.Unscoped().Delete(&dislike).Error
	return err
}

// AddVideosLikes ...
func (r *UserRepo)AddVideosLikes( c models.VideosLikes) (error){
	VideosLike :=c
	err :=r.Db.Create(&VideosLike).Error
	return  err
}

// RemoveVideosLikes ...
func (r *UserRepo) RemoveVideosLikes(id int)(error){
	like := models.VideosLikes{}
	err := r.Db.First(&like,id).Error
	if err != nil {
		return err
	}
	like.VideoID = id
	err =r.Db.Unscoped().Delete(&like).Error
	return err
}

// AddVideosDislikes ...
func (r *UserRepo)AddVideosDislikes( c models.VideosDislikes) (error){
	VideosDislike :=c
	err :=r.Db.Create(&VideosDislike).Error
	return  err
}

// RemoveVideosDislikes ...
func (r *UserRepo) RemoveVideosDislikes(id int)(error){
	dislike := models.VideosDislikes{}
	err := r.Db.First(&dislike,id).Error
	if err != nil {
		return err
	}
	dislike.VideoID = id
	err =r.Db.Unscoped().Delete(&dislike).Error
	return err
}

// AddSubscriptions ...
func (r *UserRepo)AddSubscriptions( c models.Subscriptions) (error){
	Subscription :=c
	err :=r.Db.Create(&Subscription).Error
	return  err
}

// RemoveSubscriptions ...
func (r *UserRepo) RemoveSubscriptions(id int)(error){
	Subscription := models.Subscriptions{}
	err := r.Db.First(&Subscription,id).Error
	if err != nil {
		return err
	}
	Subscription.IDSubscribed = id
	err =r.Db.Unscoped().Delete(&Subscription).Error
	return err
}