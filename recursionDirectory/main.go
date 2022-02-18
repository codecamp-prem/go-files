package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func panicReport() {
	p := recover() //Go offers a built-in recover function that can stop a program from panicking. Call "recover" and stores its return value
	if p == nil {
		return
	}
	err, ok := p.(error) // type assertion : get the underlying "error" value
	if ok {
		fmt.Println(err)
	} else {
		/* A common strategy for dealing with unanticipated panics you’re not prepared to recover from is to simply renew the panic state. Panicking again is usually appropriate because, after all, this is an unanticipated situation.*/
		panic(p)
	}
}
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
	defer panicReport()                       // Before calling code that might panic, defer a call to our new panicReport function
	scanDirectory("<path to your directory>") // "Users/<UserName>/desktop/<Directory>" in mac

}
