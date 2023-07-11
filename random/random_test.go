package random

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoissonRandom(t *testing.T) {
	fmt.Println("Generate 40 random following Poisson(200):")
	for i := 0; i < 40; i++ {
		fmt.Println(PoissonRandom(190, 210, 200))
	}
}

func TestNormalRandomRS(t *testing.T) {
	fmt.Println("Generate 40 random following Normal(10, 20), between -10 and 30, through rejection sampling:")
	for i := 0; i < 40; i++ {
		fmt.Println(NormalRandomRS(-10, 30, 10, 20))
	}
	fmt.Println("Generate 40 random following Normal(5, 2), between -5 and 15, through rejection sampling:")
	for i := 0; i < 40; i++ {
		fmt.Println(NormalRandomRS(-5, 15, 5, 2))
	}
}

func TestInnerNormalDensity(t *testing.T) {
	testCases := []struct {
		name           string
		x              float64
		miu            float64
		sigma          float64
		expectedResult float64
	}{
		{
			name:           "case(0, 0, 1)",
			x:              0,
			miu:            0,
			sigma:          1,
			expectedResult: 1 / math.Sqrt(2*math.Pi),
		},
		{
			name:           "case(10, 0, 1)",
			x:              10,
			miu:            0,
			sigma:          1,
			expectedResult: (1 / math.Sqrt(2*math.Pi)) * math.Pow(math.E, -50),
		},
		{
			name:           "case(5, 1, 2)",
			x:              5,
			miu:            1,
			sigma:          2,
			expectedResult: (1 / (2 * math.Sqrt(2*math.Pi))) * math.Pow(math.E, -2),
		},
		{
			name:           "case(7, -3, 10)",
			x:              7,
			miu:            -3,
			sigma:          10,
			expectedResult: (1 / (10 * math.Sqrt(2*math.Pi))) * math.Pow(math.E, -float64(1)/float64(2)),
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := normalDensity(testCase.x, testCase.miu, testCase.sigma)
		assert.Equal(t, testCase.expectedResult, actualResult, fmt.Sprintf("%s: result is not expected", testCase.name))
	}
}

func TestNormalRandomBM(t *testing.T) {
	fmt.Println("Generate 40 random following Normal(10, 20), between -10 and 30, through Box-Muller:")
	for i := 0; i < 40; i++ {
		fmt.Println(NormalRandomBM(-10, 30, 10, 20))
	}
	fmt.Println("Generate 40 random following Normal(5, 2), between -5 and 15, through Box-Muller:")
	for i := 0; i < 40; i++ {
		fmt.Println(NormalRandomBM(-5, 15, 5, 2))
	}
}

func TestRandomInt(t *testing.T) {
	fmt.Println("Generate 40 random in [10,20]:")
	for i := 0; i < 40; i++ {
		fmt.Println(RandomInt(10, 20))
	}
}

func TestRandomFloat64(t *testing.T) {
	fmt.Println("Generate 40 random float64 in [10,20):")
	for i := 0; i < 40; i++ {
		fmt.Println(RandomFloat64(10, 20))
	}
}

func TestRandomPickN(t *testing.T) {
	testCases := []struct {
		name string
		a    []int
		m    int
	}{
		{
			name: "case1",
			a:    []int{1, 3, 6, 70, 6, 2},
			m:    4,
		},
		{
			name: "case2",
			a:    []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			m:    3,
		},
		{
			name: "case3",
			a:    []int{1, 2, 3, 4, 5},
			m:    2,
		},
		{
			name: "case4",
			a:    []int{1, 2, 3, 4, 5, 3, 3, 3, 3, 3, 3, 3},
			m:    5,
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		actualResult := RandomPickN(testCase.a, testCase.m)
		t.Log(actualResult)
		assert.Equal(t, testCase.m, len(actualResult), fmt.Sprintf("%s: result is not expected", testCase.name))
	}

	testInt := []int{1, 2, 3, 4, 5, 3, 3, 3, 3, 3, 3, 3}
	t.Log(RandomPickN(testInt, 7))
}

func TestRandomMapPick(t *testing.T) {
	testCases := []struct {
		name string
		m    map[interface{}]interface{}
	}{
		{
			name: "case1",
			m: map[interface{}]interface{}{
				"aaa":  12,
				"sfsd": "asdfsd",
				231:    true,
				231.2:  false,
			},
		},
		{
			name: "case2",
			m: map[interface{}]interface{}{
				"231": 344,
				"sdf": "1111",
				2345:  "atre",
				1.1:   1.1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Logf("test: %s", testCase.name)
		pickedKey, pickedValue := RandomMapPick(testCase.m)
		t.Log(pickedKey, pickedValue)
	}

	t.Logf("\nStart to test the order.\n")
	testOrder := make(map[interface{}]interface{})
	testOrder["aaa"] = 12
	testOrder["sdf"] = 23
	testOrder["12"] = 6.6
	testOrder[11] = 2.2
	testOrder[23] = "adaf"
	testOrder[2.2] = true
	testOrder[3.3] = false
	testOrder[true] = "asfafdasfafd"
	testOrder[false] = "true"
	for len(testOrder) > 0 {
		pickedKey, pickedValue := RandomMapPick(testOrder)
		t.Log(pickedKey, pickedValue)
		delete(testOrder, pickedKey)
		t.Logf("The rest of the map: %v\n", testOrder)
	}

}

func TestExponentialRandom(t *testing.T) {
	fmt.Println("Generate 40 Exponential random float64 in [10,20]:")
	for i := 0; i < 40; i++ {
		fmt.Println(ExponentialRandom(3, 35, float64(1)/float64(15)))
	}
	fmt.Println("Generate 40 Exponential random float64 in [0,max]:")
	for i := 0; i < 40; i++ {
		fmt.Println(ExponentialRandom(0, math.MaxFloat64, float64(1)/float64(15)))
	}
}
