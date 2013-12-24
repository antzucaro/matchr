package matchr

import (
    "bytes"
    "strings"
)

type metaphoneResult struct {
    // the maximum number of code values to calculate
    maxLength int

    // whether to calculate an alternate
    calcAlternate bool

    // no direct modifications - only through Add()
    primary bytes.Buffer
    alternate bytes.Buffer

    // length of the private buffers
    PrimaryLength int
    AlternateLength int
}

func NewmetaphoneResult(maxLength int, calcAlternate bool) (r *metaphoneResult) {
    r = &metaphoneResult{maxLength: maxLength, calcAlternate: calcAlternate}
    return
}

func (r *metaphoneResult) Add(c1 string, c2 string) {
    if c1 != "" {
        r.primary.WriteString(c1)
        r.PrimaryLength += len(c1)
    }

    if c2 != "" && r.calcAlternate {
        r.alternate.WriteString(c2)
        r.AlternateLength += len(c2)
    }
}

func (r *metaphoneResult) IsComplete() bool {
    return r.PrimaryLength >= r.maxLength && r.AlternateLength >= r.maxLength
}

func (r *metaphoneResult) Result() (primary string, alternate string) {
    primary = r.primary.String()
    if len(primary) > r.maxLength {
        primary = primary[0:r.maxLength]
    }
    alternate = r.alternate.String()
    if len(alternate) > r.maxLength {
        alternate = alternate[0:r.maxLength]
    }
    return
}

// utility functions for checking things within a string
func isSlavoGermanic(value string) bool {
    return strings.Contains(value, "W") || strings.Contains(value, "K") ||
           strings.Contains(value, "CZ") || strings.Contains(value, "WITZ")
}

func isVowel(c rune) bool {
    switch c {
    case 'A', 'E', 'I', 'O', 'U', 'Y':
        return true
    default:
        return false
    }
}

func isSilentStart(input *String) bool {
    SILENT_START := [...]string{"GN", "KN", "PN", "WR", "PS"}

    prefix := substring(input, 0, 2)

    for _, criteria := range(SILENT_START) {
        if prefix == criteria {
            return true
        }
    }

    return false
}

func cleanInput(input string) string {
    return strings.ToUpper(strings.TrimSpace(input))
}

func charAt(value *String, index int) rune {
    if index < 0 || index >= value.RuneCount() {
        return 0
    } else {
        return rune(value.At(index))
    }
}

func contains(value *String, start int, length int, criteria string) bool {
    return substring(value, start, length) == criteria
}

// TODO: use varargs here, it makes the rest of the code cleaner
func containsAny(value *String, start int, length int, criteria []string) bool {
    substring := substring(value, start, length)
    for _, c := range(criteria) {
        if substring == c {
            return true
        }
    }
    return false
}

func substring(value *String, start int, length int) string {
    if start >= 0 && start + length <= value.RuneCount() {
        return value.Slice(start, start + length)
    } else {
        return ""
    }
}

func handleVowel(result *metaphoneResult, index int) int {
    if index == 0 {
        result.Add("A", "A")
    }

    return index + 1
}

/******************************************************************************
 * Entry handlers for letters.
 *****************************************************************************/
func handleC(input *String, result *metaphoneResult, index int) int {
    if conditionC0(input, index) {
        result.Add("K", "K")
        index += 2
    } else if index == 0 && contains(input, index, 6, "CAESAR") {
        result.Add("S", "S")
        index += 2
    } else if contains(input, index, 2, "CH") {
        index = handleCH(input, result, index)
    } else if contains(input, index, 2, "CZ") &&
             !contains(input, index - 2, 4, "WICZ") {
        result.Add("S", "X")
        index += 2
    } else if contains(input, index + 1, 3, "CIA") {
        result.Add("X", "X")
        index += 3
    } else if contains(input, index, 2, "CC") &&
        !(index == 1 && charAt(input, 0) == 'M') {
        return handleCC(input, result, index)
    } else if contains(input, index, 2, "CK") ||
              contains(input, index, 2, "CG") ||
              contains(input, index, 2, "CQ") {
        result.Add("K", "K")
        index += 2
    } else if contains(input, index, 2, "CI") ||
              contains(input, index, 2, "CE") ||
              contains(input, index, 2, "CY") {
        if contains(input, index, 3, "CIO") ||
           contains(input, index, 3, "CIE") ||
           contains(input, index, 3, "CIA") {
            result.Add("S", "X")
        } else {
            result.Add("S", "S")
        }
        index += 2
    } else {
        result.Add("K", "K")
        if contains(input, index + 1, 2, " C") ||
           contains(input, index + 1, 2, " Q") ||
           contains(input, index + 1, 2, " G") {
               index += 3
        } else if (contains(input, index + 1, 1, "C") ||
                   contains(input, index + 1, 1, "K") ||
                   contains(input, index + 1, 1, "Q")) &&
                 !(contains(input, index + 1, 2, "CE") ||
                   contains(input, index + 1, 2, "CI")){
                index += 2
        } else {
            index++
        }
    }

    return index
}

