package router
import (
	"github.com/Alaedeen/360VideoEditorAPI/helpers"
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
	r.Router.Handle("/api/v1/video", helpers.IsAuthorized(r.Handler.AddVideo)).Methods("POST") 
	r.Router.Handle("/api/v1/video", helpers.IsAuthorized(r.Handler.DeleteVideo)).Methods("DELETE")
	r.Router.Handle("/api/v1/video", helpers.IsAuthorized(r.Handler.UpdateVideo)).Methods("PUT")
	r.Router.Handle("/api/v1/comment", helpers.IsAuthorized(r.Handler.AddComment)).Methods("POST") 
	r.Router.Handle("/api/v1/comment", helpers.IsAuthorized(r.Handler.DeleteComment)).Methods("DELETE")
	r.Router.Handle("/api/v1/comment", helpers.IsAuthorized(r.Handler.UpdateComment)).Methods("PUT")
	r.Router.Handle("/api/v1/reply", helpers.IsAuthorized(r.Handler.AddReply)).Methods("POST") 
	r.Router.Handle("/api/v1/reply", helpers.IsAuthorized(r.Handler.DeleteReply)).Methods("DELETE") 
	r.Router.Handle("/api/v1/reply", helpers.IsAuthorized(r.Handler.UpdateReply)).Methods("PUT")
}