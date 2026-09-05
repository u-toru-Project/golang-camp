package chapter14

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type enrollment struct {
	studentID int
	courseID  int
	grade     *float64
}

type GradeDatabase struct {
	students    map[int]string
	courses     map[int]string
	enrollments []enrollment
}

func NewGradeDatabase() *GradeDatabase {
	return &GradeDatabase{
		students: make(map[int]string),
		courses:  make(map[int]string),
	}
}

func (s *GradeDatabase) AddStudent(id int, name string) {
	s.students[id] = name
}

func (s *GradeDatabase) AddCourse(id int, title string) {
	s.courses[id] = title
}

func (s *GradeDatabase) Enroll(studentID, courseID int, grade *float64) {
	s.enrollments = append(s.enrollments, enrollment{studentID: studentID, courseID: courseID, grade: grade})
}

func (s *GradeDatabase) GpaForStudent(studentID int) *float64 {
	sum := 0.0
	count := 0
	for _, e := range s.enrollments {
		if e.studentID == studentID && e.grade != nil {
			sum += *e.grade
			count++
		}
	}
	if count == 0 {
		return nil
	}
	value := roundAwayFromZero(sum/float64(count), 2)
	return &value
}

func (s *GradeDatabase) StudentsBelowGpa(threshold float64) []string {
	type agg struct {
		sum   float64
		count int
	}
	byStudent := make(map[int]*agg)
	for _, e := range s.enrollments {
		if e.grade == nil {
			continue
		}
		a := byStudent[e.studentID]
		if a == nil {
			a = &agg{}
			byStudent[e.studentID] = a
		}
		a.sum += *e.grade
		a.count++
	}
	names := make([]string, 0)
	for id, a := range byStudent {
		if a.count > 0 && a.sum/float64(a.count) < threshold {
			names = append(names, s.students[id])
		}
	}
	sort.Strings(names)
	return names
}

func roundAwayFromZero(value float64, digits int) float64 {
	factor := math.Pow(10, float64(digits))
	if value >= 0 {
		return math.Floor(value*factor+0.5) / factor
	}
	return math.Ceil(value*factor-0.5) / factor
}

func RunQ1407() {
	store := NewGradeDatabase()
	store.AddStudent(1, "Ann")
	store.AddCourse(1, "Math")
	grade := 3.5
	store.Enroll(1, 1, &grade)
	fmt.Printf("gpa=%v below4=%s\n", *store.GpaForStudent(1), strings.Join(store.StudentsBelowGpa(4.0), ","))
}
