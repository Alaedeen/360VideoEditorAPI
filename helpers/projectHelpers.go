package helpers

import (
	"github.com/Alaedeen/360VideoEditorAPI/models"
)

// ProjectResponseFormatter func
func ProjectResponseFormatter(result models.Project, project *models.ProjectResponse){
	project.ID= result.ID
	project.UserID=result.UserID
	project.Title=result.Title
	project.Thumbnail=result.Thumbnail
	project.AFrame=result.AFrame
	project.Video=result.Video
	project.Tag=result.Tag
	project.ShapesList=result.ShapesList
	project.TagsList=result.TagsList
	project.Shapes=make(map[string]int)
	project.Shapes["box"]=result.Box
	project.Shapes["sphere"]=result.Sphere
	project.Shapes["cylinder"]=result.Cylinder
	project.Shapes["torus"]=result.Torus
	project.Shapes["torus-knot"]=result.TorusKnot
	project.Shapes["dodecahedron"]=result.Dodecahedron
	project.Shapes["tetrahedron"]=result.Tetrahedron
	project.Shapes["image"]=result.Image
	project.Shapes["video"]=result.Video2D
	project.Shapes["text"]=result.Text
	project.Shapes["cone"]=result.Cone
}