package repository

import (
	"github.com/jinzhu/gorm"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
)

// ProjectRepository ...
type ProjectRepository interface {
	GetProjects(id uint, offset int,limit int) ([]models.Project, error)
	GetProject(id uint) (models.Project, error)
	CreateProject( p models.Project) (models.Project, error)
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

// GetProject ...
func (r *ProjectRepo) GetProject(id uint) (models.Project, error){
	var project models.Project
	var shapesList []models.AddedShapes
	var tagsList []models.AddedTags
	var tagElements []models.TagElements
	err:= r.Db.First(&project,id).Error
	r.Db.Model(&project).Related(&shapesList)
	r.Db.Model(&project).Related(&tagsList)
	project.ShapesList=shapesList
	project.TagsList=tagsList
	for index, tag := range project.TagsList {
		tagElements= tagElements[:0]
		r.Db.Model(&tag).Related(&tagElements)
		project.TagsList[index].Shapes=tagElements
	}
	return project,err
}

// CreateProject ...
func (r *ProjectRepo) CreateProject( p models.Project) (models.Project, error){
	Project :=p
	err :=r.Db.Create(&Project).Error
	return Project, err
}