package chapter03

import "fmt"

type Animal struct {
	Name string
}

type Cat struct{ Animal }

type Dog struct{ Animal }

func NewCat(name string) (*Cat, error) {
	if name == "" {
		return nil, fmt.Errorf("name must be non-empty")
	}
	return &Cat{Animal{Name: name}}, nil
}

func NewDog(name string) (*Dog, error) {
	if name == "" {
		return nil, fmt.Errorf("name must be non-empty")
	}
	return &Dog{Animal{Name: name}}, nil
}

type animalNode struct {
	data any
	next *animalNode
}

type AnimalShelter struct {
	head  *animalNode
	tail  *animalNode
	count int
}

func (s *AnimalShelter) Count() int { return s.count }

func (s *AnimalShelter) Enqueue(animal any) {
	node := &animalNode{data: animal}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
	s.count++
}

func (s *AnimalShelter) DequeueAny() any {
	if s.head == nil {
		return nil
	}
	animal := s.head.data
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	}
	s.count--
	return animal
}

func (s *AnimalShelter) DequeueCat() *Cat {
	if animal, ok := s.dequeueType(func(data any) bool {
		_, ok := data.(*Cat)
		return ok
	}); ok {
		return animal.(*Cat)
	}
	return nil
}

func (s *AnimalShelter) DequeueDog() *Dog {
	if animal, ok := s.dequeueType(func(data any) bool {
		_, ok := data.(*Dog)
		return ok
	}); ok {
		return animal.(*Dog)
	}
	return nil
}

func (s *AnimalShelter) dequeueType(match func(any) bool) (any, bool) {
	var previous *animalNode
	current := s.head
	for current != nil {
		if match(current.data) {
			if previous == nil {
				s.head = current.next
			} else {
				previous.next = current.next
			}
			if current == s.tail {
				s.tail = previous
			}
			s.count--
			return current.data, true
		}
		previous = current
		current = current.next
	}
	return nil, false
}

func RunQ306() {
	shelter := &AnimalShelter{}
	cat, _ := NewCat("Fluffy")
	dog, _ := NewDog("Sparky")
	sneezy, _ := NewCat("Sneezy")
	shelter.Enqueue(cat)
	shelter.Enqueue(dog)
	shelter.Enqueue(sneezy)
	anyAnimal := shelter.DequeueAny().(*Cat)
	fmt.Printf("any=%s, dog=%s, count=%d\n", anyAnimal.Name, shelter.DequeueDog().Name, shelter.Count())
}
