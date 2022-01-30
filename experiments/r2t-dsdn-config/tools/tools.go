package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"routing-algorithms/network"
	"routing-algorithms/random"
)

// read network from alt file generated by gt-itm, generate latency and jitter randomly
func ReadAltNet(filePath string) network.Network {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileContent []string
	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}

	splitLine2 := strings.Split(fileContent[1], " ")

	numNodes, _ := strconv.Atoi(splitLine2[0])
	numDirectedLinks, _ := strconv.Atoi(splitLine2[1])
	numUndirectedLinks := numDirectedLinks / 2

	var net network.Network
	net.Nodes = make([]network.Node, numNodes)
	net.Links = make([][]int, numNodes)
	net.Jitters = make([][]int, numNodes)
	for i := 0; i < numNodes; i++ {
		net.Links[i] = make([]int, numNodes)
		net.Jitters[i] = make([]int, numNodes)
		for j := 0; j < numNodes; j++ {
			net.Links[i][j] = -1   // init latency to -1
			net.Jitters[i][j] = -1 // init jitters to -1
		}
	}

	// read edges
	if !strings.HasPrefix(fileContent[3+numNodes+2], "EDGES") {
		panic(fmt.Errorf("%d line does not start with \"EDGES\"", 3+numNodes+2))
	}

	for i := 1; i <= numUndirectedLinks; i++ {
		edge := strings.Split(fileContent[3+numNodes+2+i], " ")
		node1, _ := strconv.Atoi(edge[0])
		node2, _ := strconv.Atoi(edge[1])
		// latency of every link is a random number following Poisson distribution, lambda=7, between [3,15]
		net.Links[node1][node2] = random.PoissonRandom(3, 15, 7)
		net.Links[node2][node1] = net.Links[node1][node2]
		// jitter of every link is a random number following Normal distribution, miu=4, sigma=6, between [1,10]
		net.Jitters[node1][node2] = int(random.NormalRandomBM(1, 10, 4, 6))
		net.Jitters[node2][node1] = net.Jitters[node1][node2]
	}

	return net
}

// write content to json file
func WriteJson(v interface{}, outPath string) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(outPath, data, 0777)
	if err != nil {
		panic(err)
	}
}

// read network from json file
func ReadJsonNet(filePath string) network.Network {
	netData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var net network.Network
	err = json.Unmarshal(netData, &net)
	if err != nil {
		panic(err)
	}
	return net
}

// use alt files in "gtitmnetworks" to generate json files in "jsonnetworks"
func GenerateAllNets(num int) {
	var netName string
	for i := 1; i <= num; i++ {
		netName = GetNetName(i)
		net := ReadAltNet("./experiments/r2t-dsdn-config/gtitmnetworks/" + netName + ".alt")
		WriteJson(net, "./experiments/r2t-dsdn-config/jsonnetworks/"+netName+".json")
	}
}

// generate num scenarios for net
func GenerateFlowsForNet(net network.Network, num int) [][]network.Flow {
	// generate num scenarios
	var flows [][]network.Flow = make([][]network.Flow, num)
	for i := 0; i < num; i++ {
		// in every scenario there are [5,20] flows
		// numFlows := random.RandomInt(5, 20)

		// there are 120 scenarios, in every scenario there are 10-30 flows, 6 scenarios have 10, 6 have 11, 6 have 12 ... 6 have 30
		// numFlows := i/1 + 10 // debug
		numFlows := i/6 + 10
		// in every scenario there are [1,numFlows-1] RT-flows
		numRTFlows := random.RandomInt(1, numFlows-1)

		flows[i] = make([]network.Flow, numFlows)
		for j := 0; j < numFlows; j++ {
			// generate source and destination
			flows[i][j].Source = random.RandomInt(0, len(net.Nodes)-1)
			flows[i][j].Destination = random.RandomInt(0, len(net.Nodes)-1)
			for flows[i][j].Source == flows[i][j].Destination {
				flows[i][j].Destination = random.RandomInt(0, len(net.Nodes)-1)
			}

			// generate desirableJitter
			flows[i][j].DesirableJitter = int(random.NormalRandomBM(15, 40, 30, 15))

			// generate data
			flows[i][j].Data = random.NormalRandomBM(30, 90, 60, 40)

			// generate deadline
			flows[i][j].Deadline = -1 // in non-RT flows, deadline is -1
			if j < numRTFlows {
				// transmission time is the sum of latency of links, if transmission time <= deadline, it can hit deadline
				// flows[i][j].Deadline = int(random.NormalRandomBM(15, 28, 22, 18))
				// flows[i][j].Deadline = random.NormalRandomBM(15, 23, 19, 3)
				// flows[i][j].Deadline = random.NormalRandomBM(7, 21, 16, 3)
				flows[i][j].Deadline = random.NormalRandomBM(18, 26, 22, 3)
			}
		}
	}
	return flows
}

