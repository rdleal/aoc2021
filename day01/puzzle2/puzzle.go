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

	fmt.Println(MeasureDepthSlidingWindowIncreases(inputs))
}

func MeasureDepthSlidingWindowIncreases(depths []int) (increases int) {
	lastWindowSum := sumSlidingWindow(depths[:3])
	depthsLen := len(depths)

	for i := 1; i+3 <= depthsLen; i++ {
		windowSum := sumSlidingWindow(depths[i : i+3])
		if windowSum > lastWindowSum {
			increases++
		}
		lastWindowSum = windowSum
	}

	return
}

func sumSlidingWindow(s []int) int {
	return s[0] + s[1] + s[2]
}
