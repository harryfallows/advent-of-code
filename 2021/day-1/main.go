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
	window := [4]int{}
	firstSum := 0
	secondSum := 0
	ok := true
	dropCount = 0

	for i := 0; i <= 3; i++ {
		ok = scanner.Scan()
		if !ok {
			log.Fatalf("Not enough inputs")
		}
		window[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("File contains non-integer value: %s", scanner.Text())
		}
		if i < 3 {
			firstSum += window[i]
		}
		if i > 0 {
			secondSum += window[i]
		}
	}

	for ok {
		if firstSum < secondSum {
			dropCount++
		}
		ok = scanner.Scan()
		if !ok {
			break
		}
		firstSum -= window[0]
		secondSum -= window[1]
		copy(window[:3], window[1:])
		window[3], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("File contains non-integer value: %s", scanner.Text())
		}
		firstSum += window[2]
		secondSum += window[3]
	}

	fmt.Printf("Answer: %d\n", dropCount)
}
