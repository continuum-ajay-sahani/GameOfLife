package main

import "fmt"

func main() {
	inputArr := get2DDemoBinaryInputArray(4, 3)
	fmt.Println("------------------------Input--------------------------")
	fmt.Println(inputArr)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------------------Output-------------------------")
	gameOfLife(inputArr)
	fmt.Println(inputArr)
	fmt.Println("-------------------------------------------------------")

}

func gameOfLife(grid [][]int) {
	rows := len(grid)
	cols := len(grid[0])
	cGrid := getCopyArray(grid)
	fmt.Println("-------------------------------------------------------")
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			count := getLiveNeighborsCount(grid, row, col)
			nextVal := getNextGenerationValue(grid[row][col], count)
			fmt.Println(fmt.Sprintf("row=%d, col=%d, count=%d, nextVal=%d", row, col, count, nextVal))
			cGrid[row][col] = nextVal
		}
	}
	fmt.Println("-------------------------------------------------------")
	copyArray(cGrid, grid)
}

func getNextGenerationValue(element, liveNeighborsCount int) int {
	if element == 1 && (liveNeighborsCount < 2 || liveNeighborsCount > 3) {
		/* Rule 1 or Rule 3 */
		element = 0
	} else if element == 0 && liveNeighborsCount == 3 {
		/* Rule 4 */
		element = 1
	}
	return element
}

func getLiveNeighborsCount(grid [][]int, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			r := row + i
			c := col + j
			if r >= 0 && r < rows && c >= 0 && c < cols {
				count += grid[r][c]
			}
		}
	}
	count -= grid[row][col]
	return count
}
