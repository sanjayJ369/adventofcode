package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	cellSize     = int32(6)
	screenWidth  = int32(300)
	screenHeight = int32(800)
	timelines    = 1
)

const (
	TypeEmpty = iota
	TypeBeam
	TypeSplitter
	TypeSource
	LeftDir
	RightDir
	StraightDir
)

type CellType int

type Cell struct {
	Type  CellType
	Split bool
	Dir   int
	Rays  int
}

func (c Cell) Symbol() string {
	if c.Rays > 0 {
		return strconv.Itoa(c.Rays)
	}
	switch c.Type {
	case TypeBeam:
		switch c.Dir {
		case LeftDir:
			return "/"
		case RightDir:
			return "\\"
		default:
			return "|"
		}
	case TypeSplitter:
		return "^"
	case TypeSource:
		return "S"
	default:
		return "."
	}
}

func (c Cell) Color() rl.Color {
	if c.Rays > 0 {
		return rl.White
	}
	switch c.Type {
	case TypeBeam:
		return rl.Orange
	case TypeSplitter:
		return rl.Blue
	case TypeSource:
		return rl.Green
	default:
		return rl.DarkGray
	}
}

func render(matrix [][]Cell) {
	n, m := len(matrix), len(matrix[0])
	size := int32(cellSize)

	for r := 1; r <= n; r++ {
		for c := 1; c <= m; c++ {
			cell := matrix[r-1][c-1]
			x := int32(c) * size
			y := int32(r) * size
			rl.DrawText(cell.Symbol(), x, y, size/2, cell.Color())
		}
	}

	msg := fmt.Sprintf("Timelines: %d", timelines)
	rl.DrawText(msg, 10, int32(n+1)*size, 24, rl.RayWhite)
}

func update(current [][]Cell) [][]Cell {
	n, m := len(current), len(current[0])
	next := make([][]Cell, n)

	// 1. Prepare Next Frame
	for r := range current {
		next[r] = make([]Cell, m)
		// Copy Splitters and Sources ONLY. Beams must move manually.
		for c := 0; c < m; c++ {
			if current[r][c].Type == TypeSplitter || current[r][c].Type == TypeSource {
				next[r][c] = current[r][c]
				next[r][c].Rays = 0
				next[r][c].Split = false
			} else {
				// Keep the road (TypeBeam) for visuals, but wipe Rays
				if current[r][c].Type == TypeBeam {
					next[r][c].Type = TypeBeam
				} else {
					next[r][c].Type = TypeEmpty
				}
				next[r][c].Rays = 0
			}
		}
	}

	// 2. Move Logic
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {

			// --- Vertical Movement (Down) ---
			if r-1 >= 0 && current[r-1][c].Type == TypeBeam && current[r-1][c].Rays > 0 {
				// CRITICAL FIX: Only move beam here if it is NOT hitting a splitter.
				// If it hits a splitter, it stops and is handled by the splitter logic below.
				if next[r][c].Type != TypeSplitter {
					next[r][c].Type = TypeBeam
					next[r][c].Dir = StraightDir
					next[r][c].Rays += current[r-1][c].Rays
				}
			}

			// --- Horizontal Movement (From Splitters) ---
			// Check Left Neighbor pushing Right
			if c-1 >= 0 && current[r][c-1].Split {
				if current[r][c-1].Rays > 0 {
					next[r][c].Type = TypeBeam
					next[r][c].Dir = RightDir
					next[r][c].Rays += current[r][c-1].Rays
				}
			}
			// Check Right Neighbor pushing Left
			if c+1 < m && current[r][c+1].Split {
				if current[r][c+1].Rays > 0 {
					next[r][c].Type = TypeBeam
					next[r][c].Dir = LeftDir
					next[r][c].Rays += current[r][c+1].Rays
				}
			}

			// --- Splitter Firing Logic ---
			if current[r][c].Type == TypeSplitter {
				// If a beam hit us from above in the PREVIOUS frame
				if r-1 >= 0 && current[r-1][c].Type == TypeBeam && current[r-1][c].Rays > 0 {
					next[r][c].Split = true
					next[r][c].Rays += current[r-1][c].Rays
					// Every split doubles the timelines for these rays.
					// Net increase = incoming rays.
					timelines += current[r-1][c].Rays
				}
			}
		}
	}
	return next
}

func main() {
	time.Sleep(1 * time.Second)
	matrixStr := utils.ReadFileAsMatrixOfString("./input.txt")
	mapFunc := func(cell string) Cell {
		switch cell {
		case "^":
			return Cell{Type: TypeSplitter}
		case "S":
			return Cell{Type: TypeSource}
		default:
			return Cell{Type: TypeEmpty}
		}
	}

	matrix := utils.PerformCell(matrixStr, mapFunc)
	n, m := len(matrix), len(matrix[0])
	screenHeight = int32(n+5) * cellSize
	screenWidth = int32(m+5) * cellSize

	// Initial Kickstart: Find S and spawn 1 beam below it
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if matrix[r][c].Type == TypeSource {
				if r+1 < n {
					matrix[r+1][c].Type = TypeBeam
					matrix[r+1][c].Rays = 1
					timelines = 1
				}
			}
		}
	}

	rl.InitWindow(screenWidth, screenHeight, "Visualizing the Beams")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	updateTime := 150 * time.Millisecond
	lastUpdate := time.Now()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if time.Since(lastUpdate) > updateTime {
			matrix = update(matrix)
			lastUpdate = time.Now()
		}
		render(matrix)

		rl.EndDrawing()
	}
}
