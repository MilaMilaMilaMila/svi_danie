package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"
)

type ProjectHandler struct {
	ProjService *services.ProjService
	PageService *services.PageService
}

func (h *ProjectHandler) CreateProj(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ProjId uuid.UUID `json:"projId"`
		Error  string    `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	ownerIdStr := r.URL.Query().Get("owner_id")
	ownerId, err := uuid.Parse(ownerIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}
	title := r.URL.Query().Get("title")
	proj := &models.Project{
		Id:      uuid.New(),
		OwnerId: ownerId,
		Title:   title,
	}

	err = h.ProjService.CreateProj(proj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.ProjId = proj.Id
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *ProjectHandler) DeleteProj(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ProjId uuid.UUID `json:"proj_id"`
		Error  string    `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	projIdStr := r.URL.Query().Get("proj_id")
	projId, err := uuid.Parse(projIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.ProjService.DeleteProj(projId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.ProjId = projId
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *ProjectHandler) GetAllProj(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Proj  []*models.Project `json:"projects"`
		Error string            `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	userIdStr := r.URL.Query().Get("user_id")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	projects, err := h.ProjService.GetAllUserProj(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Proj = projects
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *ProjectHandler) GetProj(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Proj  models.Project `json:"proj"`
		Error string         `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	projIdStr := r.URL.Query().Get("proj_id")
	projId, err := uuid.Parse(projIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	proj, err := h.ProjService.GetProj(projId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Proj = *proj
	_ = json.NewEncoder(w).Encode(resp)
}
