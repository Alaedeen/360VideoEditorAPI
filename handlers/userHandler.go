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

// UserHandler ...
type UserHandler struct {
	Repo repository.UserRepository
}

// Response ...
type Response struct {
	Code int
}

// GetUsers ...
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Users []models.User
	result,_ := h.Repo.GetUsers(Users) 
	json.NewEncoder(w).Encode(result)
}

// GetUser ...
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	var User models.User
	id, _ := strconv.Atoi(params["id"])

	result , err := h.Repo.GetUser(id,User)
	if (err != nil) {
		json.NewEncoder(w).Encode("User does not exist!")
	}else
	{
		json.NewEncoder(w).Encode(result)
	}

	
}

// CreateUser ...
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	err:=json.NewDecoder(r.Body).Decode(&User)
	if err !=nil {
		fmt.Println(err)
	}
	User.ID =uint(rand.Intn(1000000))
	result, err1:= h.Repo.CreateUser(User)
	if err1!=nil{
		json.NewEncoder(w).Encode(err1)
	}else {
		json.NewEncoder(w).Encode(result)
	}	
}

// UpdateUser ...
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Users []models.User
	id, _ := strconv.Atoi(params["id"])
	for index, item := range Users {
		if item.ID == uint(id) {
			Users = append(Users[:index], Users[index+1:]...)
			var User models.User
			_=json.NewDecoder(r.Body).Decode(&User)
			User.ID = uint(id)
			Users = append(Users, User)
			json.NewEncoder(w).Encode(User)
			return
		}
	}
	json.NewEncoder(w).Encode(Users)
}

// DeleteUser ...
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Users []models.User
	id, _ := strconv.Atoi(params["id"])
	for index, item := range Users {
		if item.ID ==  uint(id) {
			Users = append(Users[:index], Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Users)
}