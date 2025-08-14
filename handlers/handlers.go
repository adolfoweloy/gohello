package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"adolfoeloy.com/gohello/models"
	"adolfoeloy.com/gohello/services"
	"adolfoeloy.com/gohello/utils"
)

// Handler contains the URL shortener service and HTTP handlers
type Handler struct {
	shortener *services.URLShortener
}

// NewHandler creates a new handler instance
func NewHandler(shortener *services.URLShortener) *Handler {
	return &Handler{
		shortener: shortener,
	}
}

// ServeHomePage serves the main HTML page
func (h *Handler) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	html := `
<!DOCTYPE html>
<html>
<head>
    <title>URL Shortener</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .container { background: #f9f9f9; padding: 20px; border-radius: 8px; margin: 20px 0; }
        input[type="text"] { width: 70%; padding: 10px; margin: 5px; }
        button { padding: 10px 20px; background: #007cba; color: white; border: none; cursor: pointer; }
        button:hover { background: #005a8a; }
        .result { margin: 20px 0; padding: 10px; background: #e8f5e8; border-radius: 4px; }
        .error { background: #f5e8e8; }
        .url-list { margin: 20px 0; }
        .url-item { padding: 10px; border: 1px solid #ddd; margin: 5px 0; border-radius: 4px; }
    </style>
</head>
<body>
    <h1>🔗 URL Shortener</h1>
    
    <div class="container">
        <h2>Shorten a URL</h2>
        <input type="text" id="urlInput" placeholder="Enter URL to shorten (e.g., https://example.com)" />
        <button onclick="shortenURL()">Shorten</button>
        <div id="result"></div>
    </div>

    <div class="container">
        <h2>All Shortened URLs</h2>
        <button onclick="loadURLs()">Refresh List</button>
        <div id="urlList"></div>
    </div>

    <script>
        async function shortenURL() {
            const url = document.getElementById('urlInput').value;
            const resultDiv = document.getElementById('result');
            
            if (!url) {
                resultDiv.innerHTML = '<div class="result error">Please enter a URL</div>';
                return;
            }

            try {
                const response = await fetch('/shorten', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ url: url })
                });

                const data = await response.json();
                
                if (response.ok) {
                    resultDiv.innerHTML = ` + "`" + `
                        <div class="result">
                            <strong>Shortened URL:</strong> <a href="/redirect/${data.short_code}" target="_blank">http://localhost:8080/redirect/${data.short_code}</a><br>
                            <strong>Short Code:</strong> ${data.short_code}<br>
                            <strong>Original URL:</strong> ${data.original_url}
                        </div>
                    ` + "`" + `;
                    document.getElementById('urlInput').value = '';
                } else {
                    resultDiv.innerHTML = ` + "`" + `<div class="result error">Error: ${data.error}</div>` + "`" + `;
                }
            } catch (error) {
                resultDiv.innerHTML = ` + "`" + `<div class="result error">Error: ${error.message}</div>` + "`" + `;
            }
        }

        async function loadURLs() {
            try {
                const response = await fetch('/api/urls');
                const data = await response.json();
                
                const urlListDiv = document.getElementById('urlList');
                
                if (data.urls && data.urls.length > 0) {
                    let html = ` + "`" + `<p><strong>Total URLs:</strong> ${data.total}</p>` + "`" + `;
                    data.urls.forEach(url => {
                        html += ` + "`" + `
                            <div class="url-item">
                                <strong>Short Code:</strong> ${url.short_code}<br>
                                <strong>Original URL:</strong> <a href="${url.original_url}" target="_blank">${url.original_url}</a><br>
                                <strong>Short URL:</strong> <a href="/redirect/${url.short_code}" target="_blank">http://localhost:8080/redirect/${url.short_code}</a><br>
                                <strong>Clicks:</strong> ${url.click_count}<br>
                                <strong>Created:</strong> ${new Date(url.created_at).toLocaleString()}
                            </div>
                        ` + "`" + `;
                    });
                    urlListDiv.innerHTML = html;
                } else {
                    urlListDiv.innerHTML = '<p>No URLs shortened yet.</p>';
                }
            } catch (error) {
                document.getElementById('urlList').innerHTML = ` + "`" + `<p class="error">Error loading URLs: ${error.message}</p>` + "`" + `;
            }
        }

        // Load URLs on page load
        window.onload = () => loadURLs();
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// ShortenURL handles URL shortening requests
func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate URL
	if !utils.IsValidURL(req.URL) {
		h.sendErrorResponse(w, "Invalid URL format. URL must start with http:// or https://", http.StatusBadRequest)
		return
	}

	// Shorten the URL
	mapping, err := h.shortener.ShortenURL(req.URL)
	if err != nil {
		log.Printf("Error shortening URL: %v", err)
		h.sendErrorResponse(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	// Create response
	response := models.ShortenResponse{
		OriginalURL: mapping.OriginalURL,
		ShortURL:    fmt.Sprintf("http://localhost:8080/redirect/%s", mapping.ShortCode),
		ShortCode:   mapping.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectURL handles redirection from short URLs to original URLs
func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract short code from URL path
	path := strings.TrimPrefix(r.URL.Path, "/redirect/")
	if path == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	// Get original URL
	originalURL, err := h.shortener.GetOriginalURL(path)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

// ListURLs returns all shortened URLs
func (h *Handler) ListURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urls := h.shortener.ListAll()
	response := models.ListResponse{
		URLs:  urls,
		Total: len(urls),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetStats returns statistics for a specific short code
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract short code from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/stats/")
	if path == "" {
		h.sendErrorResponse(w, "Short code is required", http.StatusBadRequest)
		return
	}

	// Get statistics
	mapping, err := h.shortener.GetStats(path)
	if err != nil {
		h.sendErrorResponse(w, "Short code not found", http.StatusNotFound)
		return
	}

	response := models.StatsResponse{
		ShortCode:   mapping.ShortCode,
		OriginalURL: mapping.OriginalURL,
		ClickCount:  mapping.ClickCount,
		CreatedAt:   mapping.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// sendErrorResponse sends a JSON error response
func (h *Handler) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
