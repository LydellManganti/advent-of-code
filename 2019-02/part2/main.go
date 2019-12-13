package main

import "fmt"

// ADD - Action to Add
const ADD = 1

// MULTIPLY - Action to Multiply
const MULTIPLY = 2

// SKIP - Action to Skip
const SKIP = 99

func main() {

	for dial1 := 0; dial1 < 100; dial1++ {
		for dial2 := 0; dial2 < 100; dial2++ {
			inputs := []int{
				1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 9, 19, 1, 5, 19, 23, 1, 6, 23, 27, 1, 27, 10, 31, 1, 31, 5, 35, 2, 10, 35, 39, 1, 9, 39, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 51, 6, 55, 1, 13, 55, 59, 2, 6, 59, 63, 1, 63, 5, 67, 2, 10, 67, 71, 1, 9, 71, 75, 1, 75, 13, 79, 1, 10, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 5, 91, 95, 2, 95, 9, 99, 1, 5, 99, 103, 1, 103, 6, 107, 2, 107, 13, 111, 1, 111, 10, 115, 2, 10, 115, 119, 1, 9, 119, 123, 1, 123, 9, 127, 1, 13, 127, 131, 2, 10, 131, 135, 1, 135, 5, 139, 1, 2, 139, 143, 1, 143, 5, 0, 99, 2, 0, 14, 0,
			}
			inputs[1] = dial1
			inputs[2] = dial2
			pointer := 0
			//fmt.Printf("Length %v\n", len(inputs))

			for key, value := range inputs {
				if key == pointer && value == SKIP {
					pointer++
					fmt.Printf("Code 99 Detected! Stopping!")
					fmt.Printf("Output: %v, Dial1: %v, Dial2: %v\n", inputs[0], dial1, dial2)
					break
				}
				if key == pointer {
					//fmt.Printf("key %v, pointer %v ", key, pointer)
					//fmt.Printf("inputLine %v\n", inputs[pointer:pointer+4])
					val, pos := processLine(inputs[pointer:pointer+4], inputs)
					updateInputs(val, pos, &inputs)
					pointer = pointer + 4
				}
			}
		}
	}

	//for _, i := range inputs {
	//	fmt.Printf("%v,", i)
	//}
}

func processLine(inputLine []int, inputs []int) (int, int) {
	var action = inputLine[0]
	var var1 = inputLine[1]
	var var2 = inputLine[2]
	var output int
	var position int
	switch action {
	case ADD:
		output = inputs[var1] + inputs[var2]
		position = inputLine[3]
		//fmt.Printf("var1:%v, var2:%v, pos:%v output:%v\n", inputs[var1], inputs[var2], position, output)
	case MULTIPLY:
		output = inputs[var1] * inputs[var2]
		position = inputLine[3]
		//fmt.Printf("var1:%v, var2:%v, pos:%v, output:%v\n", inputs[var1], inputs[var2], position, output)
	case SKIP:
		//fmt.Printf("Skipping\n")
		output = 0
		position = 0
	}
	return output, position
}

func updateInputs(value int, position int, inputs *[]int) *[]int {
	(*inputs)[position] = value
	return inputs
}
