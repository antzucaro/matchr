package matchr

import "testing"

var osatests = []struct {
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
	// difference between DL and OSA. This is OSA, so it should be 3.
	{"ca", "abc", 3},
}

// OSA (Optimal String Alignment)
func TestOSA(t *testing.T) {
	for _, tt := range osatests {
		dist := OSA(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("OSA('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
