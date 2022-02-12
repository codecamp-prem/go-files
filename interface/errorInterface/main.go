package main

import (
	"fmt"
	"log"
)

// OverheatError : defined type
type OverheatError float64

func (o OverheatError) Error() string {
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
	var err error = checkTemperature(127.00, 100)
	if err != nil {
		log.Fatal(err)
	}
}
