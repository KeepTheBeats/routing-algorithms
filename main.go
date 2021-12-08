package main

import (
	"fmt"
	"r2tdsdn-configuration/network"
)

func main() {
	var net network.Network
	net.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 3, -1, 7},
		{4, 3, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net.Nodes = make([]network.Node, len(net.Links))

	var net2 network.Network
	net2.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 3, -1, 7},
		{4, 3, 0, -1, 3},
		{-1, -1, -1, 0, -1},
		{10, 7, 3, -1, 0},
	}
	net2.Nodes = make([]network.Node, len(net2.Links))

	var net3 network.Network
	net3.Links = [][]int{
		{0, -1, -1, -1, 10},
		{-1, 0, 2, -1, -1},
		{-1, 2, 0, 4, -1},
		{-1, -1, 4, 0, -1},
		{10, -1, -1, -1, 0},
	}
	net3.Nodes = make([]network.Node, len(net3.Links))
	var paths [][][]network.Path
	paths = network.Floyd(net2)
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths); j++ {
			fmt.Println(paths[i][j])
		}
		fmt.Println()
	}
	// var paths [][]network.Path
	// paths = network.Dijkstra(net, 0)
	// for i := 0; i < len(paths); i++ {
	// 	fmt.Println(paths[i])
	// }
	// fmt.Println()

	// paths = network.Dijkstra(net2, 0)
	// for i := 0; i < len(paths); i++ {
	// 	fmt.Println(paths[i])
	// }
	// fmt.Println()

	// paths = network.Dijkstra(net3, 1)
	// for i := 0; i < len(paths); i++ {
	// 	fmt.Println(paths[i])
	// }
	// fmt.Println()

}
