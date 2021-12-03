package network

// shortest paths through
func Shortest(net Network, flow Flow) []Path {
	var results []Path

	if flow.Source == flow.Destination {
		newPath := Path{
			Nodes:   []int{flow.Source},
			Latency: 0,
		}
		results = append(results, newPath)
		return results
	}

	var visited []bool = make([]bool, len(net.Nodes)) // nodes that are visited
	var path []int                                    // record of current path
	var latency int                                   // latency of current path
	visited[flow.Source] = true
	path = append(path, flow.Source)

	var dst func(Network, Flow)
	dst = func(net Network, flow Flow) {
		// visit every node from the last node on current path
		for i := 0; i < len(net.Nodes); i++ {
			nextLatency := net.Links[path[len(path)-1]][i]

			// i is visited, or there is no link to it
			if nextLatency < 0 || visited[i] {
				continue
			}

			// cannot be the shortest
			if len(results) != 0 && latency+nextLatency > results[0].Latency {
				continue
			}

			if i == flow.Destination {
				path = append(path, i)
				newPath := Path{
					Nodes:   make([]int, len(path)),
					Latency: latency + nextLatency,
				}
				copy(newPath.Nodes, path)
				if len(results) == 0 || latency+nextLatency < results[0].Latency {
					results = []Path{newPath}
				} else { // latency+nextLatency == minLatency
					results = append(results, newPath)
				}
				path = path[:len(path)-1]
				continue
			}
			visited[i] = true
			path = append(path, i)
			latency += nextLatency
			dst(net, flow)
			latency -= nextLatency
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	dst(net, flow)
	return results
}

func Dijkstra(net Network, source int) [][]Path {
	n := len(net.Nodes)
	var results [][]Path = make([][]Path, n)

	newPath := Path{
		Nodes:   []int{source},
		Latency: 0,
	}
	results[source] = []Path{newPath}

	var latencies []int = make([]int, n) // latencies from source to every node
	for i := 0; i < n; i++ {
		latencies[i] = net.Links[source][i]
	}

	var visited []bool = make([]bool, n) // already have shortest paths
	visited[source] = true

	var predecessors [][]int = make([][]int, n) // predecessors of all nodes on their shortest paths
	for i := 0; i < n; i++ {
		predecessors[i] = []int{source}
	}
	predecessors[source] = []int{-1}

	for count := 0; count < n-1; count++ { // there are n-1 nodes except source
		minNode := -1
		for i := 0; i < n; i++ { // find the node with minimum latency
			if visited[i] || latencies[i] < 0 { // lagency < 0 means unreachable
				continue
			}
			if minNode < 0 || latencies[i] < latencies[minNode] {
				minNode = i
			}
		}
		// shortest paths to miniNode is finded
		visited[minNode] = true
		paths := findPaths(minNode, source, predecessors, latencies[minNode])
		results[minNode] = paths

		// update predecessors and latencies based on miniNode
		for i := 0; i < n; i++ {
			if visited[i] || net.Links[minNode][i] < 0 {
				continue
			}
			if latencies[i] < 0 || latencies[i] > latencies[minNode]+net.Links[minNode][i] {
				latencies[i] = latencies[minNode] + net.Links[minNode][i]
				predecessors[i] = []int{minNode}
			} else if latencies[i] == latencies[minNode]+net.Links[minNode][i] {
				predecessors[i] = append(predecessors[i], minNode)
			}
		}
	}

	return results
}

func findPaths(node int, source int, predecessors [][]int, latency int) []Path {
	var paths []Path

	var stack []int
	stack = append(stack, node)

	var find func()
	find = func() {
		if stack[len(stack)-1] == source { // a path is finded
			nodes := []int{}
			for i := len(stack) - 1; i >= 0; i-- {
				nodes = append(nodes, stack[i])
			}
			newPath := Path{
				Nodes:   nodes,
				Latency: latency,
			}
			paths = append(paths, newPath)
			return
		}
		// find the predecessors of stack.top
		for _, predecessor := range predecessors[stack[len(stack)-1]] {
			stack = append(stack, predecessor)
			find()
			stack = stack[:len(stack)-1]
		}
	}
	find()

	return paths
}

// usr Disjoint Set to check whether a network is connected
func CheckConnected(net Network) bool {
	var leaders []int = make([]int, len(net.Nodes))

	var init func() = func() {
		for i := 0; i < len(leaders); i++ {
			leaders[i] = i
		}
	}
	init()

	var find func(int) int
	find = func(x int) int {
		if leaders[x] == x {
			return x
		}
		leaders[x] = find(leaders[x])
		return leaders[x]
	}

	var merge func(int, int) = func(i, j int) {
		leaders[find(i)] = leaders[find(j)]
	}
	for i := 0; i < len(net.Links); i++ {
		for j := 0; j < len(net.Links[i]); j++ {
			if net.Links[i][j] > 0 {
				merge(i, j)
			}
		}
	}

	var cnt int
	for i := 0; i < len(net.Nodes); i++ {
		if find(i) == i {
			cnt++
		}
	}
	return cnt == 1
}
