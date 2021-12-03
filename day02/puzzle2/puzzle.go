package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputs []string

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		inputs = append(inputs, scanner.Text())
	}

	fmt.Println(CalculateSubmarinePosition(inputs))
}

func CalculateSubmarinePosition(movements []string) int {
	var aim, horizontal, depth int

	for _, m := range movements {
		dir, val := parseMovement(m)

		switch dir {
		case "forward":
			horizontal += val
			depth += val * aim
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}

	return horizontal * depth
}

func parseMovement(s string) (dir string, val int) {
	f := strings.Fields(s)

	dir = f[0]
	val, _ = strconv.Atoi(f[1])

	return
}
