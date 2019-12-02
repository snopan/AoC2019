package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() string{
	scanner := bufio.NewScanner(os.Stdin)
	var text string 

	for scanner.Scan() {
		text = scanner.Text()
	}

	return text
}

func convertInt(str_arry []string) []int {
	var output []int

	for _, i := range str_arry {
		int, err := strconv.Atoi(i) 
		if (err != nil) {
			panic(err)
		}
		output = append(output, int)
	}
	return output
}

func readCodes(codes []int) {
	for i, _ := range codes {
		if (i % 4 == 0) {
			stop := opcode(codes, i)
			if (stop) {
				break
			}
		}
	}
}

func opcode(codes []int, op_index int) bool{
	switch ((codes)[op_index]) {
		case 1:
			add(codes, op_index)
			return false
		case 2: 
			multiply(codes, op_index)
			return false
		case 99: 
			// fmt.Println("Succesfully finished reading opcode")
			return true
		default: 
			// fmt.Println("Failed to read opcode")
			return true
	}
}

func add(codes []int, op_index int) {
	var read_1 = codes[op_index + 1]
	var read_2 = codes[op_index + 2]
	var write = codes[op_index + 3]

	codes[write] = (codes[read_1]) + (codes[read_2])
}

func multiply(codes []int, op_index int) {
	var read_1 = codes[op_index + 1]
	var read_2 = codes[op_index + 2]
	var write = codes[op_index + 3]

	codes[write] = (codes[read_1]) * (codes[read_2])
}

func run(codes []int, noun int, verb int) {
	codes[1] = noun
	codes[2] = verb

	readCodes(codes)
} 


func main() {
	input := readInput()
	input_arr := strings.Split(input, ",")

	for noun := 0; noun <= 99; noun ++ {
		for verb := 0; verb <= 99; verb ++ {

			codes := convertInt(input_arr) 
			run(codes, noun, verb)
			if (codes[0] == 19690720) {
				fmt.Println(100 * noun + verb)
			}
		}
	}
}