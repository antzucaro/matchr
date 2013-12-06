package matchr

import "testing"

var damlevtests = []struct {
	s1   string
	s2   string
	dist int
}{
	// insertion
	{"car", "cars", 1},
	// substitution
	{"library", "librari", 1},
	// deletion
	{"library", "librar", 1},
	// transposition
	{"library", "librayr", 1},
	// one empty, left
	{"", "library", 7},
	// one empty, right
	{"library", "", 7},
	// two empties
	{"", "", 0},
}

// Damerau Levenshtein
func TestDamerauLevenshtein(t *testing.T) {
	for _, tt := range levtests {
		dist := DamerauLevenshtein(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("DamerauLevenshtein('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
