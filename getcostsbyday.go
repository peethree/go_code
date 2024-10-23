package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	max := 0

	for i := 0; i < len(costs); i++ {
		if costs[i].day > max {
			max = costs[i].day
		}
	}

	dayCost := make([]float64, max+1)

	for i := 0; i < len(costs); i++ {
		dayCost[costs[i].day] += costs[i].value
	}

	return dayCost
}
