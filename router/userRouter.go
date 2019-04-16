package router
import (
	"github.com/gorilla/mux"
	handlers "github.com/Alaedeen/360VideoEditorAPI/handlers"
)

// RouterHandler ...
type UserRouterHandler struct {
	Router *mux.Router
	Handler handlers.UserHandler  
}

// HandleFunctions ...
func (r *UserRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints  
	r.Router.HandleFunc("/api/v1/users", r.Handler.GetUsers).Methods("GET")
	r.Router.HandleFunc("/api/v1/users/{id}", r.Handler.GetUser).Methods("GET")
	r.Router.HandleFunc("/api/v1/users", r.Handler.CreateUser).Methods("POST")
	r.Router.HandleFunc("/api/v1/users/{id}", r.Handler.UpdateUser).Methods("PUT") 
	r.Router.HandleFunc("/api/v1/users/{id}", r.Handler.DeleteUser).Methods("DELETE")
}