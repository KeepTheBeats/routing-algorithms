package network

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeNetOnePath1() Network {
	var net Network
	net.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 3, -1, 7},
		{4, 3, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 3, 4, 0, 6},
		{3, 0, 0, 10, 10},
		{4, 0, 0, 4, 9},
		{0, 10, 4, 0, 6},
		{6, 10, 9, 6, 0},
	}
	return net
}

func fakeNetOnePath2() Network {
	var net Network
	net.Links = [][]int{
		{0, 1, -1, 2},
		{1, 0, 2, 4},
		{-1, 2, 0, 7},
		{2, 4, 7, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 1, 2, 0},
		{1, 0, 3, 3},
		{2, 3, 0, 9},
		{0, 3, 9, 0},
	}
	return net
}

func fakeNetTwoPaths1() Network {
	var net Network
	net.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 2, -1, 7},
		{4, 2, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 2, 8, 2, 40},
		{2, 0, 10, 17, 8},
		{8, 10, 0, 9, 14},
		{2, 17, 9, 0, 16},
		{40, 8, 14, 16, 0},
	}
	return net
}

func fakeNetTwoPaths2() Network {
	var net Network
	net.Links = [][]int{
		{0, 1, -1, 2},
		{1, 0, 2, 4},
		{-1, 2, 0, 1},
		{2, 4, 1, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 10, 18, 4},
		{10, 0, 13, 11},
		{18, 13, 0, 8},
		{4, 11, 8, 0},
	}
	return net
}

func fakeUnconnectedNet1() Network {
	var net Network
	net.Links = [][]int{
		{0, -1, -1, -1, 10},
		{-1, 0, 2, -1, -1},
		{-1, 2, 0, 4, -1},
		{-1, -1, 4, 0, -1},
		{10, -1, -1, -1, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 4, 1, 9, 4},
		{4, 0, 16, 7, 18},
		{1, 16, 0, 15, 16},
		{9, 7, 15, 0, 7},
		{4, 18, 16, 7, 0},
	}
	return net
}

func fakeUnconnectedNet2() Network {
	var net Network
	net.Links = [][]int{
		{0, -1, -1, -1, 10},
		{-1, 0, -1, -1, -1},
		{-1, -1, 0, 4, -1},
		{-1, -1, 4, 0, -1},
		{10, -1, -1, -1, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 14, 2, 9, 7},
		{14, 0, 9, 15, 7},
		{2, 9, 0, 16, 1},
		{9, 15, 16, 0, 3},
		{7, 7, 1, 3, 0},
	}
	return net
}

func fakeUnconnectedNet3() Network {
	var net Network
	net.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 3, -1, 7},
		{4, 3, 0, -1, 3},
		{-1, -1, -1, 0, -1},
		{10, 7, 3, -1, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 2, 12, 17, 14},
		{2, 0, 5, 19, 11},
		{12, 5, 0, 0, 1},
		{17, 19, 0, 0, 6},
		{14, 11, 1, 6, 0},
	}
	return net
}

func fakeDirectedNet1() Network {
	var net Network
	net.Links = [][]int{
		{0, 3, 2, -1, -1, -1},
		{-1, 0, -1, 4, -1, -1},
		{-1, 1, 0, 2, 3, -1},
		{-1, -1, -1, 0, 2, 1},
		{-1, -1, -1, -1, 0, 2},
		{-1, -1, -1, -1, -1, 0},
	}
	net.Nodes = make([]Node, len(net.Links))
	net.Jitters = [][]int{
		{0, 10, 10, 1, 9, 6},
		{4, 0, 7, 1, 8, 2},
		{1, 10, 0, 6, 5, 7},
		{0, 3, 6, 0, 2, 7},
		{4, 9, 5, 1, 0, 4},
		{8, 2, 5, 8, 3, 0},
	}
	return net
}

