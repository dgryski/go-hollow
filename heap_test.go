package hollow

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {

	keys := []int{14, 11, 5, 9, 0, 8, 10, 3, 6, 12, 13, 4}

	var h Heap

	for _, k := range keys {
		h.Insert(E(k), k)
	}

	for h.Size() > 0 {
		m := h.FindMin()
		t.Log(m.item.(int))
		h.DeleteMin()
	}
}

type pqitem struct {
	value    interface{}
	priority int
}

type priorityQueue []*pqitem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*pqitem)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func benchmarkHeap(n int, b *testing.B) {

	x := make([]int, n)

	rand.Seed(0)

	for i := range x {
		x[i] = rand.Intn(n)
	}

	b.ResetTimer()

	var h priorityQueue

	for i := 0; i < b.N; i++ {
		for _, v := range x {
			h.Push(&pqitem{v, v})
		}

		for j := 0; j < n; j++ {
			h.Pop()
		}
	}
}

func benchmarkHollow(n int, b *testing.B) {

	x := make([]int, n)

	rand.Seed(0)

	for i := range x {
		x[i] = rand.Intn(n)
	}

	b.ResetTimer()

	var h Heap

	for i := 0; i < b.N; i++ {
		for _, v := range x {
			h.Insert(&Element{item: v}, v)
		}

		for j := 0; j < n; j++ {
			h.DeleteMin()
		}
	}
}

func BenchmarkHollow10(b *testing.B)   { benchmarkHollow(10, b) }
func BenchmarkHollow20(b *testing.B)   { benchmarkHollow(20, b) }
func BenchmarkHollow50(b *testing.B)   { benchmarkHollow(50, b) }
func BenchmarkHollow100(b *testing.B)  { benchmarkHollow(100, b) }
func BenchmarkHollow200(b *testing.B)  { benchmarkHollow(200, b) }
func BenchmarkHollow500(b *testing.B)  { benchmarkHollow(500, b) }
func BenchmarkHollow1000(b *testing.B) { benchmarkHollow(1000, b) }
func BenchmarkHollow2000(b *testing.B) { benchmarkHollow(2000, b) }
func BenchmarkHollow5000(b *testing.B) { benchmarkHollow(5000, b) }
func BenchmarkHollow1e6(b *testing.B)  { benchmarkHollow(1e6, b) }

func BenchmarkHeap10(b *testing.B)   { benchmarkHeap(10, b) }
func BenchmarkHeap20(b *testing.B)   { benchmarkHeap(20, b) }
func BenchmarkHeap50(b *testing.B)   { benchmarkHeap(50, b) }
func BenchmarkHeap100(b *testing.B)  { benchmarkHeap(100, b) }
func BenchmarkHeap200(b *testing.B)  { benchmarkHeap(200, b) }
func BenchmarkHeap500(b *testing.B)  { benchmarkHeap(500, b) }
func BenchmarkHeap1000(b *testing.B) { benchmarkHeap(1000, b) }
func BenchmarkHeap2000(b *testing.B) { benchmarkHeap(2000, b) }
func BenchmarkHeap5000(b *testing.B) { benchmarkHeap(5000, b) }
func BenchmarkHeap1e6(b *testing.B)  { benchmarkHeap(1e6, b) }
