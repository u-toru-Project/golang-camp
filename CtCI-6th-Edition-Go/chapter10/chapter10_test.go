package chapter10

import (
	"reflect"
	"sort"
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

func TestMergeAnagramsRotatedListy(t *testing.T) {
	assertMerged(t, []int{2, 3, 4, 5, 6, 8, 10, 100}, []int{1, 4, 6, 7, 7, 7}, []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 7, 7, 8, 10, 100})
	assertMerged(t, []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6})
	assertMerged(t, []int{1, 2, 3}, nil, []int{1, 2, 3})
	assertMerged(t, nil, []int{4, 5, 6}, []int{4, 5, 6})
	assertMerged(t, nil, nil, nil)
	assertMerged(t, []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6})
	assertMerged(t, []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6})
	assertMerged(t, []int{1, 1, 3}, []int{1, 2, 2}, []int{1, 1, 1, 2, 2, 3})

	assertGrouped(t, []string{"apple", "banana", "carrot", "ele", "duck", "papel", "tarroc", "cudk", "eel", "lee"})
	assertGrouped(t, []string{"dog", "cat", "bird"})
	assertGrouped(t, []string{"abc", "cab", "bac", "cba"})
	assertGrouped(t, nil)
	assertGrouped(t, []string{"only"})
	words := []string{"ab", "BA", "ba"}
	GroupAnagrams(words)
	assertAnagramsGrouped(t, words)
	if sortedKey("ab") == sortedKey("BA") {
		t.Fatal("case")
	}

	values := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	if SearchRotated(values, 8) != 3 || SearchRotated(values, 1) != 5 || SearchRotated(values, 5) != 0 {
		t.Fatal("rotated")
	}
	if SearchRotated(values, 4) != 8 || SearchRotated(values, 9) != 4 {
		t.Fatal("ends")
	}
	if SearchRotated([]int{1, 2, 3, 4, 5}, 1) != 0 || SearchRotated([]int{1, 2, 3, 4, 5}, 3) != 2 || SearchRotated([]int{1, 2, 3, 4, 5}, 5) != 4 {
		t.Fatal("sorted")
	}
	if SearchRotated(values, 10) != -1 || SearchRotated(values, 0) != -1 {
		t.Fatal("miss")
	}
	dups := []int{2, 2, 2, 3, 4, 2}
	if SearchRotated(dups, 3) != 3 || SearchRotated(dups, 4) != 4 || SearchRotated(dups, 2) < 0 || SearchRotated(dups, 5) != -1 {
		t.Fatal("dups")
	}
	if SearchRotated([]int{2, 2, 2, 2}, 2) < 0 || SearchRotated([]int{2, 2, 2, 2}, 1) != -1 {
		t.Fatal("all same")
	}
	if SearchRotated([]int{7}, 7) != 0 || SearchRotated([]int{7}, 8) != -1 || SearchRotated(nil, 1) != -1 {
		t.Fatal("tiny")
	}

	listy := NewListy(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for value := 1; value <= 9; value++ {
		if SearchListy(listy, value) != value-1 {
			t.Fatal(value)
		}
	}
	if SearchListy(listy, 0) != -1 || SearchListy(listy, 10) != -1 || SearchListy(NewListy(), 1) != -1 {
		t.Fatal("listy miss")
	}
	short := NewListy(1, 2, 3)
	if short.ElementAt(0) != 1 || short.ElementAt(2) != 3 || short.ElementAt(3) != -1 || short.ElementAt(100) != -1 || short.ElementAt(-1) != -1 {
		t.Fatal("elementAt")
	}
}

