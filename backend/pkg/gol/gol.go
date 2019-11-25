package gol

import (
	"fmt"
)

//Evolve ~ main function
func Propagate(board *[][]Cell) *[][]Cell {
	// columns := len(*board)
	// rows := len((*board)[0])
	rows := len(*board)
	columns := len((*board)[0])

	//fmt.Printf("rows: %d, columns: %d ; backingup\n", rows, columns)
	for i := range *board {
		for j := range (*board)[i] {
			(*board)[i][j].SavePreviousLife()
			(*board)[i][j].RetireCell()
		}
	}

	// Loop through every spot in our 2D array and check spots neighbors
	for x := range *board {
		for y := range (*board)[x] {
			// Add up all the states in a 3x3 surrounding grid
			neighbors := 0
			var parentsColor []string
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					r := (x + i + rows) % rows
					c := (y + j + columns) % columns
					if (*board)[(x+i+rows)%rows][(y+j+columns)%columns].PreviousLife {

						//////////////////////////////////////////////////////////////////
						if (x-r >= -1) && (x-r <= 1) && (y-c >= -1) && (y-c <= 1) {

							neighbors++
							parentsColor = append(parentsColor, (*board)[(x+i+rows)%rows][(y+j+columns)%columns].PreviousColor)

						}
					}

				}
			}
			if (*board)[x][y].PreviousLife {
				neighbors--
			}

			if ((*board)[x][y].Alive == true) && (neighbors < 2) {
				// Loneliness
				(*board)[x][y].KillCell()
				//	(*board)[x][y].ColorHex = "#d8d8d8"
				(*board)[x][y].ColorHex = ""
			} else if ((*board)[x][y].Alive == true) && (neighbors > 3) {
				// Overpopulation
				(*board)[x][y].Alive = false
				//(*board)[x][y].ColorHex = "#d8d8d8"
				(*board)[x][y].ColorHex = ""
			} else if ((*board)[x][y].Alive == false) && (neighbors == 3) {
				// Reproduction
				newColor := GetAverageColor(&parentsColor)

				(*board)[x][y].Alive = true
				(*board)[x][y].ColorHex = newColor
			} else if ((*board)[x][y].Alive == true) && (neighbors == 3 || neighbors == 2) {
			} else {
				//fmt.Printf("=========================>[%d][%d]-- irrelevant !!\n", x, y)
			}
			// else do nothing!
		}
	}
	fmt.Printf("finish propagate")
	return board
}
