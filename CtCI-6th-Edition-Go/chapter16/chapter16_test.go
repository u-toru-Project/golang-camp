package chapter16

import (
	"math"
	"reflect"
	"testing"

	"ctci/library"
)

func TestSwap(t *testing.T) {
	a, b := 5, 10
	Swap(&a, &b)
	if a != 10 || b != 5 {
		t.Fatal(a, b)
	}
	a, b = 7, 7
	Swap(&a, &b)
	if a != 7 || b != 7 {
		t.Fatal(a, b)
	}
	a, b = -3, 0
	Swap(&a, &b)
	if a != 0 || b != -3 {
		t.Fatal(a, b)
	}
}

func TestWordFrequency(t *testing.T) {
	freq := GetWordFrequency("The dog is the DOG")
	if freq["the"] != 2 || freq["dog"] != 2 || freq["is"] != 1 || len(freq) != 3 {
		t.Fatal(freq)
	}
	freq = GetWordFrequency("a  b   a")
	if freq["a"] != 2 || freq["b"] != 1 {
		t.Fatal(freq)
	}
	if _, ok := freq[""]; ok {
		t.Fatal("empty token")
	}
	if len(GetWordFrequency("")) != 0 || len(GetWordFrequency("   ")) != 0 {
		t.Fatal("empty book")
	}
	freq = GetWordFrequency("Lara spotted Lara")
	if GetFrequency(freq, "lara") != 2 || GetFrequency(freq, "LARA") != 2 || GetFrequency(freq, "spotted") != 1 {
		t.Fatal(freq)
	}
	if GetFrequency(GetWordFrequency("one two"), "three") != 0 {
		t.Fatal("missing")
	}
}

func TestSegmentIntersection(t *testing.T) {
	hit := SegmentIntersection(Point{10, 10}, Point{20, 20}, Point{20, 10}, Point{10, 20})
	if hit == nil || hit.X != 15 || hit.Y != 15 {
		t.Fatal(hit)
	}
	if SegmentIntersection(Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 1}) != nil {
		t.Fatal("parallel")
	}
	hit = SegmentIntersection(Point{10, 10}, Point{20, 10}, Point{15, 10}, Point{25, 10})
	if hit == nil || hit.Y != 10 || hit.X < 15 || hit.X > 20 {
		t.Fatal(hit)
	}
	hit = SegmentIntersection(Point{0, 0}, Point{2, 0}, Point{2, 0}, Point{2, 2})
	if hit == nil || hit.X != 2 || hit.Y != 0 {
		t.Fatal(hit)
	}
}

func TestTicTacToe(t *testing.T) {
	empty := [][]Piece{{PieceNA, PieceNA, PieceNA}, {PieceNA, PieceNA, PieceNA}, {PieceNA, PieceNA, PieceNA}}
	if HasWon(empty) != PieceNA {
		t.Fatal("empty")
	}
	rowWin := [][]Piece{{PieceX, PieceX, PieceX}, {PieceO, PieceNA, PieceO}, {PieceNA, PieceO, PieceNA}}
	if HasWon(rowWin) != PieceX {
		t.Fatal("row")
	}
	colWin := [][]Piece{{PieceX, PieceO, PieceNA}, {PieceX, PieceO, PieceX}, {PieceNA, PieceO, PieceX}}
	if HasWon(colWin) != PieceO {
		t.Fatal("col")
	}
	diag := [][]Piece{{PieceX, PieceO, PieceO}, {PieceNA, PieceX, PieceO}, {PieceNA, PieceNA, PieceX}}
	if HasWon(diag) != PieceX {
		t.Fatal("diag")
	}
	anti := [][]Piece{{PieceO, PieceO, PieceX}, {PieceNA, PieceX, PieceO}, {PieceX, PieceNA, PieceNA}}
	if HasWon(anti) != PieceX {
		t.Fatal("anti")
	}
	partial := [][]Piece{{PieceX, PieceNA, PieceNA}, {PieceNA, PieceO, PieceNA}, {PieceNA, PieceNA, PieceNA}}
	if HasWon(partial) != PieceNA {
		t.Fatal("partial")
	}
	game := NewTicTacToe(PieceX)
	game.Place(0, 2, PieceX)
	if game.GetWinner() != PieceNA {
		t.Fatal("first move")
	}
	game = NewTicTacToe(PieceX)
	game.Place(1, 1, PieceX)
	game.Place(1, 2, PieceO)
	game.Place(0, 2, PieceX)
	game.Place(0, 1, PieceO)
	game.Place(2, 0, PieceX)
	if game.GetWinner() != PieceX {
		t.Fatal("anti win")
	}
	mustPanic(t, func() { NewTicTacToe(PieceNA) })
	game = NewTicTacToe(PieceX)
	mustPanic(t, func() { game.Place(-1, 0, PieceX) })
	mustPanic(t, func() { game.Place(0, 0, PieceO) })
	game.Place(0, 0, PieceX)
	mustPanic(t, func() { game.Place(0, 0, PieceO) })
	finished := NewTicTacToe(PieceX)
	finished.Place(0, 0, PieceX)
	finished.Place(1, 0, PieceO)
	finished.Place(0, 1, PieceX)
	finished.Place(1, 1, PieceO)
	finished.Place(0, 2, PieceX)
	if finished.GetWinner() != PieceX {
		t.Fatal("row win game")
	}
	mustPanic(t, func() { finished.Place(2, 2, PieceO) })
}

