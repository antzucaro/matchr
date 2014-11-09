package matchr

import "testing"

var hamtests = []struct {
	s1   string
	s2   string
	dist int
	err  bool
}{
	{"", "", 0, false},
	{"cat", "cat", 0, false},
	{"car", "cat", 1, false},
	{"tar", "car", 1, false},
	{"xyz", "zyx", 2, false},
	{"wxyz", "zyx", 0, true},
	{"Schüßler", "Schübler", 1, false},
	{"Schüßler", "Schußler", 1, false},
}

// Hamming Distance
func TestHamming(t *testing.T) {
	for _, tt := range hamtests {
		dist, err := Hamming(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("Hamming('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}

		if tt.err && err == nil {
			t.Errorf("Hamming('%s', '%s') should throw an error", tt.s1, tt.s2)
		}
	}
}
