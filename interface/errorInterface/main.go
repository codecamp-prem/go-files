package main

import (
	"fmt"
	"log"
)

type error interface {
	Error() string
}

// OverheatError : defined type
type OverheatError float64

func (o OverheatError) Error() string { // Satisfy the error interface
	return fmt.Sprintf("Overheating by %0.2f", o)
}

func checkTemperature(actual float64, safe float64) error { // function returns ordinary error value
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err = checkTemperature(127.00, 100)
	if err != nil {
		log.Fatal(err)
	}
}
