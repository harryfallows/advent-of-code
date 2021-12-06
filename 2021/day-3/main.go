package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("File does not exist: %s", inputFile)
	}
	defer file.Close()

	fmt.Printf("--- Part One ---\n")

	scanner := bufio.NewScanner(file)
	counts := [12]int{}
	inputLen := 0

	for scanner.Scan() {
		inputLen++
		input := scanner.Text()
		splitInput := strings.Split(input, "")

		for i, bit := range splitInput {
			if bit == "1" {
				counts[i]++
			}
		}
	}

	halfInputLen := inputLen / 2
	gammaStr := ""
	epsilonStr := ""

	for _, bitCount := range counts {
		if bitCount > halfInputLen {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	fmt.Printf("Answer: %d\n", gamma*epsilon)

	fmt.Printf("--- Part Two ---\n")

	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	inputs := make([]string, 0)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	sort.Strings(inputs)

	oxygen, err := getRating(inputs, true)
	if err != nil {
		log.Fatalf(err.Error())
	}
	co2, err := getRating(inputs, false)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Answer: %d\n", oxygen*co2)
}

func getRating(inputs []string, keepMost bool) (int64, error) {
	low := 0
	high := len(inputs)
	index := 0

	for high-low > 1 && index < 12 {
		mid := ((high - low) / 2) + low
		if inputs[mid][index] == '0' {
			for inputs[mid][index] == '0' {
				mid++
			}
			if keepMost {
				high = mid
			} else {
				low = mid
			}
		} else {
			for inputs[mid][index] == '1' {
				mid--
			}
			mid++
			if keepMost {
				low = mid
			} else {
				high = mid
			}
		}
		index++
	}

	rating, err := strconv.ParseInt(inputs[low], 2, 64)
	if err != nil {
		return 0, fmt.Errorf("Error parsing rating: %v", err)
	}
	return rating, nil
}
