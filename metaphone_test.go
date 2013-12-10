package matchr

import "testing"

var metaphonetests = []struct {
	s1        string
	metaphone string
}{
	{"Ashcraft", "AXKRFT"},
	{"Ashhhcraft", "AXKRFT"},
	{"Ashcroft", "AXKRFT"},
	{"Burroughs", "BRS"},
	{"Burrows", "BRS"},
	{"Ekzampul", "EKSMPL"},
	{"Example", "EKSMPL"},
	{"Ellery", "ELR"},
	{"Euler", "ELR"},
	{"Ghosh", "KHX"},
	{"Gauss", "KS"},
	{"Gutierrez", "KTRS"},
	{"Heilbronn", "HLBRN"},
	{"Hilbert", "HLBRT"},
	{"Jackson", "JKSN"},
	{"Kant", "KNT"},
	{"Knuth", "N0"},
	{"Lee", "L"},
	{"Lukasiewicz", "LKSWKS"},
	{"Lissajous", "LSJS"},
	{"Ladd", "LT"},
	{"Lloyd", "LT"},
	{"Moses", "MSS"},
	{"O'Hara", "OHR"},
	{"Pfister", "PFSTR"},
	{"Rubin", "RBN"},
	{"Robert", "RBRT"},
	{"Rupert", "RPRT"},
	{"Soundex", "SNTKS"},
	{"Sownteks", "SNTKS"},
	{"Tymczak", "TMKSK"},
	{"VanDeusen", "FNTSN"},
	{"Washington", "WXNKTN"},
	{"Wheaton", "WTN"},
}

func TestDoubleMetaphone(t *testing.T) {
	for _, tt := range metaphonetests {
		metaphone := DoubleMetaphone(tt.s1)
		if metaphone != tt.metaphone {
			t.Errorf("DoubleMetaphone('%s') = %v, want %v", tt.s1, metaphone, tt.metaphone)
		}
	}
}
