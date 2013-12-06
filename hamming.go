package matchr

import "errors"

func Hamming(s1 string, s2 string) (distance int, err error) {
    if len(s1) != len(s2) {
        err = errors.New("Hamming distance of different sized strings.")
        return
    }

    for i, v := range(s1) {
        if s2[i] != uint8(v) {
            distance += 1
        }
    }

    return
}
