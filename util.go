package main

import "math/rand"

func copyArray(src, dest [][]int) {
	rows := len(src)
	cols := len(src[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			dest[row][col] = src[row][col]
		}
	}
}

func create2DAraay(rows, cols int) [][]int {
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}
	return arr
}

func getCopyArray(arr [][]int) [][]int {
	rows := len(arr)
	cols := len(arr[0])

	carr := create2DAraay(rows, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			carr[row][col] = arr[row][col]
		}
	}
	return carr
}

func get2DDemoBinaryInputArray(rows, cols int) [][]int {
	arr := create2DAraay(rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			arr[row][col] = rand.Intn(2)
		}
	}
	return arr
}
