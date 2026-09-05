package chapter07

import "fmt"

type Rank int

const (
	Respondent Rank = 1
	Manager    Rank = 2
	Director   Rank = 3
)

type Employee struct {
	Name string
	Rank Rank
	Busy bool
}

func NewEmployee(name string, rank Rank) *Employee {
	if name == "" {
		panic("name must be non-empty")
	}
	return &Employee{Name: name, Rank: rank}
}

func (e *Employee) TakeCall() {
	if e.Busy {
		panic("employee is busy")
	}
	e.Busy = true
}

func (e *Employee) FinishCall() { e.Busy = false }

type Call struct {
	Caller  string
	Rank    Rank
	Handler *Employee
}

func NewCall(caller string, rank Rank) *Call {
	if caller == "" {
		panic("caller must be non-empty")
	}
	return &Call{Caller: caller, Rank: rank}
}

type CallCenter struct {
	employees []*Employee
}

func NewCallCenter(employees []*Employee) *CallCenter {
	sorted := append([]*Employee{}, employees...)
	for i := range sorted {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].Rank < sorted[i].Rank {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return &CallCenter{employees: sorted}
}

func (c *CallCenter) Dispatch(call *Call) *Employee {
	if call == nil {
		panic("call is nil")
	}
	for _, employee := range c.employees {
		if employee.Rank >= call.Rank && !employee.Busy {
			employee.TakeCall()
			call.Handler = employee
			return employee
		}
	}
	return nil
}

func (c *CallCenter) Release(employee *Employee) {
	if employee == nil {
		panic("employee is nil")
	}
	employee.FinishCall()
}

func RunQ702() {
	center := NewCallCenter([]*Employee{
		NewEmployee("Alice", Respondent),
		NewEmployee("Bob", Manager),
	})
	handler := center.Dispatch(NewCall("Jane", Respondent))
	fmt.Printf("Handled by %s\n", handler.Name)
}
