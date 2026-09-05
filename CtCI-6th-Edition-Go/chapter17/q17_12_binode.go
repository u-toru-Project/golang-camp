package chapter17

import "fmt"

type BiNode struct {
	Value int
	Node1 *BiNode
	Node2 *BiNode
}

func FlattenBstToDlist(root *BiNode) *BiNode {
	if root == nil {
		return nil
	}
	head, tail := flattenHelper(root)
	head.Node1 = tail
	tail.Node2 = head
	return head
}

func flattenHelper(root *BiNode) (*BiNode, *BiNode) {
	leftHead := root
	if root.Node1 != nil {
		subHead, subTail := flattenHelper(root.Node1)
		leftHead = subHead
		subTail.Node2 = root
		root.Node1 = subTail
	}
	rightTail := root
	if root.Node2 != nil {
		subHead, subTail := flattenHelper(root.Node2)
		root.Node2 = subHead
		subHead.Node1 = root
		rightTail = subTail
	}
	return leftHead, rightTail
}

func MakeSampleTree() *BiNode {
	return &BiNode{
		Value: 20,
		Node1: &BiNode{Value: 10, Node1: &BiNode{Value: 5}},
		Node2: &BiNode{
			Value: 25,
			Node1: &BiNode{Value: 22, Node1: &BiNode{Value: 21}, Node2: &BiNode{Value: 23}},
		},
	}
}

func RunQ1712() {
	head := FlattenBstToDlist(MakeSampleTree())
	values := make([]int, 0)
	node := head
	for node != nil {
		values = append(values, node.Value)
		node = node.Node2
		if node == head {
			break
		}
	}
	fmt.Println(values)
}
