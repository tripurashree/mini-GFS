package http

import (
	"net/http"

	"mini-gfs/chunkserver/internal/storage"
)

type Handler struct {
	Store *storage.Storage
}

func NewHandler(s *storage.Storage) *Handler {
	return &Handler{Store: s}
}

// PUT /chunk/{id}
func (h *Handler) UploadChunk(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chunk/"):]

	err := h.Store.WriteChunk(id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GET /chunk/{id}
func (h *Handler) GetChunk(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chunk/"):]

	file, err := h.Store.ReadChunk(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.ServeContent(w, r, id, fileInfo.ModTime(), file)
}

// DELETE /chunk/{id}
func (h *Handler) DeleteChunk(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chunk/"):]

	err := h.Store.DeleteChunk(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	w.WriteHeader(http.StatusOK)
}