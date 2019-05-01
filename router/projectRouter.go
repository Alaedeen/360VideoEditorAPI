package router
import (
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
	r.Router.HandleFunc("/api/v1/projects", r.Handler.GetProjects).Methods("GET")
	r.Router.HandleFunc("/api/v1/project", r.Handler.GetProject).Methods("GET")
	r.Router.HandleFunc("/api/v1/project/script", r.Handler.LoadProjectScript).Methods("GET")
	r.Router.HandleFunc("/api/v1/project", r.Handler.CreateProject).Methods("POST")
	r.Router.HandleFunc("/api/v1/project", r.Handler.UpdateProject).Methods("PUT") 
	r.Router.HandleFunc("/api/v1/project", r.Handler.DeleteProject).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/project/save", r.Handler.SaveProject).Methods("PUT")
	r.Router.HandleFunc("/api/v1/sahpes", r.Handler.GetShapes).Methods("GET")
	r.Router.HandleFunc("/api/v1/fonts", r.Handler.GetFonts).Methods("GET") 
	r.Router.HandleFunc("/api/v1/project/element", r.Handler.AddElement).Methods("POST")
	r.Router.HandleFunc("/api/v1/project/element", r.Handler.DeleteElement).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/project/tag", r.Handler.AddTag).Methods("POST")
	r.Router.HandleFunc("/api/v1/project/tag", r.Handler.DeleteTag).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/project/tagElement", r.Handler.AddTagElement).Methods("POST")
	r.Router.HandleFunc("/api/v1/project/tagElement", r.Handler.DeleteTagElement).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/project/picture", r.Handler.AddPicture).Methods("POST")
	r.Router.HandleFunc("/api/v1/project/picture", r.Handler.DeletePicture).Methods("DELETE")
	r.Router.HandleFunc("/api/v1/project/video", r.Handler.AddProjectVideo).Methods("POST")
	r.Router.HandleFunc("/api/v1/project/video", r.Handler.DeleteProjectVideo).Methods("DELETE")
}