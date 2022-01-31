package main

import (
	"fmt"

	"github.com/codecamp-prem/definedTypes/newspaper"
)

func printInfo(s *newspaper.Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active? :", s.Active)
	fmt.Println("Street? :", s.Address.Street)
	fmt.Println("City? :", s.Address.City)
	fmt.Println("State? :", s.Address.State)
	fmt.Println("PostalCode? :", s.Address.PostalCode)
}

func applyDiscount(s *newspaper.Subscriber) {
	s.Rate = 4.99
}

func defaultSubscriber(name string, homeAddress map[string]string) *newspaper.Subscriber {
	s := newspaper.Subscriber{
		Name:   name,
		Rate:   7.99,
		Active: true,
		Address: newspaper.Address{
			Street:     homeAddress["Street"],
			City:       homeAddress["City"],
			State:      homeAddress["State"],
			PostalCode: homeAddress["PostalCode"],
		},
	}
	//s.Name = name
	//s.Rate = 4.99
	//s.Active = true
	return &s //return the type Subscriber struct
}

func main() {
	address1 := map[string]string{
		"Street":     "Marine Drive",
		"City":       "Mumbai",
		"State":      "Maharastra",
		"PostalCode": "04200",
	}
	subs1 := defaultSubscriber("Rocket Singh", address1) // This is no longer a struct, it's a struct pointer
	//subs1.Rate = 2.99
	applyDiscount(subs1)
	printInfo(subs1)

	address2 := map[string]string{
		"Street":     "Gurgaon",
		"City":       "Delhi",
		"State":      "New Delhi",
		"PostalCode": "00420",
	}
	subs2 := defaultSubscriber("Honey Singh", address2)
	printInfo(subs2)
}
