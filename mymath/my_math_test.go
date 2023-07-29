package mymath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDelta float64 = 0.00001

func TestUnitRound(t *testing.T) {

	testCases := []struct {
		name           string
		x              float64
		unit           float64
		expectedResult float64
	}{
		{
			name:           "case1",
			x:              2.33,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case2",
			x:              2.35,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case3",
			x:              2.36,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case4",
			x:              2.4,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case5",
			x:              2.3,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case6",
			x:              0.3398,
			unit:           0.1,
			expectedResult: 0.3,
		},
		{
			name:           "case7",
			x:              2.2398,
			unit:           0.3,
			expectedResult: 2.1,
		},
		{
			name:           "case8",
			x:              4.38732,
			unit:           0.01,
			expectedResult: 4.39,
		},
		{
			name:           "case9",
			x:              4.38432,
			unit:           0.01,
			expectedResult: 4.38,
		},
		{
			name:           "case10",
			x:              4.38432,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "case11",
			x:              4.58432,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "exact division 1",
			x:              4.58,
			unit:           0.01,
			expectedResult: 4.58,
		},
		{
			name:           "exact division 2",
			x:              4.57,
			unit:           0.01,
			expectedResult: 4.57,
		},
		{
			name:           "exact division 3",
			x:              4.5,
			unit:           0.1,
			expectedResult: 4.5,
		},
		{
			name:           "exact division 4",
			x:              4.2,
			unit:           0.1,
			expectedResult: 4.2,
		},
		{
			name:           "exact division 5",
			x:              2.2222,
			unit:           0.0001,
			expectedResult: 2.2222,
		},
		{
			name:           "exact division 6",
			x:              2.2227,
			unit:           0.0001,
			expectedResult: 2.2227,
		},
		{
			name:           "exact division 7",
			x:              4,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "exact division 8",
			x:              5,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "exact division 9",
			x:              6,
			unit:           1,
			expectedResult: 6,
		},
		{
			name:           "exact division 10",
			x:              6.354,
			unit:           0.001,
			expectedResult: 6.354,
		},
		{
			name:           "exact division 11",
			x:              6.355,
			unit:           0.001,
			expectedResult: 6.355,
		},
		{
			name:           "exact division 12",
			x:              6.356,
			unit:           0.001,
			expectedResult: 6.356,
		},
		{
			name:           "exact division 13",
			x:              6.3,
			unit:           0.1,
			expectedResult: 6.3,
		},
		{
			name:           "exact division 14",
			x:              6.4,
			unit:           0.1,
			expectedResult: 6.4,
		},
		{
			name:           "exact division 15",
			x:              6.5,
			unit:           0.1,
			expectedResult: 6.5,
		},
		{
			name:           "exact division 16",
			x:              6.6,
			unit:           0.1,
			expectedResult: 6.6,
		},
		{
			name:           "exact division 17",
			x:              6.7,
			unit:           0.1,
			expectedResult: 6.7,
		},
		{
			name:           "exact division 18",
			x:              6.8,
			unit:           0.1,
			expectedResult: 6.8,
		},
		{
			name:           "exact division 19",
			x:              6.9,
			unit:           0.1,
			expectedResult: 6.9,
		},
		{
			name:           "exact division 20",
			x:              7.0,
			unit:           0.1,
			expectedResult: 7.0,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := UnitRound(testCase.x, testCase.unit)
		assert.InDelta(t, testCase.expectedResult, actualResult, testDelta, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestUnitCeil(t *testing.T) {

	testCases := []struct {
		name           string
		x              float64
		unit           float64
		expectedResult float64
	}{
		{
			name:           "case1",
			x:              2.33,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case2",
			x:              2.35,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case3",
			x:              2.36,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case4",
			x:              2.4,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case5",
			x:              2.3,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case6",
			x:              0.3398,
			unit:           0.1,
			expectedResult: 0.4,
		},
		{
			name:           "case7",
			x:              2.2398,
			unit:           0.3,
			expectedResult: 2.4,
		},
		{
			name:           "case8",
			x:              4.38732,
			unit:           0.01,
			expectedResult: 4.39,
		},
		{
			name:           "case9",
			x:              4.38432,
			unit:           0.01,
			expectedResult: 4.39,
		},
		{
			name:           "case10",
			x:              4.38432,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "case11",
			x:              4.58432,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "exact division 1",
			x:              4.58,
			unit:           0.01,
			expectedResult: 4.58,
		},
		{
			name:           "exact division 2",
			x:              4.57,
			unit:           0.01,
			expectedResult: 4.57,
		},
		{
			name:           "exact division 3",
			x:              4.5,
			unit:           0.1,
			expectedResult: 4.5,
		},
		{
			name:           "exact division 4",
			x:              4.2,
			unit:           0.1,
			expectedResult: 4.2,
		},
		{
			name:           "exact division 5",
			x:              2.2222,
			unit:           0.0001,
			expectedResult: 2.2222,
		},
		{
			name:           "exact division 6",
			x:              2.2227,
			unit:           0.0001,
			expectedResult: 2.2227,
		},
		{
			name:           "exact division 7",
			x:              4,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "exact division 8",
			x:              5,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "exact division 9",
			x:              6,
			unit:           1,
			expectedResult: 6,
		},
		{
			name:           "exact division 10",
			x:              6.354,
			unit:           0.001,
			expectedResult: 6.354,
		},
		{
			name:           "exact division 11",
			x:              6.355,
			unit:           0.001,
			expectedResult: 6.355,
		},
		{
			name:           "exact division 12",
			x:              6.356,
			unit:           0.001,
			expectedResult: 6.356,
		},
		{
			name:           "exact division 13",
			x:              6.3,
			unit:           0.1,
			expectedResult: 6.3,
		},
		{
			name:           "exact division 14",
			x:              6.4,
			unit:           0.1,
			expectedResult: 6.4,
		},
		{
			name:           "exact division 15",
			x:              6.5,
			unit:           0.1,
			expectedResult: 6.5,
		},
		{
			name:           "exact division 16",
			x:              6.6,
			unit:           0.1,
			expectedResult: 6.6,
		},
		{
			name:           "exact division 17",
			x:              6.7,
			unit:           0.1,
			expectedResult: 6.7,
		},
		{
			name:           "exact division 18",
			x:              6.8,
			unit:           0.1,
			expectedResult: 6.8,
		},
		{
			name:           "exact division 19",
			x:              6.9,
			unit:           0.1,
			expectedResult: 6.9,
		},
		{
			name:           "exact division 20",
			x:              7.0,
			unit:           0.1,
			expectedResult: 7.0,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := UnitCeil(testCase.x, testCase.unit)
		assert.InDelta(t, testCase.expectedResult, actualResult, testDelta, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestUnitFloor(t *testing.T) {

	testCases := []struct {
		name           string
		x              float64
		unit           float64
		expectedResult float64
	}{
		{
			name:           "case1",
			x:              2.33,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case2",
			x:              2.35,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case3",
			x:              2.36,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case4",
			x:              2.4,
			unit:           0.1,
			expectedResult: 2.4,
		},
		{
			name:           "case5",
			x:              2.3,
			unit:           0.1,
			expectedResult: 2.3,
		},
		{
			name:           "case6",
			x:              0.3398,
			unit:           0.1,
			expectedResult: 0.3,
		},
		{
			name:           "case7",
			x:              2.2398,
			unit:           0.3,
			expectedResult: 2.1,
		},
		{
			name:           "case8",
			x:              4.38732,
			unit:           0.01,
			expectedResult: 4.38,
		},
		{
			name:           "case9",
			x:              4.38432,
			unit:           0.01,
			expectedResult: 4.38,
		},
		{
			name:           "case10",
			x:              4.38432,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "case11",
			x:              4.58432,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "exact division 1",
			x:              4.58,
			unit:           0.01,
			expectedResult: 4.58,
		},
		{
			name:           "exact division 2",
			x:              4.57,
			unit:           0.01,
			expectedResult: 4.57,
		},
		{
			name:           "exact division 3",
			x:              4.5,
			unit:           0.1,
			expectedResult: 4.5,
		},
		{
			name:           "exact division 4",
			x:              4.2,
			unit:           0.1,
			expectedResult: 4.2,
		},
		{
			name:           "exact division 5",
			x:              2.2222,
			unit:           0.0001,
			expectedResult: 2.2222,
		},
		{
			name:           "exact division 6",
			x:              2.2227,
			unit:           0.0001,
			expectedResult: 2.2227,
		},
		{
			name:           "exact division 7",
			x:              4,
			unit:           1,
			expectedResult: 4,
		},
		{
			name:           "exact division 8",
			x:              5,
			unit:           1,
			expectedResult: 5,
		},
		{
			name:           "exact division 9",
			x:              6,
			unit:           1,
			expectedResult: 6,
		},
		{
			name:           "exact division 10",
			x:              6.354,
			unit:           0.001,
			expectedResult: 6.354,
		},
		{
			name:           "exact division 11",
			x:              6.355,
			unit:           0.001,
			expectedResult: 6.355,
		},
		{
			name:           "exact division 12",
			x:              6.356,
			unit:           0.001,
			expectedResult: 6.356,
		},
		{
			name:           "exact division 13",
			x:              6.3,
			unit:           0.1,
			expectedResult: 6.3,
		},
		{
			name:           "exact division 14",
			x:              6.4,
			unit:           0.1,
			expectedResult: 6.4,
		},
		{
			name:           "exact division 15",
			x:              6.5,
			unit:           0.1,
			expectedResult: 6.5,
		},
		{
			name:           "exact division 16",
			x:              6.6,
			unit:           0.1,
			expectedResult: 6.6,
		},
		{
			name:           "exact division 17",
			x:              6.7,
			unit:           0.1,
			expectedResult: 6.7,
		},
		{
			name:           "exact division 18",
			x:              6.8,
			unit:           0.1,
			expectedResult: 6.8,
		},
		{
			name:           "exact division 19",
			x:              6.9,
			unit:           0.1,
			expectedResult: 6.9,
		},
		{
			name:           "exact division 20",
			x:              7.0,
			unit:           0.1,
			expectedResult: 7.0,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := UnitFloor(testCase.x, testCase.unit)
		assert.InDelta(t, testCase.expectedResult, actualResult, testDelta, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}
