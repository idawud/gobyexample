package funcs

import "testing"  

func TestPlus(t *testing.T) {
	cases := []struct {
		in []int; expected int
	}{
		{ []int{2}, 2},
		{ []int{-2, 2}, 0},
		{ []int{6, -8}, -2},
		{ []int{1, 2, 3}, 6},
	} 
	for index, c := range cases { 
		got := Plus(c.in)
		if got != c.expected {
			t.Errorf( "TEST %d:: got %d, expected % d", index + 1, got ,  c.expected)
		}
	}
}

func TestVals(t *testing.T) { 
	got_1, got_2 := Vals()
	want := [2]int{3,7}
    if got_1 != want[0] || got_2 != want[1] {
        t.Errorf("got [%d, %d] wanted %q", got_1, got_2, want)
    }
}