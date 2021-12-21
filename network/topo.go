package network

// the numbers of nodes in a network start from 0, e.g., 10 nodes are 0-9.
type Network struct {
	Nodes         []Node      `json:"nodes"`
	Links         [][]int     `json:"links"`         // latency of every link, -1 means that there is no link between two nodes
	Jitters       [][]int     `json:"jitters"`       // Jitter of every link
	FlowIndexes   [][][]int   `json:"flowindexes"`   // flows routed on every link, save indexes of flows in every scenario
	RemainIndexes [][][]bool  `json:"remainindexes"` // flows that cannot be deleted from this link
	ReservedBW    [][]float64 `json:"reservedbw"`    // reserved bandwidth of every link
}

type Node struct {
}

type Flow struct {
	Source          int     `json:"source"`
	Destination     int     `json:"destination"`
	DesirableJitter int     `json:"desirablejitter"`
	Data            float64 `json:"data"`     // the amount of data that the flow need to transmit, the unit is b(bit), in the experiment of R2T‑DSDN paper, Link data rate is 10Gb/s (1073742b/ms, 1048kb/ms, 131KB/ms)
	Deadline        float64 `json:"deadline"` // in non-RT flows, deadline is -1, unit is ms.
	Paths           []Path  `json:"paths"`    // paths chosen by routing
	Jitters         []int   `json:"jitters"`  // jitters of chosed paths
}

type Path struct {
	Nodes   []int `json:"nodes"`   // nodes on a path
	Latency int   `json:"latency"` // total latency of a path
}

type RoutingResult struct {
	Net   Network `json:"net"`
	Flows []Flow  `json:"flows"`
}
