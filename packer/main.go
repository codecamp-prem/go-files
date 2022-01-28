package main

import (
	"fmt"
	str "strings" //package

	"github.com/codecamp-prem/packer/numbers"
	"github.com/codecamp-prem/packer/strings"
	"github.com/codecamp-prem/packer/strings/greetings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greetings.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is awesome. I love Go", "Go"))
}
