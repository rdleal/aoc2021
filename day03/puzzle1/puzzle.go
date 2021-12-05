package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputs []string

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		inputs = append(inputs, scanner.Text())
	}

	fmt.Println(CalculateSubmarinePowerConsumption(inputs))
}

func CalculateSubmarinePowerConsumption(reports []string) uint64 {
	var gamaRate, epsilonRate uint64

	bitSize := len(reports[0])

	oneBitsCount := make([]uint64, bitSize)

	for _, report := range reports {
		for i, bit := range report {
			if bit == '1' {
				oneBitsCount[i] += 1
			}
		}
	}

	half := uint64(len(reports) / 2)
	for i, b := bitSize-1, uint64(1); i >= 0; i-- {
		oneCount := oneBitsCount[i]
		if oneCount > half {
			gamaRate |= b
		} else {
			epsilonRate |= b
		}

		b <<= 1
	}

	return gamaRate * epsilonRate
}
