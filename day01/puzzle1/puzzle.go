package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputs []int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		txt := scanner.Text()
		n, err := strconv.Atoi(txt)
		if err != nil {
			panic(fmt.Sprintf("input %s is not a number", txt))
		}

		inputs = append(inputs, n)
	}

	fmt.Println(MeasureDepthIncreases(inputs))
}

func MeasureDepthIncreases(depths []int) (increases int) {
	lastDepth := depths[0]
	for _, depth := range depths[1:] {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}

	return
}
