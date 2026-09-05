package chapter03

import (
	"errors"
	"reflect"
	"testing"
)

func TestMultiStack(t *testing.T) {
	for _, tc := range []struct{ stacks, size int }{{3, 6}, {1, 1}, {2, 3}} {
		assertMultiStackRoundTrip(t, tc.stacks, tc.size)
	}
	stacks, err := NewMultiStack(3, 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := stacks.Push(1, 1); !errors.As(err, new(*StackDoesNotExistError)) {
		t.Fatalf("want StackDoesNotExistError, got %v", err)
	}
	if err := stacks.Push(1, -1); !errors.As(err, new(*StackDoesNotExistError)) {
		t.Fatalf("want StackDoesNotExistError, got %v", err)
	}
	if _, err := NewMultiStack(0, 1); err == nil {
		t.Fatal("expected error")
	}
	if _, err := NewMultiStack(-1, 1); err == nil {
		t.Fatal("expected error")
	}
	if _, err := NewMultiStack(1, 0); err == nil {
		t.Fatal("expected error")
	}
}

func assertMultiStackRoundTrip(t *testing.T, numberOfStacks, stackSize int) {
	t.Helper()
	stacks, err := NewMultiStack(stackSize, numberOfStacks)
	if err != nil {
		t.Fatal(err)
	}
	for stackNum := range numberOfStacks {
		if !stacks.IsEmpty(stackNum) || stacks.IsFull(stackNum) {
			t.Fatal("initial state")
		}
		if _, err := stacks.Pop(stackNum); !errors.As(err, new(*StackEmptyError)) {
			t.Fatalf("empty pop: %v", err)
		}
		for i := 0; i < stackSize-1; i++ {
			if err := stacks.Push(i, stackNum); err != nil {
				t.Fatal(err)
			}
			peek, err := stacks.Peek(stackNum)
			if err != nil || peek != i || stacks.IsEmpty(stackNum) || stacks.IsFull(stackNum) {
				t.Fatalf("peek %d: %d %v", i, peek, err)
			}
		}
		if err := stacks.Push(999, stackNum); err != nil {
			t.Fatal(err)
		}
		if err := stacks.Push(777, stackNum); !errors.As(err, new(*StackFullError)) {
			t.Fatalf("full push: %v", err)
		}
		if stacks.IsEmpty(stackNum) || !stacks.IsFull(stackNum) {
			t.Fatal("full state")
		}
		peek, _ := stacks.Peek(stackNum)
		popped, _ := stacks.Pop(stackNum)
		if peek != 999 || popped != 999 {
			t.Fatal(peek, popped)
		}
		for i := stackSize - 2; i >= 0; i-- {
			peek, _ = stacks.Peek(stackNum)
			popped, _ = stacks.Pop(stackNum)
			if peek != i || popped != i {
				t.Fatal(i, peek, popped)
			}
		}
		if !stacks.IsEmpty(stackNum) {
			t.Fatal("should be empty")
		}
		if _, err := stacks.Peek(stackNum); !errors.As(err, new(*StackEmptyError)) {
			t.Fatalf("empty peek: %v", err)
		}
		if _, err := stacks.Pop(stackNum); !errors.As(err, new(*StackEmptyError)) {
			t.Fatalf("empty pop: %v", err)
		}
	}
}

func TestMinStack(t *testing.T) {
	stack := &MinStack{}
	if stack.Minimum() != nil {
		t.Fatal("empty min")
	}
	if _, err := stack.Pop(); err == nil {
		t.Fatal("empty pop")
	}
	stack.Push(5)
	if *stack.Minimum() != 5 {
		t.Fatal(*stack.Minimum())
	}
	stack.Push(6)
	if *stack.Minimum() != 5 {
		t.Fatal(*stack.Minimum())
	}
	stack.Push(3)
	if *stack.Minimum() != 3 {
		t.Fatal(*stack.Minimum())
	}
	stack.Push(7)
	if *stack.Minimum() != 3 {
		t.Fatal(*stack.Minimum())
	}
	stack.Push(3)
	if *stack.Minimum() != 3 {
		t.Fatal(*stack.Minimum())
	}
	_, _ = stack.Pop()
	if *stack.Minimum() != 3 {
		t.Fatal(*stack.Minimum())
	}
	_, _ = stack.Pop()
	if *stack.Minimum() != 3 {
		t.Fatal(*stack.Minimum())
	}
	_, _ = stack.Pop()
	if *stack.Minimum() != 5 {
		t.Fatal(*stack.Minimum())
	}
	stack.Push(1)
	if *stack.Minimum() != 1 {
		t.Fatal(*stack.Minimum())
	}
	_, _ = stack.Pop()
	if *stack.Minimum() != 5 {
		t.Fatal(*stack.Minimum())
	}
}

func TestSetOfStacks(t *testing.T) {
	for _, capacity := range []int{5, 1, 7} {
		assertLifo(t, capacity)
	}
	stacks, err := NewSetOfStacks(5)
	if err != nil {
		t.Fatal(err)
	}
	for i := range 35 {
		stacks.Push(i)
	}
	result := make([]int, 31)
	for i := range result {
		result[i], err = stacks.PopAt(0)
		if err != nil {
			t.Fatal(err)
		}
	}
	expected := make([]int, 31)
	for i := range expected {
		expected[i] = i + 4
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("%v", result)
	}
	if _, err := NewSetOfStacks(0); err == nil {
		t.Fatal("capacity 0")
	}
	if _, err := NewSetOfStacks(-1); err == nil {
		t.Fatal("capacity -1")
	}
	small, _ := NewSetOfStacks(2)
	small.Push(1)
	if _, err := small.PopAt(1); err == nil {
		t.Fatal("popAt out of range")
	}
}

func assertLifo(t *testing.T, capacity int) {
	t.Helper()
	stacks, err := NewSetOfStacks(capacity)
	if err != nil {
		t.Fatal(err)
	}
	n := capacity * 7
	for i := range n {
		stacks.Push(i)
	}
	result := make([]int, n)
	for i := range n {
		value := stacks.Pop()
		if value == nil {
			t.Fatal("unexpected empty")
		}
		result[i] = *value
	}
	expected := make([]int, n)
	for i := range n {
		expected[i] = n - 1 - i
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("%v", result)
	}
	if !stacks.IsEmpty() || stacks.Pop() != nil {
		t.Fatal("expected empty")
	}
}

func TestMyQueue(t *testing.T) {
	for _, sequence := range [][]int{{1, 2, 3}, {-1, 0, 1}} {
		assertQueueCount(t, sequence)
		assertQueueFront(t, sequence)
		assertQueueShift(t, sequence)
		assertQueuePeek(t, sequence)
		assertQueueRemove(t, sequence)
	}
	for _, sequence := range [][]string{{"a", "b", "c", "d", "e", "f"}} {
		assertStringQueue(t, sequence)
	}
	queue := &MyQueue[int]{}
	if !queue.IsEmpty() {
		t.Fatal("empty")
	}
	if _, err := queue.Peek(); err == nil {
		t.Fatal("peek empty")
	}
	if _, err := queue.Remove(); err == nil {
		t.Fatal("remove empty")
	}
	queue.Add(4)
	queue.Add(6)
	peek, _ := queue.Peek()
	if peek != 4 {
		t.Fatal(peek)
	}
	queue.Add(101)
	peek, _ = queue.Peek()
	if peek != 4 {
		t.Fatal(peek)
	}
	a, _ := queue.Remove()
	b, _ := queue.Remove()
	c, _ := queue.Remove()
	if a != 4 || b != 6 || c != 101 || queue.Count() != 0 {
		t.Fatal(a, b, c, queue.Count())
	}
}

func assertQueueCount(t *testing.T, sequence []int) {
	t.Helper()
	queue := &MyQueue[int]{}
	for i, value := range sequence {
		queue.Add(value)
		if queue.Count() != i+1 {
			t.Fatal(queue.Count())
		}
	}
	for i := range sequence {
		_, _ = queue.Remove()
		if queue.Count() != len(sequence)-i-1 {
			t.Fatal(queue.Count())
		}
	}
}

func assertQueueFront(t *testing.T, sequence []int) {
	t.Helper()
	queue := &MyQueue[int]{}
	for _, value := range sequence {
		queue.Add(value)
	}
	peek, _ := queue.Peek()
	if peek != sequence[0] || queue.Count() != len(sequence) {
		t.Fatal(peek, queue.Count())
	}
}

func assertQueueShift(t *testing.T, sequence []int) {
	t.Helper()
	queue := &MyQueue[int]{}
	for _, value := range sequence {
		queue.Add(value)
	}
	if queue.OldStack.Count() != 0 || queue.NewStack.Count() != len(sequence) || queue.NewStack.Peek() != sequence[len(sequence)-1] {
		t.Fatal("before shift")
	}
	queue.ShiftStacks()
	if queue.OldStack.Count() != len(sequence) || queue.NewStack.Count() != 0 || queue.OldStack.Peek() != sequence[0] {
		t.Fatal("after shift")
	}
}

func assertQueuePeek(t *testing.T, sequence []int) {
	t.Helper()
	queue := &MyQueue[int]{}
	for _, value := range sequence {
		queue.Add(value)
		peek, _ := queue.Peek()
		if peek != sequence[0] {
			t.Fatal(peek)
		}
	}
	_, _ = queue.Remove()
	peek, _ := queue.Peek()
	if peek != sequence[1] {
		t.Fatal(peek)
	}
}

func assertQueueRemove(t *testing.T, sequence []int) {
	t.Helper()
	queue := &MyQueue[int]{}
	for _, value := range sequence {
		queue.Add(value)
	}
	for _, expected := range sequence {
		got, _ := queue.Remove()
		if got != expected {
			t.Fatal(got, expected)
		}
	}
}

func assertStringQueue(t *testing.T, sequence []string) {
	t.Helper()
	queue := &MyQueue[string]{}
	for i, value := range sequence {
		queue.Add(value)
		if queue.Count() != i+1 {
			t.Fatal(queue.Count())
		}
	}
	if queue.OldStack.Count() != 0 || queue.NewStack.Peek() != sequence[len(sequence)-1] {
		t.Fatal("before shift")
	}
	queue.ShiftStacks()
	if queue.OldStack.Peek() != sequence[0] {
		t.Fatal(queue.OldStack.Peek())
	}
	peek, _ := queue.Peek()
	if peek != sequence[0] {
		t.Fatal(peek)
	}
	_, _ = queue.Remove()
	peek, _ = queue.Peek()
	if peek != sequence[1] {
		t.Fatal(peek)
	}
	rest := []string{sequence[1]}
	rest = append(rest, sequence[2:]...)
	for _, expected := range rest {
		got, _ := queue.Remove()
		if got != expected {
			t.Fatal(got, expected)
		}
	}
}

func TestSortedStack(t *testing.T) {
	cases := []struct {
		values   []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1, 4}, []int{1, 2, 3, 4}},
		{[]int{5, 5, 1}, []int{1, 5, 5}},
	}
	for _, tc := range cases {
		stack := &SortedStack{}
		for _, value := range tc.values {
			stack.Push(value)
		}
		if stack.Count() != len(tc.values) {
			t.Fatal(stack.Count())
		}
		actual := make([]int, len(tc.values))
		for i := range actual {
			var err error
			actual[i], err = stack.Pop()
			if err != nil {
				t.Fatal(err)
			}
		}
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("%v -> %v", tc.values, actual)
		}
	}
	if _, err := (&SortedStack{}).Pop(); err == nil {
		t.Fatal("empty pop")
	}
}

