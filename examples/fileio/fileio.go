// Package fileio demonstrates file input/output operations in Go
package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileManager handles file operations
type FileManager struct {
	WorkingDir string
}

// NewFileManager creates a new file manager
func NewFileManager() *FileManager {
	wd, _ := os.Getwd()
	return &FileManager{
		WorkingDir: wd,
	}
}

// DemonstrateBasicFileOperations shows basic file I/O
func DemonstrateBasicFileOperations() {
	fmt.Println("=== Basic File Operations ===")
	
	// Create a temporary directory for examples
	tempDir := "/tmp/go-file-examples"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Printf("Error creating temp dir: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up
	
	fm := &FileManager{WorkingDir: tempDir}
	
	// Writing to files
	fmt.Println("1. Writing to files:")
	fm.demonstrateWriting()
	
	// Reading from files
	fmt.Println("\n2. Reading from files:")
	fm.demonstrateReading()
	
	// File information
	fmt.Println("\n3. File information:")
	fm.demonstrateFileInfo()
	
	// Directory operations
	fmt.Println("\n4. Directory operations:")
	fm.demonstrateDirectoryOps()
	
	// File manipulation
	fmt.Println("\n5. File manipulation:")
	fm.demonstrateFileManipulation()
}

// demonstrateWriting shows different ways to write files
func (fm *FileManager) demonstrateWriting() {
	// Method 1: Using os.WriteFile (simple)
	content1 := "Hello, Go!\nThis is a simple file write example.\nLine 3 of the file."
	filePath1 := filepath.Join(fm.WorkingDir, "simple.txt")
	
	err := os.WriteFile(filePath1, []byte(content1), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Printf("  Written to %s using os.WriteFile\n", filePath1)
	
	// Method 2: Using os.Create and Write
	filePath2 := filepath.Join(fm.WorkingDir, "detailed.txt")
	file, err := os.Create(filePath2)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	lines := []string{
		"This file was created using os.Create",
		"We can write multiple lines",
		"And control the writing process",
		fmt.Sprintf("Created at: %s", time.Now().Format(time.RFC3339)),
	}
	
	for i, line := range lines {
		_, err := fmt.Fprintf(file, "%d: %s\n", i+1, line)
		if err != nil {
			fmt.Printf("Error writing line: %v\n", err)
			return
		}
	}
	fmt.Printf("  Written to %s using os.Create\n", filePath2)
	
	// Method 3: Using bufio.Writer for buffered writing
	filePath3 := filepath.Join(fm.WorkingDir, "buffered.txt")
	file3, err := os.Create(filePath3)
	if err != nil {
		fmt.Printf("Error creating buffered file: %v\n", err)
		return
	}
	defer file3.Close()
	
	writer := bufio.NewWriter(file3)
	
	for i := 1; i <= 5; i++ {
		line := fmt.Sprintf("Buffered line %d: %s\n", i, strings.Repeat("data ", i))
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing buffered line: %v\n", err)
			return
		}
	}
	
	err = writer.Flush() // Important: flush buffered data
	if err != nil {
		fmt.Printf("Error flushing writer: %v\n", err)
		return
	}
	fmt.Printf("  Written to %s using bufio.Writer\n", filePath3)
}

// demonstrateReading shows different ways to read files
func (fm *FileManager) demonstrateReading() {
	// Method 1: Read entire file
	filePath := filepath.Join(fm.WorkingDir, "simple.txt")
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("  File content (os.ReadFile):\n%s\n", string(content))
	
	// Method 2: Read line by line
	fmt.Println("  Reading line by line:")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("    Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
	
	// Method 3: Read with buffer
	fmt.Println("  Reading with buffer:")
	file2, err := os.Open(filepath.Join(fm.WorkingDir, "detailed.txt"))
	if err != nil {
		fmt.Printf("Error opening detailed file: %v\n", err)
		return
	}
	defer file2.Close()
	
	buffer := make([]byte, 64)
	for {
		n, err := file2.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading buffer: %v\n", err)
			break
		}
		fmt.Printf("    Read %d bytes: %s", n, string(buffer[:n]))
	}
}

// demonstrateFileInfo shows how to get file information
func (fm *FileManager) demonstrateFileInfo() {
	filePath := filepath.Join(fm.WorkingDir, "simple.txt")
	
	// Get file info
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	
	fmt.Printf("  File: %s\n", info.Name())
	fmt.Printf("  Size: %d bytes\n", info.Size())
	fmt.Printf("  Mode: %s\n", info.Mode())
	fmt.Printf("  Modified: %s\n", info.ModTime().Format(time.RFC3339))
	fmt.Printf("  Is directory: %t\n", info.IsDir())
	
	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("  File exists: %s\n", filePath)
	} else if os.IsNotExist(err) {
		fmt.Printf("  File does not exist: %s\n", filePath)
	} else {
		fmt.Printf("  Error checking file: %v\n", err)
	}
}

