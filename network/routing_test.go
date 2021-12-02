package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortest(t *testing.T) {
	var net1 Network
	net1.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 3, -1, 7},
		{4, 3, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net1.Nodes = make([]Node, len(net1.Links))

	var net2 Network
	net2.Links = [][]int{
		{0, 2, 4, -1, 10},
		{2, 0, 2, -1, 7},
		{4, 2, 0, 4, 3},
		{-1, -1, 4, 0, 5},
		{10, 7, 3, 5, 0},
	}
	net2.Nodes = make([]Node, len(net2.Links))

	testCases := []struct {
		name          string
		net           Network
		flow          Flow
		expectedPaths []Path
	}{
		{
			name: "case1",
			net:  net1,
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
			net:  net1,
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
			net:  net1,
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
			net:  net1,
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
			net:  net1,
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
			net:  net2,
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
			net:  net2,
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
			net:  net2,
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
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualPaths := Shortest(testCase.net, testCase.flow)
		assert.ElementsMatch(t, testCase.expectedPaths, actualPaths, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}
