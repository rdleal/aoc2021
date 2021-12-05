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

	fmt.Println(CalculateCalculateLifeSupportRating(inputs))
}

func CalculateCalculateLifeSupportRating(reports []string) uint64 {

	oxygen := filterReports(reports, 0, true)
	co2 := filterReports(reports, 0, false)

	return oxygen * co2
}

func filterReports(reports []string, bitPos uint, mostCommon bool) uint64 {
	if len(reports) == 1 {
		return uintFromBinaryString(reports[0])
	}

	bitOccurrences := make(map[byte][]string)

	for _, report := range reports {
		bit := report[bitPos]

		bitOccurrences[bit] = append(bitOccurrences[bit], report)
	}

	zeroes := bitOccurrences['0']
	ones := bitOccurrences['1']

	zeroesLen := len(zeroes)
	onesLen := len(ones)

	if zeroesLen == onesLen {
		reports = ones
		if !mostCommon {
			reports = zeroes
		}

		return filterReports(reports, bitPos+1, mostCommon)
	}

	reports = ones

	if mostCommon {
		if zeroesLen > onesLen {
			reports = zeroes
		}
	} else {
		if zeroesLen < onesLen {
			reports = zeroes
		}
	}

	return filterReports(reports, bitPos+1, mostCommon)
}

// Could've used the strconv.ParseUint(s, 2, 64), but for the sake
// of learning, I've written this function for converting a string
// representing a binary to an uint64.
func uintFromBinaryString(s string) uint64 {
	var res uint64

	biteSize := len(s) - 1

	for _, c := range s {
		if c == '1' {
			res |= 1 << biteSize
		}
		biteSize--
	}

	return res
}
