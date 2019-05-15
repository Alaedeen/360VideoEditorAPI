package repository

import (
	"strings"
	// "fmt"
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// VideoRepository ...
type VideoRepository interface {
	GetVideos(offset int,limit int) ([]models.Video, error, int)
	GetVideosByTitle(title string, offset int,limit int) ([]models.Video, error, int)
	GetVideo(id uint) (models.Video, error)
	AddVideo( v models.Video) (models.Video, error)
	DeleteVideo(id uint)(error)
	UpdateVideo(u models.Video,id uint)(error)
	AddComment( v models.Comment) (models.Comment, error)
	DeleteComment(id uint)(error)
	UpdateComment(u models.Comment,id uint)(error)
	AddReply( v models.Reply) (models.Reply, error)
	DeleteReply(id uint)(error)
	UpdateReply(u models.Reply,id uint)(error)
}

// VideoRepo ...
type VideoRepo struct {
	Db *gorm.DB
}



// GetVideos ...
func (r *VideoRepo) GetVideos(offset int,limit int) ([]models.Video, error, int){
	Videos := []models.Video{}
	video := models.Video{}
	var count int
	err :=r.Db.Offset(offset).Limit(limit).Find(&Videos).Error
	r.Db.Model(&video).Count(&count)
	return Videos, err, count
}

// GetVideosByTitle ...
func (r *VideoRepo) GetVideosByTitle(title string, offset int,limit int) ([]models.Video, error, int){
	Videos := []models.Video{}
	video := models.Video{}
	var count int
	err :=r.Db.Where("UPPER(title) LIKE ? ","%"+strings.ToUpper(title)+"%" ).Offset(offset).Limit(limit).Find(&Videos).Error
	r.Db.Model(&video).Where("UPPER(title) LIKE ? ","%"+strings.ToUpper(title)+"%" ).Count(&count)
	return Videos, err, count
}

// GetVideo ...
func (r *VideoRepo) GetVideo(id uint ) (models.Video, error){
	Video :=models.Video{}
	comments := []models.Comment{}
	var replies []models.Reply
	
	err := r.Db.First(&Video,id).Error
	r.Db.Model(&Video).Order("id desc").Related(&comments)
	
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

// UpdateVideo ...
func (r *VideoRepo) UpdateVideo(u models.Video,id uint)(error){
	Video := models.Video{}
	err := r.Db.First(&Video,id).Error
	if err != nil {
		return err
	}
	u.ID=id
	err1 :=r.Db.Model(&Video).Updates(&u).Error
	return err1

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

// UpdateComment ...
func (r *VideoRepo) UpdateComment(u models.Comment,id uint)(error){
	Comment := models.Comment{}
	err := r.Db.First(&Comment,id).Error
	if err != nil {
		return err
	}
	u.ID=id
	err1 :=r.Db.Model(&Comment).Updates(&u).Error
	return err1

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

// UpdateReply ...
func (r *VideoRepo) UpdateReply(u models.Reply,id uint)(error){
	Reply := models.Reply{}
	err := r.Db.First(&Reply,id).Error
	if err != nil {
		return err
	}
	u.ID=id
	err1 :=r.Db.Model(&Reply).Updates(&u).Error
	return err1

}