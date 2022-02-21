package prose

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) { //function will passed a pointer to testing T value
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "mango"}, want: "apple, orange, and mango"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
