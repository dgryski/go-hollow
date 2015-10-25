// Package hollow implements hollow heap
/*
   http://arxiv.org/pdf/1510.06535v1.pdf
*/
package hollow

type Node struct {
	child *Node
	next  *Node
	ep    *Node

	rank int

	item *Element
	key  int
}

type Element struct {
	item interface{}
	node *Node
}

func E(item interface{}) *Element {
	return &Element{item: item}
}

func makeNode(e *Element, k int) *Node {
	u := &Node{
		item: e,
		key:  k,
	}
	e.node = u
	return u
}

func Insert(e *Element, k int, h *Node) *Node {
	return Meld(makeNode(e, k), h)
}

func FindMin(h *Node) *Element {
	if h == nil {
		return nil
	}
	return h.item
}

func DecreaseKey(e *Element, k int, h *Node) *Node {
	u := e.node
	if u == h {
		u.key = k
		return h
	}
	v := makeNode(e, k)
	u.item = nil
	if u.rank > 2 {
		v.rank = u.rank - 2
	}
	v.child = u
	u.ep = v
	return link(v, h)
}

func DeleteMin(h *Node) *Node {
	return DeleteItem(h.item, h)
}

func DeleteItem(e *Element, h *Node) *Node {
	e.node.item = nil
	e.node = nil
	if h.item != nil {
		return h /* Non-minimum deletion */
	}

	A := make([]*Node, 64)
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

func Meld(g, h *Node) *Node {
	if h == nil {
		return g
	}

	if g == nil {
		return h
	}

	return link(g, h)
}

func link(v, w *Node) *Node {

	if v.key >= w.key {
		addChild(v, w)
		return w
	}

	addChild(w, v)
	return v
}

func addChild(v, w *Node) {
	v.next = w.child
	w.child = v
}
