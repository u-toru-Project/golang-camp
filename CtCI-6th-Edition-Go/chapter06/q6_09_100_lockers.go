package chapter06

import "fmt"

func OpenLockersAfterPasses(numLockers int) []int {
	if numLockers < 0 {
		panic("num_lockers must be non-negative")
	}
	status := make([]bool, numLockers)
	for passer := 1; passer <= numLockers; passer++ {
		for locker := passer - 1; locker < numLockers; locker += passer {
			status[locker] = !status[locker]
		}
	}
	opened := make([]int, 0)
	for i, open := range status {
		if open {
			opened = append(opened, i+1)
		}
	}
	return opened
}

func OpenLockersPerfectSquares(numLockers int) []int {
	if numLockers < 0 {
		panic("num_lockers must be non-negative")
	}
	opened := make([]int, 0)
	for i := 1; i*i <= numLockers; i++ {
		opened = append(opened, i*i)
	}
	return opened
}

func RunQ609() {
	fmt.Printf("Open lockers: %v\n", OpenLockersAfterPasses(100))
}
