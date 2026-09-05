package chapter17

import (
	"math"
	"reflect"
	"testing"
)

func mustPanic17(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func TestAddWithoutPlus(t *testing.T) {
	cases := [][3]int{{1, 1, 2}, {1, 2, 3}, {1001, 234, 1235}, {123456789, 123456789, 123456789 * 2}, {0, 0, 0}, {0, 5, 5}}
	for _, c := range cases {
		if AddWithoutPlus(c[0], c[1]) != c[2] || AddWithoutPlusRecursive(c[0], c[1]) != c[2] {
			t.Fatal(c)
		}
	}
	mustPanic17(t, func() { AddWithoutPlus(-1, 2) })
	mustPanic17(t, func() { AddWithoutPlusRecursive(1, -2) })
}

func TestShuffle(t *testing.T) {
	for _, n := range []int{0, 1, 2, 10} {
		cards := make([]int, n)
		for i := range cards {
			cards[i] = i
		}
		original := append([]int(nil), cards...)
		result := ShuffleListIteratively(cards)
		if len(result) != n {
			t.Fatal(n)
		}
		assertPerm(t, original, result)
	}
	cards := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	original := append([]int(nil), cards...)
	assertPerm(t, original, ShuffleListRecursively(cards, len(cards)-1))
	if len(ShuffleListIteratively(nil)) != 0 {
		t.Fatal("empty")
	}
	for range 20 {
		if RandomNumberGenerator(3, 3) != 3 {
			t.Fatal("bounds")
		}
	}
	mustPanic17(t, func() { RandomNumberGenerator(5, 1) })
}

func assertPerm(t *testing.T, original, shuffled []int) {
	t.Helper()
	left := append([]int(nil), original...)
	right := append([]int(nil), shuffled...)
	sortInts(left)
	sortInts(right)
	if !reflect.DeepEqual(left, right) {
		t.Fatal(original, shuffled)
	}
}

func sortInts(a []int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func TestRandomSet(t *testing.T) {
	set := NewRandomSet()
	if !set.Insert(10) || set.Insert(10) || !set.Contains(10) || set.Count() != 1 {
		t.Fatal("insert")
	}
	for _, v := range []int{1, 2, 3} {
		set.Insert(v)
	}
	if !set.Remove(2) || set.Contains(2) || set.Remove(99) || set.Count() != 3 {
		t.Fatal("remove", set.Count())
	}
	set = NewRandomSet()
	for _, v := range []int{5, 10, 15} {
		set.Insert(v)
	}
	allowed := map[int]struct{}{5: {}, 10: {}, 15: {}}
	for range 20 {
		if _, ok := allowed[set.GetRandom()]; !ok {
			t.Fatal("random")
		}
	}
	mustPanic17(t, func() { NewRandomSet().GetRandom() })
}

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		arr     []int
		n, want int
	}{
		{[]int{0, 1, 2}, 3, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, 0},
		{[]int{0, 1, 2, 3, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 0}, 9, 9},
	}
	for _, c := range cases {
		if MissingNumberXor(c.arr, c.n) != c.want || MissingNumberSum(c.arr, c.n) != c.want {
			t.Fatal(c)
		}
	}
	if MissingNumberXor([]int{0}, 1) != 1 || MissingNumberXor([]int{1}, 1) != 0 {
		t.Fatal("single")
	}
}

func TestLettersAndNumbers(t *testing.T) {
	cases := []struct {
		text, sub string
		length    int
	}{
		{"aaa2222aa22aa222aaaa", "aaa2222aa22aa222aa", 18},
		{"aaa222aa", "aaa222", 6},
		{"e4ee2a2a2a", "4ee2a2a2", 8},
		{"z0123456789a", "z0", 2},
		{"a1", "a1", 2},
	}
	for _, c := range cases {
		length, start, end := LongestBalancedSubarray(c.text)
		if length != c.length || c.text[start:end+1] != c.sub {
			t.Fatal(c, length, start, end)
		}
	}
}

