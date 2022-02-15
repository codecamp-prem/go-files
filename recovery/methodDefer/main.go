package main

import (
	"fmt"
	"log"
)

type refrigerator []string

func (r refrigerator) open() {
	fmt.Println("refrigator opening")
}

func (r refrigerator) close() {
	fmt.Println("refrigator closing")
}

func (r refrigerator) findFood(food string) error {
	r.open()
	defer r.close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found!", food)
	}
	//r.close() // No need of these when it is defer after r.open()
	return nil
}

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}

func main() {
	fridge := refrigerator{"corndog", "mushroom", "shakes"}
	for _, food := range []string{"corndog", "shakes"} {
		err := fridge.findFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
