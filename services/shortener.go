package services

import (
	"errors"
	"sync"
	"time"

	"adolfoeloy.com/gohello/models"
	"adolfoeloy.com/gohello/utils"
)

// URLShortener manages URL mappings and operations
type URLShortener struct {
	urls  map[string]*models.URLMapping // shortCode -> URLMapping
	mutex sync.RWMutex
}

// NewURLShortener creates a new URL shortener instance
func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]*models.URLMapping),
	}
}

// ShortenURL creates a short code for the given URL
func (us *URLShortener) ShortenURL(originalURL string) (*models.URLMapping, error) {
	if originalURL == "" {
		return nil, errors.New("URL cannot be empty")
	}

	// Check if URL already exists
	us.mutex.RLock()
	for _, mapping := range us.urls {
		if mapping.OriginalURL == originalURL {
			us.mutex.RUnlock()
			return mapping, nil
		}
	}
	us.mutex.RUnlock()

	// Generate a unique short code
	shortCode := us.generateUniqueShortCode()

	// Create new URL mapping
	mapping := &models.URLMapping{
		ID:          utils.GenerateID(),
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
		ClickCount:  0,
	}

	// Store the mapping
	us.mutex.Lock()
	us.urls[shortCode] = mapping
	us.mutex.Unlock()

	return mapping, nil
}

// GetOriginalURL retrieves the original URL for a given short code
func (us *URLShortener) GetOriginalURL(shortCode string) (string, error) {
	us.mutex.RLock()
	mapping, exists := us.urls[shortCode]
	us.mutex.RUnlock()

	if !exists {
		return "", errors.New("short code not found")
	}

	// Increment click count
	us.mutex.Lock()
	mapping.ClickCount++
	us.mutex.Unlock()

	return mapping.OriginalURL, nil
}

// GetStats returns statistics for a given short code
func (us *URLShortener) GetStats(shortCode string) (*models.URLMapping, error) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	mapping, exists := us.urls[shortCode]
	if !exists {
		return nil, errors.New("short code not found")
	}

	return mapping, nil
}

// ListAll returns all URL mappings
func (us *URLShortener) ListAll() []models.URLMapping {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	var urls []models.URLMapping
	for _, mapping := range us.urls {
		urls = append(urls, *mapping)
	}

	return urls
}

// generateUniqueShortCode generates a unique short code
func (us *URLShortener) generateUniqueShortCode() string {
	for {
		code := utils.GenerateShortCode(6)
		us.mutex.RLock()
		_, exists := us.urls[code]
		us.mutex.RUnlock()

		if !exists {
			return code
		}
	}
}