func handleCC(input *String, result *metaphoneResult, index int) int {
    if containsAny(input, index + 2, 1, []string{"I", "E", "H"}) &&
       !contains(input, index + 2, 2, "HU") {
        if (index == 1 && charAt(input, index - 1) == 'A') ||
           (containsAny(input, index - 1, 5, []string{"UCCEE", "UCCES"})) {
            result.Add("KS", "KS")
        } else {
            result.Add("X", "X")
        }
        index += 3
    } else {
        result.Add("K", "K")
        index += 2
    }
    return index
}

func handleCH(input *String, result *metaphoneResult, index int) int {
    if index > 0 && contains(input, index, 4, "CHAE") {
        result.Add("K", "X")
        return index + 2
    } else if conditionCH0(input, index) {
        result.Add("K", "K")
        return index + 2
    // TODO: combine this condition with the one above?
    } else if conditionCH1(input, index) {
        result.Add("K", "K")
        return index + 2
    } else {
        if index > 0 {
            if contains(input, 0, 2, "MC") {
                result.Add("K", "K")
            } else {
                result.Add("X", "K")
            }
        } else {
            result.Add("X", "X")
        }
        return index + 2
    }
}

func handleD(input *String, result *metaphoneResult, index int) int {
    if contains(input, index, 2, "DG") {
        if containsAny(input, index + 2, 1, []string{"I", "E", "Y"}) {
            result.Add("J", "J")
            index += 3
        } else {
            result.Add("TK", "TK")
            index += 2
        }
    } else if containsAny(input, index, 2, []string{"DT", "DD"}) {
        result.Add("T", "T")
        index += 2
    } else {
        result.Add("T", "T")
        index++
    }
    return index
}

func handleG(input *String, result *metaphoneResult, index int, slavoGermanic bool) int {
    if charAt(input, index + 1) == 'H' {
        index = handleGH(input, result, index)
    } else if charAt(input, index + 1) == 'N' {
        if index == 1 && isVowel(charAt(input,0)) && !slavoGermanic {
            result.Add("KN", "N")
        } else if !contains(input, index + 2, 2, "EY") && charAt(input, index + 1) != 'Y' && !slavoGermanic {
            result.Add("N", "KN")
        } else {
            result.Add("KN", "KN")
        }
        index += 2
    } else if contains(input, index + 1, 2, "LI") && !slavoGermanic {
        result.Add("KL", "L")
        index += 2
    } else if index == 0 && (charAt(input, index + 1) == 'Y' ||
        containsAny(input, index + 1, 2, []string{"ES", "EP", "EB", "EL", "EY", "IB", "IL", "IN", "IE", "EI", "ER"})) {
        result.Add("K", "J")
        index += 2
    } else if (contains(input, index + 1, 2, "ER") ||
      charAt(input, index + 1) == 'Y') &&
      !containsAny(input, 0, 6, []string{"DANGER", "RANGER", "MANGER"}) &&
      !containsAny(input, index - 1, 1, []string{"E", "I"}) &&
      !containsAny(input, index - 1, 3, []string{"RGY", "OGY"}) {
      result.Add("K", "J")
      index += 2
    } else if containsAny(input, index + 1, 1, []string{"E", "I", "Y"}) ||
        containsAny(input, index - 1, 4, []string{"AGGI", "OGGI"}){
          if containsAny(input, 0, 4, []string{"VAN ", "VON "}) ||
             contains(input, 0, 3, "SCH") ||
             contains(input, index + 1, 2, "ET") {
                result.Add("K", "K")
          } else if contains(input, index + 1, 3, "IER") {
              result.Add("J", "J")
          } else {
              result.Add("J", "K")
          }
          index += 2
    } else if charAt(input, index + 1) == 'G' {
        result.Add("K", "K")
        index += 2
    } else {
        result.Add("K", "K")
        index++
    }
    return index
}

