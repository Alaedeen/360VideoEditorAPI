package router
import (
	"github.com/Alaedeen/360VideoEditorAPI/helpers"
	"github.com/gorilla/mux"
	handlers "github.com/Alaedeen/360VideoEditorAPI/handlers"
)

// ProjectRouterHandler ...
type ProjectRouterHandler struct {
	Router *mux.Router
	Handler handlers.ProjectHandler  
}

// HandleFunctions ...
func (r *ProjectRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints  
	r.Router.Handle("/api/v1/projects", helpers.IsAuthorized(r.Handler.GetProjects)).Methods("GET")
	r.Router.Handle("/api/v1/project", helpers.IsAuthorized(r.Handler.GetProject)).Methods("GET")
	r.Router.Handle("/api/v1/project/script", helpers.IsAuthorized(r.Handler.LoadProjectScript)).Methods("GET")
	r.Router.Handle("/api/v1/project", helpers.IsAuthorized(r.Handler.CreateProject)).Methods("POST")
	r.Router.Handle("/api/v1/project", helpers.IsAuthorized(r.Handler.UpdateProject)).Methods("PUT") 
	r.Router.Handle("/api/v1/project", helpers.IsAuthorized(r.Handler.DeleteProject)).Methods("DELETE")
	r.Router.Handle("/api/v1/project/save", helpers.IsAuthorized(r.Handler.SaveProject)).Methods("PUT")
	r.Router.Handle("/api/v1/sahpes", helpers.IsAuthorized(r.Handler.GetShapes)).Methods("GET")
	r.Router.Handle("/api/v1/fonts", helpers.IsAuthorized(r.Handler.GetFonts)).Methods("GET") 
	r.Router.Handle("/api/v1/project/element", helpers.IsAuthorized(r.Handler.AddElement)).Methods("POST")
	r.Router.Handle("/api/v1/project/element", helpers.IsAuthorized(r.Handler.DeleteElement)).Methods("DELETE")
	r.Router.Handle("/api/v1/project/tag", helpers.IsAuthorized(r.Handler.AddTag)).Methods("POST")
	r.Router.Handle("/api/v1/project/tag", helpers.IsAuthorized(r.Handler.DeleteTag)).Methods("DELETE")
	r.Router.Handle("/api/v1/project/tagElement", helpers.IsAuthorized(r.Handler.AddTagElement)).Methods("POST")
	r.Router.Handle("/api/v1/project/tagElement", helpers.IsAuthorized(r.Handler.DeleteTagElement)).Methods("DELETE")
	r.Router.Handle("/api/v1/project/picture", helpers.IsAuthorized(r.Handler.AddPicture)).Methods("POST")
	r.Router.Handle("/api/v1/project/picture", helpers.IsAuthorized(r.Handler.DeletePicture)).Methods("DELETE")
	r.Router.Handle("/api/v1/project/video", helpers.IsAuthorized(r.Handler.AddProjectVideo)).Methods("POST")
	r.Router.Handle("/api/v1/project/video", helpers.IsAuthorized(r.Handler.DeleteProjectVideo)).Methods("DELETE")

	r.Router.Handle("/api/v1/uploadRequests", helpers.IsAuthorized(r.Handler.GetUploadRequests)).Methods("GET")
	r.Router.Handle("/api/v1/uploadRequests", helpers.IsAuthorized(r.Handler.DeleteUploadRequest)).Methods("DELETE")
	r.Router.Handle("/api/v1/uploadRequests", helpers.IsAuthorized(r.Handler.AddUploadRequest)).Methods("POST")
}