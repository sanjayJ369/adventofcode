package main

import (
	"aoc/utils"
	"fmt"
	"unicode"
)

// ---------------------------------
// part one
// ---------------------------------
// func main() {
// 	matrix := utils.ReadFileAsMatrixOfStringWithSeprator("./example.txt", " ")
// 	n, m := len(matrix), len(matrix[0])
// 	fmt.Println(n, m)

// 	nums := utils.PerformCell(matrix[:n-1][:], func(num string) int {
// 		return utils.StringToInt(num)
// 	})
// 	ops := matrix[n-1]

// 	res := 0
// 	for i, operator := range ops {
// 		var compute func([]int) int
// 		switch operator {
// 		case "*":
// 			compute = func(i []int) int {
// 				res := 1
// 				for _, num := range i {
// 					res *= num
// 				}
// 				return res
// 			}

// 		case "+":
// 			compute = func(i []int) int {
// 				res := 0
// 				for _, num := range i {
// 					res += num
// 				}
// 				return res
// 			}
// 		}

// 		fmt.Println(utils.GetCol(nums, i))
// 		res += compute(utils.GetCol(nums, i))

// 	}

// 	fmt.Println(res)
// }

func main() {
	matrix := utils.ReadFileAsMatrixOfString("./input.txt")
	utils.PrintMatrix(matrix)
	var nums [][]int
	n, m := len(matrix), len(matrix[0])

	// convert the rows into proper ones
	var row []int
	for i := 0; i < m; i++ {
		col := utils.GetCol(matrix, i)
		if col[n-1] != " " && len(row) > 0 {
			nums = append(nums, row[:n-1])
			row = []int{}
		}

		digit := 0
		for _, i := range col {
			if unicode.IsDigit(rune(i[0])) {
				digit *= 10
				digit += utils.StringToInt(i)
			}
		}
		row = append(row, digit)
	}
	nums = append(nums, row[:n-1])

	var operators []string
	for _, r := range matrix[n-1] {
		if r != " " {
			operators = append(operators, r)
		}
	}

	var res uint64
	for i, operator := range operators {
		var compute func([]int) int
		switch operator {
		case "*":
			compute = func(i []int) int {
				res := 1
				for _, num := range i {
					if num != 0 {
						res *= num
					}
				}
				return res
			}

		case "+":
			compute = func(i []int) int {
				res := 0
				for _, num := range i {
					res += num
				}
				return res
			}
		}
		val := uint64(compute(nums[i]))
		res += val
		fmt.Println(val, nums[i], operator)
	}

	fmt.Println(res)
}
