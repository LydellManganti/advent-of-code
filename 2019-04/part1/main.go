package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 153517
	end := 630395
	valid := 0
	for loop := start; loop <= end; loop++ {
		if isIncreasing(loop) && isDouble(loop) {
			valid++
		}
	}
	fmt.Printf("Number of valid Paswords %v\n", valid)
}

func isIncreasing(i int) bool {
	strI := strconv.Itoa(i)
	for each := range strI[:5] {
		eachInt, _ := strconv.Atoi(strI[each : each+1])
		nextInt, _ := strconv.Atoi(strI[each+1 : each+2])
		if eachInt > nextInt {
			return false
		}

	}
	return true
}

func isDouble(i int) bool {
	strI := strconv.Itoa(i)
	for each := range strI[:5] {
		currentDigit := strI[each : each+1]
		nextDigit := strI[each+1 : each+2]
		if currentDigit == nextDigit {
			return true
		}

	}
	return false
}
