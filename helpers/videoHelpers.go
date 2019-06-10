package helpers

import (
	"github.com/Alaedeen/360VideoEditorAPI/models"
)

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
	reply.Likes=*result.Likes
	reply.Dislikes=*result.Dislikes
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
	comment.Likes=*result.Likes
	comment.Dislikes=*result.Dislikes
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

// VideoResponseFormatter func
func VideoResponseFormatter(result models.Video) models.VideoResponse {
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
	video.Likes=*result.Likes
	video.Dislikes=*result.Dislikes
	video.Views=*result.Views
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