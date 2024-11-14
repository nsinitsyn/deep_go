package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	left  *node
	right *node
	key   int
	value int
}

func find(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}
	if n.key == key {
		return n, true
	}

	if key < n.key {
		return find(n.left, key)
	}
	return find(n.right, key)
}

func traversalDFS(n *node, action func(int, int)) {
	if n == nil {
		return
	}

	traversalDFS(n.left, action)
	action(n.key, n.value)
	traversalDFS(n.right, action)
}

func insert(n *node, key int, value int) {
	if key < n.key {
		if n.left == nil {
			n.left = &node{key: key, value: value}
		} else {
			insert(n.left, key, value)
		}
	} else if key > n.key {
		if n.right == nil {
			n.right = &node{key: key, value: value}
		} else {
			insert(n.right, key, value)
		}
	}
}

func findMin(n *node) *node {
	min := n
	for min.left != nil {
		min = min.left
	}
	return min
}

func remove(r *node, pointer **node) {
	if r.left == nil && r.right == nil {
		*pointer = nil
		return
	}

	if r.left != nil && r.right == nil {
		*pointer = r.left
		return
	}

	if r.left == nil && r.right != nil {
		*pointer = r.right
		return
	}

	min := findMin(r.right)
	*pointer = min
}

func findForRemove(n *node, key int) (r *node, pointer **node, found bool) {
	cur := n
	var p **node = nil
	for cur != nil {
		if key == cur.key {
			return cur, p, true
		}
		if key < cur.key {
			p = &cur.left
			cur = cur.left
			continue
		}
		if key > cur.key {
			p = &cur.right
			cur = cur.right
			continue
		}
	}
	return cur, p, false
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{nil, 0}
}

func (m *OrderedMap) Insert(key, value int) {
	m.size++
	if m.root == nil {
		m.root = &node{key: key, value: value}
		return
	}
	insert(m.root, key, value)
}

func (m *OrderedMap) Erase(key int) {
	if m.root == nil {
		return
	}
	r, pointer, ok := findForRemove(m.root, key)
	if ok {
		remove(r, pointer)
		m.size--
	}
}

func (m *OrderedMap) Contains(key int) bool {
	_, found := find(m.root, key)
	return found
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	traversalDFS(m.root, action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
