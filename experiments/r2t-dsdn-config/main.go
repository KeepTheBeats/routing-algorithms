package main

import (
	"fmt"
	"routing-algorithms/experiments/r2t-dsdn-config/tools"
)

func main() {
	numNets := 50
	tools.GenerateAllNets(numNets)
	tools.GenerateAllFlows(numNets, 100)

	nets, flows := tools.GetNetAndFlows(numNets)

	results := tools.RouteAll(nets, flows, 0.25, true)
	tools.WriteTotalResults(results)
	// results := tools.ReadTotalResults()
	fmt.Println(len(results))
	for i := 0; i < 3; i++ {
		for j := 0; j < len(results[i].Net.Nodes); j++ {
			fmt.Println(results[i].Net.FlowIndexes[j])
		}
		fmt.Println()

		for j := 0; j < len(results[i].Flows); j++ {
			fmt.Println(results[i].Flows[j])
		}
		fmt.Println()
	}

	// for j := 0; j < len(results[2].Net.Nodes); j++ {
	// 	fmt.Println(results[2].Net.FlowIndexes[j])
	// }
	// fmt.Println()

	// fmt.Println(results[2].Flows[1])
	// fmt.Println(results[2].Flows[8])
}
