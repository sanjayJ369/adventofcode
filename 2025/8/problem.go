package main

import (
	"aoc/utils"
	"container/heap"
	"fmt"
	"math"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type JunctionBox struct {
	rl.Vector3
	rawPos rl.Vector3
	idx    int // parent
	color  rl.Color
}

type Set []int

func NewSet(count int) Set {
	s := make([]int, count)
	for i := range s {
		s[i] = i
	}
	return s
}

func (s *Set) SameSet(a, b int) bool {
	return s.FindParent(a) == s.FindParent(b)
}

func (s *Set) AddEle(ele int) {
	(*s)[ele] = ele
}

func (s Set) FindParent(ele int) int {
	if s[ele] == ele {
		return ele
	}
	root := s.FindParent(s[ele])
	s[ele] = root
	return root
}

func (s Set) GetSetCount() []int {
	var res []int
	var count map[int]int = make(map[int]int)

	for i := range s {
		root := s.FindParent(i) // Find the boss for this box
		count[root]++           // Add 1 to that boss's team size
	}

	for _, c := range count {
		res = append(res, c)
	}

	slices.Sort(res)
	return res
}

func (s Set) GetSameSetElements(a int) []int {
	par := s.FindParent(a)
	var res []int
	for i := range s {
		if par == s.FindParent(i) {
			res = append(res, i)
		}
	}
	return res
}

// connect elements
func (s *Set) ConnectElements(a int, b int) {
	if a > b {
		s.ConnectElements(b, a)
		return
	}

	child := s.FindParent(b)
	(*s)[child] = a // making a pariend of b
}

var (
	boxes     []JunctionBox
	distances *utils.FloatHeap
	myset     Set
	lines     [][]rl.Vector3
	ConnCount int  = 0
	MaxConn   int  = 1000
	finished  bool = false
)

func init() {
	distances = &utils.FloatHeap{}
}

func getJunctionBoxes() {
	// problem
	matrix := utils.ReadFileAsMatrixOfFloat32WithSep("./input.txt", ",")
	for i, row := range matrix {
		boxes = append(boxes, JunctionBox{
			Vector3: rl.Vector3Scale(rl.NewVector3(row[0], row[1], row[2]), scale),
			rawPos:  rl.NewVector3(row[0], row[1], row[2]),
			idx:     i,
			color:   utils.RandomColor(),
		})
	}
	myset = NewSet(len(matrix))
}

func connect() {
	if finished || distances.Len() == 0 {
		return
	}

	for distances.Len() > 0 {
		ele := heap.Pop(distances).(utils.Item)
		box1 := ele.Contents[0].(JunctionBox)
		box2 := ele.Contents[1].(JunctionBox)

		if myset.SameSet(box1.idx, box2.idx) {
			continue
		}

		ConnCount++
		lines = append(lines, []rl.Vector3{box1.Vector3, box2.Vector3})

		if box1.idx > box2.idx {
			box1, box2 = box2, box1
		}

		targetColor := boxes[myset.FindParent(box1.idx)].color
		boxes[box2.idx].color = targetColor
		for _, idx := range myset.GetSameSetElements(box2.idx) {
			boxes[idx].color = targetColor
		}

		myset.ConnectElements(box1.idx, box2.idx)

		counts := myset.GetSetCount()
		if slices.Contains(counts, len(boxes)) {
			finished = true

			ans := int(box1.rawPos.X) * int(box2.rawPos.X)
			fmt.Println("------------------------------------------------")
			fmt.Printf("FINAL CONNECTION: %v <-> %v\n", box1.rawPos, box2.rawPos)
			fmt.Printf("PART 2 ANSWER: %d\n", ans)
			fmt.Println("------------------------------------------------")
		}

		return
	}
}

func computeDistances() {
	n := len(boxes)
	heap.Init(distances)
	for i := range n {
		for j := i + 1; j < n; j++ {

			dx := float64(boxes[i].rawPos.X - boxes[j].rawPos.X)
			dy := float64(boxes[i].rawPos.Y - boxes[j].rawPos.Y)
			dz := float64(boxes[i].rawPos.Z - boxes[j].rawPos.Z)

			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
			heap.Push(distances, utils.Item{Rank: dist, Contents: []any{boxes[i], boxes[j]}})
		}
	}

}

func ComputeFinalAnswer(counts []int) int {
	n := len(counts)
	return counts[n-1] * counts[n-2] * counts[n-3]
}