func TestShortest(t *testing.T) {
	testCases := []struct {
		name          string
		net           Network
		flow          Flow
		expectedPaths []Path
	}{
		{
			name: "case1",
			net:  fakeNetOnePath1(),
			flow: Flow{
				Source:      0,
				Destination: 1,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 1},
					Latency: 2,
				},
			},
		},
		{
			name: "case2",
			net:  fakeNetOnePath1(),
			flow: Flow{
				Source:      0,
				Destination: 2,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2},
					Latency: 4,
				},
			},
		},
		{
			name: "case3",
			net:  fakeNetOnePath1(),
			flow: Flow{
				Source:      0,
				Destination: 0,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0},
					Latency: 0,
				},
			},
		},
		{
			name: "case4",
			net:  fakeNetOnePath1(),
			flow: Flow{
				Source:      0,
				Destination: 3,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 3},
					Latency: 8,
				},
			},
		},
		{
			name: "case5",
			net:  fakeNetOnePath1(),
			flow: Flow{
				Source:      0,
				Destination: 4,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 4},
					Latency: 7,
				},
			},
		},
		{
			name: "case6",
			net:  fakeNetTwoPaths1(),
			flow: Flow{
				Source:      0,
				Destination: 2,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2},
					Latency: 4,
				},
				{
					Nodes:   []int{0, 1, 2},
					Latency: 4,
				},
			},
		},
		{
			name: "case7",
			net:  fakeNetTwoPaths1(),
			flow: Flow{
				Source:      0,
				Destination: 3,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 1, 2, 3},
					Latency: 8,
				},
				{
					Nodes:   []int{0, 2, 3},
					Latency: 8,
				},
			},
		},
		{
			name: "case8",
			net:  fakeNetTwoPaths1(),
			flow: Flow{
				Source:      0,
				Destination: 4,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 4},
					Latency: 7,
				},
				{
					Nodes:   []int{0, 1, 2, 4},
					Latency: 7,
				},
			},
		},
		{
			name: "case9",
			net:  fakeNetTwoPaths2(),
			flow: Flow{
				Source:      2,
				Destination: 0,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{2, 1, 0},
					Latency: 3,
				},
				{
					Nodes:   []int{2, 3, 0},
					Latency: 3,
				},
			},
		},
		{
			name: "case10",
			net:  fakeNetOnePath2(),
			flow: Flow{
				Source:      3,
				Destination: 2,
			},
			expectedPaths: []Path{
				{
					Nodes:   []int{3, 0, 1, 2},
					Latency: 5,
				},
			},
		},
		{
			name: "case11",
			net:  fakeUnconnectedNet1(),
			flow: Flow{
				Source:      1,
				Destination: 4,
			},
			expectedPaths: []Path{},
		},
		{
			name: "case12",
			net:  fakeUnconnectedNet1(),
			flow: Flow{
				Source:      0,
				Destination: 2,
			},
			expectedPaths: []Path{},
		},
		{
			name: "case13",
			net:  fakeUnconnectedNet3(),
			flow: Flow{
				Source:      1,
				Destination: 3,
			},
			expectedPaths: []Path{},
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualPaths := Shortest(testCase.net, testCase.flow)
		assert.ElementsMatch(t, testCase.expectedPaths, actualPaths, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestDijkstra(t *testing.T) {
	var generateExpectedPaths func(Network, int) [][]Path = func(net Network, source int) [][]Path {
		var expectedPaths [][]Path = make([][]Path, len(net.Nodes))
		for i := 0; i < len(net.Nodes); i++ {
			expectedPaths[i] = Shortest(net, Flow{Source: source, Destination: i})
		}
		return expectedPaths
	}
	testCases := []struct {
		name          string
		net           Network
		source        int
		expectedPaths [][]Path
	}{
		{
			name:   "one path case1",
			net:    fakeNetOnePath1(),
			source: 0,
			expectedPaths: [][]Path{
				[]Path{
					Path{
						Nodes:   []int{0},
						Latency: 0,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 1},
						Latency: 2,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2},
						Latency: 4,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2, 3},
						Latency: 8,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2, 4},
						Latency: 7,
					},
				},
			},
		},
		{
			name:   "one path case2",
			net:    fakeNetOnePath2(),
			source: 3,
			expectedPaths: [][]Path{
				[]Path{
					Path{
						Nodes:   []int{3, 0},
						Latency: 2,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{3, 0, 1},
						Latency: 3,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{3, 0, 1, 2},
						Latency: 5,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{3},
						Latency: 0,
					},
				},
			},
		},
		{
			name:          "one path case3",
			net:           fakeNetOnePath1(),
			source:        1,
			expectedPaths: generateExpectedPaths(fakeNetOnePath1(), 1),
		},
		{
			name:          "one path case4",
			net:           fakeNetOnePath1(),
			source:        2,
			expectedPaths: generateExpectedPaths(fakeNetOnePath1(), 2),
		},
		{
			name:          "one path case5",
			net:           fakeNetOnePath1(),
			source:        3,
			expectedPaths: generateExpectedPaths(fakeNetOnePath1(), 3),
		},
		{
			name:          "one path case6",
			net:           fakeNetOnePath1(),
			source:        4,
			expectedPaths: generateExpectedPaths(fakeNetOnePath1(), 4),
		},
		{
			name:   "two paths case1",
			net:    fakeNetTwoPaths1(),
			source: 0,
			expectedPaths: [][]Path{
				[]Path{
					Path{
						Nodes:   []int{0},
						Latency: 0,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 1},
						Latency: 2,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2},
						Latency: 4,
					},
					Path{
						Nodes:   []int{0, 1, 2},
						Latency: 4,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2, 3},
						Latency: 8,
					},
					Path{
						Nodes:   []int{0, 1, 2, 3},
						Latency: 8,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2, 4},
						Latency: 7,
					},
					Path{
						Nodes:   []int{0, 1, 2, 4},
						Latency: 7,
					},
				},
			},
		},
		{
			name:   "two paths case2 different orders",
			net:    fakeNetTwoPaths1(),
			source: 0,
			expectedPaths: [][]Path{
				[]Path{
					Path{
						Nodes:   []int{0},
						Latency: 0,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 1},
						Latency: 2,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2},
						Latency: 4,
					},
					Path{
						Nodes:   []int{0, 1, 2},
						Latency: 4,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 2, 3},
						Latency: 8,
					},
					Path{
						Nodes:   []int{0, 1, 2, 3},
						Latency: 8,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{0, 1, 2, 4},
						Latency: 7,
					},
					Path{
						Nodes:   []int{0, 2, 4},
						Latency: 7,
					},
				},
			},
		},
		{
			name:   "two paths case2",
			net:    fakeNetTwoPaths2(),
			source: 2,
			expectedPaths: [][]Path{
				[]Path{
					Path{
						Nodes:   []int{2, 1, 0},
						Latency: 3,
					},
					Path{
						Nodes:   []int{2, 3, 0},
						Latency: 3,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{2, 1},
						Latency: 2,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{2},
						Latency: 0,
					},
				},
				[]Path{
					Path{
						Nodes:   []int{2, 3},
						Latency: 1,
					},
				},
			},
		},
		{
			name:          "two paths case3",
			net:           fakeNetTwoPaths2(),
			source:        1,
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths2(), 1),
		},
		{
			name:          "two paths case4",
			net:           fakeNetTwoPaths1(),
			source:        1,
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths1(), 1),
		},
		{
			name:          "two paths case5",
			net:           fakeNetTwoPaths1(),
			source:        2,
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths1(), 2),
		},
		{
			name:          "two paths case6",
			net:           fakeNetTwoPaths1(),
			source:        3,
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths1(), 3),
		},
		{
			name:          "two paths case7",
			net:           fakeNetTwoPaths1(),
			source:        4,
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths1(), 4),
		},
		{
			name:          "unconnected case1",
			net:           fakeUnconnectedNet3(),
			source:        2,
			expectedPaths: generateExpectedPaths(fakeUnconnectedNet3(), 2),
		},
		{
			name:          "unconnected case2",
			net:           fakeUnconnectedNet2(),
			source:        0,
			expectedPaths: generateExpectedPaths(fakeUnconnectedNet2(), 0),
		},
	}
	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualPaths := Dijkstra(testCase.net, testCase.source)
		if len(testCase.expectedPaths) != len(actualPaths) {
			t.Errorf("Fail!! %s: lengths are unequal", testCase.name)
		} else {
			for i := 0; i < len(testCase.net.Nodes); i++ {
				assert.ElementsMatch(t, testCase.expectedPaths[i], actualPaths[i], fmt.Sprintf("%s, paths to node %d is not expected", testCase.name, i))
			}
		}
	}
}

