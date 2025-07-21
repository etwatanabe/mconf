package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"api/internal/models"
)

type OpenLibraryService struct {
	baseURL string
	limit int
	client *http.Client
}

func NewOpenLibraryService(baseURL string, limit int) *OpenLibraryService {
	return &OpenLibraryService{
		baseURL: baseURL,
		limit: limit,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *OpenLibraryService) SearchBooks(query string) (*models.OpenLibraryResponse, error) {
	encodedQuery := url.QueryEscape(query)
	apiURL := fmt.Sprintf("%s?q=%s&limit=%d", s.baseURL, encodedQuery, s.limit)

	resp, err := s.client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making request to OpenLibrary: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenLibrary API returned status: %d", resp.StatusCode)
	}

	var openLibResp models.OpenLibraryResponse
	if err := json.NewDecoder(resp.Body).Decode(&openLibResp); err != nil {
		return nil, fmt.Errorf("error decoding response from OpenLibrary: %w", err)
	}

	return &openLibResp, nil
}