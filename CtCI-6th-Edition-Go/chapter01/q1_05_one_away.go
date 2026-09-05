package chapter01

import "fmt"

// OneEditAway reports whether the strings differ by at most one edit.
// Time O(n), space O(1).
func OneEditAway(first, second string) bool {
	if len(first) == len(second) {
		return isOneReplaceAway(first, second)
	}
	if len(first)+1 == len(second) {
		return isOneInsertAway(first, second)
	}
	if len(first)-1 == len(second) {
		return isOneInsertAway(second, first)
	}
	return false
}

// OneEditAway2 is the single-pass variant.
func OneEditAway2(first, second string) bool {
	if abs(len(first)-len(second)) > 1 {
		return false
	}
	shorter, longer := first, second
	if len(first) > len(second) {
		shorter, longer = second, first
	}
	si, li := 0, 0
	found := false
	for li < len(longer) && si < len(shorter) {
		if shorter[si] != longer[li] {
			if found {
				return false
			}
			found = true
			if len(shorter) == len(longer) {
				si++
			}
		} else {
			si++
		}
		li++
	}
	return true
}

func isOneReplaceAway(first, second string) bool {
	found := false
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			continue
		}
		if found {
			return false
		}
		found = true
	}
	return true
}

func isOneInsertAway(shorter, longer string) bool {
	si, li := 0, 0
	for li < len(longer) && si < len(shorter) {
		if shorter[si] == longer[li] {
			si++
			li++
			continue
		}
		if si != li {
			return false
		}
		li++
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func RunQ105() {
	pairs := [][2]string{{"pale", "ple"}, {"pale", "bale"}, {"pale", "bake"}}
	for _, pair := range pairs {
		fmt.Printf("%s, %s: %t / %t\n", pair[0], pair[1], OneEditAway(pair[0], pair[1]), OneEditAway2(pair[0], pair[1]))
	}
}