// generate num continuous scenarios for net
func GenerateContinuousFlowsForNet(net network.Network, num int) [][]network.Flow {
	// generate num scenarios
	var flows [][]network.Flow = make([][]network.Flow, num)
	var maxFlowNum int = 30
	// there are 120 scenarios, in every scenario there are 10-30 flows, 6 scenarios have 10, 6 have 11, 6 have 12 ... 6 have 30

	for i := 0; i < 6; i++ { // first 6 scenarios have 30 flows
		numFlows := maxFlowNum

		flows[i] = make([]network.Flow, numFlows)
		for j := 0; j < numFlows; j++ {
			// generate source and destination
			flows[i][j].Source = random.RandomInt(0, len(net.Nodes)-1)
			flows[i][j].Destination = random.RandomInt(0, len(net.Nodes)-1)
			for flows[i][j].Source == flows[i][j].Destination {
				flows[i][j].Destination = random.RandomInt(0, len(net.Nodes)-1)
			}

			// generate desirableJitter
			flows[i][j].DesirableJitter = int(random.NormalRandomBM(15, 40, 30, 15))

			// generate data
			flows[i][j].Data = random.NormalRandomBM(30, 90, 60, 40)

			// generate deadline
			flows[i][j].Deadline = -1 // in non-RT flows, deadline is -1

			if j%3 == 0 { //in every 3 flows, there is 1 RT flow, which is similar with the experiments of original R2T-DSDN
				// transmission time is the sum of latency of links, if transmission time <= deadline, it can hit deadline
				// flows[i][j].Deadline = int(random.NormalRandomBM(15, 28, 22, 18))
				// flows[i][j].Deadline = random.NormalRandomBM(15, 23, 19, 3)
				// flows[i][j].Deadline = random.NormalRandomBM(7, 21, 16, 3)
				flows[i][j].Deadline = random.NormalRandomBM(18, 26, 22, 3)
			}
		}
	}

	for i := 6; i < num; i++ { // other scenarios
		numFlows := maxFlowNum - i/6 // in No.i scenario, there are numFlows flows
		flows[i] = make([]network.Flow, numFlows)
		copy(flows[i], flows[i-6][:len(flows[i-6])-1]) // pick the first numFlows elements of flows[i-6]
	}

	return flows
}

// read flows from json file
func ReadJsonFlows(filePath string) [][]network.Flow {
	flowData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var flows [][]network.Flow
	err = json.Unmarshal(flowData, &flows)
	if err != nil {
		panic(err)
	}
	return flows
}

// read json files of networks in "jsonnetworks" to generate json files of flows in "jsonnetworks", generate numScenario scenarios for every net
func GenerateAllFlows(numNet, numScenario int) {
	var netName string
	for i := 1; i <= numNet; i++ {
		netName = GetNetName(i)
		net := ReadJsonNet("./experiments/r2t-dsdn-config/jsonnetworks/" + netName + ".json")
		flows := GenerateContinuousFlowsForNet(net, numScenario)
		WriteJson(flows, "./experiments/r2t-dsdn-config/jsonnetworks/"+netName+"_flows.json")
	}
}

