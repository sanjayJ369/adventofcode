package utils

type Item struct {
	Rank     float64
	Contents []any
}

type FloatHeap []Item

func (f FloatHeap) Len() int {
	return len(f)
}

func (f FloatHeap) Less(i, j int) bool {
	return f[i].Rank < f[j].Rank
}

func (f FloatHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *FloatHeap) Push(x any) {
	item := x.(Item)
	*f = append(*f, item)
}

func (f *FloatHeap) Pop() any {
	old := *f
	n := len(old)
	ele := old[n-1]
	*f = old[:n-1]
	return ele
}
