// Encapsulation and Embedding
// Setter and Getter
// Method promotion allows you to easily use one type’s methods as if they
//belonged to another.
package main

import (
	"fmt"
	"log"

	"github.com/codecamp-prem/remindMe/calender"
)

func main() {
	event := calender.Event{}
	err := event.SetTitle("Parent's Anniversary!")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(02)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(07)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
