# Go Programming Basics

A comprehensive Go programming tutorial project that covers fundamental concepts and practical examples. This project is designed to help beginners understand Go programming from basic syntax to advanced topics like concurrency and error handling.

## 📚 What You'll Learn

### 1. Basic Go Syntax and Concepts
- **Variables and Data Types**: int, string, bool, arrays, slices, maps
- **Constants and Iota**: Define constants and use iota for enumerations
- **Control Structures**: if/else, for loops, switch statements
- **Functions**: Multiple return values, variadic functions, closures
- **Pointers**: Basics of pointer usage and memory management

### 2. Object-Oriented Programming
- **Structs and Methods**: Define custom types and attach methods
- **Interfaces**: Define contracts and achieve polymorphism
- **Embedding and Composition**: Code reuse through struct embedding

### 3. Error Handling
- **Error Interface**: Understanding Go's error handling philosophy
- **Custom Errors**: Creating custom error types
- **Panic and Recover**: Error recovery mechanisms

### 4. Concurrency
- **Goroutines**: Lightweight concurrent execution
- **Channels**: Communication between goroutines
- **Select Statements**: Handling multiple channel operations
- **WaitGroups**: Synchronizing concurrent operations

### 5. Practical Examples
- **CLI Calculator**: Interactive command-line calculator
- **File I/O Operations**: Reading, writing, and manipulating files
- **HTTP Client**: Making HTTP requests and handling responses
- **JSON Processing**: Marshaling and unmarshaling JSON data

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or later
- A text editor or IDE (VS Code with Go extension recommended)

### Installation
1. Clone this repository:
```bash
git clone https://github.com/adolfoweloy/gohello.git
cd gohello
```

2. Verify Go installation:
```bash
go version
```

3. Build the project:
```bash
go build ./...
```

4. Run tests:
```bash
go test ./...
```

## 📖 Project Structure

```
gohello/
├── main.go                    # Main example runner
├── go.mod                     # Go module file
├── README.md                  # This file
├── basics/                    # Basic Go concepts
│   ├── variables.go           # Variables, arrays, slices, maps
│   ├── constants.go           # Constants and iota
│   ├── control.go             # Control structures
│   ├── functions.go           # Functions and closures
│   ├── pointers.go            # Pointer basics
│   └── basics_test.go         # Unit tests
├── oop/                       # Object-oriented programming
│   ├── structs.go             # Structs, methods, embedding
│   └── interfaces.go          # Interfaces and polymorphism
├── errors/                    # Error handling
│   └── errors.go              # Error patterns and panic/recover
├── concurrency/               # Concurrency patterns
│   └── concurrency.go         # Goroutines, channels, select
├── examples/                  # Practical examples
│   ├── calculator/            # CLI calculator
│   │   ├── calculator.go      # Calculator implementation
│   │   └── calculator_test.go # Calculator tests
│   ├── fileio/               # File operations
│   │   └── fileio.go         # File I/O examples
│   ├── httpclient/           # HTTP client
│   │   └── httpclient.go     # HTTP operations
│   └── json/                 # JSON processing
│       └── json.go           # JSON marshaling/unmarshaling
└── cmd/                      # Command-line applications
    └── calculator/           # Calculator CLI
        └── main.go           # Calculator CLI entry point
```

## 🏃‍♂️ Running Examples

### All Examples
Run all examples in sequence:
```bash
go run main.go
```

### Individual Examples

#### 1. Basic Concepts
```bash
# Variables, arrays, maps
go run -c "
import 'github.com/adolfoweloy/gohello/basics'
basics.DemonstrateVariables()
basics.DemonstrateArraysAndSlices()
basics.DemonstrateMaps()
"
```

#### 2. Interactive Calculator
```bash
go run cmd/calculator/main.go
```

Example calculator session:
```
=== Go Calculator ===
Supported operations: +, -, *, /, ^, %
Commands: history, clear, help, quit
Enter expressions like: 5 + 3
=====================
calc> 10 + 5
10.00 + 5.00 = 15.00
calc> 20 * 3
20.00 * 3.00 = 60.00
calc> history
Calculation History:
-------------------
1. 10.00 + 5.00 = 15.00
2. 20.00 * 3.00 = 60.00
calc> quit
```

#### 3. File Operations
The file I/O examples demonstrate:
- Reading and writing files
- Directory operations
- File manipulation (copy, rename, delete)
- Logging to files

#### 4. HTTP Client
The HTTP client examples show:
- Making GET/POST/PUT/DELETE requests
- Handling JSON responses
- Error handling and timeouts
- Working with query parameters

#### 5. JSON Processing
JSON examples cover:
- Marshaling Go structs to JSON
- Unmarshaling JSON to Go structs
- Working with dynamic JSON data
- Custom JSON handling
- Reading/writing JSON files

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run benchmarks:
```bash
go test -bench=. ./...
```

### Test Examples

#### Basic Functions Test
```bash
go test ./basics -v
```

#### Calculator Test
```bash
go test ./examples/calculator -v
```

## 📝 Key Concepts Explained

### 1. Variables and Types
Go is statically typed with type inference:
```go
var name string = "Go"  // Explicit type
age := 25               // Type inference
var isOpen bool         // Zero value (false)
```

### 2. Slices vs Arrays
```go
// Array - fixed size
arr := [3]int{1, 2, 3}

// Slice - dynamic
slice := []int{1, 2, 3}
slice = append(slice, 4)  // Can grow
```

### 3. Error Handling
Go uses explicit error handling:
```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use result
```

### 4. Goroutines and Channels
```go
// Start a goroutine
go func() {
    // Concurrent work
}()

// Communicate via channels
ch := make(chan string)
go func() {
    ch <- "Hello"
}()
msg := <-ch  // Receive from channel
```

### 5. Interfaces
Go interfaces are satisfied implicitly:
```go
type Writer interface {
    Write([]byte) (int, error)
}

// Any type with Write method implements Writer
```

## 🎯 Best Practices Demonstrated

1. **Error Handling**: Always check errors, use custom error types when appropriate
2. **Code Organization**: Use packages to organize related functionality
3. **Testing**: Write unit tests for key functions
4. **Documentation**: Use clear comments and godoc format
5. **Concurrency**: Use goroutines and channels for concurrent operations
6. **Resource Management**: Always close files, connections, etc.

## 🔧 Tools and Commands

### Useful Go Commands
```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Get dependencies
go mod tidy

# Build executable
go build -o calculator cmd/calculator/main.go

# Cross-compile
GOOS=windows GOARCH=amd64 go build -o calculator.exe cmd/calculator/main.go
```

### IDE Setup
For VS Code, install the Go extension and add these settings:
```json
{
    "go.useLanguageServer": true,
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint"
}
```

## 📚 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 🤝 Contributing

Feel free to contribute by:
1. Adding more examples
2. Improving documentation
3. Adding tests
4. Fixing bugs
5. Suggesting improvements

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

Happy Go programming! 🎉