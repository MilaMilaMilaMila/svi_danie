package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		UserId uuid.UUID `json:"userId"`
		Error  string    `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	user := &models.User{
		Id:       uuid.New(),
		Login:    login,
		Password: password,
	}

	err := h.UserService.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.UserId = user.Id
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