func handleGH(input *String, result *metaphoneResult, index int) int {
    if index > 0 && !isVowel(charAt(input, index - 1)) {
        result.Add("K", "K")
        index += 2
    } else if index == 0 {
        if charAt(input, index + 2) == 'I' {
            result.Add("J", "J")
        } else {
            result.Add("K", "K")
        }
        index += 2
    } else if (index > 1 && containsAny(input, index - 2, 1, []string{"B", "H", "D"})) || 
              (index > 2 && containsAny(input, index - 3, 1, []string{"B", "H", "D"})) ||
              (index > 3 && containsAny(input, index - 4, 1, []string{"B", "H"})) {
        index += 2
    } else {
        if index > 2 && charAt(input, index - 1) == 'U' &&
           containsAny(input, index - 3, 1, []string{"C", "G", "L", "R", "T"}) {
            result.Add("F", "F")
        } else if index > 0 && charAt(input, index - 1) != 'I' {
            result.Add("K", "K")
        }
        index += 2
    }
    return index
}

func handleH(input *String, result *metaphoneResult, index int) int {
    if (index == 0 || isVowel(charAt(input, index - 1))) &&
       isVowel(charAt(input, index + 1)) {
        result.Add("H", "H")
        index += 2
    } else {
        index++
    }
    return index
}

func handleJ(input *String, result *metaphoneResult, index int, slavoGermanic bool) int {
    if contains(input, index, 4, "JOSE") || contains(input, 0, 4, "SAN ") {
        if (index == 0 && (charAt(input, index + 4) == ' ') ||
         input.RuneCount() == 4) || contains(input, 0, 4, "SAN ") {
             result.Add("H", "H")
        } else {
            result.Add("J", "H")
        }
        index++
    } else {
        if index == 0 && !contains(input, index, 4, "JOSE") {
            result.Add("J", "A")
        } else if isVowel(charAt(input, index - 1)) && !slavoGermanic &&
            (charAt(input, index + 1) == 'A' || charAt(input, index + 1) == 'O') {
                result.Add("J", "H")
        } else if index == (input.RuneCount() - 1) {
            result.Add("J", " ")
        } else if !containsAny(input, index + 1, 1,
          []string{"L", "T", "K", "S", "N", "M", "B", "Z"}) &&
          !containsAny(input, index - 1, 1, []string{"S", "K", "L"}){
            result.Add("J", "J")
        }

        if charAt(input, index + 1) == 'J' {
            index += 2
        } else {
            index++
        }
    }
    return index
}

func handleL(input *String, result *metaphoneResult, index int) int {
    if charAt(input, index + 1) == 'L' {
        if conditionL0(input, index) {
            result.Add("L", "L")
        } else {
            result.Add("L", "L")
        }
        index += 2
    } else {
        result.Add("L", "L")
        index++
    }
    return index
}

func handleP(input *String, result *metaphoneResult, index int) int {
    if charAt(input, index + 1) == 'H' {
        result.Add("F", "F")
        index += 2
    } else {
        result.Add("P", "P")
        if containsAny(input, index + 1, 1, []string{"P", "B"}) {
            index += 2
        } else {
            index++
        }
    }
    return index
}

func handleR(input *String, result *metaphoneResult, index int, slavoGermanic bool) int {
    if index == (input.RuneCount() - 1) && !slavoGermanic &&
        contains(input, index - 2, 2, "IE") &&
        !containsAny(input, index - 4, 2, []string{"ME", "MA"}) {
            result.Add("", "R")
    } else {
        result.Add("R", "R")
    }

    if charAt(input, index + 1) == 'R' {
        index += 2
    } else {
        index++
    }
    return index
}

func handleS(input *String, result *metaphoneResult, index int, slavoGermanic bool) int {
    if containsAny(input, index - 1, 3, []string{"ISL", "YSL"}) {
        index++
    } else if index == 0 && contains(input, index, 5, "SUGAR") {
        result.Add("X", "S")
        index++
    } else if contains(input, index, 2, "SH") {
        if containsAny(input, index + 1, 4, []string{"HEIM", "HOEK", "HOLM", "HOLZ"}) {
            result.Add("S", "S")
        } else {
            result.Add("X", "X")
        }
        index += 2
    } else if containsAny(input, index, 3, []string{"SIO", "SIA"}) ||
        contains(input, index, 4, "SIAN") {
        if slavoGermanic {
            result.Add("S", "S")
        } else {
            result.Add("S", "X")
        }
        index += 3
    } else if (index == 0 && containsAny(input, index + 1, 1, []string{"M", "N", "L", "W"})) || 
        contains(input, index + 1, 1, "Z") {
        result.Add("S", "X")
        if contains(input, index + 1, 1, "Z") {
            index += 2
        } else {
            index++
        }
    } else if contains(input, index, 2, "SC") {
        index = handleSC(input, result, index)
    } else {
        if index == input.RuneCount() - 1 &&
            containsAny(input, index - 2, 2, []string{"AI", "OI"}) {
            result.Add("", "S")
        } else {
            result.Add("S", "S")
        }

        if containsAny(input, index + 1, 1, []string{"S", "Z"}) {
            index += 2
        } else {
            index++
        }
    }
    return index
}