func TestCountOfTwos(t *testing.T) {
	cases := [][2]int{{0, 0}, {2, 1}, {20, 3}, {25, 9}, {100, 20}, {999, 300}}
	for _, c := range cases {
		if CountOfTwos(c[0]) != c[1] {
			t.Fatal(c, CountOfTwos(c[0]))
		}
	}
	for n := range 80 {
		brute := 0
		for i := 0; i <= n; i++ {
			for _, ch := range itoa(i) {
				if ch == '2' {
					brute++
				}
			}
		}
		if CountOfTwos(n) != brute || CountDigitInRange(n, 2) != brute {
			t.Fatal(n, brute, CountOfTwos(n))
		}
	}
	if CountDigitInRange(10, 1) != 2 {
		t.Fatal("digit 1")
	}
	mustPanic17(t, func() { CountDigitInRange(-1, 2) })
	mustPanic17(t, func() { CountDigitInRange(10, 10) })
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}

func TestBabyNames(t *testing.T) {
	nameCounts := map[string]int{
		"john": 10, "jon": 3, "davis": 2, "kari": 3, "johny": 11,
		"carlton": 8, "carleton": 2, "jonathan": 9, "carrie": 5,
	}
	synonyms := [][2]string{{"jonathan", "john"}, {"jon", "johny"}, {"johny", "john"}, {"kari", "carrie"}, {"carleton", "carlton"}}
	got := CountBabyNames(nameCounts, synonyms)
	want := map[string]int{"john": 33, "davis": 2, "carrie": 8, "carlton": 10}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatal(got)
		}
	}
	got = CountBabyNames(map[string]int{"a": 1}, nil)
	if got["a"] != 1 || len(got) != 1 {
		t.Fatal(got)
	}
	if len(CountBabyNames(map[string]int{}, nil)) != 0 {
		t.Fatal("empty")
	}
	mustPanic17(t, func() { CountBabyNames(map[string]int{"a": 1}, [][2]string{{"a", "missing"}}) })
}

func TestCircusTower(t *testing.T) {
	if FindMaxPeople(nil) != 0 {
		t.Fatal("empty")
	}
	if FindMaxPeople([]HeightWeight{{65, 100}, {100, 65}}) != 1 {
		t.Fatal("1")
	}
	if FindMaxPeople([]HeightWeight{{65, 100}, {65, 100}}) != 1 {
		t.Fatal("same")
	}
	if FindMaxPeople([]HeightWeight{{65, 100}, {65, 101}}) != 1 {
		t.Fatal("same h")
	}
	if FindMaxPeople([]HeightWeight{{65, 100}, {55, 40}, {75, 90}, {80, 120}}) != 3 {
		t.Fatal("3")
	}
	if FindMaxPeople([]HeightWeight{{65, 100}, {70, 150}, {56, 90}, {75, 190}, {60, 95}, {68, 110}}) != 6 {
		t.Fatal("6")
	}
}

func TestKthMultiple(t *testing.T) {
	cases := []struct {
		k    int
		want int64
	}{{1, 3}, {2, 5}, {3, 7}, {4, 9}, {5, 15}}
	for _, c := range cases {
		if GetKthMultiple(c.k) != c.want || GetKthMultipleViaHeap(c.k) != c.want {
			t.Fatal(c)
		}
	}
	if GetKthMultipleViaHeap(1000) != 82046671875 {
		t.Fatal(GetKthMultipleViaHeap(1000))
	}
	mustPanic17(t, func() { GetKthMultiple(0) })
	mustPanic17(t, func() { GetKthMultipleViaHeap(0) })
}

func TestMajority(t *testing.T) {
	if MajorityElement([]int{4, 4, 4, 4, 5, 5, 5, 5, 5}) != 5 {
		t.Fatal("5")
	}
	if MajorityElement([]int{1, 2, 3, 4, 5}) != -1 {
		t.Fatal("none")
	}
	if MajorityElement([]int{7, 7, 7, 7, 1, 2, 3}) != 7 {
		t.Fatal("7")
	}
	if MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Fatal("2")
	}
	mustPanic17(t, func() { MajorityElement(nil) })
	mustPanic17(t, func() { MajorityCandidate(nil) })
	majority := []int{7, 7, 7, 7, 1, 2, 3}
	if MajorityCandidate(majority) != 7 || !IsMajority(7, majority) || IsMajority(1, majority) {
		t.Fatal("candidate")
	}
	none := []int{1, 2, 3, 4, 5}
	if IsMajority(MajorityCandidate(none), none) {
		t.Fatal("none majority")
	}
}