func TestFactorialZeros(t *testing.T) {
	if TrailingZerosInFactorial(0) != 0 || TrailingZerosInFactorial(1) != 0 ||
		TrailingZerosInFactorial(5) != 1 || TrailingZerosInFactorial(10) != 2 ||
		TrailingZerosInFactorial(25) != 6 || TrailingZerosInFactorial(100) != 24 {
		t.Fatal("zeros")
	}
	mustPanic(t, func() { TrailingZerosInFactorial(-1) })
}

func TestSmallestDifference(t *testing.T) {
	if FindSmallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}) != 3 {
		t.Fatal("book")
	}
	if FindSmallestDifference([]int{10, 4}, []int{4, 99}) != 0 || FindSmallestDifference([]int{1}, []int{1}) != 0 {
		t.Fatal("zero")
	}
	if FindSmallestDifference([]int{-5, -1}, []int{0, 8}) != 1 || FindSmallestDifference([]int{-10, -5}, []int{-7, 4}) != 2 {
		t.Fatal("neg")
	}
	if FindSmallestDifference(nil, []int{1}) != math.MaxInt32 || FindSmallestDifference([]int{1}, nil) != math.MaxInt32 {
		t.Fatal("empty")
	}
	if GetSmallestDifferencePair(nil, []int{1}) != nil {
		t.Fatal("nil pair")
	}
	pair := GetSmallestDifferencePair([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8})
	if pair == nil || pair.First != 11 || pair.Second != 8 {
		t.Fatal(pair)
	}
}

func TestNumberMax(t *testing.T) {
	pairs := [][2]int{{10, 20}, {20, 10}, {-10, -20}, {20, -10}, {0, 10}, {-20, 0}}
	for _, p := range pairs {
		want := max(p[1], p[0])
		if GetMax(p[0], p[1]) != want {
			t.Fatal(p)
		}
	}
	if CountMaxes([]int{1, 2, 3, 2, 3, 3}) != 3 || CountMaxes([]int{5, 5, 5}) != 3 ||
		CountMaxes([]int{1}) != 1 || CountMaxes([]int{3, 1, 4, 4}) != 2 || CountMaxes(nil) != 0 {
		t.Fatal("count")
	}
}

