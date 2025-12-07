package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

// as per the problem
// it returns true
// if the num is made up of two nums
// such that num = num1 + num2
// where num1 == num2
func doesNumContainsDuplicates(num int) bool {
	numstr := strconv.Itoa(num)
	n := len(numstr)
	if n%2 != 0 {
		return false
	}

	num1str, num2str := numstr[:n/2], numstr[n/2:]
	num1, num2 := utils.StringToInt(num1str), utils.StringToInt(num2str)

	return num1 == num2
}

func isNumMadeUpOfRepeatedNums(num int, rep int) bool {
	numStr := strconv.Itoa(num)
	n := len(numStr)

	if n%rep != 0 {
		return false
	}

	dt := n / rep
	splits := []int{}
	for i := 0; i < n; i += dt {
		slice := numStr[i : i+dt]
		splits = append(splits, utils.StringToInt(slice))
	}

	ele := splits[0]
	for _, split := range splits {
		if split != ele {
			return false
		}
	}

	return true
}

// func test() {

// 	fmt.Println(isNumMadeUpOfRepeatedNums(2121212121, 5))
// 	fmt.Println(isNumMadeUpOfRepeatedNums(824824824, 3))
// 	fmt.Println(isNumMadeUpOfRepeatedNums(11, 2))

// 	os.Exit(1)
// }

func main() {
	// test()
	input := utils.ReadFileIntoLines("./input.txt")[0]
	rangesRaw := strings.Split(input, ",")
	ranges := [][]int{}

	sum := 0
	for _, r := range rangesRaw {
		r := strings.Split(r, "-")
		start, stop := r[0], r[1]
		startNum := utils.StringToInt(start)
		stopNum := utils.StringToInt(stop)

		ranges = append(ranges, []int{startNum, stopNum})
	}

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			n := len(strconv.Itoa(i))
			for j := 2; j <= n; j++ {
				if isNumMadeUpOfRepeatedNums(i, j) {
					sum += i
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
