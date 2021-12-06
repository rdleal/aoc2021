package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	var boardSize int

	var insideBoard bool
	boardIdx := -1

	boards := make([][]int, 0)

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		text := scanner.Text()
		if text == "" {
			insideBoard = true
			boardIdx++
			boards = append(boards, []int{})
			continue
		}

		if !insideBoard {
			fields := strings.Split(text, ",")
			for _, field := range fields {
				n, _ := strconv.Atoi(field)
				numbers = append(numbers, n)
			}

			continue
		}

		var board []int
		for _, field := range strings.Fields(text) {
			n, _ := strconv.Atoi(field)
			board = append(board, n)
		}

		if boardSize == 0 {
			boardSize = len(board)
		}

		boards[boardIdx] = append(boards[boardIdx], board...)
	}

	fmt.Println(CalculateBingoFinalScore(numbers, boards, boardSize))
}

func CalculateBingoFinalScore(numbers []int, boards [][]int, boardSize int) int {
	var number, unmarkedSum int
	var bingo bool

	boardColCounter := make(map[int]map[int]int)
	boardRowCounter := make(map[int]map[int]int)

loop:
	for _, number = range numbers {
		for i, board := range boards {
			colCounter, ok := boardColCounter[i]
			if !ok {
				colCounter = make(map[int]int)
				boardColCounter[i] = colCounter
			}

			rowCounter, ok := boardRowCounter[i]
			if !ok {
				rowCounter = make(map[int]int)
				boardRowCounter[i] = rowCounter
			}

			unmarkedSum = 0

			for j, row, col := 0, 0, 0; j < len(board); j++ {
				if n := board[j]; n == number {
					rowCounter[row]++
					colCounter[col]++

					if rowCounter[row] == boardSize || colCounter[col] == boardSize {
						bingo = true
					}

					board[j] = -1 // mark position with an invalid bingo number
				}

				if n := board[j]; n != -1 {
					unmarkedSum += n
				}

				col = (col + 1) % boardSize
				if col == 0 {
					row++
				}
			}

			if bingo {
				break loop
			}
		}
	}

	return number * unmarkedSum
}
