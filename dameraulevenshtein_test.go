package matchr

import "testing"

// Damerau Levenshtein
func TestDamLevInsertion(t *testing.T) {
    exp := 1
    res := DamerauLevenshtein("car", "cars")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('car', 'cars') = %v, want %v", res, exp)
    }
}

func TestDamLevSubstitution(t *testing.T) {
    exp := 1
    res := DamerauLevenshtein("library", "librari")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('library', 'librari') = %v, want %v", res, exp)
    }
}

func TestDamLevDeletion(t *testing.T) {
    exp := 1
    res := DamerauLevenshtein("library", "librar")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('library', 'librar') = %v, want %v", res, exp)
    }
}

func TestDamLevTransposition(t *testing.T) {
    exp := 1
    res := DamerauLevenshtein("library", "librayr")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('library', 'librayr') = %v, want %v", res, exp)
    }
}

func TestDamLevNullLeft(t *testing.T) {
    exp := 7
    res := DamerauLevenshtein("", "library")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('', 'library') = %v, want %v", res, exp)
    }
}

func TestDamLevNullRight(t *testing.T) {
    exp := 7
    res := DamerauLevenshtein("library", "")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('library', '') = %v, want %v", res, exp)
    }
}

func TestDamLevBothNull(t *testing.T) {
    exp := 0
    res := DamerauLevenshtein("", "")
    if  res != exp {
        t.Errorf("DamerauLevenshtein('', '') = %v, want %v", res, exp)
    }
}
