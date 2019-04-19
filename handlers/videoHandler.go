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
func replyResponseFormatter(result models.Reply, reply *models.ReplyResponse)  {
	reply.ID=result.ID
	reply.UserID=result.UserID
	reply.CommentID=result.CommentID
	reply.ReplyDate.Day=result.Day
	reply.ReplyDate.Month=result.Month
	reply.ReplyDate.Year=result.Year
	reply.NameUser=result.NameUser
	reply.ProfilePic=result.ProfilePic
	reply.Text=result.Text
	reply.Likes=result.Likes
	reply.Dislikes=result.Dislikes
}
func commentResponseFormatter(result models.Comment, comment *models.CommentResponse)  {
	comment.ID=result.ID
	comment.UserID=result.UserID
	comment.VideoID=result.VideoID
	comment.CommentDate.Day=result.Day
	comment.CommentDate.Month=result.Month
	comment.CommentDate.Year=result.Year
	comment.NameUser=result.NameUser
	comment.ProfilePic=result.ProfilePic
	comment.Text=result.Text
	comment.Likes=result.Likes
	comment.Dislikes=result.Dislikes
	var reply models.ReplyResponse
	comment.Replies= comment.Replies[:0]
	if len(result.Replies)==0 {
		comment.Replies = []models.ReplyResponse{}
	}
	for _, rep := range result.Replies {
		replyResponseFormatter(rep,&reply)
		comment.Replies = append(comment.Replies, reply)
	}
	
}
func videoResponseFormatter(result models.Video) models.VideoResponse {
	video := models.VideoResponse{}
	video.ID=result.ID
	video.UserID=result.UserID
	video.Title=result.Title
	video.UploadDate.Day=result.UploadDay
	video.UploadDate.Month=result.UploadMonth
	video.UploadDate.Year=result.UploadYear
	video.Thumbnail=result.Thumbnail
	video.Src=result.Src
	video.AFrame=result.AFrame
	video.Likes=result.Likes
	video.Dislikes=result.Dislikes
	video.Views=result.Views
	// var comment models.CommentResponse
	if result.Comments!=nil {
		rs := []models.CommentResponse{}
		for _, com := range result.Comments {
			comment := models.CommentResponse{}
			commentResponseFormatter(com,&comment)
			rs = append(rs, comment)
		}
		video.Comments = rs	
	}

	return video

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

	var videos	[]models.VideoResponse
	var video models.VideoResponse
	for _,res := range result {
		
	video = videoResponseFormatter(res)
		videos= append(videos,video)
	} 
	responseFormatter(200,"OK",videos,&response)
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
	
	video := videoResponseFormatter(result)
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
	_, err1:= h.Repo.AddComment(Comment)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","Comment Added",&response)
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
	_, err1:= h.Repo.AddReply(Reply)
	if err1!=nil{
		responseFormatter(500,"INTERNAL SERVER ERROR",err1.Error(),&response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201,"CREATED","Reply Added",&response)
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