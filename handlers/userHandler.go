package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	models "github.com/Alaedeen/360VideoEditorAPI/models"
	"github.com/Alaedeen/360VideoEditorAPI/repository"
)

// UserHandler ...
type UserHandler struct {
	Repo repository.UserRepository
}



// GetUsers ...
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Users []models.User
	var response models.Response
	result,err := h.Repo.GetUsers(Users) 
	if err ==nil {
		response.Code = 200
		response.Status= "OK"
		response.Data= result
	}else{
		response.Code = 500
		response.Status= "INTERNAL SERVER ERROR"
		response.Data= nil
	}
	json.NewEncoder(w).Encode(response)
}

// GetUser ...
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var User models.User
	var response models.Response
	id, err := strconv.Atoi(params[0])
	if err!= nil {
		response.Code = 500
		response.Status= "INTERNAL SERVER ERROR"
		response.Data= nil
	}else{
		result,err := h.Repo.GetUser(User , uint(id))
		if err!=nil {
			response.Code = 404
			response.Status= "NOT FOUND"
			response.Data= nil
		}else{
			response.Code = 200
			response.Status= "OK"
			response.Data= result
		}
	}
	json.NewEncoder(w).Encode(response)
}

// CreateUser ...
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	err:=json.NewDecoder(r.Body).Decode(&User)
	fmt.Println(User ,err)
	
}

// UpdateUser ...
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}

// DeleteUser ...
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
}

// GetUserVideos ...
func (h *UserHandler) GetUserVideos(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var User models.User
	var Videos []models.Video
	id,_ := strconv.Atoi(params[0])//error handling
	User.ID = uint(id)
	result,_ := h.Repo.GetUserVideos(User,Videos)
	
	json.NewEncoder(w).Encode(result)

	
	
}