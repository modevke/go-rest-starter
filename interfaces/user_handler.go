package interfaces

import(
	"net/http"
)

type UserHandler struct{

}

// CREATE USER
func(u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Create User"}`))
}

// FETCH USER
func(u *UserHandler) FetchUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Fetch User"}`))
}

// UPDATE USER
func(u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Update User"}`))
}

// DELETE USER
func(u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Delete User"}`))
}