package mymath

import (
	"bytes"
	"log"

	"github.com/shopspring/decimal"
)

var (
	logBuf bytes.Buffer
	logger *log.Logger
)

func init() {
	logger = log.New(&logBuf, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}

// We want to calculate  using this function, but the inaccuracy of float64 will result in some problems.
// e.g.,
// var a float64 = 2.4
// var b float64 = 0.1
// fmt.Println(a / b) // output is 23 rather than 24
// Therefore we use github.com/shopspring/decimal library to do it.
func UnitFloor(x, unit float64) float64 {
	// return math.Floor(x/unit) * unit // incorrect
	dX := decimal.NewFromFloat(x)
	dUnit := decimal.NewFromFloat(unit)

	dResult := dX.Div(dUnit).Floor().Mul(dUnit) // calculate math.Floor(x/unit) * unit
	result, exact := dResult.Float64()
	if !exact { // we do not print log here, or otherwise there will be too many logs.
		//logger.Printf("Not exact.\n x: [%g], unit: [%g], dX: [%v], dUnit: [%v], dResult: [%v], result: [%g], exact: [%v]\n", x, unit, dX, dUnit, dResult, result, exact)
		//fmt.Print(&logBuf)
	}
	return result
}

// the result of return math.Ceil(x/unit) * unit seems correct, but to be safe, we still use github.com/shopspring/decimal library here.
func UnitCeil(x, unit float64) float64 {
	// return math.Ceil(x/unit) * unit
	dX := decimal.NewFromFloat(x)
	dUnit := decimal.NewFromFloat(unit)

	dResult := dX.Div(dUnit).Ceil().Mul(dUnit) // math.Ceil(x/unit) * unit
	result, exact := dResult.Float64()
	if !exact { // we do not print log here, or otherwise there will be too many logs.
		//logger.Printf("Not exact.\n x: [%g], unit: [%g], dX: [%v], dUnit: [%v], dResult: [%v], result: [%g], exact: [%v]\n", x, unit, dX, dUnit, dResult, result, exact)
		//fmt.Print(&logBuf)
	}
	return result
}

// the result of return math.Round(x/unit) * unit seems correct, but to be safe, we still use github.com/shopspring/decimal library here.
func UnitRound(x, unit float64) float64 {
	//return math.Round(x/unit) * unit
	dX := decimal.NewFromFloat(x)
	dUnit := decimal.NewFromFloat(unit)

	dResult := dX.Div(dUnit).Round(0).Mul(dUnit) // math.Ceil(x/unit) * unit
	result, exact := dResult.Float64()
	if !exact { // we do not print log here, or otherwise there will be too many logs.
		//logger.Printf("Not exact.\n x: [%g], unit: [%g], dX: [%v], dUnit: [%v], dResult: [%v], result: [%g], exact: [%v]\n", x, unit, dX, dUnit, dResult, result, exact)
		//fmt.Print(&logBuf)
	}
	return result
}
