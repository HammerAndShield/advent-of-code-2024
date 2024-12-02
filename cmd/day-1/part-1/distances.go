package main

import "github.com/emirpasic/gods/v2/trees/binaryheap"

type Distances struct {
	LeftD  *binaryheap.Heap[int]
	RightD *binaryheap.Heap[int]
}

func (h *Distances) MinDistSum() int {
	dL, _ := h.LeftD.Pop()
	dR, _ := h.RightD.Pop()

	return CalculateDistance(dL, dR)
}

func (h *Distances) DPush(l, r int) {
	h.LeftD.Push(l)
	h.RightD.Push(r)
}

func CalculateDistance(a, b int) int {
	larger := max(a, b)
	smaller := min(a, b)
	return larger - smaller
}
