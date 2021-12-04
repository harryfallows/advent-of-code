package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	prevDepth := int(^uint(0) >> 1)
	dropCount := 0

	for scanner.Scan() {
		currDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("File contains non-integer value: %s", scanner.Text())
		}
		if currDepth > prevDepth {
			dropCount++
		}
		prevDepth = currDepth
	}

	fmt.Printf("Answer: %d\n", dropCount)

	fmt.Printf("--- Part Two ---\n")

	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	window := make([]int, 0)
	dropCount = 0

	for i := 0; scanner.Scan(); i++ {
		currDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("File contains non-integer value: %s", scanner.Text())
		}
		window = append(window, currDepth)
		if i < 3 {
			continue
		}
		if window[3] > window[0] {
			dropCount++
		}
		if i > 3 {
			window = window[1:]
		}
	}

	fmt.Printf("Answer: %d\n", dropCount)
}
