package network

// shortest paths through dfs
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

// shortest paths through Dijkstra
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
		if minNode == -1 { // all of rest nodes are unreachable
			break
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

// use Disjoint Set to check whether a network is connected
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

// shortest paths through Floyd
func Floyd(net Network) [][][]Path {
	n := len(net.Nodes)
	var paths [][][]Path = make([][][]Path, n)
	for i := 0; i < n; i++ {
		paths[i] = make([][]Path, n)
	}

	// init
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				paths[i][j] = []Path{
					{
						Nodes:   []int{i},
						Latency: 0,
					},
				}
				continue
			}
			if net.Links[i][j] > 0 {
				paths[i][j] = []Path{
					{
						Nodes:   []int{i, j},
						Latency: net.Links[i][j],
					},
				}
			}
		}
	}

	// Floyd, try every node as the relay node
	for relay := 0; relay < n; relay++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if relay == i || relay == j {
					continue // no need to relay
				}
				if len(paths[i][relay]) == 0 || len(paths[relay][j]) == 0 {
					continue // no paths for relay
				}
				if len(paths[i][j]) == 0 || paths[i][j][0].Latency > paths[i][relay][0].Latency+paths[relay][j][0].Latency {
					paths[i][j] = relayPaths(paths, i, j, relay)
				} else if paths[i][j][0].Latency == paths[i][relay][0].Latency+paths[relay][j][0].Latency {
					paths[i][j] = append(paths[i][j], relayPaths(paths, i, j, relay)...)
				}
			}
		}
	}
	return paths
}

// generate the relay paths from i to j
func relayPaths(paths [][][]Path, i, j, relay int) []Path {
	var newPaths []Path
	var newLatency int = paths[i][relay][0].Latency + paths[relay][j][0].Latency
	for _, fore := range paths[i][relay] {
		for _, back := range paths[relay][j] {
			foreNodes := make([]int, len(fore.Nodes))
			backNodes := make([]int, len(back.Nodes))
			copy(foreNodes, fore.Nodes)
			copy(backNodes, back.Nodes)
			newPath := Path{
				Nodes:   append(foreNodes[:len(foreNodes)-1], backNodes...),
				Latency: newLatency,
			}
			newPaths = append(newPaths, newPath)
		}
	}
	return newPaths
}

// k-shortest paths through Yen's Algorithm
func KShortest(net Network, flow Flow, k int) []Path {
	var A []Path
	var B pathHeap
	if k == 0 {
		return A
	}
	shortest := Dijkstra(net, flow.Source)[flow.Destination]
	if len(shortest) == 0 { // unreachable
		return A
	}

	A = append(A, minPath(shortest))
	for len(A) < k {
		prevPath := A[len(A)-1].Nodes
		// The spur node ranges from the first node to the next to last node in the previous path.
		for i := 0; i < len(prevPath)-1; i++ {
			spurNode := prevPath[i]
			rootPath := prevPath[:i+1]

			deletedLinks := make(map[[2]int]int) // map: tail, head -> latency.

			// Remove the links that are part of the previous shortest paths which share the same root path.
			for j := 0; j < len(A); j++ {
				if len(A[j].Nodes) > i && sliceEqual(A[j].Nodes[:i+1], rootPath) {
					if _, exist := deletedLinks[[2]int{A[j].Nodes[i], A[j].Nodes[i+1]}]; !exist {
						deletedLinks[[2]int{A[j].Nodes[i], A[j].Nodes[i+1]}] = net.Links[A[j].Nodes[i]][A[j].Nodes[i+1]]
						net.Links[A[j].Nodes[i]][A[j].Nodes[i+1]] = -1
					}
				}
			}

			// Remove the nodes in rootPath except spurNode, make them unreachable.
			for j := 0; j < len(rootPath)-1; j++ {
				for head := 0; head < len(net.Nodes); head++ {
					if _, exist := deletedLinks[[2]int{head, rootPath[j]}]; !exist {
						deletedLinks[[2]int{head, rootPath[j]}] = net.Links[head][rootPath[j]]
						net.Links[head][rootPath[j]] = -1
					}
				}
			}

			// Calculate the spur path from the spur node to the sink.
			spurPaths := Dijkstra(net, spurNode)[flow.Destination]

			// Add back the edges and nodes that were removed from the graph.
			for headTail, latency := range deletedLinks {
				net.Links[headTail[0]][headTail[1]] = latency
			}

			if len(spurPaths) > 0 { // spur path is found
				// put together rootPath and spurPath to build totalPath
				spurPath := minPath(spurPaths).Nodes
				var rootPathCopy []int = make([]int, len(rootPath))
				copy(rootPathCopy, rootPath)
				totalPath := append(rootPathCopy[:len(rootPathCopy)-1], spurPath...)

				var latency int
				for j := 0; j < len(totalPath)-1; j++ {
					latency += net.Links[totalPath[j]][totalPath[j+1]]
				}
				totalPathWithLatency := Path{
					Nodes:   totalPath,
					Latency: latency,
				}

				// insert totalPath to B
				if !B.contain(totalPathWithLatency) {
					B.insert(totalPathWithLatency)
				}
			}
		}
		// no more paths
		if len(B) == 0 {
			break
		}

		// let lowest cost path in B become the k-shortest path.
		// minHeap can guarantee that B[0] has the lowest cost
		A = append(A, B[0])

		// Remove the lowest cost path in B
		B.pop()
	}

	return A
}

