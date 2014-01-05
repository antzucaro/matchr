package matchr

import "testing"

var jarotests = []struct {
	s1   string
	s2   string
	dist float64
}{
	{"", "cars", 0.0},
	{"cars", "", 0.0},
	{"car", "cars", 0.9166666666666666},
	{"dixon", "dicksonx", 0.7666666666666666},
	{"martha", "marhta", 0.9444444444444445},
	{"dwayne", "duane", 0.8222222222222223},
	{"martüa", "marüta", 0.9444444444444445},
}

// Regular Jaro distance
func TestJaro(t *testing.T) {
	for _, tt := range jarotests {
		dist := Jaro(tt.s1, tt.s2)
		if dist != tt.dist {
			t.Errorf("Jaro('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}

var jarowtests = []struct {
	s1   string
	s2   string
	dist float64
}{
	{"", "cars", 0.0},
	{"cars", "", 0.0},
	{"dixon", "dicksonx", 0.8133333333333332},
	{"martha", "marhta", 0.9611111111111111},
	{"dwayne", "duane", 0.8400000000000001},
}

// Jaro-Winkler distance
func TestJaroWinkler(t *testing.T) {
	for _, tt := range jarowtests {
		dist := JaroWinkler(tt.s1, tt.s2, false)
		if dist != tt.dist {
			t.Errorf("JaroWinkler('%s', '%s') = %v, want %v", tt.s1, tt.s2, dist, tt.dist)
		}
	}
}
