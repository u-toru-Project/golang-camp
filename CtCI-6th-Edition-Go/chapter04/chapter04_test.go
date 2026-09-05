package chapter04

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"ctci/library"
)

func mustRoute(t *testing.T, fn func(map[string][]string, string, string) (bool, error), start, end string, expected bool) {
	t.Helper()
	got, err := fn(SampleGraph, start, end)
	if err != nil || got != expected {
		t.Fatalf("%s->%s: %t %v want %t", start, end, got, err, expected)
	}
}

func TestIsRoute(t *testing.T) {
	truePairs := [][2]string{{"A", "L"}, {"A", "B"}, {"H", "K"}, {"L", "D"}, {"P", "Q"}, {"Q", "P"}, {"A", "A"}}
	for _, pair := range truePairs {
		mustRoute(t, IsRoute, pair[0], pair[1], true)
		mustRoute(t, IsRouteBfs, pair[0], pair[1], true)
		mustRoute(t, IsRouteBidirectional, pair[0], pair[1], true)
	}
	falsePairs := [][2]string{{"Q", "G"}, {"R", "A"}, {"P", "B"}}
	for _, pair := range falsePairs {
		mustRoute(t, IsRoute, pair[0], pair[1], false)
		mustRoute(t, IsRouteBfs, pair[0], pair[1], false)
		mustRoute(t, IsRouteBidirectional, pair[0], pair[1], false)
	}
	if _, err := IsRouteBfs(SampleGraph, "A", "Z"); err == nil {
		t.Fatal("missing end")
	}
	if _, err := IsRoute(SampleGraph, "Z", "A"); err == nil {
		t.Fatal("missing start")
	}
	if _, err := IsRouteBidirectional(SampleGraph, "A", "Z"); err == nil {
		t.Fatal("missing end bidirectional")
	}
	graph := map[string][]string{"A": {"X"}, "B": {"X"}, "X": {}}
	if ok, _ := IsRoute(graph, "A", "B"); ok {
		t.Fatal("A should not reach B")
	}
	if ok, _ := IsRouteBfs(graph, "A", "B"); ok {
		t.Fatal("A should not reach B bfs")
	}
	if ok, _ := IsRouteBidirectional(graph, "A", "B"); ok {
		t.Fatal("false positive on common successor")
	}
	if ok, _ := IsRouteBidirectional(graph, "A", "X"); !ok {
		t.Fatal("A should reach X")
	}
}

func collectInOrder(node *library.TreeNode, values *[]int) {
	if node == nil {
		return
	}
	collectInOrder(node.Left, values)
	*values = append(*values, node.Data)
	collectInOrder(node.Right, values)
}

func TestCreateMinimalBST(t *testing.T) {
	if Create() != nil || Create([]int{}...) != nil {
		t.Fatal("empty")
	}
	root := Create(7)
	if root.Data != 7 || root.Left != nil || root.Right != nil {
		t.Fatal(root)
	}
	three := Create(1, 2, 3)
	if three.Data != 2 || three.Left.Data != 1 || three.Right.Data != 3 || three.Left.Parent != three || three.Right.Parent != three {
		t.Fatal("three node")
	}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := Create(values...)
	var inorder []int
	collectInOrder(tree, &inorder)
	if fmt.Sprint(inorder) != fmt.Sprint(values) || tree.Data != 5 {
		t.Fatal(inorder, tree.Data)
	}
	seven := Create(1, 2, 3, 4, 5, 6, 7)
	if seven.Data != 4 || seven.Height() != 3 || seven.Left.Data != 2 || seven.Right.Data != 6 {
		t.Fatal(seven.Data, seven.Height())
	}
	if seven.Left.Find(1).Parent != seven.Left || seven.Right.Find(7).Parent != seven.Right.Find(6) {
		t.Fatal("parent links")
	}
}

func createTreeFromArray(array []int) *library.TreeNode {
	if len(array) == 0 {
		return nil
	}
	root := library.NewTreeNode(array[0])
	queue := []*library.TreeNode{root}
	i := 1
	for i < len(array) {
		treeNode := queue[0]
		if treeNode.Left == nil {
			treeNode.Left = library.NewTreeNode(array[i])
			i++
			queue = append(queue, treeNode.Left)
		} else if treeNode.Right == nil {
			treeNode.Right = library.NewTreeNode(array[i])
			i++
			queue = append(queue, treeNode.Right)
		} else {
			queue = queue[1:]
		}
	}
	return root
}

func formatLevels(levels [][]*library.TreeNode) string {
	parts := make([]string, len(levels))
	for i, level := range levels {
		values := make([]string, len(level))
		for j, node := range level {
			values[j] = fmt.Sprintf("%d", node.Data)
		}
		parts[i] = strings.Join(values, ",")
	}
	return strings.Join(parts, "|")
}

