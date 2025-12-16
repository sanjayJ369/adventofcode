package utils

import (
	"fmt"
	"strings"
)

func ReadFileAsMatrixOfNums(file string) [][]int {
	data := ReadFileIntoLines(file)
	var nums [][]int
	for _, line := range data {
		numsRow := []int{}
		for _, ch := range line {
			numsRow = append(numsRow, StringToInt(string(ch)))
		}
		nums = append(nums, numsRow)
	}
	return nums
}

func ReadFileAsMatrixOfFloat64WithSep(file string, sep string) [][]float64 {
	data := ReadFileIntoLines(file)
	var nums [][]float64
	for _, line := range data {
		numsRow := []float64{}
		line := strings.Split(line, sep)
		for _, ch := range line {
			numsRow = append(numsRow, StringToFloat64(string(ch)))
		}
		nums = append(nums, numsRow)
	}
	return nums
}

func ReadFileAsMatrixOfFloat32WithSep(file string, sep string) [][]float32 {
	data := ReadFileIntoLines(file)
	var nums [][]float32
	for _, line := range data {
		numsRow := []float32{}
		line := strings.Split(line, sep)
		for _, ch := range line {
			numsRow = append(numsRow, StringToFloat32(string(ch)))
		}
		nums = append(nums, numsRow)
	}
	return nums
}

func ReadFileAsMatrixOfNumsWithSep(file string, sep string) [][]int {
	data := ReadFileIntoLines(file)
	var nums [][]int
	for _, line := range data {
		numsRow := []int{}
		line := strings.Split(line, sep)
		for _, ch := range line {
			numsRow = append(numsRow, StringToInt(string(ch)))
		}
		nums = append(nums, numsRow)
	}
	return nums
}

func ReadFileAsMatrixOfString(file string) [][]string {
	lines := ReadFileIntoLines(file)
	var res [][]string
	for _, line := range lines {
		var row []string
		for _, r := range line {
			row = append(row, string(r))
		}
		res = append(res, row)
	}
	return res
}

func PrintMatrix[T any](matrix [][]T) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func GetNeighbours[T any](matrix [][]T, row int, col int) ([]T, [][2]int) {
	dirs := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {1, 1}, {-1, 1}, {1, -1},
	}

	n, m := len(matrix), len(matrix[0])
	var result []T
	var index [][2]int

	for _, d := range dirs {
		nr, nm := row+d[0], col+d[1]
		if nr < n && nr >= 0 && nm < m && nm >= 0 {
			result = append(result, matrix[nr][nm])
			index = append(index, [2]int{nr, nm})
		}
	}

	return result, index
}

func PerformCell[T any, R any](matrix [][]T, f func(cell T) R) [][]R {
	n, m := len(matrix), len(matrix[0])

	result := make([][]R, n)
	for i := 0; i < n; i++ {
		result[i] = make([]R, m)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			result[r][c] = f(matrix[r][c])
		}
	}

	return result
}

func PerformAdjacent[T any, R any](matrix [][]T, f func(center T, neighbours []T) R) [][]R {
	n, m := len(matrix), len(matrix[0])

	result := make([][]R, n)
	for i := 0; i < n; i++ {
		result[i] = make([]R, m)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			center := matrix[row][col]
			neighbours, _ := GetNeighbours(matrix, row, col)
			result[row][col] = f(center, neighbours)
		}
	}

	return result
}

func ReadFileAsMatrixOfStringWithSeprator(file string, sep string) [][]string {
	lines := ReadFileIntoLines(file)
	var res [][]string
	for _, line := range lines {
		line = strings.Trim(line, sep)
		row := strings.Split(line, sep)
		var update []string
		for _, ele := range row {
			if len(ele) != 0 {
				update = append(update, ele)
			}
		}
		res = append(res, update)
	}
	return res
}

func GetCol[T any](matrix [][]T, col int) []T {
	n := len(matrix)
	res := make([]T, n)
	for i, row := range matrix {
		res[i] = row[col]
	}
	return res
}