// demonstrateDirectoryOps shows directory operations
func (fm *FileManager) demonstrateDirectoryOps() {
	// Create directories
	subDir := filepath.Join(fm.WorkingDir, "subdir", "nested")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}
	fmt.Printf("  Created directory: %s\n", subDir)
	
	// List directory contents
	fmt.Println("  Directory contents:")
	entries, err := os.ReadDir(fm.WorkingDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting entry info: %v\n", err)
			continue
		}
		
		typeStr := "file"
		if entry.IsDir() {
			typeStr = "dir"
		}
		fmt.Printf("    %s: %s (%d bytes)\n", typeStr, entry.Name(), info.Size())
	}
	
	// Walk directory tree
	fmt.Println("  Walking directory tree:")
	err = filepath.Walk(fm.WorkingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		relPath, _ := filepath.Rel(fm.WorkingDir, path)
		indent := strings.Repeat("  ", strings.Count(relPath, string(os.PathSeparator)))
		
		if info.IsDir() {
			fmt.Printf("    %s📁 %s/\n", indent, info.Name())
		} else {
			fmt.Printf("    %s📄 %s (%d bytes)\n", indent, info.Name(), info.Size())
		}
		
		return nil
	})
	
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
}

// demonstrateFileManipulation shows file manipulation operations
func (fm *FileManager) demonstrateFileManipulation() {
	// Copy file
	srcFile := filepath.Join(fm.WorkingDir, "simple.txt")
	dstFile := filepath.Join(fm.WorkingDir, "simple_copy.txt")
	
	if err := fm.copyFile(srcFile, dstFile); err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
	fmt.Printf("  Copied %s to %s\n", srcFile, dstFile)
	
	// Rename file
	newName := filepath.Join(fm.WorkingDir, "renamed_file.txt")
	if err := os.Rename(dstFile, newName); err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
		return
	}
	fmt.Printf("  Renamed %s to %s\n", dstFile, newName)
	
	// Append to file
	file, err := os.OpenFile(newName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file for append: %v\n", err)
		return
	}
	defer file.Close()
	
	appendText := "\nThis line was appended!"
	if _, err := file.WriteString(appendText); err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}
	fmt.Printf("  Appended text to %s\n", newName)
	
	// Delete file
	if err := os.Remove(newName); err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}
	fmt.Printf("  Deleted %s\n", newName)
}

// copyFile copies a file from src to dst
func (fm *FileManager) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()
	
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	
	// Copy file permissions
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	
	return os.Chmod(dst, sourceInfo.Mode())
}

// LogManager demonstrates logging to files
type LogManager struct {
	logFile *os.File
	logger  *bufio.Writer
}

// NewLogManager creates a new log manager
func NewLogManager(logPath string) (*LogManager, error) {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	
	return &LogManager{
		logFile: file,
		logger:  bufio.NewWriter(file),
	}, nil
}

// Log writes a log entry
func (lm *LogManager) Log(level, message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
	
	_, err := lm.logger.WriteString(logEntry)
	if err != nil {
		return err
	}
	
	return lm.logger.Flush()
}

// Close closes the log manager
func (lm *LogManager) Close() error {
	if err := lm.logger.Flush(); err != nil {
		return err
	}
	return lm.logFile.Close()
}

// DemonstrateLogging shows file-based logging
func DemonstrateLogging() {
	fmt.Println("\n=== File Logging Example ===")
	
	logPath := "/tmp/app.log"
	logger, err := NewLogManager(logPath)
	if err != nil {
		fmt.Printf("Error creating logger: %v\n", err)
		return
	}
	defer logger.Close()
	
	// Write some log entries
	logger.Log("INFO", "Application started")
	logger.Log("DEBUG", "Processing user request")
	logger.Log("WARN", "Low disk space detected")
	logger.Log("ERROR", "Failed to connect to database")
	logger.Log("INFO", "Application shutting down")
	
	fmt.Printf("  Log entries written to %s\n", logPath)
	
	// Read and display log contents
	content, err := os.ReadFile(logPath)
	if err != nil {
		fmt.Printf("Error reading log file: %v\n", err)
		return
	}
	
	fmt.Println("  Log contents:")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line != "" {
			fmt.Printf("    %s\n", line)
		}
	}
}