func TestEnglishInt(t *testing.T) {
	cases := map[int]string{
		0:             "Zero",
		1:             "One",
		19:            "Nineteen",
		20:            "Twenty",
		21:            "Twenty One",
		99:            "Ninety Nine",
		100:           "One Hundred",
		101:           "One Hundred One",
		110:           "One Hundred Ten",
		1000:          "One Thousand",
		1234:          "One Thousand Two Hundred Thirty Four",
		1_000_000:     "One Million",
		1_000_000_000: "One Billion",
		832787436:     "Eight Hundred Thirty Two Million Seven Hundred Eighty Seven Thousand Four Hundred Thirty Six",
		-15:           "Negative Fifteen",
		math.MaxInt32: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven",
		math.MinInt32: "Negative Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Eight",
	}
	for n, want := range cases {
		if got := ConvertIntToEnglish(n); got != want {
			t.Fatalf("%d: got %q want %q", n, got, want)
		}
	}
}

func TestOperations(t *testing.T) {
	pairs := [][2]int{{10, 3}, {10, -3}, {-10, 3}, {-10, -3}, {0, 5}, {-3, 12}}
	for _, p := range pairs {
		if Multiply(p[0], p[1]) != p[0]*p[1] {
			t.Fatal(p)
		}
	}
	if Divide(10, 3) != 3 || Divide(10, -3) != -4 || Divide(-10, 3) != -4 || Divide(-10, -3) != 3 ||
		Divide(7, 1) != 7 || Divide(0, 5) != 0 {
		t.Fatal("div")
	}
	if Subtract(10, 3) != 7 || Subtract(3, 10) != -7 || Subtract(3, 15) != -12 {
		t.Fatal("sub")
	}
	mustPanic(t, func() { Divide(1, 0) })
	if Negate(0) != 0 || Negate(5) != -5 || Negate(-5) != 5 {
		t.Fatal("neg")
	}
}

func TestLivingPeople(t *testing.T) {
	people := []Person{{12, 15}, {20, 90}, {10, 98}, {1, 72}, {10, 98}, {23, 82}, {13, 98}, {90, 98}, {83, 99}, {75, 94}}
	if YearWithMostLiving(people) != 90 {
		t.Fatal(YearWithMostLiving(people))
	}
	if YearWithMostLiving([]Person{{1900, 2000}, {1944, 1944}, {1944, 1944}}) != 1944 {
		t.Fatal("1944")
	}
	mustPanic(t, func() { YearWithMostLiving([]Person{{2000, 1990}}) })
}

func TestDivingBoard(t *testing.T) {
	if !reflect.DeepEqual(DivingBoardLengths(10, 10, 5), []int{50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}) {
		t.Fatal(DivingBoardLengths(10, 10, 5))
	}
	if !reflect.DeepEqual(DivingBoardLengths(10, 10, 10), []int{100}) {
		t.Fatal("equal")
	}
	if !reflect.DeepEqual(DivingBoardLengths(3, 5, 2), []int{6, 9, 12, 15}) {
		t.Fatal(DivingBoardLengths(3, 5, 2))
	}
	if !reflect.DeepEqual(DivingBoardLengths(0, 7, 3), []int{0}) {
		t.Fatal("zero")
	}
	mustPanic(t, func() { DivingBoardLengths(2, 3, 5) })
}

func TestXmlEncoding(t *testing.T) {
	cases := [][2]string{{"hello", "hello"}, {"Tom & Jerry", "Tom &amp; Jerry"}, {"<tag attr=\"x\">", "&lt;tag attr=&quot;x&quot;&gt;"}, {"it's", "it&apos;s"}}
	for _, c := range cases {
		if EncodeXml(c[0]) != c[1] || DecodeXml(c[1]) != c[0] {
			t.Fatal(c, EncodeXml(c[0]))
		}
	}
	mustPanic(t, func() { DecodeXml("&unknown;") })
}

