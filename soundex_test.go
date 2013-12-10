package matchr

import "testing"

// test cases from http://rosettacode.org/wiki/Soundex#F.23
var soundextests = []struct {
	s1      string
	soundex string
}{
	{"Ashcraft", "A261"},
	{"Ashhhcraft", "A261"},
	{"Ashcroft", "A261"},
	{"Burroughs", "B620"},
	{"Burrows", "B620"},
	{"Ekzampul", "E251"},
	{"Example", "E251"},
	{"Ellery", "E460"},
	{"Euler", "E460"},
	{"Ghosh", "G200"},
	{"Gauss", "G200"},
	{"Gutierrez", "G362"},
	{"Heilbronn", "H416"},
	{"Hilbert", "H416"},
	{"Jackson", "J250"},
	{"Kant", "K530"},
	{"Knuth", "K530"},
	{"Lee", "L000"},
	{"Lukasiewicz", "L222"},
	{"Lissajous", "L222"},
	{"Ladd", "L300"},
	{"Lloyd", "L300"},
	{"Moses", "M220"},
	{"O'Hara", "O600"},
	{"Pfister", "P236"},
	{"Rubin", "R150"},
	{"Robert", "R163"},
	{"Rupert", "R163"},
	{"Soundex", "S532"},
	{"Sownteks", "S532"},
	{"Tymczak", "T522"},
	{"VanDeusen", "V532"},
	{"Washington", "W252"},
	{"Wheaton", "W350"},
}

// Soundex
func TestSoundex(t *testing.T) {
	for _, tt := range soundextests {
		soundex := Soundex(tt.s1)
		if soundex != tt.soundex {
			t.Errorf("Soundex('%s') = %v, want %v", tt.s1, soundex, tt.soundex)
		}
	}
}
