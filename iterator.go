package skiplist

// Iterator interface
// Implements an iterator to be used to
// navigate through the skiplist
type Iterator interface {
	Next() (ok bool)
	Prev() (ok bool)
	Val() []byte
	Key() int64
}

type iterable struct {
	curr *Node
	key  int64
	val  []byte
}

func (i *iterable) Next() bool {
	next := i.curr.next()
	if next == nil {
		return false
	}
	i.curr = next
	return true
}

func (i *iterable) Prev() bool {
	prev := i.curr.prev()
	if prev == nil {
		return false
	}
	i.curr = prev
	return true
}

func (i *iterable) Val() []byte {
	if i.curr == nil {
		return nil
	}
	return i.curr.val
}

func (i *iterable) Key() int64 {
	if i.curr == nil {
		return 0
	}
	return i.curr.key
}
