package main

import (
	"aoc/utils"
	"fmt"
	"math/big"
)

// part one
// func main() {
// 	data := utils.ReadFileAsMatrixOfNums("./input.txt")
// 	joltageTotal := 0
// 	for _, row := range data {
// 		digit1, idx := utils.FindMaxAndIndex(row[:len(row)-1])
// 		digit2, _ := utils.FindMaxAndIndex(row[idx+1:])
// 		joltageTotal += (digit1 * 10) + digit2
// 		fmt.Println(digit1, digit2)
// 	}
// 	fmt.Println(joltageTotal)
// }

// part two
func main() {
	data := utils.ReadFileAsMatrixOfNums("./input.txt")
	var joltageTotal big.Int

	for _, row := range data {
		var res big.Int
		row = utils.RemoveKDigitsFromSlice(row, len(row)-12)
		resStr := utils.CombineIntSliceIntoString(row)
		res.SetString(resStr, 10)
		joltageTotal.Add(&joltageTotal, &res)
	}
	fmt.Println(joltageTotal.String())
}
