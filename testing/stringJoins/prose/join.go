package prose

import "strings"

// JoinWithCommas : joins strings with , and
func JoinWithCommas(pharses []string) string {
	if len(pharses) == 0 {
		return ""
	}
	if len(pharses) == 1 {
		return pharses[0]
	}
	if len(pharses) == 2 {
		return pharses[0] + " and " + pharses[1]
	}
	result := strings.Join(pharses[:len(pharses)-1], ", ")
	result += ", and "
	result += pharses[len(pharses)-1]
	return result
}
