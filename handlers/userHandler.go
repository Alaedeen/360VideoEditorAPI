package handlers

import (
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

func responseFormatter (code int, status string, data interface{}, response *models.Response) {
	response.Code = code
	response.Status = status
	response.Data=data
}

func userResponseFormatter(result models.User, user *models.UserResponse)  {
	user.ID=result.ID
	user.Name=result.Name
	user.Email=result.Email
	user.Password=result.Password
	user.Roles = append(user.Roles,"user")
	if result.Admin {
		user.Roles = append(user.Roles,"admin")
	}
	if result.SuperAdmin {
		user.Roles = append(user.Roles,"super admin")
	}
	user.DateOfBirth.Day=result.BirthDay
	user.DateOfBirth.Month=result.BirthMonth
	user.DateOfBirth.Year=result.BirthYear
	user.Country=result.Country
	user.Description=result.Description
	user.ProfilePic=result.ProfilePic
	user.Joined.Day=result.JoiningDay
	user.Joined.Month=result.JoiningMonth
	user.Joined.Year=result.JoiningYear
	user.Subscribers=result.Subscribers
	user.Subscriptions=result.Subscriptions
	user.VideosLikes=result.VideosLikes
	user.VideosDislikes=result.VideosDislikes
	user.CommentsLikes=result.CommentsLikes
	user.CommentsDislikes=result.CommentsDislikes
	user.RepliesLikes=result.RepliesLikes
	user.RepliesDislikes=result.RepliesDislikes
}



// GetUsers ...
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	role := r.URL.Query()["role"][0]
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
	result,err1 := h.Repo.GetUsers(role,offset,limit) 
	if err1 !=nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var users []models.UserResponse
	var user models.UserResponse
	for _,res := range result {
		user.Roles= user.Roles[:0]
		userResponseFormatter(res,&user)
		users= append(users,user)
	} 
	responseFormatter(200,"OK",users,&response)
	json.NewEncoder(w).Encode(response)
}

// GetUser ...
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])
	if err!= nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.GetUser(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var user models.UserResponse
	userResponseFormatter(result,&user)
	responseFormatter(200,"OK",user,&response)
	json.NewEncoder(w).Encode(response)
}

// GetUserBy ...
func (h *UserHandler) GetUserBy(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= r.URL.Query()
	var keys []string
	var values []interface{}
	var response models.Response
	for key,value := range params {
		keys = append(keys,key)
		val , err := strconv.Atoi(value[0])
		if err != nil {
			values = append(values, value[0])
		}else{
			values = append(values, uint(val))
		}
	}
	result,err:= h.Repo.GetUserBy(keys,values)
	if err != nil {
		responseFormatter(404,"NOT FOUND",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var user models.UserResponse
	userResponseFormatter(result,&user)
	responseFormatter(200,"OK",user,&response)
	json.NewEncoder(w).Encode(response)
}

// CreateUser ...
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result,err1 := h.Repo.CreateUser(User)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED",result.Name+" CREATED",&response)
	json.NewEncoder(w).Encode(response)
}

// DeleteUser ...
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeleteUser(uint(id))
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","USER DELETED",&response)
	json.NewEncoder(w).Encode(response)
}

// UpdateUser ...
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var User models.User
	var response models.Response
	err:=json.NewDecoder(r.Body).Decode(&User)
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
	err2 := h.Repo.UpdateUser(User,uint(id))
	if err2 !=nil {
		responseFormatter(404,"NOT FOUND",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	responseFormatter(200,"OK",User,&response)
	json.NewEncoder(w).Encode(response)
}

// GetUserVideos ...
func (h *UserHandler) GetUserVideos(w http.ResponseWriter, r *http.Request)  {
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
	params := r.URL.Query()["id"]
	var User models.User
	id,err2 := strconv.Atoi(params[0])//error handling
	if err2 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err2.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	User.ID = uint(id)
	result,err1 := h.Repo.GetUserVideos(User,offset,limit)
	if err1 !=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}

	var videos	[]models.VideoResponse
	var video models.VideoResponse
	for _,res := range result {
		
	video = videoResponseFormatter(res)
		videos= append(videos,video)
	} 
	responseFormatter(200,"OK",videos,&response)
	json.NewEncoder(w).Encode(response)
}

// AddCommentsLikes ...
func (h *UserHandler) AddCommentsLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var CommentsLikes models.CommentsLikes
	err:=json.NewDecoder(r.Body).Decode(&CommentsLikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddCommentsLikes(CommentsLikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveCommentsLikes ...
func (h *UserHandler) RemoveCommentsLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveCommentsLikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","COMMENT LIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddCommentsDislikes ...
func (h *UserHandler) AddCommentsDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var CommentsDislikes models.CommentsDislikes
	err:=json.NewDecoder(r.Body).Decode(&CommentsDislikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddCommentsDislikes(CommentsDislikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveCommentsDislikes ...
func (h *UserHandler) RemoveCommentsDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveCommentsDislikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","COMMENT DISLIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddRepliesLikes ...
func (h *UserHandler) AddRepliesLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var RepliesLikes models.RepliesLikes
	err:=json.NewDecoder(r.Body).Decode(&RepliesLikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddRepliesLikes(RepliesLikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveRepliesLikes ...
func (h *UserHandler) RemoveRepliesLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveRepliesLikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","REPLY LIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddRepliesDislikes ...
func (h *UserHandler) AddRepliesDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var RepliesDislikes models.RepliesDislikes
	err:=json.NewDecoder(r.Body).Decode(&RepliesDislikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddRepliesDislikes(RepliesDislikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveRepliesDislikes ...
func (h *UserHandler) RemoveRepliesDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveRepliesDislikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","REPLY DISLIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddVideosLikes ...
func (h *UserHandler) AddVideosLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var VideosLikes models.VideosLikes
	err:=json.NewDecoder(r.Body).Decode(&VideosLikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddVideosLikes(VideosLikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveVideosLikes ...
func (h *UserHandler) RemoveVideosLikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveVideosLikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","VIDEO LIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddVideosDislikes ...
func (h *UserHandler) AddVideosDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var VideosDislikes models.VideosDislikes
	err:=json.NewDecoder(r.Body).Decode(&VideosDislikes)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddVideosDislikes(VideosDislikes)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveVideosDislikes ...
func (h *UserHandler) RemoveVideosDislikes(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveVideosDislikes(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","VIDEO DISLIKE REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}

// AddSubscriptions ...
func (h *UserHandler) AddSubscriptions(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var Subscriptions models.Subscriptions
	err:=json.NewDecoder(r.Body).Decode(&Subscriptions)
	if err != nil {
		responseFormatter(400,"BAD REQUEST",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	err1 := h.Repo.AddSubscriptions(Subscriptions)
	if err1 != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","ADDED",&response)
	json.NewEncoder(w).Encode(response)
}

// RemoveSubscriptions ...
func (h *UserHandler) RemoveSubscriptions(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500,"INTERNAL SERVER ERROR",err.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.RemoveSubscriptions(id)
	if err1!=nil {
		responseFormatter(404,"NOT FOUND",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200,"OK","SUBSCRIPTION REMOVED",&response)
	json.NewEncoder(w).Encode(response)
}