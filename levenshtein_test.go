package matchr

import "testing"

var levtests = []struct {
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
	{"Schüßler", "Schüler", 1},
	{"Schüßler", "Schüßlers", 1},
}

// Regular Levenshtein
func TestLevenshtein(t *testing.T) {
	for _, tt := range levtests {
		dist := Levenshtein(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("Levenshtein('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
