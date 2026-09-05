package chapter14

import (
	"fmt"
	"sort"
)

type EmployeeSalaryStore struct {
	salaries []float64
}

func NewEmployeeSalaryStore() *EmployeeSalaryStore {
	return &EmployeeSalaryStore{}
}

func (s *EmployeeSalaryStore) AddSalary(salary float64) {
	s.salaries = append(s.salaries, salary)
}

func (s *EmployeeSalaryStore) MedianSalary() *float64 {
	n := len(s.salaries)
	if n == 0 {
		return nil
	}
	sorted := append([]float64(nil), s.salaries...)
	sort.Float64s(sorted)
	offset := (n - 1) / 2
	value := sorted[offset]
	return &value
}

func RunQ1403() {
	store := NewEmployeeSalaryStore()
	for _, salary := range []float64{30000.0, 50000.0, 70000.0, 90000.0} {
		store.AddSalary(salary)
	}
	fmt.Printf("median=%v\n", *store.MedianSalary())
}
