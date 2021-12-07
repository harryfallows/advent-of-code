package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"
const gridSize = 5

type boardStats struct {
	numToPos map[int][2]int
	rowSum   [gridSize]int
	colSum   [gridSize]int
}

func (b *boardStats) calculateScore(num int) int {
	unmarkedSum := 0
	for k := range b.numToPos {
		unmarkedSum += k
	}
	return unmarkedSum * num
}

func (b *boardStats) markNumber(num int) bool {
	numPos := b.numToPos[num]

	if numPos != [2]int{0, 0} {
		b.rowSum[numPos[0]-1]++
		b.colSum[numPos[1]-1]++
		delete(b.numToPos, num)

		if b.rowSum[numPos[0]-1] >= gridSize || b.colSum[numPos[1]-1] >= gridSize {
			return true
		}

	}
	return false
}

func newBoardStats() *boardStats {
	newBoard := new(boardStats)
	newBoard.numToPos = make(map[int][2]int)
	return newBoard
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("File does not exist: %s", inputFile)
	}
	defer file.Close()

	fmt.Printf("--- Part One ---\n")

	scanner := bufio.NewScanner(file)
	drawnNums, err := readDrawnNumbers(scanner)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var boards []*boardStats
	for scanner.Scan() {
		newBoard := newBoardStats()
		readBoard(scanner, newBoard)
		boards = append(boards, newBoard)
	}

	bingo := false
	bingoCount := 0
	numBoards := len(boards)
	completeBoards := make([]bool, numBoards)

	for _, num := range drawnNums {
		for i, board := range boards {
			if completeBoards[i] {
				continue
			}
			bingo = board.markNumber(num)
			if bingo {
				if bingoCount == 0 {
					fmt.Printf("Answer: %d\n", board.calculateScore(num))
					fmt.Printf("--- Part Two ---\n")
				} else if bingoCount == (numBoards - 1) {
					fmt.Printf("Answer: %d\n", board.calculateScore(num))
				}
				bingoCount++
				completeBoards[i] = true
			}
		}
	}
}

func readDrawnNumbers(scanner *bufio.Scanner) ([]int, error) {
	var drawnNums []int
	scanner.Scan()

	for _, numStr := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return []int{}, fmt.Errorf("Drawn numbers contain non-integer value: %v", err)
		}
		drawnNums = append(drawnNums, num)
	}

	return drawnNums, nil
}

func readBoard(scanner *bufio.Scanner, board *boardStats) error {
	for j := 0; j < gridSize; j++ {
		scanner.Scan()
		for i, elemStr := range strings.Fields(scanner.Text()) {
			elem, err := strconv.Atoi(elemStr)
			if err != nil {
				return fmt.Errorf("Grid contains non-integer value: %v", err)
			}
			board.numToPos[elem] = [2]int{j + 1, i + 1}
		}
	}
	return nil
}
