package chapter02

import (
	"reflect"
	"testing"

	"ctci/library"
)

func TestDeleteDups(t *testing.T) {
	cases := []struct {
		input, expected []int
	}{
		{[]int{0, 1, 0, 1, 0, 1}, []int{0, 1}},
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{7}, []int{7}},
	}
	for _, tc := range cases {
		withSet := library.CreateLinkedListFromArray(tc.input)
		withRunner := library.CreateLinkedListFromArray(tc.input)
		DeleteDupsWithSet(withSet)
		DeleteDupsWithRunner(withRunner)
		if !reflect.DeepEqual(library.LinkedListToArray(withSet), tc.expected) {
			t.Fatalf("set %v -> %v, want %v", tc.input, library.LinkedListToArray(withSet), tc.expected)
		}
		if !reflect.DeepEqual(library.LinkedListToArray(withRunner), tc.expected) {
			t.Fatalf("runner %v -> %v, want %v", tc.input, library.LinkedListToArray(withRunner), tc.expected)
		}
	}
	DeleteDupsWithSet(nil)
	DeleteDupsWithRunner(nil)
}

func TestKthToLast(t *testing.T) {
	head := library.CreateLinkedListFromArray([]int{1, 2, 3, 4, 5})
	cases := []struct {
		k, expected int
	}{{1, 5}, {2, 4}, {5, 1}}
	for _, tc := range cases {
		if KthToLast(head, tc.k).Data != tc.expected || KthToLastRecursive(head, tc.k).Data != tc.expected {
			t.Fatalf("k=%d", tc.k)
		}
	}
	short := library.CreateLinkedListFromArray([]int{1, 2, 3})
	if KthToLast(short, 0) != nil || KthToLast(short, 4) != nil || KthToLastRecursive(short, 4) != nil || KthToLast(nil, 1) != nil {
		t.Fatal("expected nil for out of range")
	}
}

func TestDeleteNode(t *testing.T) {
	head := library.CreateLinkedListFromArray([]int{1, 2, 3, 4, 5})
	if !DeleteNode(head.Next.Next) || !reflect.DeepEqual(library.LinkedListToArray(head), []int{1, 2, 4, 5}) {
		t.Fatal(library.LinkedListToArray(head))
	}
	pair := library.CreateLinkedListFromArray([]int{1, 2})
	if DeleteNode(nil) || DeleteNode(pair.Next) {
		t.Fatal("expected false")
	}
	if !reflect.DeepEqual(library.LinkedListToArray(pair), []int{1, 2}) {
		t.Fatal(library.LinkedListToArray(pair))
	}
}

func assertPartitioned(t *testing.T, head *library.LinkedListNode, pivot, expectedLength int) {
	t.Helper()
	seenPivotOrAbove := false
	length := 0
	for node := head; node != nil; node = node.Next {
		if node.Data < pivot {
			if seenPivotOrAbove {
				t.Fatalf("found %d after a value >= %d", node.Data, pivot)
			}
		} else {
			seenPivotOrAbove = true
		}
		length++
	}
	if length != expectedLength {
		t.Fatalf("length %d want %d", length, expectedLength)
	}
}

func TestPartition(t *testing.T) {
	cases := [][]int{
		{1, 3, 7, 5, 2, 9, 4},
		{3, 5, 8, 5, 10, 2, 1},
		{1, 2, 3},
		{8, 9, 10},
		{5, 5, 5},
		{4},
	}
	for _, values := range cases {
		assertPartitioned(t, Partition(library.CreateLinkedListFromArray(values), 5), 5, len(values))
		assertPartitioned(t, PartitionByPrepending(library.CreateLinkedListFromArray(values), 5), 5, len(values))
	}
	if Partition(nil, 5) != nil || PartitionByPrepending(nil, 5) != nil {
		t.Fatal("nil partition")
	}
}

func TestSumLists(t *testing.T) {
	reverse := AddListsReverse(
		library.CreateLinkedListFromArray([]int{7, 1, 6}),
		library.CreateLinkedListFromArray([]int{5, 9, 2}),
	)
	if !reflect.DeepEqual(library.LinkedListToArray(reverse), []int{2, 1, 9}) {
		t.Fatal(library.LinkedListToArray(reverse))
	}
	forward := AddListsForward(
		library.CreateLinkedListFromArray([]int{6, 1, 7}),
		library.CreateLinkedListFromArray([]int{2, 9, 5}),
	)
	if !reflect.DeepEqual(library.LinkedListToArray(forward), []int{9, 1, 2}) {
		t.Fatal(library.LinkedListToArray(forward))
	}
	padded := AddListsForward(
		library.CreateLinkedListFromArray([]int{3, 1}),
		library.CreateLinkedListFromArray([]int{5, 9, 1}),
	)
	if !reflect.DeepEqual(library.LinkedListToArray(padded), []int{6, 2, 2}) {
		t.Fatal(library.LinkedListToArray(padded))
	}
	carry := AddListsReverse(
		library.CreateLinkedListFromArray([]int{9, 9, 9}),
		library.CreateLinkedListFromArray([]int{1}),
	)
	if !reflect.DeepEqual(library.LinkedListToArray(carry), []int{0, 0, 0, 1}) {
		t.Fatal(library.LinkedListToArray(carry))
	}
}

func TestPalindrome(t *testing.T) {
	trueCases := [][]int{{0, 1, 2, 1, 0}, {1, 2, 2, 1}, {7}, {}, {1, 1}}
	for _, values := range trueCases {
		var head *library.LinkedListNode
		if len(values) > 0 {
			head = library.CreateLinkedListFromArray(values)
		}
		if !IsPalindromeWithStack(head) || !IsPalindromeRecursive(head) {
			t.Fatalf("expected palindrome %v", values)
		}
	}
	falseCases := [][]int{{0, 1, 2, 3, 4}, {1, 2, 3, 2}, {1, 2}}
	for _, values := range falseCases {
		head := library.CreateLinkedListFromArray(values)
		if IsPalindromeWithStack(head) || IsPalindromeRecursive(head) {
			t.Fatalf("expected non-palindrome %v", values)
		}
	}
}

func TestIntersection(t *testing.T) {
	first := library.CreateLinkedListFromArray([]int{-1, -2, 0, 1, 2, 3, 4})
	second := library.CreateLinkedListFromArray([]int{12, 14, 15})
	shared := first.Next.Next.Next
	second.Next.Next.Next = shared
	if FindIntersection(first, second) != shared {
		t.Fatal("expected shared node")
	}
	a := library.CreateLinkedListFromArray([]int{1, 2, 3})
	b := library.CreateLinkedListFromArray([]int{1, 2, 3})
	if FindIntersection(a, b) != nil || FindIntersection(a, nil) != nil {
		t.Fatal("expected no intersection")
	}
}

func TestFindLoopStart(t *testing.T) {
	nodes := make([]*library.LinkedListNode, 10)
	for i := range nodes {
		var prev *library.LinkedListNode
		if i > 0 {
			prev = nodes[i-1]
		}
		nodes[i] = library.NewLinkedListNode(i+1, nil, prev)
	}
	nodes[len(nodes)-1].Next = nodes[6]
	if FindLoopStart(nodes[0]) != nodes[6] {
		t.Fatal("expected loop start at 7")
	}
	head := library.CreateLinkedListFromArray([]int{1, 2, 3, 4})
	if FindLoopStart(head) != nil || FindLoopStart(nil) != nil {
		t.Fatal("expected no cycle")
	}
}
