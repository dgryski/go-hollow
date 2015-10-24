package hollow

import "testing"

func TestHeap(t *testing.T) {

	keys := []int{14, 11, 5, 9, 0, 8, 10, 3, 6, 12, 13, 4}

	var h *Node

	for _, k := range keys {
		h = Insert(E(k), k, h)
	}

	for h != nil {
		t.Log(FindMin(h))
		h = DeleteMin(h)
	}
}
