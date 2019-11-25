package gol

import "fmt"

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
			(*b)[i][j].X = j
			(*b)[i][j].Y = i
			// in case we start a new game usage
			(*b)[i][j].Alive = false

			//b[i][j].color = ""

		}
	}
	return b
}

//boardWriteTEST
func BoardWriteTest(b *[][]Cell, writeBit bool) {

	for i := range *b {
		for j := range *b {
			(*b)[i][j].X = i
			(*b)[i][j].Y = j
			// WRITE 1
			(*b)[i][j].Alive = writeBit
			(*b)[i][j].PreviousLife = writeBit

		}

	}
	//	return 1
}

//BoardClearAll to be used as a reset function
func BoardClearAll(b *[][]Cell) {

	for i := range *b {
		for j := range *b {
			(*b)[i][j].X = i
			(*b)[i][j].Y = j
			(*b)[i][j].Alive = false
			(*b)[i][j].PreviousLife = false
			(*b)[i][j].ColorHex = ""

			// WRITE 1
			// (*b)[i][j].ThisGen = "#ffffff"
			// (*b)[i][j].NextGen = "#ffffff"
			// (*b)[i][j].Now = false
			// (*b)[i][j].Next = false
		}

	}
}

// TestBoard func
func TestBoard(size int, b *[][]Cell) int {
	sum := 0
	for i := range *b {
		for j := range *b {
			(*b)[i][j].X = i
			(*b)[i][j].Y = j
			// in case we start a new game usage
			fmt.Println("checking ", i, j)
			if (*b)[i][j].Alive == true && (*b)[i][j].PreviousLife == true {
				sum++
			}
		}

	}
	return sum
}
