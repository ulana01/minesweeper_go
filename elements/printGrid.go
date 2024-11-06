package main

import "github.com/alem-platform/ap"

func (g *Game) printGrid() {
	printCells(g.width)
	// Properties of cell
	cellWidth := 7
	cellHeight := 3

	for i := 0; i < 3; i++ {
		ap.PutRune(' ')
	}
	// Print the roof of the grid
	for i := 0; i < g.width*(cellWidth+1)-1; i++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')
	// Iterate through the grid
	for i := 0; i < g.height; i++ {
		for ch := 0; ch < cellHeight; ch++ {
			// Put numbers for the borders of the grid
			if ch == 1 {
				putVertNumber(i + 1)
				if i < 9 {
					ap.PutRune(' ')
				}
			} else {
				ap.PutRune(' ')
				ap.PutRune(' ')
			}
			// Print the wall
			for j := 0; j < g.width; j++ {
				ap.PutRune('|')
				// Check each cell
				cell := g.grid[i][j]
				cellColor := reset
				if cell.isRevealed {
					if cell.isBomb {
						cellColor = red
					} else {
						cellColor = paintCell(cell.adjacentBombs)
					}
				}
				print(cellColor)
				// Print contents of each cell

				// Cell is closed
				if !cell.isRevealed {
					for x := 0; x < cellWidth; x++ {
						ap.PutRune('X')
					}
				} else if cell.isBomb {
					// Bomb is in the cell
					if ch == 1 {
						for l := 0; l < 3; l++ {
							ap.PutRune(' ')
						}
						ap.PutRune('*')
						for r := 0; r < 3; r++ {
							ap.PutRune(' ')
						}
					} else if ch == 2 {
						for g := 0; g < cellWidth; g++ {
							ap.PutRune('_')
						}
					} else {
						for g := 0; g < cellWidth; g++ {
							ap.PutRune(' ')
						}
					}
				} else { // Cell is opened and shows num of adjacent cells
					if ch == 1 && cell.adjacentBombs > 0 {
						for l := 0; l < 3; l++ {
							ap.PutRune(' ')
						}
						ap.PutRune(rune('0' + cell.adjacentBombs))
						for r := 0; r < 3; r++ {
							ap.PutRune(' ')
						}
					} else if ch == 2 {
						for g := 0; g < cellWidth; g++ {
							ap.PutRune('_')
						}
					} else {
						for g := 0; g < cellWidth; g++ {
							ap.PutRune(' ')
						}
					}
				}
				print(reset)
			}

			ap.PutRune('|')
			if ch != 3 {
				ap.PutRune('\n')
			}

		}
	}
}

func putVertNumber(n int) {
	// Print num for the vertical walls
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	if n == 0 {
		ap.PutRune('0')
		return
	}
	d := []rune{}
	for n > 0 {
		x := n % 10
		d = append(d, rune('0'+x))
		n /= 10
	}
	for i := len(d) - 1; i >= 0; i-- {
		ap.PutRune(d[i])
	}
}

func PutNumber(n int) {
	// Function to print number with putRune
	tempNum := n
	digNum := 0
	modNum := 0

	var arr []rune
	for tempNum != 0 {
		modNum = tempNum % 10
		tempNum = tempNum / 10
		arr = append(arr, rune(modNum))
		digNum++
	}

	for a := digNum; a > 0; a-- {
		ap.PutRune(rune(arr[a-1] + '0'))
	}
}

func printCells(w int) {
	// Print out contents of the opened cells
	count := 1
	for i := 0; i < 3; i++ {
		ap.PutRune(' ')
	}
	for k := 0; k < w; k++ {
		for i := 0; i < 7; i++ {
			if count <= 9 {
				if i == 3 {
					PutNumber(count)
				} else {
					ap.PutRune(' ')
				}
			} else if count <= 99 {
				if i == 2 {
					PutNumber(count)
				} else if i < 6 {
					ap.PutRune(' ')
				}
			}
		}
		count++
		ap.PutRune(' ')

	}
	ap.PutRune('\n')
}
