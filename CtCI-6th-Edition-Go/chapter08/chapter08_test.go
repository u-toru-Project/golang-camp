package chapter08

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func TestTripleStepAndRobot(t *testing.T) {
	cases := [][2]int{{0, 1}, {1, 1}, {2, 2}, {3, 4}, {4, 7}, {5, 13}, {6, 24}}
	for _, c := range cases {
		if TripleHop(c[0]) != c[1] || TripleHopDp(c[0]) != c[1] {
			t.Fatalf("%d", c[0])
		}
	}
	if TripleHop(-3) != 0 || TripleHopDp(-3) != 0 {
		t.Fatal("neg")
	}
	assertPath(t, GetPath([][]bool{{true, true}, {true, true}}), 1, 1)
	assertPath(t, GetPathMemoized([][]bool{{true, true}, {true, true}}), 1, 1)
	assertPath(t, GetPath([][]bool{{true, true}, {false, true}}), 1, 1)
	assertPath(t, GetPathMemoized([][]bool{{true, true}, {false, true}}), 1, 1)
	if GetPath([][]bool{{true, false}, {false, true}}) != nil || GetPathMemoized([][]bool{{true, false}, {false, true}}) != nil {
		t.Fatal("blocked")
	}
	if GetPath(nil) != nil || GetPath([][]bool{{}}) != nil || GetPathMemoized(nil) != nil || GetPathMemoized([][]bool{{}}) != nil {
		t.Fatal("empty")
	}
}

func assertPath(t *testing.T, path []GridPoint, endRow, endCol int) {
	t.Helper()
	if path == nil || path[0].Row != 0 || path[0].Col != 0 {
		t.Fatal("start")
	}
	last := path[len(path)-1]
	if last.Row != endRow || last.Col != endCol {
		t.Fatal("end")
	}
}

func TestMagicPowerMultiplyHanoi(t *testing.T) {
	if MagicIndex([]int{-14, -12, 0, 1, 2, 5, 9, 10, 23, 25}) != 5 {
		t.Fatal("magic")
	}
	if MagicIndex([]int{-14, -12, 0, 1, 2, 7, 9, 10, 23, 25}) != NotFound {
		t.Fatal("no magic")
	}
	if MagicIndex([]int{0, 1, 2, 3, 4}) != 2 || MagicIndex(nil) != NotFound {
		t.Fatal("more magic")
	}
	if MagicIndexNonDistinct([]int{-14, -12, 0, 1, 2, 5, 9, 10, 23, 25}) != 5 {
		t.Fatal("nd")
	}
	if MagicIndexNonDistinct([]int{-14, -12, 0, 1, 2, 7, 9, 10, 23, 25}) != NotFound {
		t.Fatal("nd miss")
	}
	if MagicIndexNonDistinct([]int{0, 1, 2, 3, 4}) != 2 || MagicIndexNonDistinct(nil) != NotFound {
		t.Fatal("nd more")
	}
	if MagicIndexNonDistinct([]int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}) != 2 {
		t.Fatal("dups")
	}

	set := []int{1, 2, 3}
	expected := formatSubsets([]string{"", "1", "1,2", "1,2,3", "1,3", "2", "2,3", "3"})
	if formatSubsetLists(GetSubsetsA(set)) != expected || formatSubsetLists(GetSubsetsB(set)) != expected || formatSubsetLists(GetSubsetsC(set)) != expected {
		t.Fatal("subsets")
	}
	if len(GetSubsetsA(nil)) != 1 || len(GetSubsetsA(nil)[0]) != 0 {
		t.Fatal("empty a")
	}
	if len(GetSubsetsB(nil)) != 1 || len(GetSubsetsB(nil)[0]) != 0 {
		t.Fatal("empty b")
	}
	if len(GetSubsetsC(nil)) != 1 || len(GetSubsetsC(nil)[0]) != 0 {
		t.Fatal("empty c")
	}
	if len(ConvertIntToSet(0, []int{1, 2, 3})) != 0 {
		t.Fatal("zero")
	}
	if !reflect.DeepEqual(ConvertIntToSet(0b101, []int{1, 2, 3}), []int{1, 3}) {
		t.Fatal("101")
	}
	if !reflect.DeepEqual(ConvertIntToSet(0b111, []int{1, 2, 3}), []int{1, 2, 3}) {
		t.Fatal("111")
	}

	for _, pair := range [][2]int{{0, 5}, {5, 6}, {28, 89}, {1234, 245334}} {
		want := pair[0] * pair[1]
		if MultiplyBitBased(pair[0], pair[1]) != want || MinProduct(pair[0], pair[1]) != want || MinProduct2(pair[0], pair[1]) != want || MinProduct3(pair[0], pair[1]) != want {
			t.Fatal(pair)
		}
	}
	if Multiply(5, 6) != 30 || Multiply(0, 9) != 0 {
		t.Fatal("multiply")
	}
	mustPanic(t, func() { MinProduct(-1, 2) })

	for n := 1; n < 10; n++ {
		towers := NewTowersOfHanoi(n)
		towers.Solve()
		want := make([]int, n)
		for i := 0; i < n; i++ {
			want[i] = n - i
		}
		if !reflect.DeepEqual(towers.GetStack(2), want) || len(towers.GetStack(0)) != 0 || len(towers.GetStack(1)) != 0 {
			t.Fatal(n, towers.GetStack(2))
		}
	}
	mustPanic(t, func() { NewMultiStack(3).GetStack(-1) })
	mustPanic(t, func() { NewMultiStack(3).GetStack(3) })
	func() {
		defer func() {
			if _, ok := recover().(*StackTooBigError); !ok {
				t.Fatal("expected StackTooBigError")
			}
		}()
		stack := NewStack(1)
		stack.Push(1)
		stack.Push(2)
	}()
}

