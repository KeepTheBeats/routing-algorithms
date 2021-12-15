package random

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate a random number that follows Poisson distribution
// P{X=k}=lambda^k/k!*e(-lambda), k=0,1,2,...
func PoissonRandom(lambda float64) int {
	// randomly generate a probability from [0.0,1.0)
	var u float64 = rand.Float64()
	for u > 0.9999999999999999 { // prevent the program from getting stuck
		u = rand.Float64()
	}

	var k int = 0
	var p float64 = math.Pow(math.E, -lambda) // the probability that X==k
	var sumP float64 = p                      // the the probability that X<=k
	for {
		if u < sumP { // section [P{X=k-1},P{X=k}) map to result value k
			return k
		}
		p *= lambda / float64(k+1) // be deduced from "P{X=k}=lambda^k/k!*e(-lambda)"
		sumP += p
		k++
	}
}
