package router
import (
	"github.com/gorilla/mux"
	handlers "github.com/Alaedeen/360VideoEditorAPI/handlers"
)

// VideoRouterHandler ...
type VideoRouterHandler struct {
	Router *mux.Router
	Handler handlers.VideoHandler
}

// HandleFunctions ...
func (r *VideoRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints  
	r.Router.HandleFunc("/api/v1/videos", r.Handler.GetVideos).Methods("GET")
	r.Router.HandleFunc("/api/v1/video", r.Handler.GetVideo).Methods("GET")
	r.Router.HandleFunc("/api/v1/video", r.Handler.AddVideo).Methods("POST") 
	r.Router.HandleFunc("/api/v1/video", r.Handler.DeleteVideo).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/comment", r.Handler.AddComment).Methods("POST") 
	r.Router.HandleFunc("/api/v1/comment", r.Handler.DeleteComment).Methods("DELETE")
	// r.Router.HandleFunc("/api/v1/reply", r.Handler.AddReply).Methods("POST") 
	// r.Router.HandleFunc("/api/v1/reply", r.Handler.DeleteReply).Methods("DELETE") 
}