package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// ProjectRepository ...
type ProjectRepository interface {
	GetProjects(id uint, offset int,limit int) ([]models.Project, error)
	
}

// ProjectRepo ...
type ProjectRepo struct {
	Db *gorm.DB
}

// GetProjects ...
func (r *ProjectRepo) GetProjects(id uint, offset int,limit int) ([]models.Project, error){
	user := models.User{}
	user.ID=id
	projects := []models.Project{}
	err:=r.Db.Model(&user).Offset(offset).Limit(limit).Related(&projects).Error
	return projects,err
}