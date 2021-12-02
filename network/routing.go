package network

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
