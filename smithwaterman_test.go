package matchr

import "testing"

var swtests = []struct {
	s1   string
	s2   string
	dist float64
}{
	// insertion
	{"car", "cars", 3.0},
	// substitution
	{"library", "librari", 6.0},
	// deletion
	{"library", "librar", 6.0},
	// transposition
	{"library", "librayr", 5.5},
	// one empty, left
	{"", "library", 7.0},
	// one empty, right
	{"library", "", 7.0},
	// two empties
	{"", "", 0.0},
	// unicode stuff!
	{"Schüßler", "Schübler", 6.0},
	{"Ant Zucaro", "Anthony Zucaro", 8.0},
	{"Schüßler", "Schüßler", 8.0},
	{"Schßüler", "Schüßler", 6.0},
	{"Schüßler", "Schüler", 6.5},
	{"Schüßler", "Schüßlers", 8.0},
}

// Smith-Waterman
func TestSmithWaterman(t *testing.T) {
	for _, tt := range swtests {
		dist := SmithWaterman(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("SmithWaterman('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
