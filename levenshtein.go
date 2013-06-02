package matchr

func min(a int, b int) (res int) {
    if a < b {
        res = a
    } else {
        res = b
    }

    return
}

func Levenshtein(s1 string, s2 string) (distance int) {
    s1Len := len(s1)
    s2Len := len(s2)

    rows := s1Len + 1
    cols := s2Len + 1

    var d1 int
    var d2 int
    var d3 int
    var i  int
    var j  int
    dist := make([]int, rows * cols)

    for i = 0; i < rows; i++ {
        dist[i * cols] = i
    }

    for j = 0; j < cols; j++ {
        dist[j] = j
    }

    for j = 1; j < cols; j++ {
        for i = 1; i < rows; i++ {
            if (s1[i - 1] == s2[j - 1]) {
                dist[(i * cols) + j] = dist[((i - 1) * cols) + (j - 1)]
            } else {
                d1 = dist[((i - 1) * cols) + j] + 1
                d2 = dist[(i * cols) + (j - 1)] + 1
                d3 = dist[((i - 1) * cols) + (j - 1)] + 1

                dist[(i * cols) + j] = min(d1, min(d2, d3))
            }
        }
    }

    distance = dist[(cols * rows) - 1]

    return
}
