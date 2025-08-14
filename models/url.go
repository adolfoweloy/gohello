package models

import "time"

// URLMapping represents a shortened URL with its metadata
type URLMapping struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
	ClickCount  int       `json:"click_count"`
}

// ShortenRequest represents the JSON payload for shortening a URL
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse represents the JSON response when shortening a URL
type ShortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	ShortCode   string `json:"short_code"`
}

// StatsResponse represents the statistics for a URL
type StatsResponse struct {
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	ClickCount  int       `json:"click_count"`
	CreatedAt   time.Time `json:"created_at"`
}

// ListResponse represents the response for listing all URLs
type ListResponse struct {
	URLs  []URLMapping `json:"urls"`
	Total int          `json:"total"`
}