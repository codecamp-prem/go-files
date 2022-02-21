package prose

import "strings"

// JoinWithCommas : joins strings with , and
func JoinWithCommas(pharses []string) string {
	result := strings.Join(pharses[:len(pharses)-1], ", ")
	result += ", and "
	result += pharses[len(pharses)-1]
	return result
}
