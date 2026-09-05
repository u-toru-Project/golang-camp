package chapter14

import (
	"fmt"
	"sort"
	"strings"
)

type apartment struct {
	id         int
	buildingID int
	unit       string
	rent       float64
}

type ApartmentStore struct {
	nextBuildingID  int
	nextApartmentID int
	buildings       map[int]string
	apartments      []apartment
}

func NewApartmentStore() *ApartmentStore {
	return &ApartmentStore{
		nextBuildingID:  1,
		nextApartmentID: 1,
		buildings:       make(map[int]string),
	}
}

func (s *ApartmentStore) AddBuilding(name string) int {
	id := s.nextBuildingID
	s.nextBuildingID++
	s.buildings[id] = name
	return id
}

func (s *ApartmentStore) AddApartment(buildingID int, unit string, rent float64) int {
	id := s.nextApartmentID
	s.nextApartmentID++
	s.apartments = append(s.apartments, apartment{id: id, buildingID: buildingID, unit: unit, rent: rent})
	return id
}

func (s *ApartmentStore) ApartmentsInBuilding(buildingID int) []string {
	units := make([]string, 0)
	for _, apt := range s.apartments {
		if apt.buildingID == buildingID {
			units = append(units, apt.unit)
		}
	}
	sort.Strings(units)
	return units
}

func RunQ1401() {
	store := NewApartmentStore()
	towerA := store.AddBuilding("Tower A")
	towerB := store.AddBuilding("Tower B")
	store.AddApartment(towerA, "101", 1200.0)
	store.AddApartment(towerA, "102", 1300.0)
	store.AddApartment(towerB, "201", 900.0)
	fmt.Println(strings.Join(store.ApartmentsInBuilding(towerA), ", "))
}
