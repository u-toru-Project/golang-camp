package chapter10

import "fmt"

type RankNode struct {
	LeftSize int
	Left     *RankNode
	Right    *RankNode
	Data     int
}

func NewRankNode(data int) *RankNode { return &RankNode{Data: data} }

func (n *RankNode) Insert(value int) {
	if value <= n.Data {
		if n.Left != nil {
			n.Left.Insert(value)
		} else {
			n.Left = NewRankNode(value)
		}
		n.LeftSize++
		return
	}
	if n.Right != nil {
		n.Right.Insert(value)
	} else {
		n.Right = NewRankNode(value)
	}
}

func (n *RankNode) GetRank(value int) int {
	if value == n.Data {
		return n.LeftSize
	}
	if value < n.Data {
		if n.Left == nil {
			return -1
		}
		return n.Left.GetRank(value)
	}
	if n.Right == nil {
		return -1
	}
	rightRank := n.Right.GetRank(value)
	if rightRank == -1 {
		return -1
	}
	return n.LeftSize + 1 + rightRank
}

type RankTracker struct {
	root *RankNode
}

func NewRankTracker() *RankTracker { return &RankTracker{} }

func (t *RankTracker) Track(number int) {
	if t.root == nil {
		t.root = NewRankNode(number)
		return
	}
	t.root.Insert(number)
}

func (t *RankTracker) GetRankOfNumber(number int) int {
	if t.root == nil {
		return -1
	}
	return t.root.GetRank(number)
}

func RunQ1010() {
	stream := []int{5, 1, 4, 4, 5, 9, 7, 13, 3}
	tracker := NewRankTracker()
	for _, number := range stream {
		tracker.Track(number)
	}
	fmt.Printf("Stream: %v\n", stream)
	for _, number := range []int{1, 3, 4, 5, 7, 9, 13, 2} {
		fmt.Printf("rank(%d) = %d\n", number, tracker.GetRankOfNumber(number))
	}
}
