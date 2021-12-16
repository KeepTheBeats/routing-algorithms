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
	if !(lambda > 0) {
		return -1 // lambda must be larger than 0
	}

	// randomly generate a probability from [0.0,1.0)
	var u float64 = rand.Float64()
	for u > 0.9999999999999999 {
		// prevent the program from getting stuck (rand.Float64() can guarantee that the result is always less than 1.0, so maybe this check is unnecessary, but I am not sure.)
		u = rand.Float64()
	}

	var k int = 0
	var p float64 = math.Pow(math.E, -lambda) // the probability that X==k
	var sumP float64 = p                      // the the probability that X<=k
	for {
		if u < sumP { // section [P{X=k-1},P{X=k}) map to result value k
			return k
		}
		// P{X=k+1} = P{X=k}*lambda/(k+1)
		// deduced from "P{X=k}=lambda^k/k!*e(-lambda)"
		p *= lambda / float64(k+1)
		sumP += p
		k++
	}
}

// generate a random number that follows Normal distribution through rejection sampling, between lowerBound and upperBound
func NormalRandomRS(lowerBound, upperBound, miu, sigma float64) float64 {
	if lowerBound > upperBound || miu < lowerBound || miu > upperBound {
		return 0
	}
	peak := normalDensity(miu, miu, sigma)
	for {
		// generate a coordinate(x,y) randomly
		x := rand.Float64()*(upperBound-lowerBound) + lowerBound
		y := rand.Float64() * (peak)
		// if coordinate(x,y) is not in the area of normal density function, reject this x, otherwise, pick this x
		if y <= normalDensity(x, miu, sigma) {
			return x
		}
	}
}

// use normal density function to get f(x) from x
func normalDensity(x float64, miu float64, sigma float64) float64 {
	return 1 / (math.Sqrt(2*math.Pi) * sigma) * math.Pow(math.E, -(math.Pow(x-miu, 2))/(2*math.Pow(sigma, 2)))
}
