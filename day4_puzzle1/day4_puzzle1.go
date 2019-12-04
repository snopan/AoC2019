package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DIGITS = 6

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
		if sameAdj(curr, prev) {
			adjacent = true
		}
	}

	return (!decrease && adjacent)
}

func lowered(curr uint8, prev uint8) bool {
	return curr < prev 
}

func sameAdj(curr uint8, prev uint8) bool {
	return curr == prev
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