func TestBisectSquares(t *testing.T) {
	first := Square{Point{2, 5}, Point{6, 5}, Point{2, 1}, Point{6, 1}}
	second := Square{Point{7, 8}, Point{9, 8}, Point{7, 6}, Point{9, 6}}
	seg := SquareCutSegment(first, second)
	if seg.Start.equal(seg.End) {
		t.Fatal(seg)
	}
	outer := Square{Point{0, 4}, Point{4, 4}, Point{0, 0}, Point{4, 0}}
	inner := Square{Point{1, 3}, Point{2, 3}, Point{1, 2}, Point{2, 2}}
	seg = SquareCutSegment(outer, inner)
	if seg.Start.X > seg.End.X {
		t.Fatal(seg)
	}
	center := outer.Center()
	if center.X != 2 || center.Y != 2 || len(outer.Sides()) != 4 {
		t.Fatal(center)
	}
	same := Square{Point{0, 2}, Point{2, 2}, Point{0, 0}, Point{2, 0}}
	other := Square{Point{-1, 3}, Point{3, 3}, Point{-1, -1}, Point{3, -1}}
	mustPanic(t, func() { SquareCutSegment(same, other) })
}

func TestBestLine(t *testing.T) {
	_, count := BestLine([]Point{{1, 1}, {4, 1}, {7, 1}, {8, 3}, {3, 3}})
	if count < 3 {
		t.Fatal(count)
	}
	line, count := BestLine([]Point{{10, 15}, {10, 16}, {10, 17}, {11, 15}})
	if !line.IsVertical || count < 3 {
		t.Fatal(line, count)
	}
}

func TestMasterMind(t *testing.T) {
	assertScore := func(sol, guess string, hits, pseudo int) {
		h, p := ScoreGuess(sol, guess)
		if h != hits || p != pseudo {
			t.Fatalf("%s %s got %d %d", sol, guess, h, p)
		}
	}
	assertScore("RGBY", "GGRR", 1, 1)
	assertScore("RGBY", "RBGY", 2, 2)
	assertScore("RRYY", "RYGY", 2, 1)
	assertScore("YYYY", "YYYY", 4, 0)
	assertScore("RGBY", "YYYY", 1, 0)
	mustPanic(t, func() { ScoreGuess("RGB", "RGBY") })
}

func TestSubSort(t *testing.T) {
	assertRange := func(arr []int, start, end int) {
		s, e := SubSort(arr)
		if s != start || e != end {
			t.Fatalf("%v got %d %d want %d %d", arr, s, e, start, end)
		}
	}
	assertRange([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}, 4, 9)
	assertRange([]int{1, 2, 3, 4, 5}, -1, -1)
	assertRange([]int{7, 12, 3, 4, 5, 6, 0}, 0, 6)
	assertRange([]int{1}, -1, -1)
	assertRange([]int{2, 1}, 0, 1)
	assertRange(nil, -1, -1)
	assertRange([]int{1, 1, 1}, -1, -1)
}

func TestContiguous(t *testing.T) {
	if MaxContiguousSum([]int{2, -8, 3, -2, 4, -10}) != 5 {
		t.Fatal("book")
	}
	if MaxContiguousSum([]int{1, 2, 3, 4, 5, 6}) != 21 {
		t.Fatal("pos")
	}
	if MaxContiguousSum([]int{-10, 100, -5, -5, -5, -5, -5, 1, 2, 3, 4, 25, -50}) != 110 {
		t.Fatal("110")
	}
	if MaxContiguousSum([]int{-3, -1, -2}) != -1 || MaxContiguousSum([]int{5}) != 5 {
		t.Fatal("neg")
	}
	mustPanic(t, func() { MaxContiguousSum(nil) })
}

func TestPatternMatcher(t *testing.T) {
	if IsMatch("aa", "a") || !IsMatch("aa", "a*") || !IsMatch("ab", ".*") || !IsMatch("aab", "c*a*b") ||
		!IsMatch("mississippi", "mis*is*ip*.") || IsMatch("abc", "d*") || !IsMatch("", ".*") || IsMatch("a", "") {
		t.Fatal("match")
	}
}

