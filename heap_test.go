package hollow

import (
	"container/heap"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestHeapSort(t *testing.T) {

	keys := []int{14, 11, 5, 9, 0, 8, 10, 3, 6, 12, 13, 4}

	var h Heap

	for _, k := range keys {
		h.Insert(E(k), k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		m := h.FindMin().item.(int)

		if m != v {
			t.Errorf("out-of-order element: got %v, want %v", m, v)
		}

		h.DeleteMin()
	}

	if h.Size() != 0 {
		t.Errorf("items remaining in heap")
	}
}

type pqitem struct {
	value    interface{}
	priority int
}

type priorityQueue []*pqitem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
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
			heap.Push(&h, &pqitem{v, v})
		}

		for j := 0; j < n; j++ {
			h[0].priority--
			heap.Fix(&h, 0)
		}

		for j := 0; j < n; j++ {
			h[0].priority--
			heap.Fix(&h, 0)
		}

		for j := 0; j < n; j++ {
			heap.Pop(&h)
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
			e := h.FindMin()
			h.DecreaseKey(e, 1)
		}

		for j := 0; j < n; j++ {
			h.DeleteMin()
		}
	}
}

func BenchmarkHollow(b *testing.B) { runBenchmarks(b, benchmarkHollow) }
func BenchmarkHeap(b *testing.B)   { runBenchmarks(b, benchmarkHeap) }

func runBenchmarks(b *testing.B, f func(n int, b *testing.B)) {
	steps := []int{10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 1e6}

	for _, n := range steps {
		b.Run(strconv.Itoa(n), func(b *testing.B) { f(n, b) })
	}
}
