package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []float64{
		110756,
		132543,
		57911,
		58262,
		119938,
		58581,
		52446,
		127591,
		132449,
		82732,
		51388,
		115723,
		67376,
		61402,
		71379,
		99264,
		54697,
		120877,
		130457,
		89519,
		92846,
		121983,
		145752,
		57606,
		136613,
		74147,
		142443,
		91993,
		66409,
		71590,
		74057,
		126005,
		103231,
		104401,
		105004,
		100771,
		60204,
		125178,
		132927,
		97615,
		116662,
		91806,
		74435,
		69993,
		77268,
		124654,
		116862,
		79505,
		132479,
		104118,
		59975,
		133267,
		71379,
		136031,
		64325,
		85017,
		149922,
		148287,
		62061,
		92790,
		81205,
		74146,
		116381,
		78975,
		66557,
		74568,
		77797,
		60262,
		111913,
		53703,
		139663,
		65642,
		90693,
		105015,
		147887,
		139533,
		97861,
		68607,
		146757,
		97707,
		148185,
		87966,
		115839,
		118377,
		71123,
		82938,
		63957,
		76062,
		144141,
		138096,
		132460,
		67338,
		142338,
		76347,
		128877,
		104797,
		104637,
		107605,
		66506,
		127296,
	}
	var totalFuel float64
	for _, each := range nums {
		var totalModuleFuel *float64 = new(float64)
		moduleFuel := calculateModuleFuel(each, totalModuleFuel)
		totalFuel += *moduleFuel
	}
	fmt.Printf("Total Fuel: %f\n", totalFuel)
}

func calculateModuleFuel(mass float64, totalModuleFuel *float64) *float64 {
	moduleFuel := math.Floor(mass/3) - 2
	if moduleFuel > 0 {
		*totalModuleFuel += moduleFuel
		calculateModuleFuel(moduleFuel, totalModuleFuel)
	}
	return totalModuleFuel
}