func TestPondSizes(t *testing.T) {
	matrix := [][]rune{
		{'w', 'h', 'l', 'w'},
		{'w', 'l', 'w', '1'},
		{'l', 'l', 'w', 'l'},
		{'w', 'l', 'w', 'l'},
	}
	sizes := GetSizes(matrix, 'w')
	if len(sizes) != 3 || sizes[0] != 2 || sizes[1] != 4 || sizes[2] != 1 {
		t.Fatal(sizes)
	}
	if len(GetSizes([][]rune{{'l', 'l'}, {'l', 'l'}}, 'w')) != 0 {
		t.Fatal("empty")
	}
	all := GetSizes([][]rune{{'w', 'w'}, {'w', 'w'}}, 'w')
	if len(all) != 1 || all[0] != 4 {
		t.Fatal(all)
	}
	if GetSizes([][]rune{{'w'}}, 'w')[0] != 1 || len(GetSizes([][]rune{{'l'}}, 'w')) != 0 {
		t.Fatal("single")
	}
}

func TestT9(t *testing.T) {
	words := T9Words("8733", []string{"tree", "used", "trend", "apple"})
	if len(words) != 2 || words[0] != "tree" || words[1] != "used" {
		t.Fatal(words)
	}
	words = T9Words("222", []string{"aaa", "aab", "aba", "abc", "cab"})
	if !reflect.DeepEqual(words, []string{"aaa", "aab", "aba", "abc", "cab"}) {
		t.Fatal(words)
	}
	if len(T9Words("2222", []string{"dog"})) != 0 {
		t.Fatal("nomatch")
	}
	mustPanic(t, func() { T9Words("102", []string{"a"}) })
}

func TestSumSwap(t *testing.T) {
	pair := SumSwap([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3})
	if pair == nil || pair.First != 4 || pair.Second != 6 {
		t.Fatal(pair)
	}
	if SumSwap([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 4}) != nil || SumSwap([]int{5, 5}, []int{3, 7}) != nil {
		t.Fatal("impossible")
	}
	first, second := []int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3}
	pair = SumSwap(first, second)
	sum1, sum2 := 0, 0
	for _, v := range first {
		sum1 += v
	}
	for _, v := range second {
		sum2 += v
	}
	if sum1-pair.First+pair.Second != sum2-pair.Second+pair.First {
		t.Fatal("not equalized")
	}
}

func TestLangtonsAnt(t *testing.T) {
	if len(Simulate(0)) != 0 {
		t.Fatal("zero")
	}
	grid := Simulate(12)
	if len(grid) < 1 {
		t.Fatal("small")
	}
	box := BoundingBox(grid)
	if box.MaxRow-box.MinRow > 5 || box.MaxCol-box.MinCol > 5 {
		t.Fatal(box)
	}
	grid = Simulate(10000)
	box = BoundingBox(grid)
	if box.MaxRow <= 10 && box.MaxCol <= 10 {
		t.Fatal(box)
	}
}

func TestRand7(t *testing.T) {
	values := []int{0, 1, 2, 3, 4}
	index := 0
	mock := func() int {
		v := values[index%len(values)]
		index++
		return v
	}
	seen := map[int]struct{}{}
	for range 50 {
		v := Rand7(mock)
		if v < 0 || v > 6 {
			t.Fatal(v)
		}
		seen[v] = struct{}{}
	}
	counts := make([]int, 7)
	for range 7000 {
		counts[Rand7(nil)]++
	}
	for i, c := range counts {
		if c < 850 || c > 1150 {
			t.Fatalf("count %d = %d", i, c)
		}
	}
	seen5 := map[int]struct{}{}
	for range 80 {
		v := Rand5()
		if v < 0 || v > 4 {
			t.Fatal(v)
		}
		seen5[v] = struct{}{}
	}
	if len(seen5) < 2 {
		t.Fatal(seen5)
	}
}

func TestPairsWithSum(t *testing.T) {
	if CountPairsWithSum([]int{4, 6, 10, 15, 16}, 21) != 1 {
		t.Fatal("21")
	}
	if CountPairsWithSum([]int{1, 1, 1, 1, 3}, 4) != 4 {
		t.Fatal("4")
	}
	if CountPairsWithSum([]int{1, 2, 3, 4}, 10) != 0 {
		t.Fatal("10")
	}
	if CountPairsWithSum([]int{-2, 0, 2, 2, 4}, 2) != 3 {
		t.Fatal("2")
	}
	if CountPairsWithSum(nil, 5) != 0 {
		t.Fatal("empty")
	}
}