// read json files in in "jsonnetworks" to get nets and flows
func GetNetAndFlows(num int) ([]network.Network, [][][]network.Flow) {
	var nets []network.Network = make([]network.Network, num)
	var flows [][][]network.Flow = make([][][]network.Flow, num)
	var netName string
	for i := 0; i < num; i++ {
		netName = GetNetName(i + 1)
		nets[i] = ReadJsonNet("./experiments/r2t-dsdn-config/jsonnetworks/" + netName + ".json")
		flows[i] = ReadJsonFlows("./experiments/r2t-dsdn-config/jsonnetworks/" + netName + "_flows.json")
	}
	return nets, flows
}

// route all flows of all nets
func RouteAll(nets []network.Network, flows [][][]network.Flow, reservedBW float64, dynamicReserve bool) []network.RoutingResult {
	var bandwidthNonRT float64
	bandwidthNonRT = bandwidth * (float64(1) - reservedBW)

	var results []network.RoutingResult
	for z := 0; z < len(nets); z++ { // for every net
		for i := 0; i < len(flows[z]); i++ { // for every scenario
			var startTime time.Time = time.Now()
			nets[z].FlowIndexes = make([][][]int, len(nets[z].Nodes))
			nets[z].RemainIndexes = make([][][]bool, len(nets[z].Nodes))
			for k := 0; k < len(nets[z].FlowIndexes); k++ {
				nets[z].FlowIndexes[k] = make([][]int, len(nets[z].Nodes))
				nets[z].RemainIndexes[k] = make([][]bool, len(nets[z].Nodes))
			}
			for j := 0; j < len(flows[z][i]); j++ { // for every flow
				paths, jitters, _ := network.R2tdsdnRouting(nets[z], flows[z][i][j], 15, 0.8, 0.2)
				// record routing results into flow
				flows[z][i][j].Paths = paths
				flows[z][i][j].Jitters = jitters
				// record routing results into net
				for _, path := range paths {
					for k := 0; k < len(path.Nodes)-1; k++ {
						smaller, bigger := smallerBigger(path.Nodes[k], path.Nodes[k+1])
						// only write to the up-right of the matrix, because graph is undirected
						nets[z].FlowIndexes[smaller][bigger] = append(nets[z].FlowIndexes[smaller][bigger], j)
					}
				}
			}
			for j := 0; j < len(nets[z].FlowIndexes); j++ {
				for k := 0; k < len(nets[z].FlowIndexes[j]); k++ {
					nets[z].RemainIndexes[j][k] = make([]bool, len(nets[z].FlowIndexes[j][k]))
				}
			}

			if dynamicReserve {
				SetFinegrainedReservedBW(&nets[z], flows[z][i])
			}

			// drop some non-RT flows because of non-RT overloaded links
			for {
				deleted := false
				for j := 0; j < len(nets[z].FlowIndexes); j++ {
					for k := j + 1; k < len(nets[z].FlowIndexes[j]); k++ { // for every link [j][k]
						if len(nets[z].FlowIndexes[j][k]) == 0 {
							continue // no flows routed on this link
						}
						// get the sum of data on this link
						var sumData float64
						for l := 0; l < len(nets[z].FlowIndexes[j][k]); l++ {
							sumData += dataPerPath(flows[z][i][nets[z].FlowIndexes[j][k][l]])
						}

						if dynamicReserve {
							bandwidthNonRT = bandwidth * (float64(1) - nets[z].ReservedBW[j][k])
						}

						if sumData <= bandwidthNonRT {
							continue // this link is not overloaded
						}

						if sumData <= bandwidth { // non-RT overloaded
							// if a non-RT flow is on this link, remove one
							var enableIndexes []int
							for l := 0; l < len(nets[z].FlowIndexes[j][k]); l++ {
								if !isRtFlow(flows[z][i][nets[z].FlowIndexes[j][k][l]]) {
									enableIndexes = append(enableIndexes, l)
								}
							}
							if len(enableIndexes) > 0 {
								index := random.RandomInt(0, len(enableIndexes)-1)
								l := enableIndexes[index]

								rmFlowFromLink(&nets[z], j, k, &flows[z][i][nets[z].FlowIndexes[j][k][l]], nets[z].FlowIndexes[j][k][l])

								deleted = true
							}
						}
					}
				}
				if !deleted {
					break // no flows need to be dropped
				}
			}

			// drop some flows because of overloaded links
			for {
				deleted := false
				for j := 0; j < len(nets[z].FlowIndexes); j++ {
					for k := j + 1; k < len(nets[z].FlowIndexes[j]); k++ { // for every link [j][k]
						if len(nets[z].FlowIndexes[j][k]) == 0 {
							continue // no flows routed on this link
						}
						// get the sum of data on this link
						var sumData float64
						for l := 0; l < len(nets[z].FlowIndexes[j][k]); l++ {
							sumData += dataPerPath(flows[z][i][nets[z].FlowIndexes[j][k][l]])
						}

						if sumData <= bandwidth {
							continue // this link is not overloaded
						}

						// overloaded, remove a flow randomly
						if len(nets[z].FlowIndexes[j][k]) != 0 {
							var l int
							if dynamicReserve {
								var enableIndexes []int
								for p := 0; p < len(nets[z].RemainIndexes[j][k]); p++ {
									if !nets[z].RemainIndexes[j][k][p] {
										enableIndexes = append(enableIndexes, p)
									}
								}
								if len(enableIndexes) == 0 {
									l = random.RandomInt(0, len(nets[z].FlowIndexes[j][k])-1)
								} else {
									index := random.RandomInt(0, len(enableIndexes)-1)
									l = enableIndexes[index]
								}
							} else {
								l = random.RandomInt(0, len(nets[z].FlowIndexes[j][k])-1)
							}
							rmFlowFromLink(&nets[z], j, k, &flows[z][i][nets[z].FlowIndexes[j][k][l]], nets[z].FlowIndexes[j][k][l])
							deleted = true
						}
					}
				}
				if !deleted {
					break // no flows need to be dropped
				}
			}

			var endTime time.Time = time.Now()
			routingTime := float64(endTime.Sub(startTime)) / float64(time.Millisecond)
			// for every scenario, record routing result
			newResult := network.RoutingResult{
				Net:         nets[z],
				Flows:       flows[z][i],
				RoutingTime: routingTime,
			}
			results = append(results, newResult)

		}
	}
	return results
}

