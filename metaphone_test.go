package matchr

import (
	"bufio"
	"compress/gzip"
	"os"
	"strings"
	"testing"
)

func TestDoubleMetaphone(t *testing.T) {
	// load gzipped corpus
	f, err := os.Open("double_metaphone_corpus.txt.gz")
	if err != nil {
		panic("Error opening file double_metaphone_corpus.txt.gz! Exiting.")
	}
	defer f.Close()

	g, err := gzip.NewReader(f)
	if err != nil {
		panic("Error with supposedly gzipped file double_metaphone_corpus.txt.gz! Exiting.")
	}

	r := bufio.NewReader(g)

	line, err := r.ReadString('\n')
	for err == nil {
		line = strings.TrimRight(line, "\n")
		v := strings.Split(line, "|")

		metaphone, alternate := DoubleMetaphone(v[0])
		if metaphone != v[1] || alternate != v[2] {
			t.Errorf("DoubleMetaphone('%s') = (%v, %v), want (%v, %v)", v[0], metaphone, alternate, v[1], v[2])
			t.FailNow()
		}

		line, err = r.ReadString('\n')
	}
}
