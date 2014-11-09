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
	// unicode stuff!
	{"Schüßler", "Schübler", 1},
	{"Schüßler", "Schußler", 1},
	{"Schüßler", "Schüßler", 0},
	{"Schßüler", "Schüßler", 1},
	{"Schüßler", "Schüler", 1},
	{"Schüßler", "Schüßlers", 1},
	// difference between DL and OSA. This is DL, so it should be 2.
	{"ca", "abc", 2},
}

// Damerau-Levenshtein
func TestDamerauLevenshtein(t *testing.T) {
	for _, tt := range damlevtests {
		dist := DamerauLevenshtein(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("DamerauLevenshtein('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
