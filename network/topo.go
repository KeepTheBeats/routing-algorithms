package network

// the numbers of nodes in a network start from 0, e.g., 10 nodes are 0-9.
type Network struct {
	Nodes   []Node  `json:"nodes"`
	Links   [][]int `json:"links"`   // latency of every link, -1 means that there is no link between two nodes
	Jitters [][]int `json:"jitters"` // Jitter of every link
}

type Node struct {
}

type Flow struct {
	Source          int
	Destination     int
	DesirableJitter int
	Data            int // the amount of data that the flow need to transmit, the unit is b(bit), in the experiment of R2T‑DSDN paper, Link data rate is 10Gb/s (1073742b/ms)
	Deadline        int // in non-RT flows, deadline is -1, unit is ms.
}

type Path struct {
	Nodes   []int // nodes on a path
	Latency int   // total latency of a path
}
