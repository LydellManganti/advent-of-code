package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 153517
	end := 630395
	//start := 111233
	//	end := 123444
	//end := 111233
	valid := 0
	for loop := start; loop <= end; loop++ {
		if isIncreasing(loop) && isDouble(loop) {
			groupings := getGroups(loop)
			if isValidGroup(groupings) {
				valid++
				fmt.Printf("Valid number %v\n", groupings)
			}
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

func getGroups(i int) [][]int {
	output := [][]int{}
	strI := strconv.Itoa(i)
	for loop := 0; loop < len(strI); loop++ {
		digitInt, _ := strconv.Atoi(strI[loop : loop+1])
		outputPos, lastGroup := getLastGroup(output)
		if len(lastGroup) == 0 {
			output = append(output, []int{digitInt})
		} else if lastGroup[len(lastGroup)-1] == digitInt {
			var updateGroup = append(lastGroup, digitInt)
			output[outputPos] = updateGroup
		} else {
			output = append(output, []int{digitInt})
		}

	}
	return output
}

func getLastGroup(groups [][]int) (int, []int) {
	if len(groups) == 0 {
		return 0, []int{}
	}
	return len(groups) - 1, groups[len(groups)-1]
}

func isValidGroup(groups [][]int) bool {
	for i := 0; i < len(groups); i++ {
		if len(groups[i]) == 2 {
			return true
		}
	}
	return false
}