func TestListOfDepths(t *testing.T) {
	if len(ListOfDepths(nil)) != 0 {
		t.Fatal("nil")
	}
	root := library.NewTreeNode(9)
	levels := ListOfDepths(root)
	if len(levels) != 1 || len(levels[0]) != 1 || levels[0][0] != root {
		t.Fatal("single")
	}
	if formatLevels(ListOfDepths(Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))) != "5|2,8|1,3,6,9|4,7,10" {
		t.Fatal(formatLevels(ListOfDepths(Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))))
	}
	if formatLevels(ListOfDepths(createTreeFromArray([]int{1, 2, 3, 4, 5}))) != "1|2,3|4,5" {
		t.Fatal(formatLevels(ListOfDepths(createTreeFromArray([]int{1, 2, 3, 4, 5}))))
	}
}

func assertBalanced(t *testing.T, root *library.TreeNode, expected bool) {
	t.Helper()
	if IsBalanced(root) != expected || IsBalancedBook(root) != expected {
		t.Fatalf("balanced=%t book=%t want %t", IsBalanced(root), IsBalancedBook(root), expected)
	}
}

func TestCheckBalanced(t *testing.T) {
	assertBalanced(t, nil, true)
	assertBalanced(t, library.NewTreeNode(1), true)
	assertBalanced(t, Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), true)
	assertBalanced(t, library.CreateMinimalBST([]int{1, 2, 3, 4, 5, 6, 7}), true)
	assertBalanced(t, createTreeFromArray([]int{1, 2, 3, 4, 5, 6, 7}), true)
	assertBalanced(t, createTreeFromArray([]int{1, 2, 3, 4}), true)
	skewed := library.NewTreeNode(2)
	skewed.SetLeftChild(library.NewTreeNode(1))
	skewed.Left.SetLeftChild(library.NewTreeNode(3))
	assertBalanced(t, skewed, false)
	root := library.NewTreeNode(1)
	root.SetLeftChild(library.NewTreeNode(2))
	root.SetRightChild(library.NewTreeNode(3))
	root.Left.SetLeftChild(library.NewTreeNode(4))
	root.Right.SetRightChild(library.NewTreeNode(5))
	root.Left.Left.SetLeftChild(library.NewTreeNode(6))
	assertBalanced(t, root, false)
}

func TestValidateBST(t *testing.T) {
	if !IsValidBST(nil) || !IsValidBST(library.NewTreeNode(1)) {
		t.Fatal("null/single")
	}
	if !IsValidBST(Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) || !IsValidBST(library.CreateMinimalBST([]int{1, 2, 3})) {
		t.Fatal("minimal")
	}
	invalid := library.NewTreeNode(2)
	invalid.SetLeftChild(library.NewTreeNode(1))
	invalid.Left.SetLeftChild(library.NewTreeNode(3))
	if IsValidBST(invalid) {
		t.Fatal("left too large")
	}
	ancestor := library.NewTreeNode(20)
	ancestor.SetLeftChild(library.NewTreeNode(10))
	ancestor.SetRightChild(library.NewTreeNode(30))
	ancestor.Right.SetLeftChild(library.NewTreeNode(15))
	if IsValidBST(ancestor) {
		t.Fatal("ancestor bound")
	}
	leftDup := library.NewTreeNode(2)
	leftDup.SetLeftChild(library.NewTreeNode(2))
	if !IsValidBST(leftDup) {
		t.Fatal("equal on left")
	}
	rightDup := library.NewTreeNode(2)
	rightDup.SetRightChild(library.NewTreeNode(2))
	if IsValidBST(rightDup) {
		t.Fatal("equal on right")
	}
}

func TestSuccessor(t *testing.T) {
	if Successor(nil) != nil {
		t.Fatal("nil")
	}
	small := Create(1, 2, 3)
	if Successor(small.Find(3)) != nil || Successor(library.NewTreeNode(1)) != nil {
		t.Fatal("max/no parent")
	}
	root := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cases := [][2]int{{3, 4}, {6, 7}, {8, 9}, {2, 3}, {1, 2}, {4, 5}, {7, 8}}
	for _, tc := range cases {
		if Successor(root.Find(tc[0])).Data != tc[1] {
			t.Fatalf("succ(%d)", tc[0])
		}
	}
	node := root.Find(1)
	for expected := 2; expected <= 10; expected++ {
		node = Successor(node)
		if node.Data != expected {
			t.Fatal(expected, node.Data)
		}
	}
	if Successor(node) != nil {
		t.Fatal("after last")
	}
}