func TestWordDistance(t *testing.T) {
	if !reflect.DeepEqual(ShortestDistancesForQueries([]string{"the", "quick", "brown", "fox", "quick"}, [][2]string{{"quick", "fox"}, {"the", "fox"}}), []int{1, 3}) {
		t.Fatal("queries")
	}
	if !reflect.DeepEqual(ShortestDistancesForQueries([]string{"a", "b", "a", "b", "a", "b"}, [][2]string{{"a", "b"}}), []int{1}) {
		t.Fatal("ab")
	}
	mustPanic17(t, func() { ShortestWordDistance("x", "y", map[string][]int{"x": {0}}) })
}

func TestBiNode(t *testing.T) {
	head := FlattenBstToDlist(MakeSampleTree())
	if head == nil {
		t.Fatal("nil")
	}
	values := make([]int, 0)
	node := head
	for range 10000 {
		values = append(values, node.Value)
		if node.Node2 == nil || node.Node2 == head {
			break
		}
		node = node.Node2
	}
	if !reflect.DeepEqual(values, []int{5, 10, 20, 21, 22, 23, 25}) {
		t.Fatal(values)
	}
	if head.Node1 == nil || head.Node1.Node2 != head || head.Node1 != head.Node1 {
		t.Fatal("circular")
	}
	if FlattenBstToDlist(nil) != nil {
		t.Fatal("empty")
	}
}

func TestRespace(t *testing.T) {
	words := []string{"help", "tech", "tips", "boot", "my", "reboot", "need", "to", "linus"}
	text := "helplinustechtipsineedtorebootmypc"
	result := Respace(text, words)
	if !containsByte(result, ' ') || stripSpaces(result) != text {
		t.Fatal(result)
	}
	if Respace("linus", []string{"linus"}) != "linus" || Respace("abc", []string{"a", "bc"}) != "a bc" {
		t.Fatal("small")
	}
}

func containsByte(s string, ch byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			return true
		}
	}
	return false
}

func stripSpaces(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			out = append(out, s[i])
		}
	}
	return string(out)
}

func TestSmallestK(t *testing.T) {
	assertBoth := func(arr []int, k int, expected []int) {
		original := append([]int(nil), arr...)
		if !reflect.DeepEqual(SmallestKHeap(arr, k), expected) || !reflect.DeepEqual(SmallestKQuickselect(arr, k), expected) {
			t.Fatal(arr, k)
		}
		if !reflect.DeepEqual(arr, original) {
			t.Fatal("mutated")
		}
	}
	assertBoth([]int{1, 2, 3, 4, 5, 6}, 4, []int{1, 2, 3, 4})
	assertBoth([]int{1, 2, 2, 2, 1, 1, 0}, 3, []int{0, 1, 1})
	assertBoth([]int{9, 8, 7}, 1, []int{7})
	assertBoth([]int{1, 2, 3}, 0, []int{})
	mustPanic17(t, func() { SmallestKHeap([]int{1, 2}, 3) })
	mustPanic17(t, func() { SmallestKQuickselect([]int{1, 2}, 3) })
	mustPanic17(t, func() { SmallestKHeap([]int{1}, -1) })
	mustPanic17(t, func() { SmallestKQuickselect([]int{1}, -1) })
}

