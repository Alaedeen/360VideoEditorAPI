package router
import (
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
	r.Router.HandleFunc("/api/v1/users", r.Handler.GetUsers).Methods("GET")
	r.Router.HandleFunc("/api/v1/user", r.Handler.GetUser).Methods("GET")
	r.Router.HandleFunc("/api/v1/userby", r.Handler.GetUserBy).Methods("GET")
	r.Router.HandleFunc("/api/v1/user/videos", r.Handler.GetUserVideos).Methods("GET")
	r.Router.HandleFunc("/api/v1/users", r.Handler.CreateUser).Methods("POST")
	r.Router.HandleFunc("/api/v1/users", r.Handler.UpdateUser).Methods("PUT") 
	r.Router.HandleFunc("/api/v1/users", r.Handler.DeleteUser).Methods("DELETE")
	
	r.Router.HandleFunc("/api/v1/commentsLikes", r.Handler.AddCommentsLikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/commentsLikes", r.Handler.RemoveCommentsLikes).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/commentsDislikes", r.Handler.AddCommentsDislikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/commentsDislikes", r.Handler.RemoveCommentsDislikes).Methods("DELETE")
	
	r.Router.HandleFunc("/api/v1/repliesLikes", r.Handler.AddRepliesLikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/repliesLikes", r.Handler.RemoveRepliesLikes).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/repliesDislikes", r.Handler.AddRepliesDislikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/repliesDislikes", r.Handler.RemoveRepliesDislikes).Methods("DELETE")

	r.Router.HandleFunc("/api/v1/videosLikes", r.Handler.AddVideosLikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/videosLikes", r.Handler.RemoveVideosLikes).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/videosDislikes", r.Handler.AddVideosDislikes).Methods("POST")
	r.Router.HandleFunc("/api/v1/videosDislikes", r.Handler.RemoveVideosDislikes).Methods("DELETE")

	r.Router.HandleFunc("/api/v1/subscriptions", r.Handler.AddSubscriptions).Methods("POST")
	r.Router.HandleFunc("/api/v1/subscriptions", r.Handler.RemoveSubscriptions).Methods("DELETE")
}