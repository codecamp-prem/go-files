// Package numbers Checks if a number is prime or not
package numbers

import "math"

// IsPrime takes int and return bool
// true if int is Prime or false if int is Not Prime
func IsPrime(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
