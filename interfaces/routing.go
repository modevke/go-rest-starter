package interfaces


import(
	// CORE LIBS
	"net/http"

	// EXTERNAL LIBS
	"github.com/gorilla/mux"

)


func Routing() http.Handler{

	// INITIALIZE ROUTER
	r := mux.NewRouter()

	// ROUTER GROUPS
	api := r.PathPrefix("/api/v1").Subrouter()


	// USER APIS
	userRoutes := api.PathPrefix("/users").Subrouter()
	userHandle := UserHandler{}

	userRoutes.HandleFunc("/create", userHandle.CreateUser).Methods("POST")
	userRoutes.HandleFunc("/fetch", userHandle.FetchUser).Methods("GET")
	userRoutes.HandleFunc("/update", userHandle.UpdateUser).Methods("PUT")
	userRoutes.HandleFunc("/delete", userHandle.DeleteUser).Methods("DELETE")

	return r

}