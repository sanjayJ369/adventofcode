package main

// import (
// 	"aoc/utils"
// 	"fmt"
// 	"time"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// var (
// 	cellSize     = int32(6)
// 	screenWidth  = int32(300)
// 	screenHeight = int32(800)
// 	splits       = 0
// 	finished     = false
// )

// const (
// 	TypeEmpty = iota
// 	TypeBeam
// 	TypeSplitter
// 	TypeSource
// 	LeftDir
// 	RightDir
// 	StringtDir
// )

// type CellType int

// type Cell struct {
// 	Type       CellType
// 	Split      bool
// 	Dir        int
// 	TimesLines []int
// }

// func (c Cell) Symbol() string {
// 	switch c.Type {
// 	case TypeBeam:
// 		switch c.Dir {
// 		case LeftDir:
// 			return "/"
// 		case RightDir:
// 			return "\\"

// 		default:
// 			return "|"
// 		}
// 	case TypeSplitter:
// 		return "^"
// 	case TypeSource:
// 		return "S"
// 	default:
// 		return "."
// 	}
// }

// func (c Cell) Color() rl.Color {
// 	switch c.Type {
// 	case TypeBeam:
// 		return rl.Orange
// 	case TypeSplitter:
// 		return rl.Blue
// 	case TypeSource:
// 		return rl.Green
// 	default:
// 		return rl.Black
// 	}
// }

// func render(matrix [][]Cell) {
// 	n, m := len(matrix), len(matrix[0])

// 	// Convert cellSize once if it's a constant
// 	size := int32(cellSize)

// 	for r := 1; r <= n; r++ {
// 		for c := 1; c <= m; c++ {
// 			cell := matrix[r-1][c-1]

// 			x := int32(c) * size
// 			y := int32(r) * size

// 			rl.DrawText(cell.Symbol(), x, y, size/2, cell.Color())
// 		}
// 	}

// 	msg := fmt.Sprintf("Splits: %d", splits)
// 	rl.DrawText(msg, 10, int32(n+1)*size, 24, rl.DarkGray)
// }

// func update(current [][]Cell) [][]Cell {
// 	n, m := len(current), len(current[0])
// 	next := make([][]Cell, n)
// 	for r := range current {
// 		next[r] = make([]Cell, m)
// 		copy(next[r], current[r])
// 	}

// 	if !finished {

// 		for r := 0; r < n; r++ {
// 			for c := 0; c < m; c++ {
// 				switch current[r][c].Type {
// 				case TypeEmpty:
// 					if r-1 >= 0 && current[r-1][c].Type == TypeBeam {
// 						next[r][c].Type = TypeBeam
// 						next[r][c].Dir = StringtDir
// 					} else if c-1 >= 0 && current[r][c-1].Split {
// 						next[r][c].Type = TypeBeam
// 						next[r][c].Dir = RightDir
// 					} else if c+1 < m && current[r][c+1].Split {
// 						next[r][c].Type = TypeBeam
// 						next[r][c].Dir = LeftDir
// 					} else if r-1 >= 0 && current[r-1][c].Type == TypeSource {
// 						next[r][c].Type = TypeBeam
// 						next[r][c].Dir = StringtDir
// 					}
// 				case TypeSplitter:
// 					if r-1 >= 0 && (current[r-1][c].Type == TypeBeam ||
// 						current[r-1][c].Type == TypeSource) && !current[r][c].Split {
// 						next[r][c].Split = true
// 						splits++ // imp
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return next
// }

// func main() {
// 	time.Sleep(5 * time.Second)
// 	matrixStr := utils.ReadFileAsMatrixOfString("./input.txt")
// 	mapFunc := func(cell string) Cell {
// 		switch cell {
// 		case "^":
// 			return Cell{
// 				Type: TypeSplitter,
// 			}
// 		case "S":
// 			return Cell{
// 				Type: TypeSource,
// 			}
// 		default:
// 			return Cell{
// 				Type: TypeEmpty,
// 			}
// 		}
// 	}

// 	matrix := utils.PerformCell(matrixStr, mapFunc)
// 	n, m := len(matrix), len(matrix[0])
// 	screenHeight = int32(m+5) * cellSize
// 	screenWidth = int32(n+1) * cellSize

// 	rl.InitWindow(screenWidth, screenHeight, "Visualizing the Beams...!!!!")
// 	defer rl.CloseWindow()
// 	rl.SetTargetFPS(60)

// 	updateTime := 100 * time.Millisecond
// 	lastUpdate := time.Now()

// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.Black)

// 		if time.Since(lastUpdate) > updateTime {
// 			matrix = update(matrix)
// 			lastUpdate = time.Now()
// 		}
// 		render(matrix)

// 		rl.EndDrawing()
// 	}
// }
