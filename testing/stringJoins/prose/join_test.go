package prose

import "testing"

func TestTwoElements(t *testing.T) { //function will passed a pointer to testing T value
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "mango"}
	if JoinWithCommas(list) != "apple, orange, and mango" {
		t.Error("didn't match expected value")
	}
}