func TestCheckConnected(t *testing.T) {
	testCases := []struct {
		name           string
		net            Network
		expectedResult bool
	}{
		{
			name:           "unconnected case1",
			net:            fakeUnconnectedNet1(),
			expectedResult: false,
		},
		{
			name:           "unconnected case2",
			net:            fakeUnconnectedNet2(),
			expectedResult: false,
		},
		{
			name:           "unconnected case3",
			net:            fakeUnconnectedNet3(),
			expectedResult: false,
		},
		{
			name:           "connected case1",
			net:            fakeNetOnePath1(),
			expectedResult: true,
		},
		{
			name:           "connected case2",
			net:            fakeNetTwoPaths1(),
			expectedResult: true,
		},
		{
			name:           "connected case3",
			net:            fakeNetTwoPaths2(),
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := CheckConnected(testCase.net)
		assert.Equal(t, testCase.expectedResult, actualResult, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestFloyd(t *testing.T) {
	var generateExpectedPaths func(Network) [][][]Path = func(net Network) [][][]Path {
		expectedPaths := [][][]Path{}
		for i := 0; i < len(net.Nodes); i++ {
			expectedPaths = append(expectedPaths, Dijkstra(net, i))
		}
		return expectedPaths
	}
	testCases := []struct {
		name          string
		net           Network
		expectedPaths [][][]Path
	}{
		{
			name: "one path case1",
			net:  fakeNetOnePath1(),
			expectedPaths: [][][]Path{
				[][]Path{
					[]Path{
						Path{
							Nodes:   []int{0},
							Latency: 0,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{0, 1},
							Latency: 2,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{0, 2},
							Latency: 4,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{0, 2, 3},
							Latency: 8,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{0, 2, 4},
							Latency: 7,
						},
					},
				},
				[][]Path{
					[]Path{
						Path{
							Nodes:   []int{1, 0},
							Latency: 2,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{1},
							Latency: 0,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{1, 2},
							Latency: 3,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{1, 2, 3},
							Latency: 7,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{1, 2, 4},
							Latency: 6,
						},
					},
				},
				[][]Path{
					[]Path{
						Path{
							Nodes:   []int{2, 0},
							Latency: 4,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{2, 1},
							Latency: 3,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{2},
							Latency: 0,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{2, 3},
							Latency: 4,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{2, 4},
							Latency: 3,
						},
					},
				},
				[][]Path{
					[]Path{
						Path{
							Nodes:   []int{3, 2, 0},
							Latency: 8,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{3, 2, 1},
							Latency: 7,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{3, 2},
							Latency: 4,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{3},
							Latency: 0,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{3, 4},
							Latency: 5,
						},
					},
				},
				[][]Path{
					[]Path{
						Path{
							Nodes:   []int{4, 2, 0},
							Latency: 7,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{4, 2, 1},
							Latency: 6,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{4, 2},
							Latency: 3,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{4, 3},
							Latency: 5,
						},
					},
					[]Path{
						Path{
							Nodes:   []int{4},
							Latency: 0,
						},
					},
				},
			},
		},
		{
			name:          "one path case2",
			net:           fakeNetOnePath2(),
			expectedPaths: generateExpectedPaths(fakeNetOnePath2()),
		},
		{
			name:          "two paths case1",
			net:           fakeNetTwoPaths1(),
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths1()),
		},
		{
			name:          "two paths case2",
			net:           fakeNetTwoPaths2(),
			expectedPaths: generateExpectedPaths(fakeNetTwoPaths2()),
		},
		{
			name:          "unconnected case1",
			net:           fakeUnconnectedNet1(),
			expectedPaths: generateExpectedPaths(fakeUnconnectedNet1()),
		},
		{
			name:          "unconnected case2",
			net:           fakeUnconnectedNet2(),
			expectedPaths: generateExpectedPaths(fakeUnconnectedNet2()),
		},
		{
			name:          "unconnected case3",
			net:           fakeUnconnectedNet3(),
			expectedPaths: generateExpectedPaths(fakeUnconnectedNet3()),
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualPaths := Floyd(testCase.net)
		if len(testCase.expectedPaths) != len(actualPaths) {
			t.Errorf("Fail!! %s: lengths are unequal", testCase.name)
		} else {
			for i := 0; i < len(testCase.net.Nodes); i++ {
				if len(testCase.expectedPaths[i]) != len(actualPaths[i]) {
					t.Errorf("Fail!! %s: lengths of paths from %d are unequal", testCase.name, i)
				} else {
					for j := 0; j < len(testCase.net.Nodes); j++ {
						assert.ElementsMatch(t, testCase.expectedPaths[i][j], actualPaths[i][j], fmt.Sprintf("%s, paths from node %d to node %d is not expected", testCase.name, i, j))
					}
				}
			}
		}
	}
}

func TestKShortest(t *testing.T) {
	testCases := []struct {
		name          string
		net           Network
		flow          Flow
		k             int
		expectedPaths []Path
	}{
		{
			name: "directed graph case1",
			net:  fakeDirectedNet1(),
			flow: Flow{Source: 0, Destination: 5},
			k:    3,
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 3, 5},
					Latency: 5,
				},
				{
					Nodes:   []int{0, 2, 4, 5},
					Latency: 7,
				},
				{
					Nodes:   []int{0, 1, 3, 5},
					Latency: 8,
				},
			},
		},
		{
			name: "undirected graph case1",
			net:  fakeNetTwoPaths1(),
			flow: Flow{Source: 0, Destination: 3},
			k:    3,
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 3},
					Latency: 8,
				},
				{
					Nodes:   []int{0, 1, 2, 3},
					Latency: 8,
				},
				{
					Nodes:   []int{0, 2, 4, 3},
					Latency: 12,
				},
			},
		},
		{
			name: "undirected graph case2",
			net:  fakeNetTwoPaths1(),
			flow: Flow{Source: 0, Destination: 3},
			k:    2,
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 3},
					Latency: 8,
				},
				{
					Nodes:   []int{0, 1, 2, 3},
					Latency: 8,
				},
			},
		},
		{
			name: "undirected graph case3",
			net:  fakeNetOnePath1(),
			flow: Flow{Source: 0, Destination: 4},
			k:    4,
			expectedPaths: []Path{
				{
					Nodes:   []int{0, 2, 4},
					Latency: 7,
				},
				{
					Nodes:   []int{0, 1, 2, 4},
					Latency: 8,
				},
				{
					Nodes:   []int{0, 1, 4},
					Latency: 9,
				},
				{
					Nodes:   []int{0, 4},
					Latency: 10,
				},
			},
		},
		{
			name:          "undirected graph case4",
			net:           fakeNetOnePath1(),
			flow:          Flow{Source: 0, Destination: 4},
			k:             4,
			expectedPaths: KShortestDFS(fakeNetOnePath1(), Flow{Source: 0, Destination: 4}, 4),
		},
		{
			name:          "undirected graph case5",
			net:           fakeNetOnePath1(),
			flow:          Flow{Source: 2, Destination: 4},
			k:             4,
			expectedPaths: KShortestDFS(fakeNetOnePath1(), Flow{Source: 2, Destination: 4}, 4),
		},
		{
			name:          "undirected graph case6",
			net:           fakeNetOnePath1(),
			flow:          Flow{Source: 2, Destination: 4},
			k:             10,
			expectedPaths: KShortestDFS(fakeNetOnePath1(), Flow{Source: 2, Destination: 4}, 10),
		},
		{
			name:          "undirected graph case7",
			net:           fakeNetTwoPaths1(),
			flow:          Flow{Source: 1, Destination: 3},
			k:             10,
			expectedPaths: KShortestDFS(fakeNetTwoPaths1(), Flow{Source: 1, Destination: 3}, 10),
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualPaths := KShortest(testCase.net, testCase.flow, testCase.k)
		assert.Equal(t, testCase.expectedPaths, actualPaths, fmt.Sprintf("%s: result is not expected", testCase.name))
	}

	var testAllFlows func(Network, int) = func(net Network, k int) {
		fmt.Printf("testAllFlows\n")
		for source := 0; source < len(net.Nodes); source++ {
			for destination := 0; destination < len(net.Nodes); destination++ {
				if source == destination {
					continue
				}
				flow := Flow{Source: source, Destination: destination}
				dfsPaths := KShortestDFS(net, flow, k)
				yenPaths := KShortest(net, flow, k)
				// if len(dfsPaths) > 6 {
				// 	fmt.Printf("dfsPaths:\n%v\n\n", dfsPaths)
				// 	fmt.Printf("yenPaths:\n%v\n\n", yenPaths)
				// }
				if len(dfsPaths) != len(yenPaths) {
					t.Errorf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\nFail!!: lengths of dfsPaths (%d) and yenPaths (%d) are unequal", net, flow, k, len(dfsPaths), len(yenPaths))
				}

				for i := 0; i < len(dfsPaths); i++ {
					if dfsPaths[i].Latency != yenPaths[i].Latency {
						t.Errorf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\nFail!!: i: %d latency is not equal\n\ndfsPaths:\n%v\n\nyenPaths:\n%v\n\n", net, flow, k, i, dfsPaths, yenPaths)
					}
					if len(dfsPaths[i].Nodes) != len(yenPaths[i].Nodes) {
						t.Errorf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\nFail!!: i: %d len(Nodes) is not equal\n\ndfsPaths:\n%v\n\nyenPaths:\n%v\n\n", net, flow, k, i, dfsPaths, yenPaths)
					}
				}

				dfsGroups := map[[2]int][]Path{}
				yenGroups := map[[2]int][]Path{}

				for _, path := range dfsPaths {
					dfsGroups[[2]int{path.Latency, len(path.Nodes)}] = append(dfsGroups[[2]int{path.Latency, len(path.Nodes)}], path)
				}
				for _, path := range yenPaths {
					yenGroups[[2]int{path.Latency, len(path.Nodes)}] = append(yenGroups[[2]int{path.Latency, len(path.Nodes)}], path)
				}

				if len(dfsGroups) != len(yenGroups) {
					t.Errorf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\nFail!!: lengths of dfsGroups (%d) and yenGroups (%d) are unequal", net, flow, k, len(dfsGroups), len(yenGroups))
				} else {
					for key, dfsGroup := range dfsGroups {
						if yenGroup, exist := yenGroups[key]; !exist {
							t.Errorf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\nFail!!: key: %v in dfsGroups does not exist in yenGroups", net, flow, k, key)
						} else {
							assert.ElementsMatch(t, dfsGroup, yenGroup, fmt.Sprintf("\nnet:\n%v\n\nflow: %v\n\nk: %d\n\ndfsGroup:\n%v\n\nyenGroup:\n%v\n\ndfsPaths:\n%v\n\nyenPaths:\n%v\n\n", net, flow, k, dfsGroup, yenGroup, dfsPaths, yenPaths))
							// still have deficiency, when more than one paths tie for the number k shortest path, dfs and yen may choose different paths, should notice this. For this condition, can print dfsPaths and yenPaths and compare them manually.
						}
					}
				}
			}
		}
	}
	testAllFlows(fakeNetOnePath1(), 4)
	testAllFlows(fakeNetOnePath1(), 6)
	testAllFlows(fakeNetOnePath1(), 7)
	testAllFlows(fakeNetOnePath2(), 1)
	testAllFlows(fakeNetOnePath2(), 2)
	testAllFlows(fakeNetOnePath2(), 3)
	testAllFlows(fakeNetOnePath2(), 5)
	testAllFlows(fakeNetOnePath2(), 10)
	testAllFlows(fakeNetTwoPaths1(), 1)
	testAllFlows(fakeNetTwoPaths1(), 3)
	testAllFlows(fakeNetTwoPaths1(), 13)
	testAllFlows(fakeNetTwoPaths1(), 15)
	testAllFlows(fakeNetTwoPaths2(), 5)
	testAllFlows(fakeNetTwoPaths2(), 10)
}

