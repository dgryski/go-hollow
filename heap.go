// Package hollow implements hollow heap
/*
   http://arxiv.org/pdf/1510.06535v1.pdf
*/
package hollow

type Heap struct {
	root *node
	size int
}

func (h *Heap) Insert(e *Element, k int) {
	h.size++
	h.root = insert(e, k, h.root)
}

func (h *Heap) FindMin() *Element {
	return findMin(h.root)
}

func (h *Heap) DecreaseKey(e *Element, k int) {
	h.root = decreaseKey(e, k, h.root)
}

func (h *Heap) DeleteItem(e *Element) {
	h.size--
	h.root = deleteItem(e, h.root)
}

func (h *Heap) Meld(g *Heap) {
	h.size += g.size
	h.root = meld(h.root, g.root)
}

func (h *Heap) DeleteMin() {
	h.size--
	h.root = deleteMin(h.root)
}

func (h *Heap) Size() int {
	return h.size
}

type node struct {
	child *node
	next  *node
	ep    *node

	rank int

	item *Element
	key  int
}

type Element struct {
	item interface{}
	node *node
}

func E(item interface{}) *Element {
	return &Element{item: item}
}

func makenode(e *Element, k int) *node {
	u := &node{
		item: e,
		key:  k,
	}
	e.node = u
	return u
}

func insert(e *Element, k int, h *node) *node {
	return meld(makenode(e, k), h)
}

func findMin(h *node) *Element {
	if h == nil {
		return nil
	}
	return h.item
}

func decreaseKey(e *Element, k int, h *node) *node {
	u := e.node
	if u == h {
		u.key = k
		return h
	}
	v := makenode(e, k)
	u.item = nil
	if u.rank > 2 {
		v.rank = u.rank - 2
	}
	v.child = u
	u.ep = v
	return link(v, h)
}

func deleteMin(h *node) *node {
	return deleteItem(h.item, h)
}

func deleteItem(e *Element, h *node) *node {
	e.node.item = nil
	e.node = nil
	if h.item != nil {
		return h /* Non-minimum deletion */
	}

	A := make([]*node, 64)
	maxRank := 0
	h.next = nil
	for h != nil { /* While L not empty */
		w := h.child
		v := h
		h = h.next
		for w != nil {
			u := w
			w = w.next
			if u.item == nil {
				if u.ep == nil {
					u.next = h
					h = u
				} else {
					if u.ep == v {
						w = nil
					} else {
						u.next = nil
					}
					u.ep = nil
				}
			} else {
				for A[u.rank] != nil {
					u = link(u, A[u.rank])
					A[u.rank] = nil
					u.rank = u.rank + 1
				}
				A[u.rank] = u
				if u.rank > maxRank {
					maxRank = u.rank
				}
			}
		}
	}
	for i := 0; i <= maxRank; i++ {
		if A[i] != nil {
			if h == nil {
				h = A[i]
			} else {
				h = link(h, A[i])
			}
			A[i] = nil
		}
	}
	return h
}

func meld(g, h *node) *node {
	if h == nil {
		return g
	}

	if g == nil {
		return h
	}

	return link(g, h)
}

func link(v, w *node) *node {

	if v.key >= w.key {
		addChild(v, w)
		return w
	}

	addChild(w, v)
	return v
}

func addChild(v, w *node) {
	v.next = w.child
	w.child = v
}