// if a flow are transmitted through n paths, data on every path is flow.Data/n
func dataPerPath(flow network.Flow) float64 {
	return flow.Data / float64(len(flow.Paths))
}

// calculate reserved bandwidth according to flows state
func DynamicReservedBW(flows []network.Flow) float64 {
	// configure reserved bandwidth dynamically according to the ratio of rtData
	var rtData, totalData float64
	for _, flow := range flows {
		if isRtFlow(flow) {
			rtData += flow.Data
		}
		totalData += flow.Data
	}
	var rtRatio float64 = rtData / totalData
	// function: y = 2/3*x + 1/15, x: rtRatio, y: reservedBW
	var reservedBW float64 = rtRatio*(float64(2)/float64(3)) + float64(1)/float64(15)
	if reservedBW > 0.6 { // max: 0.6
		reservedBW = 0.6
	}
	if reservedBW < 0.1 { // min: 0.1
		reservedBW = 0.1
	}
	return reservedBW
}

// calculate reserved bandwidth according to network, flows and routing state
func DynamicReservedBWAfterRouting(net network.Network, flows []network.Flow) float64 {
	// configure reserved bandwidth dynamically according to rtData and bandwidth of every link
	var reservedBW, totalRtData, totalBandwidth float64
	for i := 0; i < len(net.FlowIndexes); i++ {
		for j := i + 1; j < len(net.FlowIndexes[i]); j++ { // for every link [i][j]
			if len(net.FlowIndexes[i][j]) == 0 {
				continue // only consider the links with flows, including rt and non-rt flows
			}
			// get the sum of rt-data on this link
			var rtData float64
			for k := 0; k < len(net.FlowIndexes[i][j]); k++ {
				if isRtFlow(flows[net.FlowIndexes[i][j][k]]) {
					rtData += dataPerPath(flows[net.FlowIndexes[i][j][k]])
				}
			}
			totalRtData += rtData
			totalBandwidth += bandwidth
		}
	}
	reservedBW = totalRtData / totalBandwidth
	if reservedBW > 0.9 { // max: 0.6
		reservedBW = 0.9
	}
	if reservedBW < 0.1 { // min: 0.1
		reservedBW = 0.1
	}
	reservedBW *= float64(2) / float64(3)
	return reservedBW
}

