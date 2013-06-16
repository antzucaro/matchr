package matchr

func DamerauLevenshtein(s1 string, s2 string) (distance int) {
    s1Len := len(s1)
    s2Len := len(s2)
    rows := s1Len + 1
    cols := s2Len + 1

    var i, j, d1, d2, d3, d_now, cost int

    dist := make([]int, rows * cols)

    for i = 0; i < rows; i++ {
        dist[i * cols] = i
    }

    for j = 0; j < cols; j++ {
        dist[j] = j
    }

    for i = 1; i < rows; i++ {
        for j = 1; j < cols; j++ {
            if (s1[i - 1] == s2[j - 1]) {
                cost = 0
            } else {
                cost = 1
            }

            d1 = dist[((i - 1) * cols) + j] + 1
            d2 = dist[(i * cols) + (j - 1)] + 1
            d3 = dist[((i - 1) * cols) + (j - 1)] + cost

            d_now = min(d1, min(d2, d3))

            if i > 2 && j > 2 && s1[i - 1] == s2[j - 2] &&
                s1[i - 2] == s2[j - 1] {
                d1 = dist[((i - 2) * cols) + (j - 2)] + cost
                d_now = min(d_now, d1)
            }

            dist[(i * cols) + j] = d_now
        }
    }

    distance = dist[(cols * rows) - 1]

    return
}
