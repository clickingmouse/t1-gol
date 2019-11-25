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

	fmt.Printf("rows: %d, columns: %d ; backingup\n", rows, columns)
	for i := range *board {
		for j := range (*board)[i] {
			(*board)[i][j].SavePreviousLife()
			(*board)[i][j].RetireCell()
		}
	}

	// Loop through every spot in our 2D array and check spots neighbors
	for x := range *board {
		for y := range (*board)[x] {
			fmt.Printf("=================================================>propagating Cell [%d][%d]\n", y, x)
			// Add up all the states in a 3x3 surrounding grid
			neighbors := 0
			var parentsColor []string
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					//neighbors += (*board)[(x+i+columns)%columns][(y+j+rows)%rows].PreviousLife

					// if (*board)[(x+i+columns)%columns][(y+j+rows)%rows].PreviousLife {
					// 	c := (x + i + columns) % columns
					// 	r := (y + j + rows) % rows
					r := (x + i + rows) % rows
					c := (y + j + columns) % columns
					fmt.Printf("================>analyzing neighbor [%d][%d], color [%s]\n", c, r, (*board)[(x+i+rows)%rows][(y+j+columns)%columns].ColorHex)
					fmt.Printf("================>analyzing neighbor [%d][%d], color [%s]\n", r, c, (*board)[(x+i+rows)%rows][(y+j+columns)%columns].ColorHex)
					fmt.Printf("r (x+i+rows %% rows) => %d + %d +%d %% %d = [%d]\n", x, i, rows, rows, r)
					fmt.Printf("c (y+j+columns %% columns) => %d + %d +%d %% %d =[%d]\n", y, j, columns, columns, c)
					if (*board)[(x+i+rows)%rows][(y+j+columns)%columns].PreviousLife {
						fmt.Printf("[%d][%d] chcecking if valid cell\n", c, r)
						// r := (x + i + rows) % rows
						// c := (y + j + columns) % columns
						// fmt.Printf("==>analyzing [%d][%d]\n", c, r)
						// fmt.Printf("r (x+i+rows %% rows) => %d + %d +%d %% %d = [%d]\n", x, i, rows, rows, r)
						// fmt.Printf("c (y+j+columns %% columns) => %d + %d +%d %% %d =[%d]\n", y, j, columns, columns, c)
						fmt.Printf("[x=%d] [r=%d] || [y=%d] [c=%d]\n", x, r, y, c)

						//////////////////////////////////////////////////////////////////
						//						if (x+i >= 0) && (y+j >= 0) {
						if (x-r >= -1) && (x-r <= 1) && (y-c >= -1) && (y-c <= 1) {
							//if math.Abs(float64((x+i+columns)%columns)) < float64(columns) {
							fmt.Printf("found live neighbout at cell [%d][%d]\n", c, r)

							neighbors++
							fmt.Printf("saving color dna [%s]\n", (*board)[(x+i+rows)%rows][(y+j+columns)%columns].PreviousColor)
							parentsColor = append(parentsColor, (*board)[(x+i+rows)%rows][(y+j+columns)%columns].PreviousColor)
							// add parents dna here

						}
					}

				}
			}
			fmt.Printf("----------------------total neighbours is %d\n", neighbors)
			// A little trick to subtract the current cell's state since
			// we added it in the above loop
			//neighbors -= (*board)[x][y].PreviousLife
			if (*board)[x][y].PreviousLife {
				neighbors--
			}
			// Rules of Life
			//fmt.Println("evolving cell#", x, y, (*board)[x][y].Alive, neighbors)
			fmt.Printf("evolving cell #[%d][%d] wasAlive? %t, aliveNeighbors:%d\n", x, y, (*board)[x][y].Alive, neighbors)
			//fmt.println("")
			if ((*board)[x][y].Alive == true) && (neighbors < 2) {
				// Loneliness
				//(*board)[x][y].Alive = false
				fmt.Printf("=========================>[%d][%d]-- UnderPopulation !!\n", x, y)
				(*board)[x][y].KillCell()
				//	(*board)[x][y].ColorHex = "#d8d8d8"
				(*board)[x][y].ColorHex = ""
			} else if ((*board)[x][y].Alive == true) && (neighbors > 3) {
				// Overpopulation
				fmt.Printf("=========================>[%d][%d]-- Over crowding !!\n", x, y)
				(*board)[x][y].Alive = false
				//(*board)[x][y].ColorHex = "#d8d8d8"
				(*board)[x][y].ColorHex = ""
			} else if ((*board)[x][y].Alive == false) && (neighbors == 3) {
				// Reproduction
				fmt.Printf("=========================>[%d][%d]-- Reproduction !!\n", x, y)
				fmt.Printf("colors dna : %+v", parentsColor)
				newColor := GetAverageColor(&parentsColor)
				fmt.Printf("spawned color : %+v", newColor)

				(*board)[x][y].Alive = true
				(*board)[x][y].ColorHex = newColor
			} else if ((*board)[x][y].Alive == true) && (neighbors == 3 || neighbors == 2) {
				fmt.Printf("=========================>[%d][%d]-- survives !!\n", x, y)
			} else {
				fmt.Printf("=========================>[%d][%d]-- irrelevant !!\n", x, y)
			}
			// else do nothing!
		}
	}

	return board
}
