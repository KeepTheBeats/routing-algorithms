package network

import (
	"fmt"
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