func TestPermsParensPaint(t *testing.T) {
	assertPerms(t, "", []string{""})
	assertPerms(t, "a", []string{"a"})
	assertPerms(t, "ab", []string{"ab", "ba"})
	assertPerms(t, "str", []string{"str", "srt", "tsr", "trs", "rst", "rts"})
	if GetPerms(nil) != nil {
		t.Fatal("nil")
	}
	aa := "aa"
	if len(GetPerms(&aa)) != 2 {
		t.Fatal("aa")
	}
	if InsertCharAt("ab", 'X', 0) != "Xab" || InsertCharAt("ab", 'X', 1) != "aXb" || InsertCharAt("ab", 'X', 2) != "abX" {
		t.Fatal("insert")
	}
	assertDupPerms(t, "", []string{""})
	assertDupPerms(t, "a", []string{"a"})
	assertDupPerms(t, "aaf", []string{"aaf", "afa", "faa"})
	assertDupPerms(t, "abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"})
	result := PrintPerms("aab")
	if len(result) != 3 || len(unique(result)) != 3 {
		t.Fatal(result)
	}

	parenCases := []struct {
		n    int
		want []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, c := range parenCases {
		formatted := formatStrings(c.want)
		if formatStrings(GenerateParenthesesPermutationsBruteForce(c.n)) != formatted {
			t.Fatal("brute", c.n)
		}
		if formatStrings(GenerateParenthesesPermutationsRecursive1(c.n)) != formatted {
			t.Fatal("rec1", c.n)
		}
		if formatStrings(GenerateParenthesesPermutationsRecursive2(c.n)) != formatted {
			t.Fatal("rec2", c.n)
		}
	}
	mustPanic(t, func() { GenerateParenthesesPermutationsRecursive1(-1) })
	if !IsMatchedParentheses([]rune{'(', ')'}) || !IsMatchedParentheses([]rune("(())()")) {
		t.Fatal("matched")
	}
	if IsMatchedParentheses([]rune{')', '('}) || IsMatchedParentheses([]rune{'(', '(', ')'}) || IsMatchedParentheses([]rune{'(', 'a', ')'}) {
		t.Fatal("unmatched")
	}
	arr := []rune("abc")
	if !NextPermutation(arr) || string(arr) != "acb" {
		t.Fatal(string(arr))
	}
	if NextPermutation([]rune("cba")) {
		t.Fatal("cba")
	}
	addResults := []string{}
	AddParen(&addResults, 2, 2, make([]rune, 4), 0)
	if formatStrings(addResults) != formatStrings([]string{"(())", "()()"}) {
		t.Fatal(addResults)
	}

	screen := [][]int{{1, 2, 5}, {2, 2, 4}, {2, 8, 6}}
	PaintFill(screen, 1, 1, 3)
	if !reflect.DeepEqual(screen, [][]int{{1, 3, 5}, {3, 3, 4}, {3, 8, 6}}) {
		t.Fatal(screen)
	}
	single := [][]int{{7}}
	PaintFill(single, 0, 0, 9)
	if single[0][0] != 9 {
		t.Fatal(single)
	}
	same := [][]int{{1, 1}, {1, 1}}
	PaintFill(same, 0, 0, 1)
	if !reflect.DeepEqual(same, [][]int{{1, 1}, {1, 1}}) {
		t.Fatal(same)
	}
	mustPanic(t, func() { PaintFill(nil, 0, 0, 1) })
	mustPanic(t, func() { PaintFill([][]int{{1}}, 1, 0, 2) })
	flood := [][]int{{2, 2, 1}, {2, 8, 1}}
	FloodFill(flood, 0, 0, 2, 9)
	if !reflect.DeepEqual(flood, [][]int{{9, 9, 1}, {9, 8, 1}}) {
		t.Fatal(flood)
	}
	FloodFill(flood, -1, 0, 9, 0)
	if flood[0][0] != 9 {
		t.Fatal("oob")
	}
}

