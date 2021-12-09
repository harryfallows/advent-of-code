package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"
const maxCoord = 999

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("File does not exist: %s", inputFile)
	}
	defer file.Close()

	fmt.Printf("--- Part One ---\n")

	scanner := bufio.NewScanner(file)
	ventMap := make([][]int, maxCoord)
	for i := range ventMap {
		ventMap[i] = make([]int, maxCoord)
	}
	overlaps := 0

	for n := 0; n < 2; n++ {
		if n == 1 {
			file.Seek(0, io.SeekStart)
			scanner = bufio.NewScanner(file)
			fmt.Printf("Answer: %d\n", overlaps)
			fmt.Printf("--- Part Two ---\n")
		}

		for scanner.Scan() {
			inputLine := strings.Fields(scanner.Text())

			startCoords, err := readCoords(inputLine[0])
			if err != nil {
				log.Panicf(err.Error())
			}
			endCoords, err := readCoords(inputLine[2])
			if err != nil {
				log.Panicf(err.Error())
			}

			xDiff := endCoords[0] - startCoords[0]
			xSign := 1
			if xDiff < 0 {
				xSign = -1
			}
			xDiff = int(math.Abs(float64(xDiff)))

			yDiff := endCoords[1] - startCoords[1]
			ySign := 1
			if yDiff < 0 {
				ySign = -1
			}
			yDiff = int(math.Abs(float64(yDiff)))

			if n == 0 {
				if xDiff == 0 {
					for j := 0; j < yDiff+1; j++ {
						ventMap[startCoords[0]][(j*ySign)+startCoords[1]]++

						if ventMap[startCoords[0]][(j*ySign)+startCoords[1]] == 2 {
							overlaps++
						}
					}
				} else if yDiff == 0 {
					for i := 0; i < xDiff+1; i++ {
						ventMap[(i*xSign)+startCoords[0]][startCoords[1]]++

						if ventMap[(i*xSign)+startCoords[0]][startCoords[1]] == 2 {
							overlaps++
						}
					}
				}
			} else if n == 1 {
				if xDiff == yDiff {
					for k := 0; k < xDiff+1; k++ {
						ventMap[(k*xSign)+startCoords[0]][(k*ySign)+startCoords[1]]++

						if ventMap[(k*xSign)+startCoords[0]][(k*ySign)+startCoords[1]] == 2 {
							overlaps++
						}
					}
				}
			}
		}
	}
	fmt.Printf("Answer: %d\n", overlaps)
}

func readCoords(coordsStr string) ([2]int, error) {
	coordsSplit := strings.Split(coordsStr, ",")
	var coords [2]int
	for i := range coordsSplit {
		var err error
		coords[i], err = strconv.Atoi(coordsSplit[i])
		if err != nil {
			return [2]int{0, 0}, fmt.Errorf("Non-integer coordinate provided: %v", err)
		}
	}
	return coords, nil
}