func TestSparseSortMissingDups(t *testing.T) {
	strings := []string{"apple", "", "", "banana", "", "", "", "carrot", "duck", "", "", "eel", "", "flower"}
	assertBoth(t, strings, "apple", 0)
	assertBoth(t, strings, "banana", 3)
	assertBoth(t, strings, "carrot", 7)
	assertBoth(t, strings, "duck", 8)
	assertBoth(t, strings, "eel", 11)
	assertBoth(t, strings, "flower", 13)
	assertBoth(t, []string{"apple", "", "", "banana", "", "carrot"}, "missing", -1)
	assertBoth(t, []string{"apple", "", "", "banana", "", "carrot"}, "", -1)
	assertBoth(t, nil, "apple", -1)
	assertBoth(t, []string{"", "", "ball", "", "cat", ""}, "ball", 2)
	assertBoth(t, []string{"", "", "ball", "", "cat", ""}, "cat", 4)
	assertBoth(t, []string{"", "", "ball", "", "cat", ""}, "apple", -1)
	assertBoth(t, []string{"", "", "ball", "", "cat", ""}, "dog", -1)
	assertBoth(t, []string{"only"}, "only", 0)
	assertBoth(t, []string{"only"}, "nope", -1)
	assertBoth(t, []string{"", "", ""}, "a", -1)
	assertBoth(t, []string{}, "a", -1)

	assertSorted(t, []int{})
	assertSorted(t, []int{1})
	assertSorted(t, []int{3, 1, 2})
	assertSorted(t, []int{5, 4, 3, 2, 1, 0})
	assertSorted(t, []int{10, -1, 7, 7, 2})
	if !reflect.DeepEqual(ExternalMergeSort([]int{3, 1, 2}, 1), []int{1, 2, 3}) {
		t.Fatal("chunk 1")
	}
	mustPanic(t, func() { ExternalMergeSort([]int{1, 2}, 0) })

	assertMissing(t, []int{1, 2, 3}, 0)
	assertMissing(t, []int{0, 1, 3}, 2)
	assertMissing(t, []int{0, 2, 3, 4, 5}, 1)
	nums := make([]int, 0, 999)
	for i := range 1000 {
		if i != 777 {
			nums = append(nums, i)
		}
	}
	if MissingIntBitVector(nums) != 777 || MissingIntSum(nums) != 777 {
		t.Fatal("777")
	}
	mustPanic(t, func() { MissingIntBitVector(nil) })
	mustPanic(t, func() { MissingIntBitVector([]int{}) })
	mustPanic(t, func() { MissingIntBitVector([]int{0, 0}) })
	mustPanic(t, func() { MissingIntBitVector([]int{0, 5}) })

	if !reflect.DeepEqual(FindDuplicates([]int{1, 2, 3, 2, 4, 3, 3}), []int{2, 3, 3}) {
		t.Fatal("dups")
	}
	if !reflect.DeepEqual(FindDuplicates([]int{5, 5}), []int{5}) {
		t.Fatal("5")
	}
	if len(FindDuplicates([]int{1, 2, 3, 4})) != 0 || len(FindDuplicates([]int{7})) != 0 || len(FindDuplicates(nil)) != 0 {
		t.Fatal("none")
	}
	if !reflect.DeepEqual(FindDuplicates([]int{1, MaxValue, 1, MaxValue}), []int{1, MaxValue}) {
		t.Fatal("max")
	}
	bits := NewBitSet(MaxValue)
	if bits.Get(0) || bits.Get(31) || bits.Get(32) || bits.Get(MaxValue-1) {
		t.Fatal("unset")
	}
	bits.Set(0)
	bits.Set(31)
	bits.Set(32)
	bits.Set(MaxValue - 1)
	if !bits.Get(0) || !bits.Get(31) || !bits.Get(32) || !bits.Get(MaxValue-1) || bits.Get(1) {
		t.Fatal("set")
	}
}

