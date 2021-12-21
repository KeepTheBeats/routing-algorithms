package main

import (
	"routing-algorithms/experiments/r2t-dsdn-config/tools"
)

func main() {
	numNets := 50
	tools.GenerateAllNets(numNets)
	tools.GenerateAllFlows(numNets, 126)

	// debug
	// numNets := 10
	// tools.GenerateAllNets(numNets)
	// tools.GenerateAllFlows(numNets, 33)

	nets, flows := tools.GetNetAndFlows(numNets)

	results := tools.RouteAll(nets, flows, 0.10, false)
	// tools.WriteTotalResults(results)
	tools.OutputData(results, "10")

	// results := tools.ReadTotalResults()
	// fmt.Println(len(results))
	// for i := 44; i < 47; i++ {
	// 	for j := 0; j < len(results[i].Net.Nodes); j++ {
	// 		fmt.Println(results[i].Net.FlowIndexes[j])
	// 	}
	// 	fmt.Println()

	// 	for j := 0; j < len(results[i].Flows); j++ {
	// 		fmt.Println(results[i].Flows[j])
	// 	}
	// 	fmt.Println()
	// }

	results = tools.RouteAll(nets, flows, 0.25, false)
	tools.OutputData(results, "25")
	results = tools.RouteAll(nets, flows, 0.40, false)
	tools.OutputData(results, "40")
	results = tools.RouteAll(nets, flows, 0.40, true)
	tools.OutputData(results, "dynamic")
}