func TestAnimalShelter(t *testing.T) {
	shelter := &AnimalShelter{}
	fluffy, _ := NewCat("Fluffy")
	sparky, _ := NewDog("Sparky")
	sneezy, _ := NewCat("Sneezy")
	shelter.Enqueue(fluffy)
	shelter.Enqueue(sparky)
	shelter.Enqueue(sneezy)
	if shelter.Count() != 3 {
		t.Fatal(shelter.Count())
	}

	oldest := &AnimalShelter{}
	oldest.Enqueue(fluffy)
	oldest.Enqueue(sparky)
	animal := oldest.DequeueAny()
	cat, ok := animal.(*Cat)
	if !ok || cat.Name != "Fluffy" || oldest.Count() != 1 {
		t.Fatal(animal, oldest.Count())
	}

	dogsFirst := &AnimalShelter{}
	rex, _ := NewDog("Rex")
	mittens, _ := NewCat("Mittens")
	dogsFirst.Enqueue(rex)
	dogsFirst.Enqueue(mittens)
	dog := dogsFirst.DequeueDog()
	if dog == nil || dog.Name != "Rex" || dogsFirst.Count() != 1 {
		t.Fatal(dog, dogsFirst.Count())
	}

	skip := &AnimalShelter{}
	fido, _ := NewDog("Fido")
	skip.Enqueue(rex)
	skip.Enqueue(mittens)
	skip.Enqueue(fido)
	gotCat := skip.DequeueCat()
	if gotCat == nil || gotCat.Name != "Mittens" || skip.Count() != 2 {
		t.Fatal(gotCat, skip.Count())
	}

	onlyDog := &AnimalShelter{}
	onlyDog.Enqueue(rex)
	if onlyDog.DequeueCat() != nil {
		t.Fatal("expected nil cat")
	}
	if _, err := NewCat(""); err == nil {
		t.Fatal("empty name")
	}
}
