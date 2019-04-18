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
	DeleteVideo(id uint)(error)
	AddComment( v models.Comment) (models.Comment, error)
	DeleteComment(id uint)(error)
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
func (r *VideoRepo) DeleteVideo(id uint)(error){
	video := models.Video{}
	err := r.Db.First(&video,id).Error
	if err != nil {
		return err
	}
	
	video.ID=id
	err =r.Db.Delete(&video).Error
	return err
}

// AddComment ...
func (r *VideoRepo) AddComment(b models.Comment) (models.Comment, error){
	Comment :=b
	err :=r.Db.Create(&Comment).Error
	return Comment, err
}

// DeleteComment ...
func (r *VideoRepo) DeleteComment(id uint)(error){
	Comment := models.Comment{}
	err := r.Db.First(&Comment,id).Error
	if err != nil {
		return err
	}
	
	Comment.ID=id
	err =r.Db.Delete(&Comment).Error
	return err
}