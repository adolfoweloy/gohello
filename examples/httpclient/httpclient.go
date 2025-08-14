// Package httpclient demonstrates HTTP client operations in Go
package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient wraps the standard HTTP client with additional functionality
type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}
}

// SetHeader sets a default header for all requests
func (hc *HTTPClient) SetHeader(key, value string) {
	hc.headers[key] = value
}

// SetTimeout sets the client timeout
func (hc *HTTPClient) SetTimeout(timeout time.Duration) {
	hc.client.Timeout = timeout
}

// GET performs a GET request
func (hc *HTTPClient) GET(endpoint string) (*http.Response, error) {
	url := hc.buildURL(endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	hc.addHeaders(req)
	return hc.client.Do(req)
}

// POST performs a POST request with JSON data
func (hc *HTTPClient) POST(endpoint string, data interface{}) (*http.Response, error) {
	url := hc.buildURL(endpoint)
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	hc.addHeaders(req)
	
	return hc.client.Do(req)
}

// PUT performs a PUT request with JSON data
func (hc *HTTPClient) PUT(endpoint string, data interface{}) (*http.Response, error) {
	url := hc.buildURL(endpoint)
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	hc.addHeaders(req)
	
	return hc.client.Do(req)
}

// DELETE performs a DELETE request
func (hc *HTTPClient) DELETE(endpoint string) (*http.Response, error) {
	url := hc.buildURL(endpoint)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	
	hc.addHeaders(req)
	return hc.client.Do(req)
}

// buildURL builds the full URL from base URL and endpoint
func (hc *HTTPClient) buildURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}
	
	baseURL := strings.TrimSuffix(hc.baseURL, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")
	
	if baseURL == "" {
		return endpoint
	}
	
	return baseURL + "/" + endpoint
}

// addHeaders adds default headers to the request
func (hc *HTTPClient) addHeaders(req *http.Request) {
	for key, value := range hc.headers {
		req.Header.Set(key, value)
	}
}

// Response represents an HTTP response with additional helper methods
type Response struct {
	*http.Response
	body []byte
}

// NewResponse creates a Response from http.Response
func NewResponse(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return &Response{
		Response: resp,
		body:     body,
	}, nil
}

// Body returns the response body as bytes
func (r *Response) Body() []byte {
	return r.body
}

// String returns the response body as string
func (r *Response) String() string {
	return string(r.body)
}

// JSON unmarshals the response body into the provided interface
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

// IsSuccess returns true if the status code indicates success (2xx)
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// DemonstrateBasicHTTP shows basic HTTP operations
func DemonstrateBasicHTTP() {
	fmt.Println("=== Basic HTTP Client Operations ===")
	
	// Create a client
	client := NewHTTPClient("")
	client.SetHeader("User-Agent", "Go-HTTP-Example/1.0")
	client.SetTimeout(10 * time.Second)
	
	// Example 1: GET request to a public API
	fmt.Println("1. GET request example:")
	demonstrateGET(client)
	
	// Example 2: POST request simulation
	fmt.Println("\n2. POST request example:")
	demonstratePOST(client)
	
	// Example 3: Error handling
	fmt.Println("\n3. Error handling:")
	demonstrateErrorHandling(client)
	
	// Example 4: Query parameters
	fmt.Println("\n4. Query parameters:")
	demonstrateQueryParams(client)
}