func TestMatrixRankPeaks(t *testing.T) {
	book := [][]int{
		{15, 30, 50, 70, 73},
		{35, 40, 100, 102, 120},
		{36, 42, 105, 110, 125},
		{46, 51, 106, 111, 130},
		{48, 55, 109, 140, 150},
	}
	for _, target := range []int{15, 40, 110, 150, 42, 73} {
		assertFound(t, book, target)
	}
	for _, target := range []int{16, 41, 0, 151} {
		assertMissingMatrix(t, book, target)
	}
	small := [][]int{{1, 2}, {3, 4}}
	assertFound(t, small, 1)
	assertFound(t, small, 2)
	assertFound(t, small, 3)
	assertFound(t, small, 4)
	assertMissingMatrix(t, small, 5)
	assertFound(t, [][]int{{7}}, 7)
	assertMissingMatrix(t, [][]int{{7}}, 8)
	assertMissingMatrix(t, nil, 1)
	assertMissingMatrix(t, [][]int{{}}, 1)
	coordinate := FindElement2(book, 110)
	if coordinate == nil || coordinate.Row != 2 || coordinate.Column != 3 {
		t.Fatal(coordinate)
	}
	if FindElement2(book, 16) != nil {
		t.Fatal("16")
	}

	tracker := track(5, 1, 4, 4, 5, 9, 7, 13, 3)
	if tracker.GetRankOfNumber(1) != 0 || tracker.GetRankOfNumber(3) != 1 || tracker.GetRankOfNumber(4) != 3 {
		t.Fatal("rank")
	}
	if tracker.GetRankOfNumber(5) != 5 || tracker.GetRankOfNumber(7) != 6 || tracker.GetRankOfNumber(9) != 7 || tracker.GetRankOfNumber(13) != 8 {
		t.Fatal("rank more")
	}
	if NewRankTracker().GetRankOfNumber(1) != -1 || track(5, 1, 4).GetRankOfNumber(2) != -1 || track(5, 1, 4).GetRankOfNumber(10) != -1 {
		t.Fatal("missing rank")
	}
	if track(2, 2, 2).GetRankOfNumber(2) != 2 {
		t.Fatal("ties")
	}
	first := track(1, 2)
	second := track(10)
	if first.GetRankOfNumber(1) != 0 || first.GetRankOfNumber(2) != 1 || first.GetRankOfNumber(10) != -1 {
		t.Fatal("first")
	}
	if second.GetRankOfNumber(10) != 0 || second.GetRankOfNumber(1) != -1 {
		t.Fatal("second")
	}
	if track(1, 2, 3, 4).GetRankOfNumber(1) != 0 || track(1, 2, 3, 4).GetRankOfNumber(4) != 3 {
		t.Fatal("asc")
	}
	if track(4, 3, 2, 1).GetRankOfNumber(1) != 0 || track(4, 3, 2, 1).GetRankOfNumber(4) != 3 {
		t.Fatal("desc")
	}

	assertAll(t, []int{48, 40, 31, 62, 28, 21, 64, 40, 23, 17})
	assertAll(t, []int{1, 2, 3, 4, 5})
	assertAll(t, []int{5, 4, 3, 2, 1})
	assertAll(t, nil)
	assertAll(t, []int{7})
	assertAll(t, []int{2, 1})
	assertAll(t, []int{1, 2, 3})
	assertAll(t, []int{5, 5, 5, 5})
	assertAll(t, []int{1, 1, 2, 2, 3})
}

func assertMerged(t *testing.T, aValues, b, expected []int) {
	t.Helper()
	a := append(append([]int{}, aValues...), make([]int, len(b))...)
	Merge(a, b, len(aValues), len(b))
	if !reflect.DeepEqual(a, expected) && !(len(a) == 0 && len(expected) == 0) {
		t.Fatalf("got %v want %v", a, expected)
	}
}

func assertGrouped(t *testing.T, words []string) {
	t.Helper()
	hashed := append([]string{}, words...)
	sorted := append([]string{}, words...)
	GroupAnagrams(hashed)
	SortByAnagram(sorted)
	assertSameMultiset(t, words, hashed)
	assertSameMultiset(t, words, sorted)
	assertAnagramsGrouped(t, hashed)
	assertAnagramsGrouped(t, sorted)
}