func TestBuildOrder(t *testing.T) {
	projects := []string{"a", "b", "c", "d", "e", "f", "g"}
	dependencies := [][2]string{{"d", "g"}, {"a", "e"}, {"b", "e"}, {"c", "a"}, {"f", "a"}, {"b", "a"}, {"f", "c"}, {"f", "b"}}
	order, err := DetermineBuildOrder(projects, dependencies)
	if err != nil || len(order) != len(projects) {
		t.Fatal(order, err)
	}
	index := map[string]int{}
	for i, project := range order {
		index[project] = i
	}
	for _, dep := range dependencies {
		if index[dep[0]] >= index[dep[1]] {
			t.Fatalf("%s should precede %s in %v", dep[0], dep[1], order)
		}
	}
	if _, err := DetermineBuildOrder([]string{"a", "b"}, [][2]string{{"a", "b"}, {"b", "a"}}); err == nil {
		t.Fatal("cycle")
	}
	empty, err := DetermineBuildOrder([]string{}, [][2]string{})
	if err != nil || len(empty) != 0 {
		t.Fatal(empty, err)
	}
	noDeps, err := DetermineBuildOrder([]string{"a", "b"}, [][2]string{})
	if err != nil || len(noDeps) != 2 {
		t.Fatal(noDeps, err)
	}
	seen := map[string]bool{"a": false, "b": false}
	for _, project := range noDeps {
		seen[project] = true
	}
	if !seen["a"] || !seen["b"] {
		t.Fatal(noDeps)
	}
	if _, err := DetermineBuildOrder([]string{"a"}, [][2]string{{"a", "z"}}); err == nil {
		t.Fatal("unknown project")
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := Create(1, 2, 3, 4, 5)
	if LowestCommonAncestor(nil, root, root) != nil || LowestCommonAncestor(root, nil, root) != nil || LowestCommonAncestor(root, root, nil) != nil {
		t.Fatal("nulls")
	}
	if LowestCommonAncestor(root, root.Find(1), library.NewTreeNode(99)) != nil {
		t.Fatal("missing")
	}
	full := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cases := [][3]int{{1, 3, 2}, {4, 7, 5}, {7, 9, 8}, {1, 10, 5}, {2, 4, 2}, {8, 10, 8}, {5, 7, 5}}
	for _, tc := range cases {
		lca := LowestCommonAncestor(full, full.Find(tc[0]), full.Find(tc[1]))
		if lca.Data != tc[2] {
			t.Fatalf("lca(%d,%d)=%d want %d", tc[0], tc[1], lca.Data, tc[2])
		}
	}
	tiny := Create(1, 2, 3)
	two := tiny.Find(2)
	if LowestCommonAncestor(tiny, two, two) != two {
		t.Fatal("same node")
	}
}

func formatSequences(sequences [][]int) string {
	parts := make([]string, len(sequences))
	for i, sequence := range sequences {
		values := make([]string, len(sequence))
		for j, value := range sequence {
			values[j] = fmt.Sprintf("%d", value)
		}
		parts[i] = strings.Join(values, ",")
	}
	sort.Strings(parts)
	return strings.Join(parts, ";")
}

func TestAllSequences(t *testing.T) {
	nullSeq := AllSequences(nil)
	if len(nullSeq) != 1 || len(nullSeq[0]) != 0 {
		t.Fatal(nullSeq)
	}
	if formatSequences(AllSequences(library.NewTreeNode(1))) != "1" {
		t.Fatal(formatSequences(AllSequences(library.NewTreeNode(1))))
	}
	if formatSequences(AllSequences(Create(1, 2, 3))) != "2,1,3;2,3,1" {
		t.Fatal(formatSequences(AllSequences(Create(1, 2, 3))))
	}
	skewed := library.NewTreeNode(1)
	skewed.SetRightChild(library.NewTreeNode(2))
	skewed.Right.SetRightChild(library.NewTreeNode(3))
	if formatSequences(AllSequences(skewed)) != "1,2,3" {
		t.Fatal(formatSequences(AllSequences(skewed)))
	}
}

func TestContainsTree(t *testing.T) {
	if !ContainsTree(nil, nil) || !ContainsTree(library.NewTreeNode(1), nil) {
		t.Fatal("null subtree")
	}
	if ContainsTree(nil, library.NewTreeNode(1)) {
		t.Fatal("null tree")
	}
	root := Create(1, 2, 3)
	if !ContainsTree(root, Create(1, 2, 3)) || !ContainsTree(library.NewTreeNode(4), library.NewTreeNode(4)) || ContainsTree(library.NewTreeNode(4), library.NewTreeNode(5)) {
		t.Fatal("identical")
	}
	tree := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if !ContainsTree(tree, Create(6, 7, 8, 9, 10)) || !ContainsTree(tree, tree.Find(7)) {
		t.Fatal("subtree")
	}
	if ContainsTree(tree, Create(7, 8, 9, 10)) || ContainsTree(tree, Create(1, 2, 3)) {
		t.Fatal("structure differs")
	}
}

func buildRandomSample() *RandomBinarySearchTree {
	tree := &RandomBinarySearchTree{}
	for _, key := range []int{20, 9, 25, 5, 12, 11, 14} {
		tree.Insert(key)
	}
	return tree
}

func TestRandomNode(t *testing.T) {
	tree := &RandomBinarySearchTree{}
	for _, key := range []int{20, 9, 25, 5, 12} {
		tree.Insert(key)
	}
	node, err := tree.GetNode(12)
	if err != nil || node.Data != 12 || tree.Root.Size != 5 {
		t.Fatal(node, err, tree.Root.Size)
	}
	sample := buildRandomSample()
	if err := sample.Delete(12); err != nil || sample.Root.Size != 6 {
		t.Fatal(err, sample.Root.Size)
	}
	if _, err := sample.GetNode(12); err == nil {
		t.Fatal("deleted key still present")
	}
	remaining := map[int]struct{}{20: {}, 9: {}, 25: {}, 5: {}, 11: {}, 14: {}}
	seen := map[int]struct{}{}
	for i := 0; i < sample.Root.Size; i++ {
		seen[GetIthNode(sample.Root, i).Data] = struct{}{}
	}
	if len(seen) != len(remaining) {
		t.Fatal(seen)
	}
	for key := range remaining {
		if _, ok := seen[key]; !ok {
			t.Fatal(key)
		}
	}
	counts := map[int]int{}
	trials := 6000
	for range trials {
		got, err := sample.GetRandomNode()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := remaining[got.Data]; !ok {
			t.Fatal(got.Data)
		}
		counts[got.Data]++
	}
	if len(counts) != len(remaining) {
		t.Fatal(counts)
	}
	expected := trials / len(remaining)
	for _, count := range counts {
		if absInt(count-expected) >= expected*25/100 {
			t.Fatalf("distribution %v", counts)
		}
	}
	if _, err := (&RandomBinarySearchTree{}).GetRandomNode(); err == nil {
		t.Fatal("empty")
	}
	missing := &RandomBinarySearchTree{}
	missing.Insert(1)
	if err := missing.Delete(99); err == nil || missing.Root.Size != 1 {
		t.Fatal(err, missing.Root.Size)
	}
}

func TestPathsWithSum(t *testing.T) {
	cases := []struct {
		root     *library.TreeNode
		target   int
		expected int
	}{
		{CreateSampleTree(), 8, 5},
		{CreateSampleTree(), 6, 2},
		{nil, 0, 0},
		{nil, 8, 0},
		{library.NewTreeNode(8), 8, 1},
		{library.NewTreeNode(8), 0, 0},
		{library.NewTreeNode(0), 0, 1},
	}
	neg := library.NewTreeNode(1)
	neg.SetLeftChild(library.NewTreeNode(-1))
	cases = append(cases, struct {
		root     *library.TreeNode
		target   int
		expected int
	}{neg, 0, 1}, struct {
		root     *library.TreeNode
		target   int
		expected int
	}{neg, 1, 1})
	for _, tc := range cases {
		if CountPathsWithSum(tc.root, tc.target) != tc.expected || CountPathsWithSumOptimized(tc.root, tc.target) != tc.expected {
			t.Fatalf("target %d got %d / %d want %d", tc.target, CountPathsWithSum(tc.root, tc.target), CountPathsWithSumOptimized(tc.root, tc.target), tc.expected)
		}
	}
}

func TestReplaceNode(t *testing.T) {
	if Replace(nil, 11) != nil {
		t.Fatal("nil")
	}
	root := Create(1, 2, 3)
	newRoot := Replace(root, 9)
	if root.Data != 2 || newRoot.Data != 9 || newRoot.Left != root.Left || newRoot.Right != root.Right || newRoot == root {
		t.Fatal("root replace")
	}
	full := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	originalEight := full.Find(8)
	replacedRoot := Replace(full.Find(6), 11)
	replaced := replacedRoot.Right.Left
	if full.Find(6).Data != 6 || replaced.Data != 11 || replaced.Parent.Data != 8 || replacedRoot == full || replacedRoot.Right == originalEight {
		t.Fatal("path copy")
	}
	two := full.Find(2)
	seven := full.Find(7)
	nine := full.Find(9)
	shared := Replace(full.Find(6), 11)
	if shared.Left != two || shared.Right.Right != nine || shared.Right.Left.Right != seven {
		t.Fatal("shared subtrees")
	}
}
