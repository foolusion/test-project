package numbers

import(
	"testing"
)

type numberTest struct {
	in, out int
}

var doubleTests = []numberTest{
	numberTest{1, 2},
	numberTest{2,4},
	numberTest{-5, -10},
}

var squareTests = []numberTest{
	numberTest{1, 1},
	numberTest{2,4},
	numberTest{-5, 25},
}

func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}

func TestSquare(t *testing.T) {
	for _, dt := range squareTests {
		v := Square(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}