func TestLongestWord(t *testing.T) {
	got := LongestCompositeWord([]string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"})
	if got == nil || *got != "dogwalker" {
		t.Fatal(got)
	}
	if LongestCompositeWord([]string{"cat", "banana", "dog"}) != nil {
		t.Fatal("none")
	}
	got = LongestCompositeWord([]string{"cat", "banana", "dog", "nana", "catbanana", "walk", "walker", "dogwalker"})
	if got == nil || *got != "catbanana" {
		t.Fatal(got)
	}
	if LongestCompositeWord(nil) != nil {
		t.Fatal("empty")
	}
}

func TestMasseuse(t *testing.T) {
	if FindBestSchedule([]int{30, 15, 60, 75, 45, 15, 15, 45}) != 180 {
		t.Fatal("180a")
	}
	if FindBestSchedule([]int{30, 15, 60, 15, 45, 15, 45}) != 180 {
		t.Fatal("180b")
	}
	if FindBestSchedule([]int{30, 15, 15, 60}) != 90 {
		t.Fatal("90")
	}
	if FindBestSchedule(nil) != 0 || FindBestSchedule([]int{40}) != 40 || FindBestSchedule([]int{10, 20}) != 20 {
		t.Fatal("short")
	}
	mustPanic17(t, func() { FindBestSchedule([]int{1, -2}) })
}

func TestMultiSearch(t *testing.T) {
	found := MultiSearch("mississippi", []string{"i", "is", "pp", "ms"})
	if len(found) != 3 || !reflect.DeepEqual(found["i"], []int{1, 4, 7, 10}) || !reflect.DeepEqual(found["is"], []int{1, 4}) || !reflect.DeepEqual(found["pp"], []int{8}) {
		t.Fatal(found)
	}
	if _, ok := found["ms"]; ok {
		t.Fatal("ms")
	}
	if len(MultiSearch("abc", []string{"z"})) != 0 || len(MultiSearch("", []string{"a"})) != 0 {
		t.Fatal("empty")
	}
}

func TestMinWindow(t *testing.T) {
	w := MinWindowString("75902135791158897", "159")
	if w == nil || w.Left != 7 || w.Right != 10 {
		t.Fatal(w)
	}
	if MinWindowString("abc", "x") != nil {
		t.Fatal("null")
	}
	w = MinWindowInts([]int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, []int{1, 5, 9})
	if w == nil || w.Left != 7 || w.Right != 10 {
		t.Fatal(w)
	}
	mustPanic17(t, func() { MinWindowString("abc", "") })
	mustPanic17(t, func() { MinWindowInts([]int{1, 2, 3}, nil) })
}

func TestMissingTwo(t *testing.T) {
	assertBoth := func(arr []int, n, a, b int) {
		x1, y1 := MissingTwoSumSquares(arr, n)
		x2, y2 := MissingTwoXorBit(arr, n)
		if x1 != a || y1 != b || x2 != a || y2 != b {
			t.Fatal(arr, n, x1, y1, x2, y2)
		}
	}
	assertBoth([]int{0, 1}, 3, 2, 3)
	assertBoth([]int{1, 2, 3, 4, 5, 6, 7, 8}, 9, 0, 9)
	assertBoth([]int{0, 1, 2, 3, 5, 6, 7, 8}, 9, 4, 9)
	n := 12
	for a := 0; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			arr := make([]int, 0, n-1)
			for x := 0; x <= n; x++ {
				if x != a && x != b {
					arr = append(arr, x)
				}
			}
			assertBoth(arr, n, a, b)
		}
	}
	mustPanic17(t, func() { MissingTwoSumSquares([]int{0}, 1) })
	mustPanic17(t, func() { MissingTwoXorBit([]int{0}, 1) })
}

func TestMedian(t *testing.T) {
	assertMedians := func(stream []int, expected []float64) {
		actual := MediansAfterEach(stream)
		if len(actual) != len(expected) {
			t.Fatal(actual)
		}
		for i := range expected {
			if math.Abs(expected[i]-actual[i]) >= 1e-9 {
				t.Fatal(expected, actual)
			}
		}
	}
	assertMedians([]int{5}, []float64{5})
	assertMedians([]int{5, 15, 1}, []float64{5, 10, 5})
	assertMedians([]int{0, 1, 2, 3, 4, 5}, []float64{0, 0.5, 1, 1.5, 2, 2.5})
	assertMedians([]int{5, 15}, []float64{5, 10})
	mustPanic17(t, func() { NewMedianFinder().Median() })
}

func TestHistogram(t *testing.T) {
	if FindVolume([]int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 5, 0, 1, 0, 0, 0}) != 26 {
		t.Fatal("book")
	}
	if FindVolume(nil) != 0 || FindVolume([]int{1}) != 0 || FindVolume([]int{1, 2}) != 0 {
		t.Fatal("short")
	}
	if FindVolume([]int{2, 0, 2}) != 2 || FindVolume([]int{3, 0, 0, 2, 0, 4}) != 10 {
		t.Fatal("trap")
	}
	mustPanic17(t, func() { FindVolume([]int{1, -1, 1}) })
}

func TestWordTransformer(t *testing.T) {
	assertPath := func(source, target string, dict, expected []string) {
		actual := WordTransformer(source, target, dict)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatal(source, target, actual, expected)
		}
	}
	assertPath("bit", "dog", []string{"but", "put", "big", "pot", "pog", "dog", "lot"}, []string{"bit", "but", "put", "pot", "pog", "dog"})
	assertPath("damp", "like", []string{"damp", "lime", "limp", "lamp", "like"}, []string{"damp", "lamp", "limp", "lime", "like"})
	assertPath("a", "a", []string{"a"}, []string{"a"})
	if WordTransformer("hit", "cog", []string{"hot", "dot", "dog"}) != nil {
		t.Fatal("unreachable")
	}
	mustPanic17(t, func() { WordTransformer("a", "ab", []string{"a", "ab"}) })
}

