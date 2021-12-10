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
	net.Jitters = [][]int{
		{0, 2, 8, 2, 40},
		{2, 0, 10, 17, 8},
		{8, 10, 0, 9, 14},
		{2, 17, 9, 0, 16},
		{40, 8, 14, 16, 0},
	}

	var paths1 [][]network.Path = network.Dijkstra(net, 0)
	for i := 0; i < len(paths1); i++ {
		fmt.Println(paths1[i])
	}
	fmt.Println()

	var paths2 []network.Path = network.KShortest(net, network.Flow{Source: 0, Destination: 4}, 10)
	fmt.Println()
	for i := 0; i < len(paths2); i++ {
		fmt.Println(paths2[i])
	}
	var paths3 []network.Path
	var jitters []int
	var reliabilities []float64
	paths3, jitters, reliabilities = network.R2tdsdnRouting(net, network.Flow{Source: 0, Destination: 4, DesirableJitter: 40}, 10, 0.6, 0.3)
	fmt.Println("Paths: -------------")
	network.PrintPaths(paths3)
	for i := 0; i < len(paths3); i++ {
		fmt.Printf("Path: %v, jitters: %v, reliabilities: %v\n", paths3, jitters, reliabilities)
	}

	// var net1 network.Network
	// net1.Links = [][]int{
	// 	{0, 3, 2, -1, -1, -1},
	// 	{-1, 0, -1, 4, -1, -1},
	// 	{-1, 1, 0, 2, 3, -1},
	// 	{-1, -1, -1, 0, 2, 1},
	// 	{-1, -1, -1, -1, 0, 2},
	// 	{-1, -1, -1, -1, -1, 0},
	// }
	// net1.Nodes = make([]network.Node, len(net1.Links))
	// fmt.Println()
	// network.GenerateDirectedJitters(net1, 0, 10)
}
