# URL Shortener in Go

A simple URL shortener service built with Go that demonstrates fundamental concepts including HTTP servers, JSON handling, and concurrent programming.

## Features

- 🔗 Shorten long URLs to 6-character codes
- 📊 Track click statistics for each URL
- 🌐 Simple web interface for easy testing
- 🔄 REST API for programmatic access
- 🔒 Thread-safe concurrent operations
- 📝 In-memory storage (perfect for learning)

## Quick Start

1. **Clone and Setup**
   ```bash
   git clone <your-repo-url>
   cd url-shortener
   go mod tidy
   ```

2. **Run the Server**
   ```bash
   go run main.go
   ```

3. **Access the Application**
   - Web Interface: http://localhost:8080
   - API Base URL: http://localhost:8080/api

## API Endpoints

### Shorten a URL
```bash
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/very/long/url"
}
```

**Response:**
```json
{
  "original_url": "https://example.com/very/long/url",
  "short_url": "http://localhost:8080/redirect/abc123",
  "short_code": "abc123"
}
```

### Redirect to Original URL
```bash
GET /redirect/{shortCode}
```
Redirects to the original URL and increments click count.

### List All URLs
```bash
GET /api/urls
```

**Response:**
```json
{
  "urls": [
    {
      "id": "...",
      "original_url": "https://example.com",
      "short_code": "abc123",
      "created_at": "2024-01-01T12:00:00Z",
      "click_count": 5
    }
  ],
  "total": 1
}
```

### Get URL Statistics
```bash
GET /api/stats/{shortCode}
```

**Response:**
```json
{
  "short_code": "abc123",
  "original_url": "https://example.com",
  "click_count": 5,
  "created_at": "2024-01-01T12:00:00Z"
}
```

## Project Structure

```
url-shortener/
├── main.go              # Server setup and routing
├── handlers/
│   └── handlers.go      # HTTP request handlers
├── models/
│   └── url.go          # Data structures
├── services/
│   └── shortener.go    # Business logic
├── utils/
│   └── generator.go    # Utility functions
├── go.mod              # Go module
└── README.md           # This file
```

## Learning Objectives

This project demonstrates:

- **HTTP Server**: Creating and configuring HTTP servers
- **Routing**: Handling different HTTP methods and URL patterns
- **JSON**: Marshaling and unmarshaling JSON data
- **Structs**: Defining and using Go structs
- **Maps**: Using maps for data storage
- **Concurrency**: Thread-safe operations with sync.RWMutex
- **Error Handling**: Proper Go error handling patterns
- **Package Organization**: Structuring Go projects
- **Environment Variables**: Configuration management

## Testing the Service

### Using curl

1. **Shorten a URL:**
   ```bash
   curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://golang.org"}'
   ```

2. **Test the redirect:**
   ```bash
   curl -L http://localhost:8080/redirect/abc123
   ```

3. **Get statistics:**
   ```bash
   curl http://localhost:8080/api/stats/abc123
   ```

### Using the Web Interface

1. Open http://localhost:8080 in your browser
2. Enter a URL in the input field
3. Click "Shorten" to create a short URL
4. Click "Refresh List" to see all URLs and their statistics

## Configuration

Set the `PORT` environment variable to change the server port:

```bash
PORT=3000 go run main.go
```

## Next Steps for Learning

1. **Add Persistence**: Replace in-memory storage with a database
2. **Add Authentication**: Implement user accounts and API keys
3. **Add Validation**: More comprehensive URL validation
4. **Add Caching**: Implement Redis for better performance
5. **Add Metrics**: Prometheus metrics for monitoring
6. **Add Tests**: Write comprehensive unit and integration tests
7. **Add Docker**: Containerize the application
8. **Add Rate Limiting**: Prevent abuse with rate limiting

## Dependencies

This project uses only Go's standard library, making it perfect for learning core concepts without external dependencies.