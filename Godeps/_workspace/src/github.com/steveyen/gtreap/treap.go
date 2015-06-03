package gtreap

type Treap struct {
	compare Compare
	root    *node
}

// Compare returns an integer comparing the two items
// lexicographically. The result will be 0 if a==b, -1 if a < b, and
// +1 if a > b.
type Compare func(a, b interface{}) int

// Item can be anything.
type Item interface{}

type node struct {
	item     Item
	priority int
	left     *node
	right    *node
}

func NewTreap(c Compare) *Treap {
	return &Treap{compare: c, root: nil}
}

func (t *Treap) Min() Item {
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return n.item
		}
		n = n.left
	}
	return nil
}

func (t *Treap) Max() Item {
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return n.item
		}
		n = n.right
	}
	return nil
}

func (t *Treap) Get(target Item) Item {
	n := t.root
	for {
		if n == nil {
			break
		}
		c := t.compare(target, n.item)
		if c < 0 {
			n = n.left
		} else if c > 0 {
			n = n.right
		} else {
			return n.item
		}
	}
	return nil
}

func (t *Treap) Upsert(item Item, itemPriority int) *Treap {
	r := t.union(t.root, &node{item: item, priority: itemPriority})
	return &Treap{compare: t.compare, root: r}
}

func (t *Treap) union(this *node, that *node) *node {
	if this == nil {
		return that
	}
	if that == nil {
		return this
	}
	if this.priority > that.priority {
		left, middle, right := t.split(that, this.item)
		if middle == nil {
			return &node{
				item:     this.item,
				priority: this.priority,
				left:     t.union(this.left, left),
				right:    t.union(this.right, right),
			}
		}
		return &node{
			item:     middle.item,
			priority: middle.priority,
			left:     t.union(this.left, left),
			right:    t.union(this.right, right),
		}
	}
	// We don't use middle because the "that" has precendence.
	left, _, right := t.split(this, that.item)
	return &node{
		item:     that.item,
		priority: that.priority,
		left:     t.union(left, that.left),
		right:    t.union(right, that.right),
	}
}

// Splits a treap into two treaps based on a split item "s".
// The result tuple-3 means (left, X, right), where X is either...
// nil - meaning the item s was not in the original treap.
// non-null - returning the node that had item s.
// The tuple-3's left result has items < s,
// and the tuple-3's right result has items > s.
func (t *Treap) split(n *node, s Item) (*node, *node, *node) {
	if n == nil {
		return nil, nil, nil
	}
	c := t.compare(s, n.item)
	if c == 0 {
		return n.left, n, n.right
	}
	if c < 0 {
		left, middle, right := t.split(n.left, s)
		return left, middle, &node{
			item:     n.item,
			priority: n.priority,
			left:     right,
			right:    n.right,
		}
	}
	left, middle, right := t.split(n.right, s)
	return &node{
		item:     n.item,
		priority: n.priority,
		left:     n.left,
		right:    left,
	}, middle, right
}

func (t *Treap) Delete(target Item) *Treap {
	left, _, right := t.split(t.root, target)
	return &Treap{compare: t.compare, root: t.join(left, right)}
}

// All the items from this are < items from that.
func (t *Treap) join(this *node, that *node) *node {
	if this == nil {
		return that
	}
	if that == nil {
		return this
	}
	if this.priority > that.priority {
		return &node{
			item:     this.item,
			priority: this.priority,
			left:     this.left,
			right:    t.join(this.right, that),
		}
	}
	return &node{
		item:     that.item,
		priority: that.priority,
		left:     t.join(this, that.left),
		right:    that.right,
	}
}

type ItemVisitor func(i Item) bool

// Visit items greater-than-or-equal to the pivot.
func (t *Treap) VisitAscend(pivot Item, visitor ItemVisitor) {
	t.visitAscend(t.root, pivot, visitor)
}

func (t *Treap) visitAscend(n *node, pivot Item, visitor ItemVisitor) bool {
	if n == nil {
		return true
	}
	if t.compare(pivot, n.item) <= 0 {
		if !t.visitAscend(n.left, pivot, visitor) {
			return false
		}
		if !visitor(n.item) {
			return false
		}
	}
	return t.visitAscend(n.right, pivot, visitor)
}
