package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return lines
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

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func StringSliceToInts(lines []string) []int {
	res := []int{}
	for _, line := range lines {
		res = append(res, StringToInt(line))
	}
	return res
}

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

func CombineIntSliceIntoString(nums []int) string {
	var res strings.Builder
	for _, n := range nums {
		res.WriteString(strconv.Itoa(n))
	}
	return res.String()
}

// returns minimumNumber, it's index
func FindMinAndIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, -1
	}

	minNum := 9
	minIdx := -1

	for i, n := range nums {
		if n < minNum {
			minNum = n
			minIdx = i
		}
	}

	return minNum, minIdx
}

func RemoveKDigitsFromSlice(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		nums = RemoveDigitFromSliceToGetLargetNumber(nums)
	}
	return nums
}

// RemoveDigitFromSliceToGetLargetNumber
// remove a digit from a number which is split into []int
// to get the largest number after removing a digit
func RemoveDigitFromSliceToGetLargetNumber(nums []int) []int {
	n := len(nums)
	for i, n := range nums[:n-1] {
		if n < nums[i+1] {
			return append(nums[:i], nums[i+1:]...)
		}
	}

	return nums[:n-1]
}

// returns maximumNumber, it's index
func FindMaxAndIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, -1
	}

	maxNum := nums[0]
	maxIdx := 0

	for i, n := range nums {
		if n > maxNum {
			maxNum = n
			maxIdx = i
		}
	}

	return maxNum, maxIdx
}