// set reserved bandwidth for every link of the network, according to network, flows and routing state
func SetReservedBW(net *network.Network, flows []network.Flow) {
	net.ReservedBW = make([][]float64, len(net.Nodes))
	for i := 0; i < len(net.ReservedBW); i++ {
		net.ReservedBW[i] = make([]float64, len(net.Nodes))
	}
	// set reserved bandwidth to every link according to rtData and bandwidth
	var reservedBW float64
	for i := 0; i < len(net.FlowIndexes); i++ {
		for j := i + 1; j < len(net.FlowIndexes[i]); j++ { // for every link [i][j]
			// get the sum of rt-data on this link
			var rtData float64
			for k := 0; k < len(net.FlowIndexes[i][j]); k++ {
				if isRtFlow(flows[net.FlowIndexes[i][j][k]]) {
					rtData += dataPerPath(flows[net.FlowIndexes[i][j][k]])
				}
			}
			reservedBW = rtData / bandwidth

			// switch {
			// case reservedBW <= 0.4:
			// 	reservedBW = reservedBW
			// case reservedBW > 0.4 && reservedBW <= 0.8:
			// 	reservedBW = 0.4 + (reservedBW-0.4)*0.5
			// case reservedBW > 0.8 && reservedBW < 1:
			// 	reservedBW = 0.4 + (0.8-0.4)*0.5 + (reservedBW-0.8)*0.25
			// default:
			// 	reservedBW = 0.4 + (0.8-0.4)*0.5 + (1-0.8)*0.25
			// }
			if reservedBW > 0.4 {
				reservedBW = 0.4
			}
			reservedBW *= 1.1
			// switch {
			// case reservedBW <= 0.4:
			// 	reservedBW = reservedBW
			// default:
			// 	reservedBW = 0.4 + (reservedBW-0.4)*0.5
			// }
			// if reservedBW > 0.4+(0.8-0.4)*0.5 {
			// 	reservedBW = 0.4 + (0.8-0.4)*0.5
			// }
			net.ReservedBW[i][j] = reservedBW
		}
	}
}

// set finegrained reserved bandwidth for every link of the network, according to network, flows and routing state
func SetFinegrainedReservedBW(net *network.Network, flows []network.Flow) {
	net.ReservedBW = make([][]float64, len(net.Nodes))
	for i := 0; i < len(net.ReservedBW); i++ {
		net.ReservedBW[i] = make([]float64, len(net.Nodes))
	}
	// set finegrained reserved bandwidth to every link according to rtData and bandwidth
	var reservedBW float64
	for i := 0; i < len(net.FlowIndexes); i++ {
		for j := i + 1; j < len(net.FlowIndexes[i]); j++ { // for every link [i][j]
			var rtDatas []float64
			var rtIndexes []int

			for k := 0; k < len(net.FlowIndexes[i][j]); k++ {
				if isRtFlow(flows[net.FlowIndexes[i][j][k]]) {
					rtDatas = append(rtDatas, dataPerPath(flows[net.FlowIndexes[i][j][k]]))
					rtIndexes = append(rtIndexes, k)
				}
			} // len(rtDatas) == len(rtIndexes) == num of rt flows on link i,j

			var maxSum float64
			var pick []bool
			maxSum, pick = maxElements(rtDatas, bandwidth*0.6)

			reservedBW = maxSum / bandwidth
			for k := 0; k < len(pick); k++ {
				if pick[k] {
					net.RemainIndexes[i][j][rtIndexes[k]] = true // elements in rtIndexes are keys of net.FlowIndexes[i][j] and net.RemainIndexes[i][j]
				}
			}

			net.ReservedBW[i][j] = reservedBW
		}
	}
}

