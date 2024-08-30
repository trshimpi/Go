package main

import (
	"fmt"
	"os"
)

func main() {

	// Open the file "example.txt" for reading
	// os.Open returns a file descriptor (f) and an error (err)
	f, err := os.Open("example.txt")

	// Check if there was an error opening the file
	if err != nil {
		// If there was an error, panic and stop the program
		panic(err)
	}

	// Get file information using the Stat() method
	// This returns a FileInfo struct and an error
	fileInfo, err := f.Stat()

	// Check if there was an error getting file information
	if err != nil {
		// If there was an error, panic and stop the program
		panic(err)
	}

	// Print various file attributes using methods of the FileInfo struct
	fmt.Println("file name", fileInfo.Name())             // Name of the file
	fmt.Println("file size", fileInfo.Size())             // Size of the file in bytes
	fmt.Println("file permissions", fileInfo.Mode())      // File permissions
	fmt.Println("file modified time", fileInfo.ModTime()) // Last modification time
	fmt.Println("file is dir", fileInfo.IsDir())          // Whether it's a directory or not

}
