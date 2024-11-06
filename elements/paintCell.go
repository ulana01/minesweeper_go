package main

func paintCell(num int) string {
	// Choose color based on number of adjacent bombs
	switch num {
	case 1:
		return green
	case 2:
		return cyan
	case 3:
		return blue
	case 4:
		return yellow
	case 5:
		return purple
	case 6:
		return BrightBlackBG
	case 7:
		return BrightRedBG
	default:
		return black
	}
}
