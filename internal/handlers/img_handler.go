package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"svi_danie/internal/repositories/models"
	"svi_danie/internal/services"

	"github.com/google/uuid"
)

type ImgHandler struct {
	ImgService *services.ImgService
}

func (h *ImgHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ImageUrl string `json:"imgUrl"`
		Error    string `json:"error"`
	}

	resp := &Response{}
	defer json.NewEncoder(w).Encode(resp)
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		return
	}

	imgFile, _, err := r.FormFile("img")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		return
	}

	imgBytes, err := io.ReadAll(imgFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Error = err.Error()
		return
	}

	createdImg := models.Img{
		Id:   uuid.New(),
		Data: imgBytes,
	}
	err = h.ImgService.CreateImage(createdImg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Error = err.Error()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp.ImageUrl = r.Host + "/" + createdImg.Id.String()
}

func (h *ImgHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	ImgIdStr := r.URL.Query().Get("img_id")
	imgId, err := uuid.Parse(ImgIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	img, err := h.ImgService.GetImageById(imgId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(img.Data))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(img.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
