package chapter15

import "fmt"

type Enumerator interface {
	MoveNext() bool
}

type yieldTask struct {
	remaining int
}

func (t *yieldTask) MoveNext() bool {
	if t.remaining <= 0 {
		return false
	}
	t.remaining--
	return true
}

func TaskA() Enumerator {
	return &yieldTask{remaining: 2}
}

func TaskB() Enumerator {
	return &yieldTask{remaining: 1}
}

func CooperativeScheduler(tasks []Enumerator) []string {
	type item struct {
		name string
		gen  Enumerator
	}
	queue := make([]item, 0, len(tasks))
	for i, task := range tasks {
		queue = append(queue, item{name: fmt.Sprintf("t%d", i), gen: task})
	}
	log := make([]string, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.gen.MoveNext() {
			log = append(log, cur.name)
			queue = append(queue, cur)
		} else {
			log = append(log, cur.name+":done")
		}
	}
	return log
}

func RunQ1502() {
	log := CooperativeScheduler([]Enumerator{TaskA(), TaskB()})
	fmt.Println(log)
}
