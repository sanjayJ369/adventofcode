package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	cellSize     = int32(5)
	screenWidth  = int32(300)
	screenHeight = int32(800)
)

func main() {
	matrix := utils.ReadFileAsMatrixOfString("./input.txt")
	n, m := len(matrix), len(matrix[0])
	screenHeight = int32(m)*cellSize + 100 // 100 for displaying data and all
	screenWidth = int32(n) * cellSize

	rl.InitWindow(screenWidth, screenHeight, "Visualizing the Scan")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	scanSpeed := 10 * time.Microsecond
	lastUpdate := time.Now()
	finished := false
	totalCount := 0
	prevScanCount := 0
	r, c := 0, 0
	scans := 1
	stepsPerLoop := 500

	var heighlateIndexes [][2]int

	for !rl.WindowShouldClose() {

		// update phase
		heighlateIndexes = [][2]int{}
		prevRow, prevCol := r, c
		if !finished && time.Since(lastUpdate) > scanSpeed {
			for i := 0; i < stepsPerLoop; i++ {

				if matrix[r][c] == "@" {
					neighbours, indexes := utils.GetNeighbours(matrix, r, c)
					count := 0

					for _, neigh := range neighbours {
						if neigh == "@" || neigh == "X" {
							count++
						}

						if count > 4 {
							break
						}
					}

					if count < 4 {
						matrix[r][c] = "X"
						totalCount++
					}

					heighlateIndexes = indexes
					lastUpdate = time.Now()
				}

				// going to the next scan
				c++
				if c >= m {
					c = 0
					r++
					if r >= n {
						// prep for next scan
						scans += 1
						r, c = 0, 0

						if totalCount == prevScanCount {
							finished = true
						} else {
							prevScanCount = totalCount
						}
						// clear current rolls
						for i := 0; i < n; i++ {
							for j := 0; j < n; j++ {
								if matrix[i][j] == "X" {
									matrix[i][j] = "."
								}
							}
						}
						break
					}
				}
			}
		}

		// draw the matrix
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		xpos := int32(0)
		ypos := int32(0)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var color rl.Color

				switch matrix[i][j] {
				case "@":
					color = rl.Blue
				case ".":
					color = rl.Gray
				case "X":
					color = rl.Green
				default:
					color = rl.Black
				}

				if slices.Contains(heighlateIndexes, [2]int{i, j}) {
					color = rl.Red
				}

				if i == r && j == c {
					color = rl.SkyBlue
				}

				if i == prevRow && j == prevCol {
					color = rl.Black
				}

				rl.DrawRectangle(xpos, ypos, int32(cellSize)-1, int32(cellSize)-1, color)
				xpos += int32(cellSize)
			}
			xpos = 0
			ypos += int32(cellSize)
		}

		textRolls := fmt.Sprintf("%d rolls", totalCount)
		textScanns := fmt.Sprintf("%d", scans)
		rl.DrawText(textRolls, 10, cellSize*int32(n), 24, rl.Black)
		rl.DrawText(textScanns, 10, cellSize*int32(n)+24, 24, rl.Black)

		if finished {
			rl.DrawText("SCAN COMPLETE", 10, int32(n)*cellSize+70, 20, rl.DarkGreen)
		}

		rl.EndDrawing()
	}

}
