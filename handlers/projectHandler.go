package handlers

import (
	"github.com/Alaedeen/360VideoEditorAPI/helpers"
	"io/ioutil"
	"time"
	"os"
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


// GetProjects ...
func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var responseWithCount models.ResponseWithCount
	offset,err0 := strconv.Atoi(r.URL.Query()["offset"][0])
	if err0 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err0.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	limit , err:= strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	id , err1:= strconv.Atoi(r.URL.Query()["id"][0])
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	result,count,err2 := h.Repo.GetProjects(uint(id),offset,limit) 
	if err2 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	responseFormatter(200,"OK",result,&response)
	responseWithCount.Response=response
	responseWithCount.Count=count
	json.NewEncoder(w).Encode(responseWithCount)
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
	helpers.ProjectResponseFormatter(result,&project)
	responseFormatter(200,"OK",project,&response)
	json.NewEncoder(w).Encode(response)
}

// LoadProjectScript ...
func (h *ProjectHandler) LoadProjectScript(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	fileName := r.URL.Query()["file"][0]
	var response models.Response
	var project models.Script
	script, err := ioutil.ReadFile("assets/project/videos/script/"+fileName)
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err = json.Unmarshal([]byte(script), &project)
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK",project ,&response)
	json.NewEncoder(w).Encode(response)

	
}

