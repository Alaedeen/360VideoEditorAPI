package handlers

import (
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
	result,err := h.Repo.GetVideos() 
	if err !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK",result,&response)
	json.NewEncoder(w).Encode(response)
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
	responseFormatter(200,"OK",result,&response)
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