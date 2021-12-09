package main

import (
	"fmt"

	"routing-algorithms/network"
)

func main() {
	var net network.Network
	net.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 2, -1, 7},
		{4, 2, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net.Nodes = make([]network.Node, len(net.Links))

	var paths1 [][]network.Path = network.Dijkstra(net, 0)
	for i := 0; i < len(paths1); i++ {
		fmt.Println(paths1[i])
	}
	fmt.Println()

	var paths2 []network.Path = network.KShortest(net, network.Flow{Source: 1, Destination: 3}, 10)
	fmt.Println()
	for i := 0; i < len(paths2); i++ {
		fmt.Println(paths2[i])
	}
	var paths3 []network.Path = network.KShortestDFS(net, network.Flow{Source: 1, Destination: 3}, 10)
	fmt.Println()
	for i := 0; i < len(paths3); i++ {
		fmt.Println(paths3[i])
	}
}
