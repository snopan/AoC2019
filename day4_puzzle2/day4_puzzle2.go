package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for scanner.Scan() {
		text = scanner.Text()
	}

	return text
}

func checkPass(number string) bool {
	decrease, adjacent := false, false


	for i := 1; i < len(number); i ++ {
		curr, prev := number[i], number[i-1]

		if lowered(curr, prev) {
			decrease = true
		}

		if !adjacent && sameAdj(curr, prev){
			prev_edge, next_edge := false, false

			if i-2 >= 0 {
				prev_x2 := number[i-2]
				prev_edge = !sameAdj(curr, prev_x2)
			} else {
				prev_edge = true
			}

			if i+1 < len(number) {
				next := number[i+1]
				next_edge = !sameAdj(curr, next)
			} else {
				next_edge = true
			}

			adjacent = prev_edge && next_edge
		}
	}

	return (!decrease && adjacent)
}

func lowered(curr uint8, prev uint8) bool {
	return curr < prev 
}

func sameAdj(val_1 uint8, val_2 uint8) bool {
	return val_1 == val_2
}

func main() {
	var lower, upper int

	input := readInput()
	fmt.Sscanf(input, "%d-%d", &lower, &upper)

	combinations := 0
	for pass := lower; pass <= upper; pass ++ {
		
		str_pass := strconv.Itoa(pass)

		if checkPass(str_pass) {
			combinations ++
		}
	}

	fmt.Println(combinations)
}