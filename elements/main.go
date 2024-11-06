package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

const (
	reset         = "\033[0m"
	black         = "\033[40m"
	red           = "\033[41m"
	yellow        = "\033[43m"
	blue          = "\033[44m"
	white         = "\033[47m"
	green         = "\033[42m"
	cyan          = "\033[46m"
	purple        = "\033[45m"
	BrightBlackBG = "\033[100m"
	BrightRedBG   = "\033[101m"
	BrightGreenBG = "\033[102m"
)

type Cell struct {
	isBomb        bool
	isRevealed    bool
	adjacentBombs int
}

type Game struct {
	grid       [][]Cell
	height     int
	width      int
	bombCount  int
	movesCount int
}

func main() {
	// 1. input and choose mode (validations.go)
	choice := checkMode()
	// 2. creategrid

	println("Enter height and width: ")

	var input1 []rune
	var char1 rune
	for {
		_, err := fmt.Scanf("%c", &char1)
		if err != nil || char1 == '\n' {
			break
		}
		input1 = append(input1, char1)
	}

	var rowArr1, colArr1 []int
	seenNonSpace1 := false

	for i := 0; i < len(input1); i++ {
		if input1[i] != ' ' {
			if !seenNonSpace1 {
				rowArr1 = append(rowArr1, int(input1[i]-'0'))
				// seenNonSpace1 = true
			} else {
				colArr1 = append(colArr1, int(input1[i]-'0'))
			}
		} else if len(rowArr1) > 0 && !seenNonSpace1 {
			seenNonSpace1 = true

			for i < len(input1) && input1[i] == ' ' {
				i++
			}
			if i < len(input1) && input1[i] != ' ' {
				colArr1 = append(colArr1, int(input1[i]-'0'))
			}
		}
	}

	h := 0
	w := 0

	for i := 0; i < len(rowArr1); i++ {
		h = h*10 + rowArr1[i]
	}
	for i := 0; i < len(colArr1); i++ {
		w = w*10 + colArr1[i]
	}

	// fmt.Scanf("%d %d", &h, &w)
	if checkGrid(h, w) == false {
		return
	}

	var game *Game

	if choice == 1 {
		game = NewGame(h, w)
		println("Please enter grid:")
		for i := 0; i < h; i++ {
			var row string
			fmt.Scanf("%s", &row)

			if len(row) != w {
				println("ERROR: Elements out of bound")
				return
			}
			if !checkForDotAndBomb(row) {
				println("ERROR: Invalid elements")
				return
			}

			for j, ch := range row {
				if ch == '*' {

					game.grid[i][j].isBomb = true
					game.bombCount++
				}
			}
		}

		if game.bombCount < 2 {
			println("ERROR: Create at least 2 bombs")
			return
		} else if game.bombCount >= game.height*game.width {
			println("ERROR: No empty cells, too many bombs you died")
			return
		}

		game.neighbourBombs()
	} else if choice == 2 {
		game = NewGame(h, w)
	}

	firstMove := true
	fMove := true
	for {
		game.printGrid()
		// var col int
		// var row int

		println("Enter coordinates")

		// n, err := fmt.Scanf("%d %d", &row, &col)
		// if err != nil || n != 2 {
		// 	println("ERROR : use correct input")
		// 	var input []rune
		// 	var char rune
		// 	for {
		// 		_, err := fmt.Scanf("%c", &char)
		// 		if err != nil || char == '\n' {
		// 			break
		// 		}
		// 		input = append(input, char)

		// 	}
		// }

		var input []rune
		var char rune
		for {
			_, err := fmt.Scanf("%c", &char)
			if err != nil || char == '\n' {
				break
			}
			input = append(input, char)
		}

		var rowArr, colArr []int
		seenNonSpace := false

		for i := 0; i < len(input); i++ {
			if input[i] != ' ' {
				if !seenNonSpace {
					rowArr = append(rowArr, int(input[i]-'0'))
				} else {
					colArr = append(colArr, int(input[i]-'0'))
				}
			} else if len(rowArr) > 0 && !seenNonSpace {
				seenNonSpace = true

				for i < len(input) && input[i] == ' ' {
					i++
				}
				if i < len(input) && input[i] != ' ' {
					colArr = append(colArr, int(input[i]-'0'))
				}
			}
		}

		row := 0
		col := 0
		for i := 0; i < len(rowArr); i++ {
			row = row*10 + rowArr[i]
		}
		for i := 0; i < len(colArr); i++ {
			col = col*10 + colArr[i]
		}

		row--
		col--

		// var input string

		if row < 0 || row >= game.height || col < 0 || col >= game.width {
			println("Coordinates don't exist")
			continue
		}

		if firstMove {
			if choice == 2 {
				game.placeBombs(col, row)
			}
			firstMove = false
		}
		if fMove == true {
			if game.grid[row][col].isBomb == true {
				replaceMine(game.grid, row, col)
				game.neighbourBombs()

			}
		}
		fMove = false

		hitBomb := game.openCell(col, row)

		game.movesCount++

		if hitBomb {
			for i := 0; i < game.height; i++ {
				for j := 0; j < game.width; j++ {
					if game.grid[i][j].isBomb == true {
						game.grid[i][j].isRevealed = true
					}
				}
			}
			game.printGrid()
			println("Game Over!")
			stats(game.height, game.width, game.bombCount, game.movesCount)
			return
		}

		if game.isGameWon() {
			game.printGrid()
			println("You Win!")
			stats(game.height, game.width, game.bombCount, game.movesCount)
			return
		}
	}
}

func stats(height, width, bombCount, movesCount int) {
	print("Your statistics:\nField size:")
	PutNumber(height)
	ap.PutRune('x')
	PutNumber(width)
	ap.PutRune('\n')
	print("Number of bombs: ")
	PutNumber(bombCount)
	ap.PutRune('\n')
	print("Number of moves: ")
	PutNumber(movesCount)
	ap.PutRune('\n')
}

func Contains(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return false
		}
	}

	return true
}
