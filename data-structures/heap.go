package main

import "golang.org/x/exp/constraints"

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

type Heap[T constraints.Ordered] struct {
	nodes []T
}

func (h *Heap[T]) Push(el T) {
	h.nodes = append(h.nodes, el)
	i := len(h.nodes) - 1
	for ; h.nodes[i] > h.nodes[parent(i)]; i = parent(i) {
		h.swap(i, parent(i))
	}
}

func (h *Heap[T]) Pop() (el T) {
	el = h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.rearrange(0)
	return
}

func (h *Heap[T]) rearrange(i int) {
	max := i
	left, right, size := leftChild(i), rightChild(i), len(h.nodes)

	if left < size && h.nodes[left] > h.nodes[max] {
		max = left
	}

	if right < size && h.nodes[right] > h.nodes[max] {
		max = right
	}

	if max != i {
		h.swap(i, max)
		h.rearrange(max)
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *Heap[T]) Size() int {
	return len(h.nodes)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}
