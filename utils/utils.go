package utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
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

func StringToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func StringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func StringToFloat32(str string) float32 {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return float32(num)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func RandomColor() rl.Color {
	return rl.NewColor(uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 200+uint8(rand.Intn(56)))
}
