package matchr

import "testing"

var lcstests = []struct {
	s1     string
	s2     string
	length int
}{
	// match beginning
	{"cans", "can", 3},
	// match end
	{"ebay", "bay", 3},
	// gap in the middle
	{"coins", "cons", 4},
	// one empty, left
	{"", "hello", 0},
	// one empty, right
	{"goodbye", "", 0},
	// two empties
	{"", "", 0},
	// unicode stuff!
	{"Schüßler", "Schüßler", 8},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tt := range lcstests {
		length := LongestCommonSubsequence(tt.s1, tt.s2)
		if length != tt.length {
			t.Errorf("LongestCommonSubsequence('%s', '%s') = %v, want %v", tt.s1, tt.s2, length, tt.length)
		}
	}
}
