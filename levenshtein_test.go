package matchr

import "testing"

func TestLevInsertion(t *testing.T) {
    exp := 1
    res := Levenshtein("car", "cars")
    if  res != exp {
        t.Errorf("Levenshtein('car', 'cars') = %v, want %v", res, exp)
    }
}

func TestLevSubstitution(t *testing.T) {
    exp := 1
    res := Levenshtein("library", "librari")
    if  res != exp {
        t.Errorf("Levenshtein('library', 'librari') = %v, want %v", res, exp)
    }
}

func TestLevDeletion(t *testing.T) {
    exp := 1
    res := Levenshtein("library", "librar")
    if  res != exp {
        t.Errorf("Levenshtein('library', 'librar') = %v, want %v", res, exp)
    }
}

func TestLevNullLeft(t *testing.T) {
    exp := 7
    res := Levenshtein("", "library")
    if  res != exp {
        t.Errorf("Levenshtein('', 'library') = %v, want %v", res, exp)
    }
}

func TestLevNullRight(t *testing.T) {
    exp := 7
    res := Levenshtein("library", "")
    if  res != exp {
        t.Errorf("Levenshtein('library', '') = %v, want %v", res, exp)
    }
}

func TestLevBothNull(t *testing.T) {
    exp := 0
    res := Levenshtein("", "")
    if  res != exp {
        t.Errorf("Levenshtein('', '') = %v, want %v", res, exp)
    }
}
