package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path) // Get the slice with directory content
	if err != nil {
		panic(err) // using inbuilt panic function to show what went wrong. All the recursive calls to scanDirectory exit
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // join directory and file name
		if file.IsDir() {                            // if this is a subdirectory
			scanDirectory(filePath) // recursively call scanDirectory, this time with the subdirectory path
		} else {
			fmt.Println(filePath) // just print regular file
		}
	}
	return nil
}

func main() {
	scanDirectory("<path to your directory>") // "Users/<UserName>/desktop/<Directory>" in mac

}
