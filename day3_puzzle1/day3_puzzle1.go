package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Move struct {
	dir byte
	steps int
}

type Interval struct {
	dir byte
	start Point
	end Point
}

func readInput() []string{
	scanner := bufio.NewScanner(os.Stdin)
	var output []string

	for scanner.Scan() {
		text := scanner.Text()
		output = append(output, text)
	}

	return output
}

func formatPath(path []string) []Move {
	output := make([]Move, len(path))

	for i, value := range path {
		var move Move

		fmt.Sscanf(value, "%c%d", &move.dir, &move.steps)
		output[i] = move
	}

	return output
}

func pathInterval(path []Move) []Interval {
	output := make([]Interval, len(path)) 

	for i, value := range path {
		var interval Interval
		interval.dir = value.dir

		if (i != 0) {
			interval.start = output[i-1].end
		}

		interval.end.x = interval.start.x
		interval.end.y = interval.start.y
		switch (value.dir) {
			case 'R':
				interval.end.x += float64(value.steps)
				break
			case 'L':
				interval.end.x -= float64(value.steps)
				break
			case 'U':
				interval.end.y += float64(value.steps)
				break
			case 'D':
				interval.end.y -= float64(value.steps)
				break
		}

		output[i] = interval
	}

	return output
}

func checkInterval(path_1 []Interval, path_2 []Interval) {
	for _, intv_1 := range path_1 {
		for _, intv_2 := range path_2 {
			found := checkIntersect(
				intv_1.start, 
				intv_1.end,
				intv_2.start,
				intv_2.end)

			if (found) {
				intersect := findInt(
				intv_1.start, 
				intv_1.end,
				intv_2.start,
				intv_2.end)
				dist := manhatDist(intersect.x, intersect.y)
				fmt.Println(intersect, dist)
			}
		}
	}
}

func onSegment(p Point, q Point, r Point) bool {
	return (q.x <= math.Max(p.x, r.x) && q.x >= math.Min(p.x, r.x) &&
		q.y <= math.Max(p.y, r.y) && q.x >= math.Min(p.y, r.y))
}

func orientation(p Point, q Point, r Point) int{
	val := (q.y - p.y) * (r.x - q.x) -
		   (q.x - p.x) * (r.y - q.y)

	if (val == 0) {
		return 0
	}

	if (val > 0) {
		return 1
	} else {
		return 2
	}
}

func checkIntersect(p1 Point, q1 Point, p2 Point, q2 Point) bool{
	// reference: https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/

	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	if (o1 != o2 && o3 != o4) {
		return true
	}

	if (o1 == 0 && onSegment(p1, p2, q1)) {
		return true
	}

	if (o2 == 0 && onSegment(p1, q2, q1)) {
		return true
	}

	if (o3 == 0 && onSegment(p2, p1, q2)) {
		return true
	}

	if (o4 == 0 && onSegment(p2, q1, q2)) {
		return true
	}

	return false
}

func findInt(A Point, B Point, C Point, D Point) (Point) {
	// reference: https://www.geeksforgeeks.org/program-for-point-of-intersection-of-two-lines/
	a1 := B.y - A.y
	b1 := A.x - B.x
	c1 := a1*(A.x) + b1*(A.y)

	a2 := D.y - C.y
	b2 := C.x - D.x
	c2 := a2*(C.x) + b2*(C.y)

	m := a1*b2 - a2*b1

	var point Point

	point.x = (b2*c1 - b1*c2)/m
	point.y = (a1*c2 - a2*c1)/m
	return point
}

func manhatDist(x float64, y float64) float64{
	return math.Abs(x) + math.Abs(y)
}

func main() {
	input := readInput()
	path_1 := strings.Split(input[0], ",")
	path_2 := strings.Split(input[1], ",")

	fpath_1 := formatPath(path_1)
	fpath_2 := formatPath(path_2)


	ipath_1 := pathInterval(fpath_1)
	ipath_2 := pathInterval(fpath_2)

	checkInterval(ipath_1, ipath_2)

}