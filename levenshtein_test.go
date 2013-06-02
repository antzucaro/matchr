package matchr

import "testing"

func TestInsertion(t *testing.T) {
    exp := 1
    res := Levenshtein("car", "cars")
    if  res != exp {
        t.Errorf("Levenshtein('car', 'cars') = %v, want %v", res, exp)
    }
}

func TestSubstitution(t *testing.T) {
    exp := 1
    res := Levenshtein("library", "librari")
    if  res != exp {
        t.Errorf("Levenshtein('library', 'librari') = %v, want %v", res, exp)
    }
}

func TestNullLeft(t *testing.T) {
    exp := 7
    res := Levenshtein("", "library")
    if  res != exp {
        t.Errorf("Levenshtein('', 'library') = %v, want %v", res, exp)
    }
}

func TestNullRight(t *testing.T) {
    exp := 7
    res := Levenshtein("library", "")
    if  res != exp {
        t.Errorf("Levenshtein('library', '') = %v, want %v", res, exp)
    }
}

func TestBothNull(t *testing.T) {
    exp := 0
    res := Levenshtein("", "")
    if  res != exp {
        t.Errorf("Levenshtein('', '') = %v, want %v", res, exp)
    }
}
