package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"api/internal/models"
	"api/internal/services"
)

type BookHandler struct {
	openLibraryService *services.OpenLibraryService
}

func NewBookHandler(openLibraryService *services.OpenLibraryService) *BookHandler {
	return &BookHandler{
		openLibraryService: openLibraryService,
	}
}

func (h *BookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	log.Printf("Received request for query: %s", query)

	openLibResp, err := h.openLibraryService.SearchBooks(query)
	if err != nil {
		log.Printf("Error searching books: %v", err)
		http.Error(w, "Error fetching data from external API", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d books for query: %s", openLibResp.NumFound, query)

	response := models.APIResponse{
		Query: query,
		Books: openLibResp.Docs,
		Total: openLibResp.NumFound,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
