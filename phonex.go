package matchr

import "bytes"

type phonexResult struct {
    result bytes.Buffer
    length int
    last string
}

func (r *phonexResult) Add(c string) {
    if c != r.last || c == "0" {
        r.result.WriteString(c)
        r.length += len(c)
        r.last = c
    }
}

func (r *phonexResult) Length() int {
    return r.length
}

func (r *phonexResult) String() string {
    return r.result.String()
}

func preProcess(input []rune) ([]rune) {
    output := make([]rune, 0, len(input))

    // 0. Remove all non-ASCII characters
    for  _, v := range(input) {
        if v >= 65 && v <= 90 {
            output = append(output, v)
        }
    }

    // 1. Remove all trailing 'S' characters at the end of the name
    if output[len(output)-1] == 'S' {
        output = append(output[:len(output)-1], output[len(output):]...)
    }

    // 2. Convert leading letter pairs as follows
    //    KN -> N, PH -> F, WR -> R
    if len(output) > 1 {
        switch string(output[0:2]) {
        case "KN":
            output = append(output[:0], output[1:]...)
        case "PH":
            output = append(output[:0], output[1:]...)
            output[0] = 'F'
        case "WR":
            output = append(output[:0], output[1:]...)
        }
    }

    // 3. Convert leading single letters as follows:
    //    H         -> Remove
    //    E,I,O,U,Y -> A
    //    P         -> B
    //    V         -> F
    //    K,Q       -> C
    //    J         -> G
    //    Z         -> S
    switch output[0] {
    case 'H':
        output = append(output[:0], output[1:]...)
    case 'E', 'I', 'O', 'U', 'Y':
        output[0] = 'A'
    case 'P':
        output[0] = 'B'
    case 'V':
        output[0] = 'F'
    case 'K', 'Q':
        output[0] = 'C'
    case 'J':
        output[0] = 'G'
    case 'Z':
        output[0] = 'S'
    }

    return output
}

func Phonex(s1 string) (string){
    if len(s1) == 0 {
        return ""
    }

    // preprocess
    s1 = cleanInput(s1)

    // convert to a slice of character code points for easy indexing
    input := preProcess([]rune(s1))

    result := phonexResult{}

    // keep the first letter
    if len(input) > 0 {
        result.Add(string(input[0]))
    }

    index := 1
    for result.Length() < 4 && index < len(input) {

        switch input[index] {
        case 'A', 'E', 'H', 'I', 'O', 'U', 'W', 'Y':
            index++
        case 'B', 'F', 'P', 'V':
            result.Add("1")
            index++
        case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
            result.Add("2")
            index++
        case 'D', 'T':
            if index == len(input) - 1 || (index + 1 < len(input) &&
               input[index + 1] != 'C') {
                result.Add("3")
            }
            index++
        case 'L':
            if index + 1 < len(input) && !isVowel(input[index + 1]) && 
               index != (len(input) - 1) {
                result.Add("4")
            }
            index++
        case 'M', 'N':
            result.Add("5")
            if index + 1 < len(input) &&
               (input[index + 1] == 'D' || input[index + 1] == 'G') {
                index += 2
            } else {
                index++
            }
        case 'R':
            if index + 1 < len(input) &&
               (!isVowel(input[index + 1]) && index != (len(input) - 1)){
                result.Add("6")
            }
            index++
        default:
            index++
        }
    }

    for result.Length() < 4 {
        result.Add("0")
    }

    return result.String()
}
