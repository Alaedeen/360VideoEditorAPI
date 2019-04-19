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
	Subscriptions 		[]Subscriptions		`json:"subscriptions"`
	VideosLikes 		[]VideosLikes		`json:"videosLikes"`
	VideosDislikes 		[]VideosDislikes	`json:"videosDislikes"`
	CommentsLikes 		[]CommentsLikes		`json:"commentsLikes"`
	CommentsDislikes 	[]CommentsDislikes	`json:"commentsDislikes"`
	RepliesLikes 		[]RepliesLikes		`json:"repliesLikes"`
	RepliesDislikes 	[]RepliesDislikes	`json:"repliesDislikes"`
}

// Date Struct
type Date struct {
	Day 	int		`json:"day"`
	Month 	string 	`json:"month"`
	Year 	int		`json:"year"`
}