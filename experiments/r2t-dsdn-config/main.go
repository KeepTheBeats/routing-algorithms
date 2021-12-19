package main

import (
	"fmt"
	"routing-algorithms/experiments/r2t-dsdn-config/tools"
)

func main() {
	tools.GenerateAllNets(50)
	tools.GenerateAllFlows(50, 100)
	flows := tools.ReadJsonFlows("./experiments/r2t-dsdn-config/jsonnetworks/net04_flows.json")

	for i := 0; i < len(flows); i++ {
		fmt.Println("scenario", i+1)
		for j := 0; j < len(flows[i]); j++ {
			fmt.Println(flows[i][j])
		}
	}
}
