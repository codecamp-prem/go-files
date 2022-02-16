package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path) // Get the slice with directory content
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // join directory and file name
		if file.IsDir() {                            // if this is a subdirectory
			err := scanDirectory(filePath) // recursively call scanDirectory, this time with the subdirectory path
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath) // just print regular file
		}
	}
	return nil
}

func main() {
	err := scanDirectory("<full path to your folder>")
	if err != nil {
		log.Fatal(err)
	}
}
