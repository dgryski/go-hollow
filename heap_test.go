package hollow

import "testing"

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
