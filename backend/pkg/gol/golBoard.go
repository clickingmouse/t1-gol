package gol

// CreateNewBoard creates a new board with r X c dimension
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

//InitBoard initializes the X & Y valeus of board, and each sell is not alive
func InitBoard(b *[][]Cell) *[][]Cell {
	for i := range *b {
		for j := range (*b)[i] {
			(*b)[i][j].X = i
			(*b)[i][j].Y = j
			// in case we start a new game usage
			(*b)[i][j].Alive = false

			//b[i][j].color = ""

		}
	}
	return b
}
