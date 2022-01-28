package strings

// Reverses the strings

/*
Since strings in GO are immutable, we first convert the sting to a mutable array
to perform the reverse operation on that, and then re-cast to a string
*/

func Reverse(s string) string {
	runes := []rune(s)
	reversedRunes := reverseRunes(runes)
	return string(reversedRunes)
}
