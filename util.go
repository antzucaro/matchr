package matchr

import "math"

// min of two integers
func min(a int, b int) (res int) {
    if a < b {
        res = a
    } else {
        res = b
    }

    return
}

// is this string index outside of the ASCII numeric code points?
func NaN(c uint8) bool {
	return ((c > 57) || (c < 48))
}

// Round a float64 to the given precision
//
// http://play.golang.org/p/S654PxAe_N
//
// (via Rory McGuire at 
// https://groups.google.com/forum/#!topic/golang-nuts/ITZV08gAugI)
func Round(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}
