package matchr

import "testing"

// test cases from http://rosettacode.org/wiki/phonex#F.23
var phonextests = []struct {
	s1     string
	phonex string
}{
	{"123 testsss", "T230"},
	{"24/7 test", "T230"},
	{"A", "A000"},
	{"Lee", "L000"},
	{"Kuhne", "C500"},
	{"Meyer-Lansky", "M452"},
	{"Oepping", "A150"},
	{"Daley", "D400"},
	{"Dalitz", "D432"},
	{"Duhlitz", "D432"},
	{"Dull", "D400"},
	{"De Ledes", "D430"},
	{"Sandemann", "S500"},
	{"Schüßler", "S460"},
	{"Schmidt", "S530"},
	{"Sinatra", "S536"},
	{"Heinrich", "A562"},
	{"Hammerschlag", "A524"},
	{"Williams", "W450"},
	{"Wilms", "W500"},
	{"Wilson", "W250"},
	{"Worms", "W500"},
	{"Zedlitz", "S343"},
	{"Zotteldecke", "S320"},
	{"ZYX test", "S232"},
	{"Scherman", "S500"},
	{"Schurman", "S500"},
	{"Sherman", "S500"},
	{"Shermansss", "S500"},
	{"Shireman", "S650"},
	{"Shurman", "S500"},
	{"Euler", "A460"},
	{"Ellery", "A460"},
	{"Hilbert", "A130"},
	{"Heilbronn", "A165"},
	{"Gauss", "G000"},
	{"Ghosh", "G200"},
	{"Knuth", "N300"},
	{"Kant", "C530"},
	{"Lloyd", "L430"},
	{"Ladd", "L300"},
	{"Lukasiewicz", "L200"},
	{"Lissajous", "L200"},
	{"Ashcraft", "A261"},
	{"Philip", "F410"},
	{"Fripp", "F610"},
	{"Czarkowska", "C200"},
	{"Hornblower", "A514"},
	{"Looser", "L260"},
	{"Wright", "R230"},
	{"Phonic", "F520"},
	{"Quickening", "C250"},
	{"Kuickening", "C250"},
	{"Joben", "G150"},
	{"Zelda", "S300"},
	{"S", "0000"},
	{"H", "0000"},
	{"", "0000"},
}

// phonex
func TestPhonex(t *testing.T) {
	for _, tt := range phonextests {
		phonex := Phonex(tt.s1)
		if phonex != tt.phonex {
			t.Errorf("Phonex('%s') = %v, want %v", tt.s1, phonex, tt.phonex)
		}
	}
}
