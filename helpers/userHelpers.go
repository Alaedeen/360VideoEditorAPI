package helpers

import(
	models "github.com/Alaedeen/360VideoEditorAPI/models"
	"crypto/sha1"
)

// UserResponseFormatter func
func UserResponseFormatter(result models.User, user *models.UserResponse)  {
	user.ID=result.ID
	user.Name=result.Name
	user.Email=result.Email
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
	user.Subscriptions = []int{}
	for _,subscription := range result.Subscriptions {
		user.Subscriptions = append(user.Subscriptions,subscription.IDSubscribed)
	}
	user.VideosLikes = []int{}
	for _,VideoLike := range result.VideosLikes {
		user.VideosLikes = append(user.VideosLikes,VideoLike.VideoID)
	}
	user.VideosDislikes = []int{}
	for _,VideoDislike := range result.VideosDislikes {
		user.VideosDislikes = append(user.VideosDislikes,VideoDislike.VideoID)
	}
	user.CommentsLikes = []models.CommentsLikesResponse{}
	for _,CommentLike := range result.CommentsLikes {
		var commentLike models.CommentsLikesResponse
		commentLike.VideoID=CommentLike.VideoID
		commentLike.CommentID=CommentLike.CommentID
		user.CommentsLikes = append(user.CommentsLikes,commentLike)
	}
	user.CommentsDislikes = []models.CommentsDislikesResponse{}
	for _,CommentDislike := range result.CommentsDislikes {
		var commentDislike models.CommentsDislikesResponse
		commentDislike.VideoID=CommentDislike.VideoID
		commentDislike.CommentID=CommentDislike.CommentID
		user.CommentsDislikes = append(user.CommentsDislikes,commentDislike)
	}
	user.RepliesLikes = []models.RepliesLikesResponse{}
	for _,ReplyLike := range result.RepliesLikes {
		var replyLike models.RepliesLikesResponse
		replyLike.VideoID=ReplyLike.VideoID
		replyLike.CommentID=ReplyLike.CommentID
		replyLike.ReplyID=ReplyLike.ReplyID
		user.RepliesLikes = append(user.RepliesLikes,replyLike)
	}
	user.RepliesDislikes = []models.RepliesDislikesResponse{}
	for _,ReplyDislike := range result.RepliesDislikes {
		var replyDislike models.RepliesDislikesResponse
		replyDislike.VideoID=ReplyDislike.VideoID
		replyDislike.CommentID=ReplyDislike.CommentID
		replyDislike.ReplyID=ReplyDislike.ReplyID
		user.RepliesDislikes = append(user.RepliesDislikes,replyDislike)
	}
}

// UserRequestFormatter func
func UserRequestFormatter(request models.UserRequest, user *models.User){
	if request.Password!="" {
		crypt := sha1.New()
		crypt.Write([]byte(request.Password))
		user.Password=crypt.Sum(nil)
	}
	user.Name=request.Name
	user.Email=request.Email
	user.Admin=request.Admin
	user.SuperAdmin=request.SuperAdmin
	user.BirthDay=request.BirthDay
	user.BirthMonth=request.BirthMonth
	user.BirthYear=request.BirthYear
	user.Country=request.Country
	user.Description=request.Description
	user.ProfilePic=request.ProfilePic
	user.JoiningDay=request.JoiningDay
	user.JoiningMonth=request.JoiningMonth
	user.JoiningYear=request.JoiningYear
	user.Subscribers=request.Subscribers
}