// pick some elements with max sum less than limit
func maxElements(elements []float64, limit float64) (float64, []bool) {
	var remain []bool = make([]bool, len(elements))
	var max float64

	var dp func(limit float64, index int, pick []bool) (float64, []bool)
	dp = func(limit float64, index int, pick []bool) (float64, []bool) {
		var pickCopy []bool = make([]bool, len(pick))
		copy(pickCopy, pick)
		if limit < 0 { // sum cannot be larger than limit
			return -100000000, pickCopy
		}
		if index >= len(elements) { // traversed all
			return 0, pickCopy
		}
		// two choices, choose elements[index] or not choose
		a, picka := dp(limit-elements[index], index+1, pickCopy)
		b, pickb := dp(limit, index+1, pickCopy)
		a = elements[index] + a
		picka[index] = true
		if a > b { // pick the choice with larger sum
			return a, picka
		}
		return b, pickb
	}

	max, remain = dp(limit, 0, remain)

	return max, remain
}

func isRtFlow(flow network.Flow) bool {
	return flow.Deadline >= 0
}

func rmFlowFromLink(net *network.Network, head, tail int, flow *network.Flow, flowIndex int) {
	// find a path containing this link
	var removeIndex int = -1
FINDINDEX:
	for index, path := range flow.Paths {
		for i := 0; i < len(path.Nodes)-1; i++ {
			if (head == path.Nodes[i] && tail == path.Nodes[i+1]) || (tail == path.Nodes[i] && head == path.Nodes[i+1]) {
				removeIndex = index
				break FINDINDEX
			}
		}
	}
	// remove flow index of all links
	for i := 0; i < len(flow.Paths[removeIndex].Nodes)-1; i++ {
		smaller, bigger := smallerBigger(flow.Paths[removeIndex].Nodes[i], flow.Paths[removeIndex].Nodes[i+1])
		j := 0
		for j < len(net.FlowIndexes[smaller][bigger]) {
			if net.FlowIndexes[smaller][bigger][j] == flowIndex {
				break
			}
			j++
		}
		net.FlowIndexes[smaller][bigger] = append(net.FlowIndexes[smaller][bigger][:j], net.FlowIndexes[smaller][bigger][j+1:]...)
		net.RemainIndexes[smaller][bigger] = append(net.RemainIndexes[smaller][bigger][:j], net.RemainIndexes[smaller][bigger][j+1:]...)
	}

	// remove path of the flow
	flow.Paths = append(flow.Paths[:removeIndex], flow.Paths[removeIndex+1:]...)
	flow.Jitters = append(flow.Jitters[:removeIndex], flow.Jitters[removeIndex+1:]...)
}

func smallerBigger(a, b int) (int, int) {
	smaller := a
	bigger := b
	if bigger < smaller {
		smaller, bigger = bigger, smaller
	}
	return smaller, bigger
}

// get net name from index
func GetNetName(i int) string {
	if i < 10 {
		return fmt.Sprintf("net0%d", i)
	} else {
		return fmt.Sprintf("net%d", i)
	}
}

// write routing results to json files
func WriteResults(nets []network.Network, flows [][][]network.Flow) {
	WriteJson(nets, "./experiments/r2t-dsdn-config/jsonnetworks/nets_results.json")
	WriteJson(flows, "./experiments/r2t-dsdn-config/jsonnetworks/flows_results.json")
}

