package chapter10

import "fmt"

func Merge(a, b []int, lastA, lastB int) {
	indexA := lastA - 1
	indexB := lastB - 1
	indexMerged := lastA + lastB - 1
	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA]
			indexA--
		} else {
			a[indexMerged] = b[indexB]
			indexB--
		}
		indexMerged--
	}
}

func RunQ1001() {
	a := []int{2, 3, 4, 5, 6, 8, 10, 100, 0, 0, 0, 0, 0, 0}
	b := []int{1, 4, 6, 7, 7, 7}
	Merge(a, b, 8, 6)
	fmt.Println(a)
}