func TestInnerR2tdsdnRel(t *testing.T) {
	testCases := []struct {
		name           string
		path           Path
		switchRel      float64
		linkRel        float64
		expectedResult float64
	}{
		{
			name:           "case1",
			path:           Path{Nodes: []int{1, 4, 0, 2, 3}, Latency: 25},
			switchRel:      0.8,
			linkRel:        0.2,
			expectedResult: math.Pow(0.8, 5) * math.Pow(0.2, 4),
		},
		{
			name:           "case2",
			path:           Path{Nodes: []int{1, 2, 0, 4, 3}, Latency: 24},
			switchRel:      0.8,
			linkRel:        0.2,
			expectedResult: math.Pow(0.8, 5) * math.Pow(0.2, 4),
		},
		{
			name:           "case3",
			path:           Path{Nodes: []int{1, 6, 8, 4, 3, 8}, Latency: 15},
			switchRel:      0.5,
			linkRel:        0.3,
			expectedResult: math.Pow(0.5, 6) * math.Pow(0.3, 5),
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := r2tdsdnRel(testCase.path, testCase.switchRel, testCase.linkRel)
		assert.Equal(t, testCase.expectedResult, actualResult, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestInnerR2tdsdnJitter(t *testing.T) {
	testCases := []struct {
		name           string
		net            Network
		path           Path
		expectedResult int
	}{
		{
			name:           "case1",
			net:            fakeNetOnePath1(),
			path:           Path{Nodes: []int{1, 2, 3}, Latency: 7},
			expectedResult: 4,
		},
		{
			name:           "case2",
			net:            fakeNetOnePath1(),
			path:           Path{Nodes: []int{1, 0, 2, 3}, Latency: 10},
			expectedResult: 11,
		},
		{
			name:           "case3",
			net:            fakeNetOnePath1(),
			path:           Path{Nodes: []int{1, 4, 2, 3}, Latency: 14},
			expectedResult: 23,
		},
		{
			name:           "case4",
			net:            fakeNetTwoPaths1(),
			path:           Path{Nodes: []int{1, 2, 0, 4, 3}, Latency: 21},
			expectedResult: 74,
		},
		{
			name:           "case5",
			net:            fakeNetTwoPaths1(),
			path:           Path{Nodes: []int{1, 4, 0, 2, 3}, Latency: 25},
			expectedResult: 65,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := r2tdsdnJitter(testCase.net, testCase.path)
		assert.Equal(t, testCase.expectedResult, actualResult, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestR2tdsdnRouting(t *testing.T) {
	testCases := []struct {
		name               string
		net                Network
		flow               Flow
		k                  int
		switchRel, linkRel float64
		expectedPaths      []Path
	}{
		{
			name:      "case1",
			net:       fakeNetOnePath1(),
			flow:      Flow{Source: 1, Destination: 3, DesirableJitter: 50},
			k:         10,
			switchRel: 0.8,
			linkRel:   0.2,
			expectedPaths: []Path{
				Path{Nodes: []int{1, 2, 3}, Latency: 7},
				Path{Nodes: []int{1, 4, 3}, Latency: 12},
			},
		},
		{
			name:      "case2",
			net:       fakeNetOnePath1(),
			flow:      Flow{Source: 1, Destination: 3, DesirableJitter: 16},
			k:         10,
			switchRel: 0.8,
			linkRel:   0.2,
			expectedPaths: []Path{
				Path{Nodes: []int{1, 2, 3}, Latency: 7},
				Path{Nodes: []int{1, 4, 3}, Latency: 12},
			},
		},
		{
			name:      "case3",
			net:       fakeNetOnePath1(),
			flow:      Flow{Source: 1, Destination: 3, DesirableJitter: 15},
			k:         10,
			switchRel: 0.8,
			linkRel:   0.2,
			expectedPaths: []Path{
				Path{Nodes: []int{1, 2, 3}, Latency: 7},
			},
		},
		{
			name:      "case4",
			net:       fakeNetOnePath1(),
			flow:      Flow{Source: 1, Destination: 3, DesirableJitter: 50},
			k:         4,
			switchRel: 0.8,
			linkRel:   0.2,
			expectedPaths: []Path{
				Path{Nodes: []int{1, 2, 3}, Latency: 7},
				Path{Nodes: []int{1, 4, 3}, Latency: 12},
			},
		},
		{
			name:      "case5",
			net:       fakeNetOnePath1(),
			flow:      Flow{Source: 1, Destination: 3, DesirableJitter: 50},
			k:         3,
			switchRel: 0.8,
			linkRel:   0.2,
			expectedPaths: []Path{
				Path{Nodes: []int{1, 2, 3}, Latency: 7},
			},
		},
		{
			name:      "case6",
			net:       fakeNetTwoPaths1(),
			flow:      Flow{Source: 0, Destination: 4, DesirableJitter: 50},
			k:         10,
			switchRel: 0.6,
			linkRel:   0.3,
			expectedPaths: []Path{
				Path{Nodes: []int{0, 4}, Latency: 10},
			},
		},
		{
			name:      "case7",
			net:       fakeNetTwoPaths1(),
			flow:      Flow{Source: 0, Destination: 4, DesirableJitter: 39},
			k:         10,
			switchRel: 0.6,
			linkRel:   0.3,
			expectedPaths: []Path{
				Path{Nodes: []int{0, 2, 4}, Latency: 7},
				Path{Nodes: []int{0, 1, 4}, Latency: 9},
			},
		},
		{
			name:      "case8",
			net:       fakeNetTwoPaths1(),
			flow:      Flow{Source: 0, Destination: 4, DesirableJitter: 40},
			k:         3,
			switchRel: 0.6,
			linkRel:   0.3,
			expectedPaths: []Path{
				Path{Nodes: []int{0, 2, 4}, Latency: 7},
				Path{Nodes: []int{0, 1, 4}, Latency: 9},
			},
		},
		{
			name:      "case9",
			net:       fakeNetTwoPaths1(),
			flow:      Flow{Source: 0, Destination: 4, DesirableJitter: 21},
			k:         10,
			switchRel: 0.6,
			linkRel:   0.3,
			expectedPaths: []Path{
				Path{Nodes: []int{0, 1, 4}, Latency: 9},
			},
		},
		{
			name:          "case10",
			net:           fakeNetTwoPaths1(),
			flow:          Flow{Source: 0, Destination: 4, DesirableJitter: 21},
			k:             2,
			switchRel:     0.6,
			linkRel:       0.3,
			expectedPaths: []Path{},
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)

		actualResult, jitters, reliabilities := R2tdsdnRouting(testCase.net, testCase.flow, testCase.k, testCase.switchRel, testCase.linkRel)
		assert.ElementsMatch(t, testCase.expectedPaths, actualResult, fmt.Sprintf("%s: result is not expected", testCase.name))
		if len(actualResult) != len(jitters) || len(actualResult) != len(reliabilities) {
			t.Errorf("Fail!! %s: lengths of actualResult, jitters, reliabilities are unequal", testCase.name)
		} else {
			for i := 0; i < len(actualResult); i++ {
				assert.Equal(t, jitters[i], r2tdsdnJitter(testCase.net, actualResult[i]), fmt.Sprintf("%s: jitters[%d] is not expected", testCase.name, i))
				assert.Equal(t, reliabilities[i], r2tdsdnRel(actualResult[i], testCase.switchRel, testCase.linkRel), fmt.Sprintf("%s: reliabilities[%d] is not expected", testCase.name, i))
			}
		}
	}
}
