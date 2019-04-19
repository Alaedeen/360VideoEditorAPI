package repository

import (
	// "fmt"
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
	AddReply( v models.Reply) (models.Reply, error)
	DeleteReply(id uint)(error)
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
	
	Video.Comments=comments
	for index := 0; index < len(comments); index++ {
		replies = replies[:0]
		r.Db.Model(&comments[index]).Related(&replies)
		Video.Comments[index].Replies=replies
	}
	
	
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

// AddReply ...
func (r *VideoRepo) AddReply(b models.Reply) (models.Reply, error){
	Reply :=b
	err :=r.Db.Create(&Reply).Error
	return Reply, err
}

// DeleteReply ...
func (r *VideoRepo) DeleteReply(id uint)(error){
	Reply := models.Reply{}
	err := r.Db.First(&Reply,id).Error
	if err != nil {
		return err
	}
	
	Reply.ID=id
	err =r.Db.Delete(&Reply).Error
	return err
}