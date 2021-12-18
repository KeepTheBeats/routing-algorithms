package main

import (
	"fmt"
	"routing-algorithms/experiments/r2t-dsdn-config/tools"
)

func main() {
	tools.GenerateAllNets(50)
	var netName string
	for i := 1; i <= 50; i++ {
		if i < 10 {
			netName = fmt.Sprintf("net0%d", i)
		} else {
			netName = fmt.Sprintf("net%d", i)
		}
		net := tools.ReadJsonNet("./experiments/r2t-dsdn-config/jsonnetworks/" + netName + ".json")
		fmt.Println(net)
	}
}
