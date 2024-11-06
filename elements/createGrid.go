package main

func (g *Game) initializeGrid() {
	// Create a slice for the grid
	g.grid = make([][]Cell, g.height)
	for i := range g.grid {
		g.grid[i] = make([]Cell, g.width)
	}
}

func NewGame(height, width int) *Game {
	// Create new struct for the game grid
	game := &Game{
		height: height,
		width:  width,
	}
	game.initializeGrid()
	return game
}

func (g *Game) openCell(x, y int) bool {
	// Validate cell coordinates
	if x < 0 || x >= g.width || y < 0 || y >= g.height || g.grid[y][x].isRevealed {
		return false
	}

	g.grid[y][x].isRevealed = true
	// If bomb is opened, return true
	if g.grid[y][x].isBomb {
		return true
	}
	// Count adjacent bombs
	if g.grid[y][x].adjacentBombs == 0 {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				g.openCell(x+dx, y+dy)
			}
		}
	}

	return false
}

func (g *Game) isGameWon() bool {
	// If no bombs were opened and all empty cells are opened, game is won
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if !g.grid[y][x].isRevealed && !g.grid[y][x].isBomb {
				return false
			}
		}
	}
	return true
}