func handleSC(input *String, result *metaphoneResult, index int) int {
    if charAt(input, index + 2) == 'H' {
        if containsAny(input, index + 3, 2, []string{"OO", "ER", "EN", "UY", "ED", "EM"}) {
            if containsAny(input, index + 3, 2, []string{"ER", "EN"}) {
                result.Add("X", "SK")
            } else {
                result.Add("SK", "SK")
            }
        } else {
            if index == 0 && !isVowel(charAt(input, 3)) && charAt(input, 3) != 'W' {
                result.Add("X", "S")
            } else {
                result.Add("X", "X")
            }
        }
    } else if containsAny(input, index + 2, 1, []string{"I", "E", "Y"}){
        result.Add("S", "S")
    } else {
        result.Add("SK", "SK")
    }
    index += 3

    return index
}

func handleT(input *String, result *metaphoneResult, index int) int {
    if contains(input, index, 4, "TION") {
        result.Add("X", "X")
        index += 3
    } else if containsAny(input, index, 3, []string{"TIA", "TCH"}) {
        result.Add("X", "X")
        index += 3
    } else if contains(input, index, 2, "TH") || contains(input, index, 3, "TTH") {
        if containsAny(input, index + 2, 2, []string{"OM", "AM"}) ||
           containsAny(input, 0, 4, []string{"VAN ", "VON "}) ||
           contains(input, 0, 3, "SCH") {
            result.Add("T" ,"T")
        } else {
            result.Add("0", "T")
        }
        index += 2
    } else {
        result.Add("T" ,"T")
        if containsAny(input, index + 1, 1, []string{"T", "D"}) {
            index += 2
        } else {
            index++
        }
    }
    return index
}

func handleW(input *String, result *metaphoneResult, index int) int {
    if contains(input, index, 2, "WR") {
        result.Add("R", "R")
        index += 2
    } else {
        if index == 0 && (isVowel(charAt(input, index + 1)) ||
          contains(input, index, 2, "WH")) {
            if isVowel(charAt(input, index + 1)) {
                result.Add("A", "F")
            } else {
                result.Add("A", "A")
            }
            index++
        } else if (index == input.RuneCount() - 1 && isVowel(charAt(input, index - 1))) || 
        containsAny(input, index - 1, 5, []string{"EWSKI", "EWSKY", "OWSKI", "OWSKY"}) ||
        contains(input, 0, 3, "SCH") {
            result.Add("", "F")
            index++
        } else if containsAny(input, index, 4, []string{"WICZ", "WITZ"}) {
            result.Add("TS", "FX")
            index += 4
        } else {
            index++
        }
    }
    return index
}

func handleX(input *String, result *metaphoneResult, index int) int {
    if index == 0 {
        result.Add("S", "S")
        index++
    } else {
        if !((index == input.RuneCount() - 1) && 
        (containsAny(input, index -3, 3, []string{"IAU", "EAU"}) ||
        containsAny(input, index - 2, 2, []string{"AU", "OU"}))) {
            result.Add("KS", "KS")
        }

        if containsAny(input, index + 1, 1, []string{"C", "X"}) {
            index += 2
        } else {
            index++
        }
    }
    return index
}

func handleZ(input *String, result *metaphoneResult, index int, slavoGermanic bool) int {
    if charAt(input, index + 1) == 'H' {
        result.Add("J", "J")
    } else {
        if containsAny(input, index + 1, 2, []string{"ZO", "ZI", "ZA"}) ||
        (slavoGermanic && (index > 0 && charAt(input, index - 1) != 'T')) {
            result.Add("S", "TS")
        } else {
            result.Add("S", "S")
        }
    }

    if charAt(input, index + 1) == 'Z' {
        index += 2
    } else {
        index++
    }
    return index
}

/******************************************************************************
 * Complex conditional handlers for letters
 *****************************************************************************/
func conditionC0(input *String, index int) bool {
    if contains(input, index, 4, "CHIA") {
        return true
    } else if index <= 1 {
        return false
    } else if isVowel(charAt(input, index - 2)) {
        return false
    } else if !contains(input, index - 1, 3, "ACH") {
        return false
    } else {
        c := charAt(input, index + 2)
        return (c != 'I' && c != 'E') ||
              (contains(input, index - 2, 6, "BACHER") ||
               contains(input, index - 2, 6, "MACHER"))
    }
}

