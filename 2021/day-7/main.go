// I am embarassed by this solution, but I'm tired and I'm going to bed

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	if !scanner.Scan() {
		log.Panicf("Input file %s is empty", inputFile)
	}
	crabPositionsSplit := strings.Split(scanner.Text(), ",")

	posCounts := make(map[int]int)
	minPos := int(^uint(0) >> 1)
	maxPos := -minPos - 1

	for i := range crabPositionsSplit {
		crabPos, err := strconv.Atoi(crabPositionsSplit[i])
		if err != nil {
			log.Panicf("Non-integer crab position")
		}
		posCounts[crabPos]++
		if crabPos < minPos {
			minPos = crabPos
		}
		if crabPos > maxPos {
			maxPos = crabPos
		}
	}

	minFuelRequired := int(^uint(0) >> 1)
	for i := minPos; i <= maxPos; i++ {
		fuelRequired := calculateFuel(&posCounts, i)
		if fuelRequired < minFuelRequired {
			minFuelRequired = fuelRequired
		} else {
			break
		}
	}

	fmt.Printf("Answer: %d\n", minFuelRequired)

	fmt.Printf("--- Part Two ---\n")

	minFuelRequired = int(^uint(0) >> 1)
	for i := minPos; i <= maxPos; i++ {
		fuelRequired := calculateFuel2(&posCounts, i)
		if fuelRequired < minFuelRequired {
			minFuelRequired = fuelRequired
		} else {
			break
		}
	}

	fmt.Printf("Answer: %d\n", minFuelRequired)
}

func calculateFuel(posCounts *map[int]int, pos int) int {
	fuelRequired := 0

	for k, v := range *posCounts {
		fuelRequired += v * int(math.Abs(float64(k-pos)))
	}

	return fuelRequired
}

func calculateFuel2(posCounts *map[int]int, pos int) int {
	fuelRequired := 0

	for k, v := range *posCounts {
		posDiff := int(math.Abs(float64(k - pos)))
		fuelRequired += v * int((0.5*math.Pow(float64(posDiff), 2))+(0.5*float64(posDiff)))
	}

	return fuelRequired
}