// whether two slices are equal
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// minHeap for Path
type pathHeap []Path

// adjust minHeap from up to down
func (h pathHeap) shiftDown(start, end int) {
	var dad int = start
	var son int = dad*2 + 1

	for son <= end { // only compare when son is in range
		if son+1 <= end && pathLess(h[son+1], h[son]) { // choose the bigger son
			son++
		}
		if !pathLess(h[son], h[dad]) {
			break // adjustment is finished
		}
		h[dad], h[son] = h[son], h[dad]
		dad = son
		son = dad*2 + 1
	}
	return
}

// adjust minHeap from down to up
func (h pathHeap) shiftUp(start int) {
	var son int = start
	var dad int = (son - 1) / 2
	for dad >= 0 {
		if !pathLess(h[son], h[dad]) {
			break // // adjustment is finished
		}
		h[dad], h[son] = h[son], h[dad]
		son = dad
		dad = (son - 1) / 2
	}
	return
}

// build a minHeap from a Path slice
func buildHeap(h []Path) {
	// init heap from the last non-leaf node
	for i := len(h)/2 - 1; i >= 0; i-- {
		pathHeap(h).shiftDown(i, len(h)-1)
	}
	return
}

// insert p to h following the rules of minHeap
func (h *pathHeap) insert(p Path) {
	*h = append(*h, p)     // insert
	h.shiftUp(len(*h) - 1) // adjust
	return
}

// remove the minimum element
func (h *pathHeap) pop() {
	(*h)[0] = (*h)[len(*h)-1] // move the last element to the first
	*h = (*h)[:len(*h)-1]     // remove the last element
	h.shiftDown(0, len(*h)-1) // adjust
	return
}

// whether h contains p
func (h pathHeap) contain(p Path) bool {
	for i := 0; i < len(h); i++ {
		if sliceEqual(h[i].Nodes, p.Nodes) {
			return true
		}
	}
	return false
}

// whether p1 < p2
func pathLess(p1, p2 Path) bool {
	if p1.Latency < p2.Latency {
		return true
	}
	if p1.Latency > p2.Latency {
		return false
	}
	return len(p1.Nodes) < len(p2.Nodes)
}

// return the shortest path in a Path slice
func minPath(paths []Path) Path {
	var min Path
	if len(paths) == 0 {
		return min
	}
	min = paths[0]
	for i := 1; i < len(paths); i++ {
		if pathLess(paths[i], min) {
			min = paths[i]
		}
	}
	return min
}

// k-shortest paths through dfs (can make test cases for Yen's Algorithm)
func KShortestDFS(net Network, flow Flow, k int) []Path {
	var paths pathHeap

	n := len(net.Nodes)
	var visited []bool = make([]bool, n) // nodes that are visited
	var path []int                       // record of current path
	var latency int                      // latency of current path

	visited[flow.Source] = true
	path = append(path, flow.Source)

	// use dfs find all paths from source to destination
	var dfs func()
	dfs = func() {
		if path[len(path)-1] == flow.Destination { // find a path
			newPath := Path{
				Nodes:   make([]int, len(path)),
				Latency: latency,
			}
			copy(newPath.Nodes, path)
			paths.insert(newPath)
			return
		}
		for i := 0; i < n; i++ {
			nextLatency := net.Links[path[len(path)-1]][i]
			if nextLatency < 0 { // no link
				continue
			}
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, i)
			latency += nextLatency
			dfs()
			path = path[:len(path)-1]
			visited[i] = false
			latency -= nextLatency
		}
	}
	dfs()

	// pick the k-shortest paths
	var shortestKPaths []Path
	for i := 0; i < k && len(paths) > 0; i++ {
		shortestKPaths = append(shortestKPaths, paths[0])
		paths.pop()
	}
	return shortestKPaths
}
