package matchr

import "testing"

var nysiistests = []struct {
	s1     string
	nysiis string
}{
	{"knight", "NAGT"},
	{"mitchell", "MATCAL"},
	{"o'daniel", "ODANAL"},
	{"brown sr", "BRANSR"},
	{"browne III", "BRAN"},
	{"browne IV", "BRANAV"},
	{"O'Banion", "OBANAN"},
	{"Mclaughlin", "MCLAGL"},
	{"McCormack", "MCARNA"},
	{"Chapman", "CAPNAN"},
	{"Silva", "SALV"},
	{"McDonald", "MCDANA"},
	{"Lawson", "LASAN"},
	{"Jacobs", "JACAB"},
	{"Greene", "GRAN"},
	{"O'Brien", "OBRAN"},
	{"Morrison", "MARASA"},
	{"Larson", "LARSAN"},
	{"Willis", "WAL"},
	{"Mackenzie", "MCANSY"},
	{"Carr", "CAR"},
	{"Lawrence", "LARANC"},
	{"Matthews", "MAT"},
	{"Richards", "RACARD"},
	{"Bishop", "BASAP"},
	{"Franklin", "FRANCL"},
	{"McDaniel", "MCDANA"},
	{"Harper", "HARPAR"},
	{"Lynch", "LYNC"},
	{"Watkins", "WATCAN"},
	{"Carlson", "CARLSA"},
	{"Wheeler", "WALAR"},
	{"Louis XVI", "LASXV"},
}

// NYSIIS
func TestNYIIS(t *testing.T) {
	for _, tt := range nysiistests {
		nysiis := NYSIIS(tt.s1)
		if nysiis != tt.nysiis {
			t.Errorf("NYSIIS('%s') = %v, want %v", tt.s1, nysiis, tt.nysiis)
		}
	}
}
