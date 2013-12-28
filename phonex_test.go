package matchr

import "testing"

// test cases from http://rosettacode.org/wiki/phonex#F.23
var phonextests = []struct {
	s1      string
	phonex string
}{
    {"123 testsss" , "T232"},
    {"24/7 test" , "T230"},
    {"A" , "A000"},
    {"Lee" , "L000"},
    {"Kuhne" , "C500"},
    {"Meyer-Lansky" , "M652"},
    {"Oepping" , "A150"},
    {"Daley" , "D000"},
    {"Dalitz" , "D320"},
    {"Duhlitz" , "D320"},
    {"De Ledes" , "D300"},
    {"Sandemann" , "S500"},
    {"Schüßler" , "S200"},
    {"Schmidt" , "S253"},
    {"Sinatra" , "S530"},
    {"Heinrich" , "E520"},
    {"Hammerschlag" , "A562"},
    {"Williams" , "W450"},
    {"Wilms" , "W450"},
    {"Wilson" , "W425"},
    {"Worms" , "W650"},
    {"Zedlitz" , "S320"},
    {"Zotteldecke" , "S343"},
    {"ZYX test" , "S232"},
    {"Scherman" , "S265"},
    {"Schurman" , "S265"},
    {"Sherman" , "S650"},
    {"Shireman" , "S500"},
    {"Shurman" , "S650"},
    {"Euler" , "A000"},
    {"Ellery" , "A400"},
    {"Hilbert" , "I416"},
    {"Heilbronn" , "E415"},
    {"Gauss" , "G200"},
    {"Ghosh" , "G200"},
    {"Knuth" , "N300"},
    {"Kant" , "C530"},
    {"Lloyd" , "L300"},
    {"Ladd" , "L300"},
    {"Lukasiewicz" , "L200"},
    {"Lissajous" , "L200"},
    {"Ashcraft" , "A213"},
    {"Philip" , "F100"},
    {"Fripp" , "F100"},
    {"Czarkowska" , "C262"},
    {"Hornblower" , "O651"},
    {"Looser" , "L200"},
    {"Wright" , "R230"},
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
