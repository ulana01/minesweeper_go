package main

import "fmt"

func checkMode() int {
	var choice int
	for {
		println("Choose mode (1 - custom, 2 - random): ")
		n, err := fmt.Scanf("%d", &choice)

		if err != nil || n != 1 {
			println("Invalid input, please enter a number.")
			var invalidInput string
			fmt.Scanf("%s", &invalidInput)
			continue
		}

		if choice == 1 || choice == 2 {
			return choice
		} else {
			println("Mode doesn't exist")
		}
	}
}

func checkGrid(h, w int) bool {
	if h < 3 || w < 3 {
		println("ERROR: Invalid grid size. Height and width can't be less than 3.")
		return false
	}
	if h > 99 || w > 99 {
		println("ERROR: Invalid grid size.")
		return false
	}
	return true
}

func checkForDotAndBomb(row string) bool {
	for i := 0; i < len(row); i++ {
		if row[i] != '.' && row[i] != '*' {
			return false
		}
	}
	return true
}
