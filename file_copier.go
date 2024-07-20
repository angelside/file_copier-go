package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// go run file_copier.go source.txt ./test/source.txt

func main() {
	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	// Create the destination directory recursively if it does not exist
	destinationDir := filepath.Dir(destinationFile)
	err := os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open the source file
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(destinationFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destination.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File copied successfully!")
}
