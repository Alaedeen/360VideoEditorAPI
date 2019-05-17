package router
import (
	"github.com/Alaedeen/360VideoEditorAPI/helpers"
	"github.com/gorilla/mux"
	handlers "github.com/Alaedeen/360VideoEditorAPI/handlers"
)

// UserRouterHandler ...
type UserRouterHandler struct {
	Router *mux.Router
	Handler handlers.UserHandler  
}

// HandleFunctions ...
func (r *UserRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints  
	r.Router.Handle("/api/v1/users", helpers.IsAuthorized(r.Handler.GetUsers)).Methods("GET")
	r.Router.Handle("/api/v1/usersbyname", helpers.IsAuthorized(r.Handler.GetUsersByName)).Methods("GET")
	r.Router.HandleFunc("/api/v1/user", r.Handler.GetUser).Methods("GET")
	r.Router.Handle("/api/v1/userby", helpers.IsAuthorized(r.Handler.GetUserBy)).Methods("GET")
	r.Router.HandleFunc("/api/v1/login", r.Handler.Login).Methods("GET")
	r.Router.HandleFunc("/api/v1/user/videos", r.Handler.GetUserVideos).Methods("GET")
	r.Router.Handle("/api/v1/user/pictures", helpers.IsAuthorized(r.Handler.GetUserPictures)).Methods("GET")
	r.Router.Handle("/api/v1/user/projectVideos", helpers.IsAuthorized(r.Handler.GetUserProjectVideos)).Methods("GET")
	r.Router.HandleFunc("/api/v1/users", r.Handler.CreateUser).Methods("POST")
	r.Router.Handle("/api/v1/users", helpers.IsAuthorized(r.Handler.UpdateUser)).Methods("PUT") 
	r.Router.Handle("/api/v1/users", helpers.IsAuthorized(r.Handler.DeleteUser)).Methods("DELETE")
	
	r.Router.Handle("/api/v1/commentsLikes", helpers.IsAuthorized(r.Handler.AddCommentsLikes)).Methods("POST")
	r.Router.Handle("/api/v1/commentsLikes", helpers.IsAuthorized(r.Handler.RemoveCommentsLikes)).Methods("DELETE")
	r.Router.Handle("/api/v1/commentsDislikes", helpers.IsAuthorized(r.Handler.AddCommentsDislikes)).Methods("POST")
	r.Router.Handle("/api/v1/commentsDislikes", helpers.IsAuthorized(r.Handler.RemoveCommentsDislikes)).Methods("DELETE")
	
	r.Router.Handle("/api/v1/repliesLikes", helpers.IsAuthorized(r.Handler.AddRepliesLikes)).Methods("POST")
	r.Router.Handle("/api/v1/repliesLikes", helpers.IsAuthorized(r.Handler.RemoveRepliesLikes)).Methods("DELETE")
	r.Router.Handle("/api/v1/repliesDislikes", helpers.IsAuthorized(r.Handler.AddRepliesDislikes)).Methods("POST")
	r.Router.Handle("/api/v1/repliesDislikes", helpers.IsAuthorized(r.Handler.RemoveRepliesDislikes)).Methods("DELETE")

	r.Router.Handle("/api/v1/videosLikes", helpers.IsAuthorized(r.Handler.AddVideosLikes)).Methods("POST")
	r.Router.Handle("/api/v1/videosLikes", helpers.IsAuthorized(r.Handler.RemoveVideosLikes)).Methods("DELETE")
	r.Router.Handle("/api/v1/videosDislikes", helpers.IsAuthorized(r.Handler.AddVideosDislikes)).Methods("POST")
	r.Router.Handle("/api/v1/videosDislikes", helpers.IsAuthorized(r.Handler.RemoveVideosDislikes)).Methods("DELETE")

	r.Router.Handle("/api/v1/subscriptions", helpers.IsAuthorized(r.Handler.AddSubscriptions)).Methods("POST")
	r.Router.Handle("/api/v1/subscriptions", helpers.IsAuthorized(r.Handler.RemoveSubscriptions)).Methods("DELETE")

	r.Router.HandleFunc("/api/v1/user/reset_password", r.Handler.ResetPassword).Methods("PUT")
}