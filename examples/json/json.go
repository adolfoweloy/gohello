// Package jsonexample demonstrates JSON marshaling and unmarshaling in Go
package jsonexample

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Person represents a person with various data types
type Person struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	IsActive    bool      `json:"is_active"`
	Salary      float64   `json:"salary,omitempty"`     // omitempty means don't include if zero value
	Department  string    `json:"department,omitempty"`
	JoinDate    time.Time `json:"join_date"`
	Skills      []string  `json:"skills"`
	Address     Address   `json:"address"`
	Metadata    Meta      `json:"metadata,omitempty"`
}

// Address represents a physical address
type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
	Country  string `json:"country"`
}

// Meta represents additional metadata
type Meta struct {
	CreatedBy   string                 `json:"created_by"`
	CreatedAt   time.Time             `json:"created_at"`
	Tags        []string              `json:"tags,omitempty"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
}

// Company represents a company with employees
type Company struct {
	Name        string    `json:"name"`
	Founded     int       `json:"founded"`
	Employees   []Person  `json:"employees"`
	Headquarters Address  `json:"headquarters"`
	Revenue     float64   `json:"revenue"`
	IsPublic    bool      `json:"is_public"`
}

// DemonstrateBasicJSON shows basic JSON operations
func DemonstrateBasicJSON() {
	fmt.Println("=== Basic JSON Operations ===")
	
	// Create sample data
	person := Person{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
		IsActive:  true,
		Salary:    75000.50,
		Department: "Engineering",
		JoinDate:  time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
		Skills:    []string{"Go", "Python", "Docker", "Kubernetes"},
		Address: Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94102",
			Country: "USA",
		},
		Metadata: Meta{
			CreatedBy: "admin",
			CreatedAt: time.Now(),
			Tags:      []string{"developer", "senior"},
			Preferences: map[string]interface{}{
				"theme":           "dark",
				"notifications":   true,
				"max_work_hours":  8,
			},
		},
	}
	
	fmt.Println("1. Marshaling (Go -> JSON):")
	demonstrateMarshaling(person)
	
	fmt.Println("\n2. Unmarshaling (JSON -> Go):")
	demonstrateUnmarshaling()
	
	fmt.Println("\n3. Working with arrays:")
	demonstrateArrays()
	
	fmt.Println("\n4. Custom JSON handling:")
	demonstrateCustomJSON()
}

// demonstrateMarshaling shows different ways to marshal JSON
func demonstrateMarshaling(person Person) {
	// Basic marshaling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("  Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("  Compact JSON (%d bytes):\n  %s\n", len(jsonData), jsonData)
	
	// Pretty printing
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("  Error marshaling with indent: %v\n", err)
		return
	}
	fmt.Printf("  Pretty JSON (%d bytes):\n%s\n", len(prettyJSON), prettyJSON)
	
	// Marshal specific fields only
	simpleData := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}{
		Name:  person.FirstName + " " + person.LastName,
		Email: person.Email,
		Age:   person.Age,
	}
	
	simpleJSON, err := json.MarshalIndent(simpleData, "", "  ")
	if err != nil {
		fmt.Printf("  Error marshaling simple data: %v\n", err)
		return
	}
	fmt.Printf("  Simple structure:\n%s\n", simpleJSON)
}

// demonstrateUnmarshaling shows different ways to unmarshal JSON
func demonstrateUnmarshaling() {
	// Sample JSON data
	jsonString := `{
		"id": 2,
		"first_name": "Jane",
		"last_name": "Smith",
		"email": "jane.smith@example.com",
		"age": 28,
		"is_active": true,
		"salary": 82000,
		"department": "Marketing",
		"join_date": "2021-03-10T00:00:00Z",
		"skills": ["JavaScript", "React", "Node.js", "MongoDB"],
		"address": {
			"street": "456 Oak Ave",
			"city": "Austin",
			"state": "TX",
			"zip_code": "78701",
			"country": "USA"
		}
	}`
	
	// Unmarshal into struct
	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Printf("  Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("  Unmarshaled person:\n")
	fmt.Printf("    Name: %s %s\n", person.FirstName, person.LastName)
	fmt.Printf("    Email: %s\n", person.Email)
	fmt.Printf("    Age: %d\n", person.Age)
	fmt.Printf("    Department: %s\n", person.Department)
	fmt.Printf("    Skills: %v\n", person.Skills)
	fmt.Printf("    Address: %s, %s, %s %s\n", 
		person.Address.Street, person.Address.City, person.Address.State, person.Address.ZipCode)
	
	// Unmarshal into map for dynamic handling
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Printf("  Error unmarshaling to map: %v\n", err)
		return
	}
	
	fmt.Printf("  Dynamic access:\n")
	fmt.Printf("    ID: %.0f\n", data["id"].(float64)) // JSON numbers are float64
	fmt.Printf("    Name: %s %s\n", data["first_name"], data["last_name"])
	fmt.Printf("    Skills: %v\n", data["skills"])
	
	// Access nested objects
	if address, ok := data["address"].(map[string]interface{}); ok {
		fmt.Printf("    City: %s\n", address["city"])
	}
}

// demonstrateArrays shows working with JSON arrays
func demonstrateArrays() {
	// Create sample company with employees
	company := Company{
		Name:    "TechCorp Inc.",
		Founded: 2015,
		Revenue: 50000000.00,
		IsPublic: true,
		Headquarters: Address{
			Street:  "100 Tech Blvd",
			City:    "Silicon Valley",
			State:   "CA",
			ZipCode: "94301",
			Country: "USA",
		},
		Employees: []Person{
			{
				ID:         1,
				FirstName:  "Alice",
				LastName:   "Johnson",
				Email:      "alice@techcorp.com",
				Age:        32,
				IsActive:   true,
				Department: "Engineering",
				Skills:     []string{"Go", "Rust", "Docker"},
			},
			{
				ID:         2,
				FirstName:  "Bob",
				LastName:   "Wilson",
				Email:      "bob@techcorp.com",
				Age:        29,
				IsActive:   true,
				Department: "Product",
				Skills:     []string{"JavaScript", "React", "Node.js"},
			},
			{
				ID:         3,
				FirstName:  "Carol",
				LastName:   "Brown",
				Email:      "carol@techcorp.com",
				Age:        35,
				IsActive:   false,
				Department: "Sales",
				Skills:     []string{"CRM", "Salesforce", "Communication"},
			},
		},
	}
	
	// Marshal company
	companyJSON, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		fmt.Printf("  Error marshaling company: %v\n", err)
		return
	}
	
	fmt.Printf("  Company JSON (first 500 chars):\n%s...\n", 
		string(companyJSON)[:min(500, len(companyJSON))])
	
	// Unmarshal and process employees
	var unmarshaled Company
	err = json.Unmarshal(companyJSON, &unmarshaled)
	if err != nil {
		fmt.Printf("  Error unmarshaling company: %v\n", err)
		return
	}
	
	fmt.Printf("  Company: %s (founded %d)\n", unmarshaled.Name, unmarshaled.Founded)
	fmt.Printf("  Employees (%d total):\n", len(unmarshaled.Employees))
	
	for i, emp := range unmarshaled.Employees {
		status := "Active"
		if !emp.IsActive {
			status = "Inactive"
		}
		fmt.Printf("    %d. %s %s (%s) - %s - %v\n", 
			i+1, emp.FirstName, emp.LastName, emp.Department, status, emp.Skills)
	}
}

// CustomTime demonstrates custom JSON marshaling/unmarshaling
type CustomTime struct {
	time.Time
}

// MarshalJSON implements custom JSON marshaling for CustomTime
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format("2006-01-02"))
}

// UnmarshalJSON implements custom JSON unmarshaling for CustomTime
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	
	ct.Time = t
	return nil
}

// Employee demonstrates custom JSON handling
type Employee struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	HireDate CustomTime `json:"hire_date"`
	Salary   int        `json:"-"`              // Excluded from JSON
	Internal bool       `json:"internal,string"` // Convert bool to string
}

// demonstrateCustomJSON shows custom JSON handling
func demonstrateCustomJSON() {
	employee := Employee{
		ID:       100,
		Name:     "David Miller",
		HireDate: CustomTime{time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC)},
		Salary:   90000, // This will be excluded
		Internal: true,
	}
	
	// Marshal with custom handling
	jsonData, err := json.MarshalIndent(employee, "", "  ")
	if err != nil {
		fmt.Printf("  Error marshaling employee: %v\n", err)
		return
	}
	
	fmt.Printf("  Employee with custom JSON:\n%s\n", jsonData)
	
	// Unmarshal back
	var newEmployee Employee
	err = json.Unmarshal(jsonData, &newEmployee)
	if err != nil {
		fmt.Printf("  Error unmarshaling employee: %v\n", err)
		return
	}
	
	fmt.Printf("  Unmarshaled employee:\n")
	fmt.Printf("    ID: %d\n", newEmployee.ID)
	fmt.Printf("    Name: %s\n", newEmployee.Name)
	fmt.Printf("    Hire Date: %s\n", newEmployee.HireDate.Format("2006-01-02"))
	fmt.Printf("    Internal: %t\n", newEmployee.Internal)
	fmt.Printf("    Salary: %d (excluded from JSON)\n", newEmployee.Salary)
}

// DemonstrateFileJSON shows reading and writing JSON files
func DemonstrateFileJSON() {
	fmt.Println("\n=== JSON File Operations ===")
	
	// Create sample data
	config := map[string]interface{}{
		"app_name":    "Go JSON Example",
		"version":     "1.0.0",
		"debug":       true,
		"port":        8080,
		"database": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"username": "admin",
			"database": "app_db",
		},
		"features": []string{"auth", "logging", "metrics"},
		"limits": map[string]interface{}{
			"max_connections": 100,
			"timeout":         30,
		},
	}
	
	// Write JSON to file
	filename := "/tmp/config.json"
	fmt.Printf("1. Writing JSON to file: %s\n", filename)
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("  Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	
	if err := encoder.Encode(config); err != nil {
		fmt.Printf("  Error encoding JSON: %v\n", err)
		return
	}
	
	fmt.Printf("  JSON written successfully\n")
	
	// Read JSON from file
	fmt.Printf("\n2. Reading JSON from file:\n")
	
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("  Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	var loadedConfig map[string]interface{}
	decoder := json.NewDecoder(file)
	
	if err := decoder.Decode(&loadedConfig); err != nil {
		fmt.Printf("  Error decoding JSON: %v\n", err)
		return
	}
	
	fmt.Printf("  Loaded configuration:\n")
	fmt.Printf("    App: %s v%s\n", loadedConfig["app_name"], loadedConfig["version"])
	fmt.Printf("    Debug: %t\n", loadedConfig["debug"])
	fmt.Printf("    Port: %.0f\n", loadedConfig["port"].(float64))
	
	if db, ok := loadedConfig["database"].(map[string]interface{}); ok {
		fmt.Printf("    Database: %s@%s:%.0f/%s\n", 
			db["username"], db["host"], db["port"].(float64), db["database"])
	}
	
	if features, ok := loadedConfig["features"].([]interface{}); ok {
		fmt.Printf("    Features: ")
		for i, feature := range features {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", feature)
		}
		fmt.Println()
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}