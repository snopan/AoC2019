package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func convertFuel(mass int) int{
	fuel := math.Floor(float64(mass)/3) - 2
	return int(fuel)
}

func readLine(reader *bufio.Reader) string{
	str, err := reader.ReadString('\n')
	if (err != nil) {
		return "error"
	}
	str = strings.TrimSuffix(str, "\n")
	return str
} 

func main() {
	reader := bufio.NewReader(os.Stdin)
	var total_fuel int

	for {
		line := readLine(reader)
		if (line == "error") {
			break
		}
		mass, _ := strconv.Atoi(line)
		total_fuel += convertFuel(mass)
	}
	
	fmt.Println(total_fuel)
}