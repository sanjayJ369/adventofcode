package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileIntoLines("./input.txt")

	dail := 50
	count := 0

	for _, line := range lines {
		prev := count
		prevDial := dail

		num := utils.StringToInt(line[1:])
		if num >= 100 {
			count += num / 100 // part two: counting clicks during rotation
			num = num % 100
		}

		if line[0] == 'R' {
			dail += num
		} else {
			dail -= num
		}

		if dail < 0 {
			dail += 100
			if dail != 0 && prevDial != 0 {
				count += 1
			}
		} else if dail >= 100 {
			dail -= 100
			if dail != 0 && prevDial != 0 {
				count += 1
			}
		}

		if dail == 0 {
			count += 1
		}

		fmt.Println(count - prev)
	}

	fmt.Println(count)

}
