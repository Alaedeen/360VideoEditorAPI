package models

import (
	"github.com/jinzhu/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Name				string 				`json:"name"`
	Email				string 				`json:"email"`
	Password			string 				`json:"password"`
	Admin				bool 				`json:"admin"`
	SuperAdmin			bool 				`json:"superAdmin"`
	BirthDay 			int					`json:"birthDay"`
	BirthMonth 			string				`json:"birthMonth"`
	BirthYear 			int					`json:"birthYear"`
	Country 			string				`json:"countryOfResidence"`
	Description 		string				`json:"description"`
	ProfilePic 			string				`json:"profilePic"`
	JoiningDay 			int					`json:"joiningDay"`
	JoiningMonth 		string				`json:"joiningMonth"`
	JoiningYear 		int					`json:"joiningYear"`
	Subscribers 		int					`json:"subscribers"`
	Videos				[]Video				`json:"videos"`
	Subscriptions 		[]Subscriptions		`json:"subscriptions"`
	VideosLikes 		[]VideosLikes		`json:"videosLikes"`
	VideosDislikes 		[]VideosDislikes	`json:"videosDislikes"`
	CommentsLikes 		[]CommentsLikes		`json:"commentsLikes"`
	CommentsDislikes 	[]CommentsDislikes	`json:"commentsDislikes"`
	RepliesLikes 		[]RepliesLikes		`json:"repliesLikes"`
	RepliesDislikes 	[]RepliesDislikes	`json:"repliesDislikes"`
	ProjectVideos		[]Video2D			`json:"projectVideos"`
	ProjectPictures		[]Picture			`json:"projectPictures"`
}