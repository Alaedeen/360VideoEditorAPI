package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// VideoRepository ...
type VideoRepository interface {
	GetVideos(v []models.Video) ([]models.Video, error)
	GetVideo(id int, v models.Video) (models.Video, error)
	CreateVideo( v models.Video) (models.Video, error)
	DeleteVideo()
}

// VideoRepo ...
type VideoRepo struct {
	Db *gorm.DB
}



// GetVideos ...
func (r *VideoRepo) GetVideos(v []models.Video) ([]models.Video, error){
	 Videos := v

	r.Db.Find(&Videos)
	
	return Videos, nil
}

// GetVideo ...
func (r *VideoRepo) GetVideo(id int, b models.Video) (models.Video, error){
	Video :=b

	err := r.Db.Where("id = ?", uint(id)).First(&Video).Error

	return Video,err
}

// CreateVideo ...
func (r *VideoRepo) CreateVideo(b models.Video) (models.Video, error){
	Video :=b
	err :=r.Db.Create(&Video).Error
	return Video, err
}

// DeleteVideo ...
func (r *VideoRepo) DeleteVideo(){

}