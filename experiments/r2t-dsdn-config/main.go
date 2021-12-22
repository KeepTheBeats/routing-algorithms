package main

import (
	"fmt"
	"routing-algorithms/experiments/r2t-dsdn-config/tools"
	"routing-algorithms/network"
)

func main() {
	numNets := 50
	tools.GenerateAllNets(numNets)
	tools.GenerateAllFlows(numNets, 126)

	// debug
	// numNets := 1
	// tools.GenerateAllNets(numNets)
	// tools.GenerateAllFlows(numNets, 11)

	nets, flows := tools.GetNetAndFlows(numNets)

	var results []network.RoutingResult

	// fmt.Println("-----------5--------------")
	// results = tools.RouteAll(nets, flows, 0.05, false)
	// // tools.WriteTotalResults(results)
	// tools.OutputData(results, "5")

	fmt.Println("-----------10--------------")
	results = tools.RouteAll(nets, flows, 0.10, false)
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

	fmt.Println("-----------15--------------")
	results = tools.RouteAll(nets, flows, 0.15, false)
	tools.OutputData(results, "15")
	// fmt.Println("-----------20--------------")
	// results = tools.RouteAll(nets, flows, 0.20, false)
	// tools.OutputData(results, "20")

	fmt.Println("-----------25--------------")
	results = tools.RouteAll(nets, flows, 0.25, false)
	tools.OutputData(results, "25")

	// fmt.Println("-----------30--------------")
	// results = tools.RouteAll(nets, flows, 0.30, false)
	// tools.OutputData(results, "30")
	// fmt.Println("-----------40--------------")
	// results = tools.RouteAll(nets, flows, 0.40, false)
	// tools.OutputData(results, "40")
	fmt.Println("-----------dynamic--------------")
	results = tools.RouteAll(nets, flows, 0.2, true)
	tools.OutputData(results, "dynamic")
}
