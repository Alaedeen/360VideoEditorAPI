package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
	"github.com/Alaedeen/360VideoEditorAPI/repository"
)

// ProjectHandler ...
type ProjectHandler struct {
	Repo repository.ProjectRepository
}

func projectResponseFormatter(result models.Project, project *models.ProjectResponse){
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

// GetProjects ...
func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	offset,err0 := strconv.Atoi(r.URL.Query()["offset"][0])
	if err0 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err0.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	limit , err:= strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	id , err1:= strconv.Atoi(r.URL.Query()["id"][0])
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err2 := h.Repo.GetProjects(uint(id),offset,limit) 
	if err2 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK",result,&response)
	json.NewEncoder(w).Encode(response)
}

// GetProject ...
func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])
	if err!= nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.GetProject(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var project models.ProjectResponse
	projectResponseFormatter(result,&project)
	responseFormatter(200,"OK",project,&response)
	json.NewEncoder(w).Encode(response)
}