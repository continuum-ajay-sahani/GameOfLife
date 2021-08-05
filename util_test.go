package main

import "testing"

func TestCreate2DArray(t *testing.T) {
	rows := 4
	cols := 3

	arr := create2DAraay(rows, cols)
	if rows != len(arr) {
		t.Errorf("expected rows= %d but found %d", rows, len(arr))
	}
	if cols != len(arr[0]) {
		t.Errorf("expected cols= %d but found %d", cols, len(arr))
	}
}

func TestGetCopyArray(t *testing.T) {
	arr := [][]int{
		{1, 0, 1},
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, 0},
	}
	carr := getCopyArray(arr)

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[0]); col++ {
			if carr[row][col] != arr[row][col] {
				t.Errorf("expected value %d at [%d, %d] but found = %d", arr[row][col], row, col, carr[row][col])
			}
		}
	}
}

func TestCopyArray(t *testing.T) {
	src := [][]int{
		{1, 0, 1},
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, 0},
	}

	dest := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	copyArray(src, dest)

	for row := 0; row < len(src); row++ {
		for col := 0; col < len(src[0]); col++ {
			if src[row][col] != dest[row][col] {
				t.Errorf("expected value %d at [%d, %d] but found = %d", src[row][col], row, col, dest[row][col])
			}
		}
	}
}