func assertAnagramsGrouped(t *testing.T, words []string) {
	t.Helper()
	seen := map[string]struct{}{}
	currentKey := ""
	hasCurrent := false
	for _, word := range words {
		key := sortedKey(word)
		if hasCurrent && key == currentKey {
			continue
		}
		if _, ok := seen[key]; ok {
			t.Fatalf("group %s not contiguous", key)
		}
		seen[key] = struct{}{}
		currentKey = key
		hasCurrent = true
	}
}

func assertSameMultiset(t *testing.T, expected, actual []string) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatalf("len %d %d", len(expected), len(actual))
	}
	remaining := append([]string{}, expected...)
	for _, value := range actual {
		found := -1
		for i, v := range remaining {
			if v == value {
				found = i
				break
			}
		}
		if found < 0 {
			t.Fatalf("unexpected %s", value)
		}
		remaining = append(remaining[:found], remaining[found+1:]...)
	}
}

func assertSameIntMultiset(t *testing.T, expected, actual []int) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatalf("len %d %d", len(expected), len(actual))
	}
	remaining := append([]int{}, expected...)
	for _, value := range actual {
		found := -1
		for i, v := range remaining {
			if v == value {
				found = i
				break
			}
		}
		if found < 0 {
			t.Fatalf("unexpected %d", value)
		}
		remaining = append(remaining[:found], remaining[found+1:]...)
	}
}

func assertBoth(t *testing.T, strings []string, value string, expected int) {
	t.Helper()
	if SearchSparse(strings, value) != expected || SearchSparseIterative(strings, value) != expected {
		t.Fatalf("%s got %d/%d want %d", value, SearchSparse(strings, value), SearchSparseIterative(strings, value), expected)
	}
}

func assertSorted(t *testing.T, data []int) {
	t.Helper()
	expected := append([]int{}, data...)
	sort.Ints(expected)
	got := ExternalMergeSort(data, 2)
	if !reflect.DeepEqual(got, expected) && !(len(got) == 0 && len(expected) == 0) {
		t.Fatalf("got %v want %v", got, expected)
	}
}

func assertMissing(t *testing.T, nums []int, expected int) {
	t.Helper()
	if MissingIntBitVector(nums) != expected || MissingIntSum(nums) != expected {
		t.Fatal(nums)
	}
}

func assertFound(t *testing.T, matrix [][]int, target int) {
	t.Helper()
	if !FindElement(matrix, target) {
		t.Fatal(target)
	}
	coordinate := FindElement2(matrix, target)
	if coordinate == nil || matrix[coordinate.Row][coordinate.Column] != target {
		t.Fatal(target, coordinate)
	}
}

func assertMissingMatrix(t *testing.T, matrix [][]int, target int) {
	t.Helper()
	if FindElement(matrix, target) || FindElement2(matrix, target) != nil {
		t.Fatal(target)
	}
}

func track(numbers ...int) *RankTracker {
	tracker := NewRankTracker()
	for _, number := range numbers {
		tracker.Track(number)
	}
	return tracker
}

func assertAll(t *testing.T, original []int) {
	t.Helper()
	assertPeaksAndValleys(t, original, SortValleyPeak)
	assertPeaksAndValleys(t, original, SortValleyPeak2)
	assertPeaksAndValleys(t, original, SortValleyPeak3)
}

func assertPeaksAndValleys(t *testing.T, original []int, sortFn func([]int)) {
	t.Helper()
	values := append([]int{}, original...)
	sortFn(values)
	assertSameIntMultiset(t, original, values)
	if !isPeaksAndValleys(values) {
		t.Fatalf("not peaks and valleys: %v", values)
	}
}

func isPeaksAndValleys(values []int) bool {
	for i := 1; i < len(values)-1; i++ {
		isPeak := values[i] >= values[i-1] && values[i] >= values[i+1]
		isValley := values[i] <= values[i-1] && values[i] <= values[i+1]
		if !isPeak && !isValley {
			return false
		}
	}
	return true
}
