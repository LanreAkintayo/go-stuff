package main

import "fmt"

type DataPlan struct {
	id string
	size int
	price float64
}


func OptimizeDataPlans(dataPlans []DataPlan) map[int]DataPlan {
	optimized := make(map[int]DataPlan)

	for _, currentPlan := range dataPlans {
		existingDataPlan, exists := optimized[currentPlan.size]

		if exists {
			// Compare and retain the one that has the higher price
			if currentPlan.price > existingDataPlan.price {
				optimized[currentPlan.size] = currentPlan
			}
		} else {
			optimized[currentPlan.size] = currentPlan
		}
	}

	return optimized
}
   
func main(){
	// Get a list of data plans
	dataPlans := []DataPlan {
		{"p1", 1, 300.0},
		{"p2", 5, 1500.0},
		{"p3", 1, 450.0},
		{"p4", 10, 3000.0},
		{"p5", 5, 1200.0},
	} 

	optimizedDataPlans  := OptimizeDataPlans(dataPlans)

	fmt.Println(optimizedDataPlans)
}