# Go Hello - Comprehensive Golang Basics Project

A beginner-friendly project that demonstrates fundamental Go programming concepts with practical examples. This project is designed to help developers understand Go fundamentals through well-commented, executable code examples.

## 🎯 Learning Objectives

This project covers the essential concepts every Go developer should know:

### 1. **Basic Go Syntax & Concepts**
- Variables and data types (int, string, bool, arrays, slices, maps)
- Constants and iota
- Control structures (if/else, for loops, switch statements)
- Functions with multiple return values
- Pointers basics

### 2. **Object-Oriented Programming in Go**
- Structs and methods
- Interfaces and implementation
- Embedding and composition

### 3. **Error Handling**
- Error interface implementation
- Custom errors
- Panic and recover

### 4. **Concurrency**
- Goroutines basics
- Channels for communication
- Select statements
- WaitGroups

### 5. **Practical Examples**
- CLI calculator
- File I/O operations
- HTTP client
- JSON marshaling/unmarshaling

## 🗂️ Project Structure

```
gohello/
├── README.md                    # This file
├── go.mod                       # Module definition
├── main.go                      # Interactive menu to run examples
├── basics/                      # Basic Go concepts
│   ├── variables.go            # Variables and data types
│   ├── constants.go            # Constants and iota
│   ├── control.go              # Control structures
│   ├── functions.go            # Functions and multiple returns
│   ├── pointers.go             # Pointer basics
│   └── *_test.go               # Unit tests
├── oop/                        # Object-oriented concepts
│   ├── structs.go              # Structs and methods
│   ├── interfaces.go           # Interfaces
│   ├── embedding.go            # Embedding and composition
│   └── *_test.go               # Unit tests
├── errors/                     # Error handling
│   ├── basic_errors.go         # Basic error handling
│   ├── custom_errors.go        # Custom error types
│   ├── panic_recover.go        # Panic and recover
│   └── *_test.go               # Unit tests
├── concurrency/                # Concurrency concepts
│   ├── goroutines.go           # Goroutines basics
│   ├── channels.go             # Channel communication
│   ├── select_stmt.go          # Select statements
│   ├── waitgroups.go           # WaitGroups
│   └── *_test.go               # Unit tests
└── examples/                   # Practical examples
    ├── calculator/             # CLI calculator
    ├── fileio/                 # File operations
    ├── httpclient/             # HTTP client
    └── json/                   # JSON operations
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/adolfoweloy/gohello.git
   cd gohello
   ```

2. Run the interactive menu:
   ```bash
   go run main.go
   ```

3. Or run specific examples:
   ```bash
   # Run calculator example
   go run examples/calculator/main.go
   
   # Run file I/O example
   go run examples/fileio/main.go
   
   # Run HTTP client example
   go run examples/httpclient/main.go
   
   # Run JSON example
   go run examples/json/main.go
   ```

### Running Tests
Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./basics
go test ./oop
go test ./errors
go test ./concurrency
```

## 📚 Learning Path

We recommend following this learning path:

1. **Start with Basics** (`basics/` package)
   - Understand variables, data types, and constants
   - Learn control structures and functions
   - Grasp pointer concepts

2. **Object-Oriented Concepts** (`oop/` package)
   - Learn about structs and methods
   - Understand interfaces
   - Explore embedding and composition

3. **Error Handling** (`errors/` package)
   - Basic error handling patterns
   - Creating custom errors
   - Understanding panic and recover

4. **Concurrency** (`concurrency/` package)
   - Start with goroutines
   - Learn channel communication
   - Understand select statements and WaitGroups

5. **Practice with Examples** (`examples/` directory)
   - Build a calculator
   - Work with files
   - Make HTTP requests
   - Handle JSON data

## 🛠️ Code Quality

This project follows Go best practices:
- **Formatted code**: All code is formatted with `go fmt`
- **Linted code**: Code passes `go vet` checks
- **Well-tested**: Comprehensive unit tests with examples
- **Documented**: Clear comments explaining concepts
- **Idiomatic Go**: Following Go conventions and patterns

## 🤝 Contributing

Feel free to contribute by:
- Adding more examples
- Improving documentation
- Fixing bugs
- Suggesting new features

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 📖 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

Happy coding! 🎉