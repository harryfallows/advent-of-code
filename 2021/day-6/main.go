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
const partOneDays = 80
const partTwoDays = 256

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("File does not exist: %s", inputFile)
	}
	defer file.Close()

	fmt.Printf("--- Part One ---\n")

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Panicf("Input file %s is empty", inputFile)
	}
	startingFish := strings.Split(scanner.Text(), ",")

	fishCount := 0
	fishMap := make(map[int]int)

	for i := range startingFish {
		fishTimer, err := strconv.Atoi(startingFish[i])
		if err != nil {
			log.Panicf("Non-integer fish timer")
		}

		fishMap[fishTimer]++
		fishCount++
	}

	for i := 0; i < partTwoDays; i++ {
		timerTemp := fishMap[8]

		for j := 8; j > 0; j-- {
			timerTempNext := fishMap[j-1]
			fishMap[j-1] = timerTemp
			timerTemp = timerTempNext
		}

		fishMap[8] = timerTemp
		fishMap[6] += timerTemp
		fishCount += timerTemp

		if i == partOneDays-1 {
			fmt.Printf("Answer: %d\n", fishCount)
			fmt.Printf("--- Part Two ---\n")
		}
	}
	fmt.Printf("Answer: %d\n", fishCount)
}