func TestLruCache(t *testing.T) {
	cache := NewLruCache(2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	if cache.Get("a") != 1 {
		t.Fatal("a")
	}
	cache.Put("c", 3)
	mustPanic(t, func() { cache.Get("b") })
	if cache.Get("c") != 3 {
		t.Fatal("c")
	}
	cache = NewLruCache(2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("a", 10)
	cache.Put("c", 3)
	if cache.Get("a") != 10 {
		t.Fatal("update")
	}
	mustPanic(t, func() { cache.Get("b") })
	mustPanic(t, func() { NewLruCache(0) })
}

func TestCalculator(t *testing.T) {
	eq := func(expected, actual float64) {
		if math.Abs(expected-actual) >= 1e-9 {
			t.Fatalf("expected %v got %v", expected, actual)
		}
	}
	eq(-27, Calculate("2 -6 - 7 * 8 / 2 + 5"))
	eq(23.5, Calculate("2*3+5/6*3+15"))
	eq(7, Calculate("1+2*3"))
	eq(5, Calculate("10/2"))
	eq(7.0, Calculate("3.5*2"))
	eq(0, Calculate("0+0"))
	eq(42, Calculate("42"))
	for _, input := range []string{"__import__('os')", "1+1; print(1)", "().__class__", "1+open('/etc/passwd')", "eval(1)", "1**2", "1+2a", "", "   ", "+", "1+", "+1"} {
		mustPanic(t, func() { Calculate(input) })
	}
	mustPanic(t, func() { Calculate("1/0") })
	mustPanic(t, func() { Collapse(1, 2, "%") })
	tokens := ParseEquation("2 + 3.5*4")
	if len(tokens) != 5 || tokens[0].(float64) != 2.0 || tokens[1].(rune) != '+' || tokens[2].(float64) != 3.5 || tokens[3].(rune) != '*' || tokens[4].(float64) != 4.0 {
		t.Fatal(tokens)
	}
	mustPanic(t, func() { ParseEquation("2 + x") })
}

func TestBoggle(t *testing.T) {
	matrix := [][]rune{
		{'b', 'c', 'e', 'p'},
		{'e', 'o', 'r', 'o'},
		{'e', 'm', 't', 'n'},
		{'s', 'e', 'a', 'i'},
	}
	trie := library.NewTrie([]string{"bee", "become", "seem", "certain", "top", "missing"})
	words := Solve(matrix, trie.Root)
	set := map[string]struct{}{}
	for _, w := range words {
		set[w] = struct{}{}
	}
	for _, w := range []string{"bee", "become", "seem", "certain", "top"} {
		if _, ok := set[w]; !ok {
			t.Fatal("missing", w, words)
		}
	}
	if _, ok := set["missing"]; ok || len(words) != 5 {
		t.Fatal(words)
	}
	small := [][]rune{{'a', 'b'}, {'c', 'd'}}
	words = Solve(small, library.NewTrie([]string{"ab", "aba"}).Root)
	set = map[string]struct{}{}
	for _, w := range words {
		set[w] = struct{}{}
	}
	if _, ok := set["ab"]; !ok {
		t.Fatal(words)
	}
	if _, ok := set["aba"]; ok {
		t.Fatal(words)
	}
	if len(Solve([][]rune{{'z'}}, library.NewTrie([]string{"abc"}).Root)) != 0 {
		t.Fatal("empty")
	}
	words = Solve([][]rune{{'a', 'b'}}, library.NewTrie([]string{"a", "b", "c"}).Root)
	set = map[string]struct{}{}
	for _, w := range words {
		set[w] = struct{}{}
	}
	if _, ok := set["a"]; !ok {
		t.Fatal(words)
	}
	if _, ok := set["b"]; !ok {
		t.Fatal(words)
	}
	if _, ok := set["c"]; ok || len(words) != 2 {
		t.Fatal(words)
	}
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}
