package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines [][4]int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		txt := scanner.Text()

		s := strings.Replace(txt, " -> ", ",", -1)

		line := [4]int{}
		for i, n := range strings.Split(s, ",") {
			line[i], _ = strconv.Atoi(n)
		}

		lines = append(lines, line)
	}

	fmt.Println(CalculateLinesIntersection(lines))
}

func CalculateLinesIntersection(vents [][4]int) int {
	visitedPoints := make(map[string]int)
	overlapedPoints := make(map[string]struct{})

	for _, line := range vents {
		x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

		orientationX := 1
		orientationY := 1

		if x1 > x2 {
			orientationX = -1
		}

		if y1 > y2 {
			orientationY = -1
		}

		i := x1
		j := y1

		switch {
		case x1 == x2: // vertical line
			for ; (y2-j)*orientationY >= 0; j += 1 * orientationY {
				point := fmt.Sprintf("%d,%d", x1, j)
				visitedPoints[point]++

				if visitedPoints[point] >= 2 {
					overlapedPoints[point] = struct{}{}
				}
			}
		case y1 == y2: // horizontal line
			for ; (x2-i)*orientationX >= 0; i += 1 * orientationX {
				point := fmt.Sprintf("%d,%d", i, y1)
				visitedPoints[point]++

				if visitedPoints[point] >= 2 {
					overlapedPoints[point] = struct{}{}
				}
			}
		}
	}

	return len(overlapedPoints)
}
