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
	// r.Router.HandleFunc("/api/v1/project", r.Handler.GetProject).Methods("GET")
	// r.Router.HandleFunc("/api/v1/sahpes", r.Handler.GetShapes).Methods("GET")
	// r.Router.HandleFunc("/api/v1/fonts", r.Handler.GetFonts).Methods("GET") 
	// r.Router.HandleFunc("/api/v1/project", r.Handler.CreateProject).Methods("POST")
	// r.Router.HandleFunc("/api/v1/project", r.Handler.UpdateProject).Methods("PUT") 
	// r.Router.HandleFunc("/api/v1/project", r.Handler.DeleteProject).Methods("DELETE")
	// r.Router.HandleFunc("/api/v1/element", r.Handler.AddElement).Methods("POST")
	// r.Router.HandleFunc("/api/v1/element", r.Handler.DeleteElement).Methods("DELETE")
	// r.Router.HandleFunc("/api/v1/tag", r.Handler.AddTag).Methods("POST")
	// r.Router.HandleFunc("/api/v1/tag", r.Handler.DeleteTag).Methods("DELETE")
	// r.Router.HandleFunc("/api/v1/tagElement", r.Handler.AddTagElement).Methods("POST")
	// r.Router.HandleFunc("/api/v1/tagElement", r.Handler.DeleteTagElement).Methods("DELETE")
}