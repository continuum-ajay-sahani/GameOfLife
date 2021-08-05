package main

import "testing"

func TestGetLiveNeighborsCount(t *testing.T) {
	arr := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	cArr := [][]int{
		{1, 1, 2},
		{3, 5, 3},
		{1, 3, 2},
		{2, 3, 2},
	}

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[0]); col++ {
			count := getLiveNeighborsCount(arr, row, col)
			if count != cArr[row][col] {
				t.Errorf("expected value %d at [%d, %d] but found = %d", cArr[row][col], row, col, count)
			}
		}
	}
}

func TestGetNextGenerationValue(t *testing.T) {
	arr := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	count := [][]int{
		{1, 1, 2},
		{3, 5, 3},
		{1, 3, 2},
		{2, 3, 2},
	}

	next := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[0]); col++ {
			nextVal := getNextGenerationValue(arr[row][col], count[row][col])
			if nextVal != next[row][col] {
				t.Errorf("expected value %d at [%d, %d] but found = %d", next[row][col], row, col, nextVal)
			}
		}
	}
}

func TestGameOfLife(t *testing.T) {
	arr := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	next := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}

	gameOfLife(arr)

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[0]); col++ {
			if arr[row][col] != next[row][col] {
				t.Errorf("expected value %d at [%d, %d] but found = %d", next[row][col], row, col, arr[row][col])
			}
		}
	}
}
