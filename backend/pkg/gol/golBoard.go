package gol

// CreateNewBoard creates a new board with w X h dimension or h X w. clarify row & col later
// row x columns = width x h
func CreateNewBoard(r, c int) *[][]Cell {
	a := make([]Cell, c*r)
	m := make([][]Cell, r)
	lo, hi := 0, c
	for i := range m {
		m[i] = a[lo:hi:hi]
		lo, hi = hi, hi+c
	}
	return &m
}
