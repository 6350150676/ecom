package user

import (
	"encoding/json"
	"go/types"
	"net/http"

	"github.com/gorilla/mux"
	utils "github.com/sikozonpc/ecom/Utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: implement login logic
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// TODO: implement register logic
	// get JSON payload
	var payload types.RegisterUserPayload
	if err:=utils.ParseJson(r,payload) err!=nil{
		utils.writeError(w, http.StatusBadrequest(),err )
	}
	// check if the user exists
	// if it doesnt we create the new user

}
