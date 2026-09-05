package chapter17

import "fmt"

func FindBestSchedule(appointments []int) int {
	for _, length := range appointments {
		if length < 0 {
			panic("appointment lengths must be non-negative")
		}
	}
	n := len(appointments)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return appointments[0]
	}
	dp := make([]int, n+1)
	dp[n-1] = appointments[n-1]
	maxSoFar := appointments[n-1]
	for i := n - 2; i >= 0; i-- {
		take := appointments[i] + dp[i+2]
		skip := dp[i+1]
		dp[i] = max(take, skip)
		if dp[i] > maxSoFar {
			maxSoFar = dp[i]
		}
	}
	return maxSoFar
}

func RunQ1716() {
	fmt.Println(FindBestSchedule([]int{30, 15, 60, 75, 45, 15, 15, 45}))
}
