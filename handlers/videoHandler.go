package handlers

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
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
	err:=json.NewDecoder(r.Body).Decode(&Video)
	if err !=nil {
		fmt.Println(err)
	}
	Video.ID =uint(rand.Intn(1000000))
	result, err1:= h.Repo.CreateVideo(Video)
	if err1!=nil{
		json.NewEncoder(w).Encode(err1)
	}else {
		json.NewEncoder(w).Encode(result)
	}	
}


// DeleteVideo ...
func (h *VideoHandler) DeleteVideo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Videos []models.Video
	id, _ := strconv.Atoi(params["id"])
	for index, item := range Videos {
		if item.ID ==  uint(id) {
			Videos = append(Videos[:index], Videos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Videos)
}