// demonstrateGET shows GET request example
func demonstrateGET(client *HTTPClient) {
	// Using httpbin.org for testing (a popular HTTP testing service)
	resp, err := client.GET("https://httpbin.org/json")
	if err != nil {
		fmt.Printf("  Error making GET request: %v\n", err)
		return
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		fmt.Printf("  Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Content-Type: %s\n", response.Header.Get("Content-Type"))
	fmt.Printf("  Response body (first 200 chars): %s...\n", 
		truncateString(response.String(), 200))
	
	// Parse JSON response
	var jsonData map[string]interface{}
	if err := response.JSON(&jsonData); err != nil {
		fmt.Printf("  Error parsing JSON: %v\n", err)
		return
	}
	
	fmt.Printf("  JSON keys: %v\n", getMapKeys(jsonData))
}

// demonstratePOST shows POST request example
func demonstratePOST(client *HTTPClient) {
	// Sample data to send
	postData := map[string]interface{}{
		"name":    "John Doe",
		"email":   "john@example.com",
		"age":     30,
		"active":  true,
		"tags":    []string{"developer", "golang", "backend"},
	}
	
	resp, err := client.POST("https://httpbin.org/post", postData)
	if err != nil {
		fmt.Printf("  Error making POST request: %v\n", err)
		return
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		fmt.Printf("  Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Success: %t\n", response.IsSuccess())
	
	// Parse the response to see what was sent
	var responseData struct {
		JSON map[string]interface{} `json:"json"`
		URL  string                 `json:"url"`
	}
	
	if err := response.JSON(&responseData); err != nil {
		fmt.Printf("  Error parsing response JSON: %v\n", err)
		return
	}
	
	fmt.Printf("  Sent to URL: %s\n", responseData.URL)
	fmt.Printf("  Received data: %v\n", responseData.JSON)
}

// demonstrateErrorHandling shows error handling patterns
func demonstrateErrorHandling(client *HTTPClient) {
	// Try to access a non-existent endpoint
	resp, err := client.GET("https://httpbin.org/status/404")
	if err != nil {
		fmt.Printf("  Network error: %v\n", err)
		return
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		fmt.Printf("  Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Is success: %t\n", response.IsSuccess())
	
	if !response.IsSuccess() {
		fmt.Printf("  Request failed with status code: %d\n", response.StatusCode)
	}
	
	// Try with an invalid URL
	resp, err = client.GET("invalid-url")
	if err != nil {
		fmt.Printf("  Invalid URL error: %v\n", err)
	}
}

// demonstrateQueryParams shows how to work with query parameters
func demonstrateQueryParams(client *HTTPClient) {
	// Build URL with query parameters
	baseURL := "https://httpbin.org/get"
	params := url.Values{}
	params.Add("name", "Alice")
	params.Add("city", "New York")
	params.Add("lang", "go")
	params.Add("lang", "python") // Multiple values for same key
	
	fullURL := baseURL + "?" + params.Encode()
	fmt.Printf("  Request URL: %s\n", fullURL)
	
	resp, err := client.GET(fullURL)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
		return
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		fmt.Printf("  Error reading response: %v\n", err)
		return
	}
	
	var responseData struct {
		Args map[string]interface{} `json:"args"`
		URL  string                 `json:"url"`
	}
	
	if err := response.JSON(&responseData); err != nil {
		fmt.Printf("  Error parsing JSON: %v\n", err)
		return
	}
	
	fmt.Printf("  Parsed query parameters: %v\n", responseData.Args)
}

// APIClient demonstrates a more realistic API client
type APIClient struct {
	httpClient *HTTPClient
	apiKey     string
}

// NewAPIClient creates a new API client
func NewAPIClient(baseURL, apiKey string) *APIClient {
	client := NewHTTPClient(baseURL)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	client.SetHeader("Accept", "application/json")
	
	return &APIClient{
		httpClient: client,
		apiKey:     apiKey,
	}
}

// User represents a user in our API
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Website  string `json:"website,omitempty"`
}

// GetUsers fetches all users
func (api *APIClient) GetUsers() ([]User, error) {
	resp, err := api.httpClient.GET("/users")
	if err != nil {
		return nil, err
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		return nil, err
	}
	
	if !response.IsSuccess() {
		return nil, fmt.Errorf("API request failed: %s", response.Status)
	}
	
	var users []User
	if err := response.JSON(&users); err != nil {
		return nil, err
	}
	
	return users, nil
}

// GetUser fetches a specific user by ID
func (api *APIClient) GetUser(userID int) (*User, error) {
	endpoint := fmt.Sprintf("/users/%d", userID)
	resp, err := api.httpClient.GET(endpoint)
	if err != nil {
		return nil, err
	}
	
	response, err := NewResponse(resp)
	if err != nil {
		return nil, err
	}
	
	if !response.IsSuccess() {
		return nil, fmt.Errorf("API request failed: %s", response.Status)
	}
	
	var user User
	if err := response.JSON(&user); err != nil {
		return nil, err
	}
	
	return &user, nil
}

// DemonstrateAPIClient shows a realistic API client example
func DemonstrateAPIClient() {
	fmt.Println("\n=== API Client Example ===")
	
	// Using JSONPlaceholder - a free fake JSON API for testing
	apiClient := NewAPIClient("https://jsonplaceholder.typicode.com", "fake-api-key")
	
	// Get all users
	fmt.Println("1. Fetching all users:")
	users, err := apiClient.GetUsers()
	if err != nil {
		fmt.Printf("  Error fetching users: %v\n", err)
		return
	}
	
	fmt.Printf("  Found %d users:\n", len(users))
	for i, user := range users[:3] { // Show first 3
		fmt.Printf("    %d. %s (%s) - %s\n", i+1, user.Name, user.Username, user.Email)
	}
	fmt.Printf("    ... and %d more\n", len(users)-3)
	
	// Get specific user
	fmt.Println("\n2. Fetching specific user:")
	user, err := apiClient.GetUser(1)
	if err != nil {
		fmt.Printf("  Error fetching user: %v\n", err)
		return
	}
	
	fmt.Printf("  User details:\n")
	fmt.Printf("    ID: %d\n", user.ID)
	fmt.Printf("    Name: %s\n", user.Name)
	fmt.Printf("    Email: %s\n", user.Email)
	fmt.Printf("    Username: %s\n", user.Username)
	if user.Website != "" {
		fmt.Printf("    Website: %s\n", user.Website)
	}
}

// Helper functions

// truncateString truncates a string to the specified length
func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}

// getMapKeys returns the keys of a map[string]interface{}
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}