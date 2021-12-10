package network

// the numbers of nodes in a network start from 0, e.g., 10 nodes are 0-9.
type Network struct {
	Nodes   []Node
	Links   [][]int // latency of every link, -1 means that there is no link between two nodes
	Jitters [][]int // Jitter of every link
}

type Node struct {
}

type Flow struct {
	Source          int
	Destination     int
	DesirableJitter int
}

type Path struct {
	Nodes   []int // nodes on a path
	Latency int   // total latency of a path
}
