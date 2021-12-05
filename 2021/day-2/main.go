package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	hPos := 0
	depth := 0

	for scanner.Scan() {
		command := scanner.Text()
		splitCommand := strings.Split(command, " ")

		direction := splitCommand[0]
		distance, err := strconv.Atoi(splitCommand[1])
		if err != nil {
			log.Fatalf("Command distance is a non-integer value: %s", splitCommand[1])
		}

		switch direction {
		case "forward":
			hPos += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	fmt.Printf("Answer: %d\n", hPos*depth)

	fmt.Printf("--- Part Two ---\n")

	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	hPos = 0
	depth = 0
	aim := 0

	for scanner.Scan() {
		command := scanner.Text()
		splitCommand := strings.Split(command, " ")

		direction := splitCommand[0]
		distance, err := strconv.Atoi(splitCommand[1])
		if err != nil {
			log.Fatalf("Command distance is a non-integer value: %s", splitCommand[1])
		}

		switch direction {
		case "forward":
			hPos += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	fmt.Printf("Answer: %d\n", hPos*depth)
}
