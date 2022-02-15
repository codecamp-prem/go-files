package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func openFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing File")
	file.Close()
}

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := openFile(fileName)
	if err != nil {
		return nil, err
	}
	defer closeFile(file) // defer right after openFile()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err //closeFile() will also work here
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err() // closeFile() will also work here
	}
	return numbers, nil
}

func main() {
	numbers, err := getFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum : %0.2f\n", sum)
}
