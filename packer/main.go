// Have custom numbers and strings Package
/*
1. numbers package check if the int is prime or not
2. strings package reverse runes and it's subpackage greeting have
constant that greets
*/
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