func TestCoinsQueensBoxesBool(t *testing.T) {
	if CoinCombinations(0) != 1 || CoinCombinations(1) != 1 || CoinCombinations(5) != 2 || CoinCombinations(10) != 4 {
		t.Fatal("coins")
	}
	if CoinCombinationsIterative(0) != 1 || CoinCombinationsIterative(1) != 1 || CoinCombinationsIterative(10) != 4 || CoinCombinationsIterative(100) != 242 {
		t.Fatal("iter")
	}
	if CoinCombinationsIterative(4, []int{1, 2}) != 3 || CoinCombinations(4, []int{1, 2}) != 3 {
		t.Fatal("custom")
	}
	mustPanic(t, func() { CoinCombinationsIterative(-1) })

	board := emptyBoard(8)
	if !PositionIsValid(0, 0, board) || !PositionIsValid(7, 7, board) {
		t.Fatal("valid empty")
	}
	board[4][3] = 'Q'
	if PositionIsValid(7, 3, board) {
		t.Fatal("col")
	}
	board[4][3] = '.'
	board[0][0] = 'Q'
	if PositionIsValid(5, 5, board) {
		t.Fatal("diag")
	}
	board[0][0] = '.'
	board[2][7] = 'Q'
	if PositionIsValid(3, 6, board) {
		t.Fatal("diag2")
	}
	if len(Queens(1)) != 1 || len(Queens(4)) != 2 {
		t.Fatal("counts")
	}
	expected := formatBoards([][][]rune{
		{{'.', 'Q', '.', '.'}, {'.', '.', '.', 'Q'}, {'Q', '.', '.', '.'}, {'.', '.', 'Q', '.'}},
		{{'.', '.', 'Q', '.'}, {'Q', '.', '.', '.'}, {'.', '.', '.', 'Q'}, {'.', 'Q', '.', '.'}},
	})
	if formatBoards(Queens(4)) != expected {
		t.Fatal("boards")
	}
	mustPanic(t, func() { Queens(0) })
	mustPanic(t, func() { Queens(13) })

	if TallestStack(nil) != 0 || TallestStack([]Box{NewBox(3, 2, 1)}) != 3 {
		t.Fatal("boxes")
	}
	if TallestStack([]Box{NewBox(3, 2, 1), NewBox(5, 4, 1)}) != 5 {
		t.Fatal("no stack")
	}
	if TallestStack([]Box{NewBox(3, 2, 1), NewBox(6, 5, 4)}) != 9 {
		t.Fatal("stack")
	}
	mustPanic(t, func() { NewBox(0, 1, 1) })

	if Evaluate("1^0|0|1", false) != 2 || Evaluate("0&0&0&1^1|0", true) != 10 {
		t.Fatal("eval")
	}
	if Evaluate("1", true) != 1 || Evaluate("0", true) != 0 || Evaluate("1&0", true) != 0 || Evaluate("1&0", false) != 1 {
		t.Fatal("simple")
	}
	mustPanic(t, func() { Evaluate("1+", true) })
	if !StringToBool("1") || StringToBool("0") {
		t.Fatal("bool")
	}
	mustPanic(t, func() { StringToBool("2") })
	if CountWays("1^0|0|1", false, map[string]int{}) != 2 || CountWays("", true, map[string]int{}) != 0 || CountWays("1", true, map[string]int{}) != 1 {
		t.Fatal("ways")
	}
}

func formatSubsetLists(subsets [][]int) string {
	parts := make([]string, len(subsets))
	for i, subset := range subsets {
		strs := make([]string, len(subset))
		for j, v := range subset {
			strs[j] = itoa(v)
		}
		parts[i] = strings.Join(strs, ",")
	}
	return formatSubsets(parts)
}

func formatSubsets(subsets []string) string {
	cp := append([]string{}, subsets...)
	sort.Strings(cp)
	return strings.Join(cp, ";")
}

func formatStrings(values []string) string {
	cp := append([]string{}, values...)
	sort.Strings(cp)
	return strings.Join(cp, ";")
}

func assertPerms(t *testing.T, value string, expected []string) {
	t.Helper()
	v := value
	if formatStrings(expected) != formatStrings(GetPerms(&v)) || formatStrings(expected) != formatStrings(GetPerms2(value)) {
		t.Fatal(value)
	}
}

func assertDupPerms(t *testing.T, value string, expected []string) {
	t.Helper()
	if formatStrings(expected) != formatStrings(PrintPerms(value)) {
		t.Fatal(value, PrintPerms(value))
	}
}

func unique(values []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0)
	for _, v := range values {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func emptyBoard(n int) [][]rune {
	board := make([][]rune, n)
	for i := range n {
		board[i] = make([]rune, n)
		for j := range n {
			board[i][j] = '.'
		}
	}
	return board
}

func formatBoards(boards [][][]rune) string {
	parts := make([]string, len(boards))
	for i, board := range boards {
		rows := make([]string, len(board))
		for r, row := range board {
			rows[r] = string(row)
		}
		parts[i] = strings.Join(rows, "/")
	}
	sort.Strings(parts)
	return strings.Join(parts, "|")
}

func itoa(v int) string {
	return strconv.Itoa(v)
}
