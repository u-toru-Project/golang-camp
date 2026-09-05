package chapter02

import (
	"fmt"

	"ctci/library"
)

// IsPalindromeWithStack reports whether the list is a palindrome.
// Time O(n), space O(n).
func IsPalindromeWithStack(head *library.LinkedListNode) bool {
	fast, slow := head, head
	stack := make([]int, 0)
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Data)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != slow.Data {
			return false
		}
		slow = slow.Next
	}
	return true
}

type recurseResult struct {
	node         *library.LinkedListNode
	isPalindrome bool
}

// IsPalindromeRecursive reports whether the list is a palindrome.
// Time O(n), space O(n).
func IsPalindromeRecursive(head *library.LinkedListNode) bool {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	return isPalindromeRecursive(head, length).isPalindrome
}

func isPalindromeRecursive(head *library.LinkedListNode, length int) recurseResult {
	if head == nil || length <= 0 {
		return recurseResult{node: head, isPalindrome: true}
	}
	if length == 1 {
		return recurseResult{node: head.Next, isPalindrome: true}
	}
	result := isPalindromeRecursive(head.Next, length-2)
	if !result.isPalindrome || result.node == nil {
		return result
	}
	return recurseResult{
		node:         result.node.Next,
		isPalindrome: head.Data == result.node.Data,
	}
}

func RunQ206() {
	palindrome := library.CreateLinkedListFromArray([]int{0, 1, 2, 1, 0})
	notPalindrome := library.CreateLinkedListFromArray([]int{0, 1, 2, 3, 4})
	fmt.Printf("%s: %t / %t\n", palindrome.PrintForward(), IsPalindromeWithStack(palindrome), IsPalindromeRecursive(palindrome))
	fmt.Printf("%s: %t / %t\n", notPalindrome.PrintForward(), IsPalindromeWithStack(notPalindrome), IsPalindromeRecursive(notPalindrome))
}
