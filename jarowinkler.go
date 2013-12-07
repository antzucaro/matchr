package matchr

func jaroWinklerBase(s1 string, s2 string,
	longTolerance bool, winklerize bool) (distance float64) {

	s1Length := len(s1)
	s2Length := len(s2)

	if s1Length == 0 || s2Length == 0 {
		return
	}

	minLength := 0
	if s1Length > s2Length {
		minLength = s1Length
	} else {
		minLength = s2Length
	}

	searchRange := minLength
	searchRange = (searchRange / 2) - 1
	if searchRange < 0 {
		searchRange = 0
	}
	var lowLim, hiLim, transCount, commonChars int
	var i, j, k int

	s1Flag := make([]bool, s1Length+1)
	s2Flag := make([]bool, s2Length+1)

	// find the common chars within the acceptable range
	commonChars = 0
	for i, _ = range s1 {
		if i >= searchRange {
			lowLim = i - searchRange
		} else {
			lowLim = 0
		}

		if (i + searchRange) <= (s2Length - 1) {
			hiLim = i + searchRange
		} else {
			hiLim = s2Length - 1
		}

		for j := lowLim; j <= hiLim; j++ {
			if !s2Flag[j] && s2[j] == s1[i] {
				s2Flag[j] = true
				s1Flag[i] = true
				commonChars++

				break
			}
		}
	}

	// if we have nothing in common at this point, nothing else can be done
	if commonChars == 0 {
		return
	}

	// otherwise we count the transpositions
	k = 0
	transCount = 0
	for i, _ := range s1 {
		if s1Flag[i] {
			for j = k; j < s2Length; j++ {
				if s2Flag[j] {
					k = j + 1
					break
				}
			}
			if s1[i] != s2[j] {
				transCount++
			}
		}
	}
	transCount /= 2

	// adjust for similarities in nonmatched characters
	distance = float64(commonChars)/float64(s1Length) +
		float64(commonChars)/float64(s2Length) +
		(float64(commonChars-transCount))/float64(commonChars)
	distance /= 3.0

	// give more weight to already-similar strings
	if winklerize && distance > 0.7 {

		// the first 4 characters in common
		if minLength >= 4 {
			j = 4
		} else {
			j = minLength
		}

		for i = 0; i < j && s1[i] == s2[i] && NaN(s1[i]); i++ {
		}

		if i > 0 {
			distance += float64(i) * 0.1 * (1.0 - distance)
		}

		if longTolerance && (minLength > 4) && (commonChars > i+1) &&
			(2*commonChars >= minLength+i) {
			if NaN(s1[0]) {
				distance += (1.0 - distance) * (float64(commonChars-i-1) /
					(float64(s1Length) + float64(s2Length) - float64(i*2) + 2))
			}
		}
	}

	return
}

func Jaro(s1 string, s2 string) (distance float64) {
	return jaroWinklerBase(s1, s2, false, false)
}

func JaroWinkler(s1 string, s2 string, longTolerance bool) (distance float64) {
	return jaroWinklerBase(s1, s2, longTolerance, true)
}
