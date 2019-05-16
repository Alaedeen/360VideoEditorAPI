package handlers

import (
	"github.com/Alaedeen/360VideoEditorAPI/helpers"
	"encoding/json"
	"net/http"
	"strconv"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
	"github.com/Alaedeen/360VideoEditorAPI/repository"
)

// VideoHandler ...
type VideoHandler struct {
	Repo repository.VideoRepository
}


// GetVideos ...
func (h *VideoHandler) GetVideos(w http.ResponseWriter, r *http.Request)  {
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
	result,err1,count := h.Repo.GetVideos(offset,limit)
	if err1 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}

	var videos	[]models.VideoResponse
	var video models.VideoResponse
	for _,res := range result {
		
	video = helpers.VideoResponseFormatter(res)
		videos= append(videos,video)
	} 
	responseFormatter(200,"OK",videos,&response)
	responseWithCount.Response=response
	responseWithCount.Count=count
	json.NewEncoder(w).Encode(responseWithCount)
}

// GetVideosByTitle ...
func (h *VideoHandler) GetVideosByTitle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var responseWithCount models.ResponseWithCount
	title := r.URL.Query()["title"][0]
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
	result,err1,count := h.Repo.GetVideosByTitle(title,offset,limit)
	if err1 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		responseWithCount.Response=response
		responseWithCount.Count=0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}

	var videos	[]models.VideoResponse
	var video models.VideoResponse
	for _,res := range result {
		
	video = helpers.VideoResponseFormatter(res)
		videos= append(videos,video)
	} 
	responseFormatter(200,"OK",videos,&response)
	responseWithCount.Response=response
	responseWithCount.Count=count
	json.NewEncoder(w).Encode(responseWithCount)
}

// GetVideo ...
func (h *VideoHandler) GetVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := r.URL.Query() //Get params
	id, err := strconv.Atoi(params["id"][0]) 
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result , err1 := h.Repo.GetVideo(uint(id))
	if (err1 != nil) {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}

	video := helpers.VideoResponseFormatter(result)
	responseFormatter(200,"OK",video,&response)
	json.NewEncoder(w).Encode(response)
}

// AddVideo ...
func (h *VideoHandler) AddVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Video models.Video
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Video)
	if err !=nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err1:= h.Repo.AddVideo(Video)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Title+" Added",&response)
	json.NewEncoder(w).Encode(response)

}


// DeleteVideo ...
func (h *VideoHandler) DeleteVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := r.URL.Query() //Get params
	id, err := strconv.Atoi(params["id"][0]) 
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteVideo(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","Video DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// UpdateVideo ...
func (h *VideoHandler) UpdateVideo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var Video models.Video
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Video)
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
	err2 := h.Repo.UpdateVideo(Video,uint(id))
	if err2 !=nil {
		responseFormatter(404,"NOT FOUND",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	responseFormatter(200,"OK",Video,&response)
	json.NewEncoder(w).Encode(response)
}

// AddComment ...
func (h *VideoHandler) AddComment(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Comment models.Comment
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Comment)
	if err !=nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	comm, err1:= h.Repo.AddComment(Comment)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}

	responseFormatter(201,"CREATED",comm.ID,&response)
	json.NewEncoder(w).Encode(response)
 
}


// DeleteComment ... 
func (h *VideoHandler) DeleteComment(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := r.URL.Query() //Get params
	id, err := strconv.Atoi(params["id"][0]) 
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	} 
	err1 := h.Repo.DeleteComment(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","Comment DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// UpdateComment ...
func (h *VideoHandler) UpdateComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var Comment models.Comment
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Comment)
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
	err2 := h.Repo.UpdateComment(Comment,uint(id))
	if err2 !=nil {
		responseFormatter(404,"NOT FOUND",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	responseFormatter(200,"OK",Comment,&response)
	json.NewEncoder(w).Encode(response)
}

// AddReply ...
func (h *VideoHandler) AddReply(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Reply models.Reply
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Reply)
	if err !=nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	rep, err1:= h.Repo.AddReply(Reply)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",rep.ID,&response)
	json.NewEncoder(w).Encode(response)
 
}


// DeleteReply ... 
func (h *VideoHandler) DeleteReply(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := r.URL.Query() //Get params
	id, err := strconv.Atoi(params["id"][0]) 
	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	} 
	err1 := h.Repo.DeleteReply(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","Reply DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// UpdateReply ...
func (h *VideoHandler) UpdateReply(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var Reply models.Reply
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&Reply)
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
	err2 := h.Repo.UpdateReply(Reply,uint(id))
	if err2 !=nil {
		responseFormatter(404,"NOT FOUND",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	responseFormatter(200,"OK",Reply,&response)
	json.NewEncoder(w).Encode(response)
}