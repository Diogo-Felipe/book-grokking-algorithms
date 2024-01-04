package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	var graph = map[string]map[string]float64{
		"start": {
			"a": 6,
			"b": 2,
		},
		"a": {
			"end": 1,
		},
		"b": {
			"a": 3,
			"b": 5,
		},
	}

	var costs map[string]float64
	costs = make(map[string]float64)
	costs["a"] = 6
	costs["b"] = 2
	costs["end"] = math.Inf(1)

	var parents map[string]string
	parents = map[string]string{}
	parents["a"] = "start"
	parents["b"] = "start"

	endCost := djikstraLowestCost(graph, costs, parents)
	fmt.Printf("The cost to get at end is: %v\n", endCost)
}

func djikstraLowestCost(graph map[string]map[string]float64, costs map[string]float64, parents map[string]string) float64 {
	var processed []string

	node := getLowestCost(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for neighbor, neighborCost := range neighbors {
			newCost := cost + neighborCost

			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}

		processed = append(processed, node)
		node = getLowestCost(costs, processed)
	}

	return costs["end"]
}

func getLowestCost(costs map[string]float64, processed []string) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""

	for index, node := range costs {
		if node < lowestCost && !slices.Contains(processed, index) {
			lowestCost = node
			lowestCostNode = index
		}
	}

	return lowestCostNode
}
