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

// CreateProject ...
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Project models.Project
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Project)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Project.Box=0
	Project.Sphere=0
	Project.Cone=0
	Project.Cylinder=0
	Project.Torus=0
	Project.TorusKnot=0
	Project.Dodecahedron=0
	Project.Tetrahedron=0
	Project.Image=0
	Project.Tag=0
	Project.Video2D=0
	Project.Text=0
	result,err1 := h.Repo.CreateProject(Project)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Title+" CREATED",&response)
	json.NewEncoder(w).Encode(response)
}

// UpdateProject ...
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var Project models.Project
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Project)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	} 
	id, err1 := strconv.Atoi(params[0])
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err2 := h.Repo.UpdateProject(Project,uint(id))
	if err2 !=nil {
		responseFormatter(404,"NOT FOUND",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	responseFormatter(200,"OK",Project,&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteProject ...
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteProject(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","PROJECT DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// GetShapes ...
func (h *ProjectHandler) GetShapes(w http.ResponseWriter, r *http.Request)  {
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
	result,err2 := h.Repo.GetShapes(offset,limit) 
	if err2 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK",result,&response)
	json.NewEncoder(w).Encode(response)
}

// GetFonts ...
func (h *ProjectHandler) GetFonts(w http.ResponseWriter, r *http.Request)  {
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
	result,err2 := h.Repo.GetFonts(offset,limit) 
	if err2 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK",result,&response)
	json.NewEncoder(w).Encode(response)
}

// AddElement ...
func (h *ProjectHandler) AddElement(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Element models.AddedShapes
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Element)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.AddElement(Element)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Type+" ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteElement ...
func (h *ProjectHandler) DeleteElement(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteElement(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","ELEMENT DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddTag ...
func (h *ProjectHandler) AddTag(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Tag models.AddedTags
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Tag)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.AddTag(Tag)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.IDTag+" ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteTag ...
func (h *ProjectHandler) DeleteTag(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteTag(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","TAG DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddTagElement ...
func (h *ProjectHandler) AddTagElement(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var TagElement models.TagElements
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&TagElement)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.AddTagElement(TagElement)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Type+" ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteTagElement ...
func (h *ProjectHandler) DeleteTagElement(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteTagElement(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","TAG ELEMENT DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddPicture ...
func (h *ProjectHandler) AddPicture(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Picture models.Picture
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Picture)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.AddPicture(Picture)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Type+" ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeletePicture ...
func (h *ProjectHandler) DeletePicture(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeletePicture(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","PICTURE DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddProjectVideo ...
func (h *ProjectHandler) AddProjectVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Video2D models.Video2D
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Video2D)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.AddProjectVideo(Video2D)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Type+" ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteProjectVideo ...
func (h *ProjectHandler) DeleteProjectVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteProjectVideo(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","VIDEO DELETED",&response)
	json.NewEncoder(w).Encode(response)
}