func conditionCH0(input *String, index int) bool {
    if index != 0 {
        return false
    } else if !containsAny(input, index + 1, 5, []string{"HARAC", "HARIS"}) &&
        !containsAny(input, index + 1, 3, []string{"HOR", "HYM", "HIA", "HEM"}) {
            return false
    } else if contains(input, 0, 5, "CHORE") {
        return false
    } else {
        return true
    }
}

func conditionCH1(input *String, index int) bool {
    // good god this is ugly
    return (containsAny(input, 0, 4, []string{"VAN ", "VON "}) || contains(input, 0, 3, "SCH")) ||
    containsAny(input, index - 2, 6, []string{"ORCHES", "ARCHIT", "ORCHID"}) ||
    containsAny(input, index + 2, 1, []string{"T", "S"}) ||
    ((containsAny(input, index - 1, 1, []string{"A", "O", "U", "E"}) || index == 0) &&
     (containsAny(input, index + 2, 1, []string{"L", "R", "N", "M", "B", "H", "F", "V", "W", " "}) || index + 1 == input.RuneCount() - 1))
}

func conditionL0(input *String, index int) bool {
    if index == (input.RuneCount() - 3) &&
       containsAny(input, index - 1, 4, []string{"ILLO", "ILLA", "ALLE"}) {
        return true
    } else if (containsAny(input, input.RuneCount() - 2, 2, []string{"AS", "OS"}) ||
      containsAny(input, input.RuneCount() - 1, 1, []string{"A", "O"})) &&
      (contains(input, index - 1, 4, "ALLE")) {
        return true
    } else {
        return false
    }
}

func conditionM0(input *String, index int) bool {
    if charAt(input, index + 1) == 'M' {
        return true
    }

    return contains(input, index - 1, 3, "UMB") &&
       ((index + 1) == (input.RuneCount()-1) ||
       contains(input, index + 2, 2, "ER"))
}

func DoubleMetaphone(s1 string) (string, string){
    // trim, upper space
    s1 = cleanInput(s1)
    // structure to traverse the string by code point, not byte
    input := NewString(s1)

    slavoGermanic := isSlavoGermanic(s1)

    // where we are in the string
    index := 0

    if isSilentStart(input) {
        index += 1
    }

    result := NewmetaphoneResult(4, true)

    for !result.IsComplete() && index <= (input.RuneCount() - 1) {
        c := rune(input.At(index))
        switch c {
        case 'A', 'E', 'I', 'O', 'U', 'Y':
            index = handleVowel(result, index)
        case 'B':
            result.Add("P", "P")
            if charAt(input, index + 1) == 'B' {
                index += 2
            } else {
                index++
            }
        case 'Ç':
            result.Add("S", "S")
            index++
        case 'C':
            index = handleC(input, result, index)
        case 'D':
            index = handleD(input, result, index)
        case 'F':
            result.Add("F", "F")
            if charAt(input, index + 1) == 'F' {
                index += 2
            } else {
                index++
            }
        case 'G':
            index = handleG(input, result, index, slavoGermanic)
        case 'H':
            index = handleH(input, result, index)
        case 'J':
            index = handleJ(input, result, index, slavoGermanic)
        case 'K':
            result.Add("K", "K")
            if charAt(input, index + 1) == 'K' {
                index += 2
            } else {
                index++
            }
        case 'L':
            index = handleL(input, result, index)
        case 'M':
            result.Add("M", "M")
            if conditionM0(input, index) {
                index += 2
            } else {
                index++
            }
        case 'N':
            result.Add("N", "N")
            if charAt(input, index + 1) == 'N' {
                index += 2
            } else {
                index++
            }
        case 'Ñ':
            result.Add("N", "N")
            index++
        case 'P':
            index = handleP(input, result, index)
        case 'Q':
            result.Add("K", "Q")
            if charAt(input, index + 1) == 'Q' {
                index += 2
            } else {
                index++
            }
        case 'R':
            index = handleR(input, result, index, slavoGermanic)
        case 'S':
            index = handleS(input, result, index, slavoGermanic)
        case 'T':
            index = handleT(input, result, index)
        case 'V':
            result.Add("F", "F")
            if charAt(input, index + 1) == 'V' {
                index += 2
            } else {
                index++
            }
        case 'W':
            index = handleW(input, result, index)
        case 'X':
            index = handleX(input, result, index)
        case 'Z':
            index = handleZ(input, result, index, slavoGermanic)
        default:
            index++
        }

    }

    return result.Result()
}
