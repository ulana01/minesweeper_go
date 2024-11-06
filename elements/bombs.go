package main

import (
	"math/rand"
)

func replaceMine(board [][]Cell, row, col int) {
	for {
		// Generate random coordinate
		newRow := rand.Intn(len(board))
		newCol := rand.Intn(len(board[0]))

		// If cell is empty, move the bomb here
		if board[newRow][newCol].isBomb == false {
			board[newRow][newCol].isBomb = true
			board[row][col].isBomb = false
			break
		}
	}
}

func (g *Game) neighbourBombs() {
	// Iterate through the grid
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if !g.grid[y][x].isBomb {
				count := 0 // num of bombs
				// Check cells in all directions using DFS
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dx == 0 && dy == 0 {
							continue
						}
						// If bomb is in the cell, increase the count
						nx, ny := x+dx, y+dy
						if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height && g.grid[ny][nx].isBomb {
							count++
						}
					}
				}
				g.grid[y][x].adjacentBombs = count
			}
		}
	}
}

func (g *Game) placeBombs(firstMoveX, firstMoveY int) {
	g.bombCount = ((g.height * g.width) / 5) // 1/5 of cells are bombs
	if g.bombCount < 2 {
		g.bombCount = 2 // Minimum num of bombs is 2
	}
	// Place bombs randomly
	for i := 0; i < g.bombCount; i++ {
		for {
			x, y := rand.Intn(g.width), rand.Intn(g.height)
			// If cell is empty and it's not the first move, bomb is set
			if !g.grid[y][x].isBomb && (x != firstMoveX || y != firstMoveY) {
				g.grid[y][x].isBomb = true
				break
			}
		}
	}

	g.neighbourBombs()
}
