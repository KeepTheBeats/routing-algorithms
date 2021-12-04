package main

import (
	"fmt"
	"r2tdsdn-configuration/network"
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
	var paths [][][]network.Path
	paths = network.Floyd(net)
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths); j++ {
			fmt.Println(paths[i][j])
		}
		fmt.Println()
	}
	// var paths [][]network.Path
	// paths = network.Dijkstra(net, 4)
	// for i := 0; i < len(paths); i++ {
	// 	fmt.Println(paths[i])
	// }
	// fmt.Println()
}
