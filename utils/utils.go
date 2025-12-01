package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
