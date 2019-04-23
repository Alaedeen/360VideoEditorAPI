package models



// Response Struct
type Response struct {
	Code	int 		`json:"code"`
	Status	string 		`json:"status"`
	Data	interface{}	`json:"data"`
}

// UserResponse Struct
type UserResponse struct {
	ID					uint				`json:"id"`
	Name				string 				`json:"name"`
	Email				string 				`json:"email"`
	Password			string 				`json:"password"`
	Roles				[]string 			`json:"roles"`
	DateOfBirth 		Date				`json:"dateOfBirth"`
	Country 			string				`json:"countryOfResidence"`
	Description 		string				`json:"description"`
	ProfilePic 			string				`json:"profilePic"`
	Joined				Date				`json:"joined"`
	Subscribers 		int					`json:"subscribers"`
	Videos				[]Video				`json:"videos"`
	Subscriptions 		[]int				`json:"subscriptions"`
	VideosLikes 		[]int				`json:"videosLikes"`
	VideosDislikes 		[]int				`json:"videosDislikes"`
	CommentsLikes 		[]CommentsLikesResponse		`json:"commentsLikes"`
	CommentsDislikes 	[]CommentsDislikesResponse	`json:"commentsDislikes"`
	RepliesLikes 		[]RepliesLikesResponse		`json:"repliesLikes"`
	RepliesDislikes 	[]RepliesDislikesResponse	`json:"repliesDislikes"`
}

// CommentsLikesResponse struct
type CommentsLikesResponse struct {
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
}

// CommentsDislikesResponse struct
type CommentsDislikesResponse struct {
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
}

// RepliesLikesResponse struct
type RepliesLikesResponse struct {
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
	ReplyID 	int `json:"idReply"`
}

// RepliesDislikesResponse struct
type RepliesDislikesResponse struct {
	VideoID		int `json:"idVideo"`
	CommentID	int `json:"idComment"`
	ReplyID 	int `json:"idReply"`
}

// VideoResponse Struct
type VideoResponse struct {
	ID				uint				`json:"vidId"`
	UserID			int 				`json:"userId"`
	Title			string 				`json:"title"`
	UploadDate 		Date				`json:"uploadDate"`
	Thumbnail 		string				`json:"thumbnail"`
	Src 			string				`json:"src"`
	AFrame 			string				`json:"aFrame"`
	Likes 			int					`json:"likes"`
	Dislikes 		int					`json:"dislikes"`
	Views 			int					`json:"views"`
	Comments		[]CommentResponse	`json:"comments"`
}

// CommentResponse Struct
type CommentResponse struct {
	ID				uint			`json:"idComment"`
	UserID			int 			`json:"idUser"`
	VideoID			int 			`json:"videoId"`
	NameUser		string 			`json:"nameUser"`
	ProfilePic		string 			`json:"profilePic"`
	Text			string			`json:"text"`
	CommentDate 	Date			`json:"date"`
	Likes 			int				`json:"likes"`
	Dislikes 		int				`json:"dislikes"`
	Replies			[]ReplyResponse	`json:"replies"`
}

// ReplyResponse Struct
type ReplyResponse struct {
	ID				uint	`json:"idReply"`
	UserID			int 	`json:"idUser"`
	CommentID		int 	`json:"commentId"`
	NameUser		string 	`json:"nameUser"`
	ProfilePic		string 	`json:"profilePic"`
	Text			string	`json:"text"`
	ReplyDate 		Date	`json:"date"`
	Likes 			int		`json:"likes"`
	Dislikes 		int		`json:"dislikes"`
}

// Date Struct
type Date struct {
	Day 	int		`json:"day"`
	Month 	string 	`json:"month"`
	Year 	int		`json:"year"`
}