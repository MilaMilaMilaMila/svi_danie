package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/google/uuid"
)

type PageHandler struct {
	PageService *services.PageService
}

func (h *PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		PageId uuid.UUID `json:"pageId"`
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

	projIdStr := r.URL.Query().Get("project_id")
	projId, err := uuid.Parse(projIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	title := r.URL.Query().Get("title")
	// Читаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Создаем переменную типа json.RawMessage и декодируем тело запроса в нее
	var data json.RawMessage
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// Теперь data содержит сырой JSON
	log.Printf("Raw JSON: %s", data)

	//imgs := h.SaveImgs(w, r)

	page := &models.Page{
		Id:        uuid.New(),
		OwnerId:   ownerId,
		ProjectId: projId,
		Title:     title,
		Data:      data,
	}
	err = h.PageService.CreatePage(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.PageId = page.Id
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		PageId uuid.UUID `json:"page_id"`
		Error  string    `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	pageIdStr := r.URL.Query().Get("page_id")
	pageId, err := uuid.Parse(pageIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.PageService.DeletePage(pageId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.PageId = pageId
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *PageHandler) EditPage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		PageId uuid.UUID `json:"pageId"`
		Error  string    `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	ownerIdStr := r.URL.Query().Get("owner_id")
	ownerId, err := uuid.Parse(ownerIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	projIdStr := r.URL.Query().Get("project_id")
	projId, err := uuid.Parse(projIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	title := r.URL.Query().Get("title")
	// Читаем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Создаем переменную типа json.RawMessage и декодируем тело запроса в нее
	var data json.RawMessage
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// Теперь data содержит сырой JSON
	log.Printf("Raw JSON: %s", data)

	//imgs := h.SaveImgs(w, r)

	page := &models.Page{
		Id:        id,
		OwnerId:   ownerId,
		ProjectId: projId,
		Title:     title,
		Data:      data,
	}
	err = h.PageService.UpdatePage(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.PageId = page.Id
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *PageHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Page  models.Page `json:"page"`
		Error string      `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	pageIdStr := r.URL.Query().Get("page_id")
	pageId, err := uuid.Parse(pageIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	page, err := h.PageService.GetPage(pageId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Page = *page
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *PageHandler) GetAllProjectPages(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Proj  []*models.Page `json:"pages"`
		Error string         `json:"error"`
	}

	resp := &Response{}
	w.Header().Set("Content-Type", "application/json")

	projIdStr := r.URL.Query().Get("project_id")
	projId, err := uuid.Parse(projIdStr)
	if err != nil {
		resp.Error = err.Error()
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	projects, err := h.PageService.GetAllProjectPages(projId)
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
