package chapter14

import "testing"

func TestApartmentsInBuilding(t *testing.T) {
	store := NewApartmentStore()
	towerA := store.AddBuilding("Tower A")
	towerB := store.AddBuilding("Tower B")
	store.AddApartment(towerA, "101", 1200.0)
	store.AddApartment(towerA, "102", 1300.0)
	store.AddApartment(towerB, "201", 900.0)
	units := store.ApartmentsInBuilding(towerA)
	if stringsJoin(units, ",") != "101,102" {
		t.Fatal(units)
	}
}

func TestUsersSpeaking(t *testing.T) {
	store := NewLanguageStore()
	store.AddUser(1, "Alice")
	store.AddUser(2, "Bob")
	store.SetLanguages(1, []string{"en", "fr"})
	store.SetLanguages(2, []string{"en"})
	if stringsJoin(store.UsersSpeaking("en"), ",") != "Alice,Bob" {
		t.Fatal(store.UsersSpeaking("en"))
	}
	if stringsJoin(store.UsersSpeaking("fr"), ",") != "Alice" {
		t.Fatal(store.UsersSpeaking("fr"))
	}
}

func TestMedianSalary(t *testing.T) {
	store := NewEmployeeSalaryStore()
	for _, salary := range []float64{30000.0, 50000.0, 70000.0, 90000.0} {
		store.AddSalary(salary)
	}
	if store.MedianSalary() == nil || *store.MedianSalary() != 50000.0 {
		t.Fatal(store.MedianSalary())
	}
	empty := NewEmployeeSalaryStore()
	if empty.MedianSalary() != nil {
		t.Fatal("expected nil")
	}
}

func TestRankOf(t *testing.T) {
	store := NewScoreboardStore()
	store.UpsertScore("a", 100)
	store.UpsertScore("b", 200)
	store.UpsertScore("c", 150)
	if *store.RankOf("b") != 1 || *store.RankOf("c") != 2 {
		t.Fatal("ranks")
	}
	if store.RankOf("missing") != nil {
		t.Fatal("unknown player")
	}
}

func TestSearchDocuments(t *testing.T) {
	store := NewDocumentSearchStore()
	store.AddDocument("Python Tips", "use pytest")
	store.AddDocument("SQL Guide", "joins and indexes")
	hits := store.SearchDocuments("pytest")
	if len(hits) != 1 || hits[0].Id != 1 || hits[0].Title != "Python Tips" {
		t.Fatal(hits)
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		store.SearchDocuments("")
	}()
}

func TestAdjustInventory(t *testing.T) {
	store := NewInventoryStore()
	if store.AdjustInventory("A1", 10) != 10 || store.AdjustInventory("A1", -3) != 7 {
		t.Fatal("qty")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		store.AdjustInventory("A1", -100)
	}()
}

func TestGradeDatabase(t *testing.T) {
	store := NewGradeDatabase()
	store.AddStudent(1, "Ann")
	store.AddCourse(1, "Math")
	grade := 3.5
	store.Enroll(1, 1, &grade)
	if store.GpaForStudent(1) == nil || *store.GpaForStudent(1) != 3.5 {
		t.Fatal(store.GpaForStudent(1))
	}
	if stringsJoin(store.StudentsBelowGpa(4.0), ",") != "Ann" {
		t.Fatal(store.StudentsBelowGpa(4.0))
	}
	empty := NewGradeDatabase()
	empty.AddStudent(1, "Ann")
	if empty.GpaForStudent(1) != nil {
		t.Fatal("expected nil gpa")
	}
}

func stringsJoin(values []string, sep string) string {
	if len(values) == 0 {
		return ""
	}
	out := values[0]
	for i := 1; i < len(values); i++ {
		out += sep + values[i]
	}
	return out
}