// CreateProject ...
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Project models.Project
	var response models.Response
	
	dt := time.Now().Format("01-02-2006 15:04:05")	
	
	r.ParseMultipartForm(10 << 20)
	//upload the video 
	file, handler, err5 := r.FormFile("videoProject")
	var fileType string
    if err5 != nil {
        responseFormatter(400,"BAD REQUEST",err5.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()
	fileType = handler.Header["Content-Type"][0]
	fileType= fileType[6:]
	videoFile, err3 := ioutil.TempFile("assets/project/videos", "video_*_"+dt+"." + fileType)
	if err3 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err3.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer videoFile.Close()
	fileBytes, err4 := ioutil.ReadAll(file)
	if err4 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err4.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	videoFile.Write(fileBytes)

	// //upload the video thumbnail
	// file1, handler1, err2 := r.FormFile("thumbnail")
	// var fileType1 string
    // if err2 != nil {
    //     responseFormatter(400,"BAD REQUEST",err2.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// defer file1.Close()
	// fileType1 = handler1.Header["Content-Type"][0]
	// fileType1= fileType1[6:]
	// thumbnail, err := ioutil.TempFile("assets/project/videos/thumbnails", "thumbnail_*_"+dt+"." + fileType1)
	// if err != nil {
	// 	responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// defer thumbnail.Close()
	// fileBytes1, err0 := ioutil.ReadAll(file1)
	// if err0 != nil {
	// 	responseFormatter(500,"INTERNAL SERVER ERROR",err0.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// thumbnail.Write(fileBytes1)
	

	aFrame, err := os.Create("assets/project/videos/script/aframe_"+r.Form["userId"][0]+"_"+dt+".json")
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	aFrame.Close()
	var err1 error
	Project.UserID,err1=strconv.Atoi(r.Form["userId"][0])
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Project.Title=r.Form["title"][0]
	Project.AFrame="aframe_"+r.Form["userId"][0]+"_"+dt+".json"
	Project.Video= videoFile.Name()[22:]	
	// Project.Thumbnail= thumbnail.Name()[33:]
	Project.Thumbnail= "thumbnail.jpg"	
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

// SaveProject ...
func (h *ProjectHandler) SaveProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/json")
	var response models.Response
	name := r.URL.Query()["name"][0]
	err := os.Remove("assets/project/videos/script/"+name)
    if err != nil {
        responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	aFrame, err := os.Create("assets/project/videos/script/"+name)
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer aFrame.Close()
	jsonData, err := ioutil.ReadAll(r.Body) 

 	if err != nil {
 		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
 	}
	_,err = aFrame.WriteString(string(jsonData))
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}

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
	responseFormatter(201,"CREATED",result.ID,&response)
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
	responseFormatter(201,"CREATED",result.ID,&response)
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
	responseFormatter(201,"CREATED",result.ID,&response)
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

	dt := time.Now().Format("01-02-2006 15:04:05")	
	
	r.ParseMultipartForm(10 << 20)
	//upload the picture 
	file, handler, err := r.FormFile("picture")
	var fileType string
    if err != nil {
        responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()
	fileType = handler.Header["Content-Type"][0]
	fileType= fileType[6:]
	pictureFile, err3 := ioutil.TempFile("assets/project/projectPictures", "picture_*_"+dt+"." + fileType)
	if err3 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err3.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer pictureFile.Close()
	fileBytes, err4 := ioutil.ReadAll(file)
	if err4 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err4.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	pictureFile.Write(fileBytes)

	Picture.UserID,err=strconv.Atoi(r.Form["userId"][0])
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Picture.Type=r.Form["type"][0]
	Picture.Ratio,err=strconv.ParseFloat(r.Form["ratio"][0],64)
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Picture.Src= "http://localhost:8000/assets/project/projectPictures/" +  pictureFile.Name()[31:]

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
	
	dt := time.Now().Format("01-02-2006 15:04:05")	
	
	r.ParseMultipartForm(10 << 20)
	//upload the video 
	file, handler, err := r.FormFile("projectVideo")
	var fileType string
    if err != nil {
        responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()
	fileType = handler.Header["Content-Type"][0]
	fileType= fileType[6:]
	videoFile, err3 := ioutil.TempFile("assets/project/projectVideos", "projectVideo_*_"+dt+"." + fileType)
	if err3 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err3.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer videoFile.Close()
	fileBytes, err4 := ioutil.ReadAll(file)
	if err4 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err4.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	videoFile.Write(fileBytes)

	// //upload the video thumbnail
	// file1, handler1, err2 := r.FormFile("thumbnail")
	// var fileType1 string
    // if err2 != nil {
    //     responseFormatter(400,"BAD REQUEST",err2.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// defer file1.Close()
	// fileType1 = handler1.Header["Content-Type"][0]
	// fileType1= fileType1[6:]
	// thumbnail, err := ioutil.TempFile("assets/project/projectVideos/thumbnails", "projVidThumb_*_"+dt+"." + fileType1)
	// if err != nil {
	// 	responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// defer thumbnail.Close()
	// fileBytes1, err0 := ioutil.ReadAll(file1)
	// if err0 != nil {
	// 	responseFormatter(500,"INTERNAL SERVER ERROR",err0.Error(),&response)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// thumbnail.Write(fileBytes1)

	Video2D.UserID,err=strconv.Atoi(r.Form["userId"][0])
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Video2D.Type=r.Form["type"][0]
	Video2D.Ratio,err=strconv.ParseFloat(r.Form["ratio"][0],64)
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	Video2D.Src= "http://localhost:8000/assets/project/projectVideos/"+videoFile.Name()[29:]
	// Video2D.Thumbnail= thumbnail.Name()[40:]
	Video2D.Thumbnail=  "http://localhost:8000/assets/project/projectVideos/thumbnails/thumbnail.png"
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

// GetUploadRequests ...
func (h *ProjectHandler) GetUploadRequests(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var responseWithCount models.ResponseWithCount
	offset,err0 := strconv.Atoi(r.URL.Query()["offset"][0])
	if err0 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err0.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	limit , err:= strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	result,count,err2 := h.Repo.GetUploadRequests(offset,limit) 
	if err2 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	responseFormatter(200,"OK",result,&response)
	responseWithCount.Response=response
	responseWithCount.Count=count
	json.NewEncoder(w).Encode(responseWithCount)
}

// DeleteUploadRequest ...
func (h *ProjectHandler) DeleteUploadRequest(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := r.URL.Query() //Get params
	id, err := strconv.Atoi(params["id"][0]) 
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteUploadRequest(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","REQUEST DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddUploadRequest ...
func (h *ProjectHandler) AddUploadRequest(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var UploadRequest models.UploadRequest
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&UploadRequest)
	if err !=nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err1:= h.Repo.AddUploadRequest(UploadRequest)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Title+" Added",&response)
	json.NewEncoder(w).Encode(response)

}