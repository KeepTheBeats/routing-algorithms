package network

// the numbers of nodes in a network start from 0, e.g., 10 nodes are 0-9.
type Network struct {
	Links [][]int // latency of every link, -1 means that there is no link between two nodes
	Nodes []Node
}

type Node struct {
}

type Flow struct {
	Source      int
	Destination int
}

type Path struct {
	Nodes   []int // nodes on a path
	Latency int   // total latency of a path
}