// read routing results from json files
func ReadResults() (nets []network.Network, flows [][][]network.Flow) {
	netsData, err := ioutil.ReadFile("./experiments/r2t-dsdn-config/jsonnetworks/nets_results.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(netsData, &nets)
	if err != nil {
		panic(err)
	}
	flowsData, err := ioutil.ReadFile("./experiments/r2t-dsdn-config/jsonnetworks/flows_results.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(flowsData, &flows)
	if err != nil {
		panic(err)
	}
	return
}

// write results with struct []network.RoutingResult
func WriteTotalResults(results []network.RoutingResult, suffix string) {
	WriteJson(results, "./experiments/r2t-dsdn-config/jsonnetworks/total_results_"+suffix+".json")
}

// read results with struct []network.RoutingResult
func ReadTotalResults() (results []network.RoutingResult) {
	resultsData, err := ioutil.ReadFile("./experiments/r2t-dsdn-config/jsonnetworks/total_results.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resultsData, &results)
	if err != nil {
		panic(err)
	}
	return
}

// write Average Deadline Hit Ratio (%) to file
func OutputADHR(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		// calculate Average Deadline Hit Ratio
		var deadlineHitRatios []float64
		for _, result := range categories[k] { // for every scenario
			deadlineHitRatios = append(deadlineHitRatios, deadlineHitRatio(result))
		}
		averageRatio := average(deadlineHitRatios)
		outputs += fmt.Sprintf("%d %v", k, averageRatio*100)
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}

	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/deadline_hit_ratio_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write Average Delay of Flows (ms) to file
func OutputADoF(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		var averageLagencies []float64 // average latency of all flows in this category

		for _, result := range categories[k] { // for every scenario
			var weightedTotalLatency, totalData float64 // weighted total latency and data of flows
			for i := 0; i < len(result.Flows); i++ {    // for every flow
				var latency float64 // latency of this flow
				if len(result.Flows[i].Paths) == 0 {
					latency = 60 // dropped, set the latency is very large
				} else {
					var latencies []float64 // latencies of all paths of this flow

					for j := 0; j < len(result.Flows[i].Paths); j++ { // for every path
						latencies = append(latencies, float64(result.Flows[i].Paths[j].Latency))
					}
					latency = average(latencies) + result.RoutingTime/float64(k)
				}
				weightedTotalLatency += latency * result.Flows[i].Data
				totalData += result.Flows[i].Data
			}
			weightedLatency := weightedTotalLatency / totalData
			averageLagencies = append(averageLagencies, weightedLatency)
		}
		averagelagency := average(averageLagencies)
		outputs += fmt.Sprintf("%d %v", k, averagelagency)
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}
	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/average_delay_of_flows_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write Average Packet Drop Ratio (%) to file
func OutputAPDR(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		var dropRates []float64
		for _, result := range categories[k] { // for every scenario
			var dropData, totalData float64
			for i := 0; i < len(result.Flows); i++ {
				if len(result.Flows[i].Paths) == 0 {
					dropData += result.Flows[i].Data
				}
				totalData += result.Flows[i].Data
			}
			dropRate := dropData / totalData
			dropRates = append(dropRates, dropRate)
		}
		averageDropRate := average(dropRates)
		outputs += fmt.Sprintf("%d %v", k, averageDropRate*100)
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}
	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/average_packet_drop_ratio_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write Average Non-RT Packet Drop Ratio (%) to file
func OutputANRPDR(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		var dropRates []float64
		for _, result := range categories[k] { // for every scenario
			var dropData, totalData float64
			for i := 0; i < len(result.Flows); i++ {
				if !isRtFlow(result.Flows[i]) {
					if len(result.Flows[i].Paths) == 0 {
						dropData += result.Flows[i].Data
					}
					totalData += result.Flows[i].Data
				}
			}
			dropRate := dropData / totalData
			dropRates = append(dropRates, dropRate)
		}
		averageDropRate := average(dropRates)
		outputs += fmt.Sprintf("%d %v", k, averageDropRate*100)
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}
	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/average_non_rt_packet_drop_ratio_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write Average RT Packet Drop Ratio (%) to file
func OutputARPDR(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		var dropRates []float64
		for _, result := range categories[k] { // for every scenario
			var dropData, totalData float64
			for i := 0; i < len(result.Flows); i++ {
				if isRtFlow(result.Flows[i]) {
					if len(result.Flows[i].Paths) == 0 {
						dropData += result.Flows[i].Data
					}
					totalData += result.Flows[i].Data
				}
			}
			dropRate := dropData / totalData
			dropRates = append(dropRates, dropRate)
		}
		averageDropRate := average(dropRates)
		outputs += fmt.Sprintf("%d %v", k, averageDropRate*100)
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}
	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/average_rt_packet_drop_ratio_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write Average Routing Time of Flows (ms) to file
func OutputARToF(results []network.RoutingResult, suffix string) {
	// categorize results
	var categories map[int][]network.RoutingResult = categorize(results)
	var keys []int
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var outputs string
	for index, k := range keys { // for every category, i.e. scenario group
		var routingTimes []float64
		for _, result := range categories[k] { // for every scenario
			routingTimes = append(routingTimes, result.RoutingTime)
		}
		averageRoutingTime := average(routingTimes)
		outputs += fmt.Sprintf("%d %v", k, averageRoutingTime/float64(k))
		if index != len(keys)-1 {
			outputs += fmt.Sprint("\n")
		}
	}
	err := ioutil.WriteFile("./experiments/r2t-dsdn-config/jsonnetworks/average_routing_time_of_flows_"+suffix+".data", []byte(outputs), 0777)
	if err != nil {
		panic(err)
	}
}

// write all data to files
func OutputData(results []network.RoutingResult, suffix string) {
	OutputADHR(results, suffix)   // Average Deadline Hit Ratio (%)
	OutputADoF(results, suffix)   // Average Delay of Flows (ms)
	OutputAPDR(results, suffix)   // Average Packet Drop Ratio (%)
	OutputARToF(results, suffix)  // Average Routing Time of Flows (ms)
	OutputANRPDR(results, suffix) // Average Non-RT Packet Drop Ratio (%)
	OutputARPDR(results, suffix)  // Average RT Packet Drop Ratio (%)
}

// categorize results according to the number of flows in every scenario
func categorize(results []network.RoutingResult) map[int][]network.RoutingResult {
	var categories map[int][]network.RoutingResult = make(map[int][]network.RoutingResult)

	for i := 0; i < len(results); i++ {
		categories[len(results[i].Flows)] = append(categories[len(results[i].Flows)], results[i])
	}

	return categories
}

// calculate deadline hit ratio of a scenario
func deadlineHitRatio(result network.RoutingResult) float64 {
	var rtFlows []network.Flow
	for i := 0; i < len(result.Flows); i++ {
		if isRtFlow(result.Flows[i]) {
			rtFlows = append(rtFlows, result.Flows[i])
		}
	}

	averageRoutingTime := result.RoutingTime / float64(len(result.Flows))

	var hitData, totalData float64
	for i := 0; i < len(rtFlows); i++ {
		if deadlineHit(rtFlows[i], averageRoutingTime) {
			hitData += rtFlows[i].Data
		}
		totalData += rtFlows[i].Data
	}
	return hitData / totalData
}

// whether the deadline of an RT-flow hit
func deadlineHit(flow network.Flow, routingTime float64) bool {
	if len(flow.Paths) == 0 {
		return false
	}
	var lagencies []float64
	for i := 0; i < len(flow.Paths); i++ {
		lagencies = append(lagencies, float64(flow.Paths[i].Latency))
	}
	return flow.Deadline >= average(lagencies)+routingTime
}

func average(nums []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}