func TestMaxBlackSquare(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 1},
	}
	result := MaxBlackSquare(matrix)
	if result == nil {
		t.Fatal("nil")
	}
	side := result.BottomRight.Row - result.TopLeft.Row + 1
	if side != result.BottomRight.Column-result.TopLeft.Column+1 || side != 2 {
		t.Fatal(result)
	}
	for row := result.TopLeft.Row; row <= result.BottomRight.Row; row++ {
		for col := result.TopLeft.Column; col <= result.BottomRight.Column; col++ {
			if matrix[row][col] != 1 {
				t.Fatal("not black")
			}
		}
	}
	if MaxBlackSquare([][]int{{0, 0}, {0, 0}}) != nil || MaxBlackSquare(nil) != nil || MaxBlackSquare([][]int{{}}) != nil {
		t.Fatal("empty")
	}
}

func TestMaxSubmatrix(t *testing.T) {
	if MaxSubmatrixSum([][]int{{1, 2, 3, 4, 5}, {2, 3, -5000, 5, 6}, {10, 20, 30, 40, 50}}) != 150 {
		t.Fatal("150")
	}
	if MaxSubmatrixSum([][]int{{1, -2, -1, 4}, {1, -1, 1, 1}, {0, -1, -1, 1}, {0, 0, 1, 1}}) != 7 {
		t.Fatal("7")
	}
	if MaxSubmatrixSum([][]int{{-1}}) != -1 {
		t.Fatal("-1")
	}
	mustPanic17(t, func() { MaxSubmatrixSum(nil) })
	mustPanic17(t, func() { MaxSubmatrixSum([][]int{{}}) })
}

func TestWordRectangle(t *testing.T) {
	words := []string{"area", "lead", "wall", "lady", "ball"}
	result := FindLargestWordRectangle(words)
	if result == nil || len(result[0]) != len(result) {
		t.Fatal(result)
	}
	set := map[string]struct{}{}
	for _, w := range words {
		set[w] = struct{}{}
	}
	for _, row := range result {
		if _, ok := set[row]; !ok {
			t.Fatal(row)
		}
	}
	side := len(result)
	for col := range side {
		colWord := make([]byte, side)
		for row := range side {
			colWord[row] = result[row][col]
		}
		if _, ok := set[string(colWord)]; !ok {
			t.Fatal(string(colWord))
		}
	}
	if FindLargestWordRectangle([]string{"ab", "cd", "ef"}) != nil || FindLargestWordRectangle(nil) != nil {
		t.Fatal("null")
	}
}

func TestSparseSimilarity(t *testing.T) {
	docs := []Document{
		{13, []int{14, 15, 100, 9, 3}},
		{16, []int{32, 1, 9, 3, 5}},
		{19, []int{15, 29, 2, 6, 8, 7}},
		{24, []int{7, 10}},
	}
	got := SparseSimilarity(docs)
	if len(got) != 3 {
		t.Fatal(got)
	}
	eq := func(a, b float64) bool { return math.Abs(a-b) < 1e-9 }
	if got[0].Id1 != 13 || got[0].Id2 != 16 || !eq(got[0].Similarity, 2.0/8) {
		t.Fatal(got[0])
	}
	if got[1].Id1 != 13 || got[1].Id2 != 19 || !eq(got[1].Similarity, 1.0/10) {
		t.Fatal(got[1])
	}
	if got[2].Id1 != 19 || got[2].Id2 != 24 || !eq(got[2].Similarity, 1.0/7) {
		t.Fatal(got[2])
	}
	a := map[int]int{1: 1, 2: 2}
	b := map[int]int{2: 2, 3: 1}
	if !eq(SparseDotSimilarity(a, b), 4/(math.Sqrt(5)*math.Sqrt(5))) {
		t.Fatal(SparseDotSimilarity(a, b))
	}
	if SparseDotSimilarity(map[int]int{1: 1}, map[int]int{}) != 0 {
		t.Fatal("zero")
	}
}
