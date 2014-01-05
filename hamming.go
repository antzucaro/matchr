package matchr

import "errors"

func Hamming(s1 string, s2 string) (distance int, err error) {
    // index by code point, not byte
    r1 := []rune(s1)
    r2 := []rune(s2)

    if len(r1) != len(r2) {
        err = errors.New("Hamming distance of different sized strings.")
        return
    }

    for i, v := range(r1) {
        if r2[i] != v {
            distance += 1
        }
    }
    return
}
