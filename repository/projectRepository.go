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
	UpdateProject(p models.Project,id uint)(error)
	DeleteProject(id uint)(error)
	GetShapes(offset int,limit int) ([]models.Shape, error)
	GetFonts(offset int,limit int) ([]models.Font, error)
	AddElement( e models.AddedShapes) (models.AddedShapes, error)
	DeleteElement(id uint)(error)
	AddTag( e models.AddedTags) (models.AddedTags, error)
	DeleteTag(id uint)(error)
	AddTagElement( e models.TagElements) (models.TagElements, error)
	DeleteTagElement(id uint)(error)
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

// UpdateProject ...
func (r *ProjectRepo) UpdateProject(p models.Project,id uint)(error){
	project := models.Project{}
	err := r.Db.First(&project,id).Error
	if err != nil {
		return err
	}
	p.ID=id
	err1 :=r.Db.Model(&project).Updates(&p).Error
	return err1
}

// DeleteProject ...
func (r *ProjectRepo) DeleteProject(id uint)(error) {
	project := models.Project{}
	err := r.Db.First(&project,id).Error
	if err != nil {
		return err
	}
	project.ID=id
	err =r.Db.Delete(&project).Error
	return err
}

// GetShapes ...
func (r *ProjectRepo) GetShapes(offset int,limit int) ([]models.Shape, error){
	Shapes := []models.Shape{}

	err :=r.Db.Offset(offset).Limit(limit).Find(&Shapes).Error
	
	return Shapes, err
}

// GetFonts ...
func (r *ProjectRepo) GetFonts(offset int,limit int) ([]models.Font, error){
	Fonts := []models.Font{}

	err :=r.Db.Offset(offset).Limit(limit).Find(&Fonts).Error
	
	return Fonts, err
}

// AddElement ...
func (r *ProjectRepo) AddElement( e models.AddedShapes) (models.AddedShapes, error){
	Element :=e
	err :=r.Db.Create(&Element).Error
	return Element, err
}

// DeleteElement ...
func (r *ProjectRepo) DeleteElement(id uint)(error) {
	element := models.AddedShapes{}
	err := r.Db.First(&element,id).Error
	if err != nil {
		return err
	}
	element.ID=id
	err =r.Db.Delete(&element).Error
	return err
}

// AddTag ...
func (r *ProjectRepo) AddTag( e models.AddedTags) (models.AddedTags, error){
	Tag :=e
	err :=r.Db.Create(&Tag).Error
	return Tag, err
}

// DeleteTag ...
func (r *ProjectRepo) DeleteTag(id uint)(error) {
	Tag := models.AddedTags{}
	err := r.Db.First(&Tag,id).Error
	if err != nil {
		return err
	}
	Tag.ID=id
	err =r.Db.Delete(&Tag).Error
	return err
}

// AddTagElement ...
func (r *ProjectRepo) AddTagElement( e models.TagElements) (models.TagElements, error){
	TagElement :=e
	err :=r.Db.Create(&TagElement).Error
	return TagElement, err
}

// DeleteTagElement ...
func (r *ProjectRepo) DeleteTagElement(id uint)(error) {
	TagElement := models.TagElements{}
	err := r.Db.First(&TagElement,id).Error
	if err != nil {
		return err
	}
	TagElement.ID=id
	err =r.Db.Delete(&TagElement).Error
	return err
}