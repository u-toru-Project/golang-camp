package library

import "testing"

func TestLinkedListFromArray(t *testing.T) {
	head := CreateLinkedListFromArray([]int{1, 2, 3})
	if head.PrintForward() != "1->2->3" {
		t.Fatalf("got %s", head.PrintForward())
	}
	if got := LinkedListToArray(head); len(got) != 3 || got[0] != 1 || got[2] != 3 {
		t.Fatalf("to array %v", got)
	}
	clone := head.Clone()
	if clone.PrintForward() != "1->2->3" || clone == head {
		t.Fatal("clone")
	}
	if CreateLinkedListFromArray(nil) != nil {
		t.Fatal("empty list")
	}
}

func TestTreeNode(t *testing.T) {
	root := CreateMinimalBST([]int{1, 2, 3, 4, 5})
	if root == nil || !root.IsBST() || root.Find(3) == nil || root.Find(99) != nil {
		t.Fatal("minimal bst")
	}
	if root.Height() < 2 {
		t.Fatal("height")
	}
	node := NewTreeNode(10)
	node.InsertInOrder(5)
	node.InsertInOrder(15)
	if !node.IsBST() || node.Find(5) == nil || node.Size != 3 {
		t.Fatal("insert in order")
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie([]string{"cat", "car", "dog"})
	if !trie.Contains("ca", false) || !trie.Contains("cat", true) || trie.Contains("ca", true) {
		t.Fatal("contains")
	}
	if trie.Contains("bird", false) {
		t.Fatal("missing word")
	}
}
