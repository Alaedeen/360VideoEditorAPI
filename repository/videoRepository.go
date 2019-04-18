package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// VideoRepository ...
type VideoRepository interface {
	GetVideos() ([]models.Video, error)
	GetVideo(id uint) (models.Video, error)
	AddVideo( v models.Video) (models.Video, error)
	DeleteVideo()
}

// VideoRepo ...
type VideoRepo struct {
	Db *gorm.DB
}



// GetVideos ...
func (r *VideoRepo) GetVideos() ([]models.Video, error){
	 Videos := []models.Video{}

	r.Db.Find(&Videos)
	
	return Videos, nil
}

// GetVideo ...
func (r *VideoRepo) GetVideo(id uint ) (models.Video, error){
	Video :=models.Video{}
	comments := []models.Comment{}
	var replies []models.Reply
	
	err := r.Db.First(&Video,id).Error
	r.Db.Model(&Video).Related(&comments)
	// for _,comment := range comments {
	// 	fmt.Println(comment)
	// 	r.Db.Model(&comment).Related(&replies)
	// 	comment.Replies=replies
	// }
	for index := 0; index < len(comments); index++ {
		r.Db.Model(&comments[index]).Related(&replies)
		comments[index].Replies=replies
	}
	Video.Comments=comments
	return Video,err
}

// AddVideo ...
func (r *VideoRepo) AddVideo(b models.Video) (models.Video, error){
	Video :=b
	err :=r.Db.Create(&Video).Error
	return Video, err
}

// DeleteVideo ...
func (r *VideoRepo) DeleteVideo(){

}