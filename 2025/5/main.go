package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")
	var ranges [][2]uint64
	var ingredients []uint64

	i := 0
	for _, line := range lines {
		if line == "" {
			i++
			break
		}
		var r [2]uint64
		digits := strings.Split(line, "-")
		r[0] = utils.StringToUint64(digits[0])
		r[1] = utils.StringToUint64(digits[1])
		ranges = append(ranges, r)
		i++
	}

	for ; i < len(lines); i++ {
		ingredients = append(ingredients, utils.StringToUint64(lines[i]))
	}

	count := 0
	for _, i := range ingredients {
		for _, r := range ranges {
			if r[0] <= i && r[1] >= i {
				count++
				break
			}
		}
	}

	fmt.Println(count)

	// for part two
	// sort the ranges
	// combine them
	sortRangesRunc := func(a, b [2]uint64) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			} else if a[1] > b[1] {
				return 1
			} else {
				return 0
			}
		} else {
			return 0
		}
	}

	slices.SortFunc(ranges, sortRangesRunc)
	combinedRanges := [][2]uint64{}
	for _, r := range ranges {
		if len(combinedRanges) == 0 {
			combinedRanges = append(combinedRanges, r)
		} else {
			// check if they overlap
			n := len(combinedRanges)
			last := combinedRanges[n-1]
			if r[0] <= last[1] {
				if r[1] > last[1] {
					combinedRanges[n-1][1] = r[1]
				}
			} else {
				combinedRanges = append(combinedRanges, r)
			}
		}
	}

	var fresh uint64 = 0
	for _, r := range combinedRanges {
		fresh += r[1] - r[0] + 1
	}
	fmt.Println(fresh)
}
