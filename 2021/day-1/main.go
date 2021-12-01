package main

import (
    "log"
    "strconv"
)

const inputFile = "input.txt"

func main() {
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("File does not exist: %s", inputFile)
    }
    defer file.Close()

    prevDepth := int(^uint(0) >> 1)
    dropCount := 0

    for scanner.Scan() {
        currDepth, err := strconv.Atoi(scanner.Text())
        if currDepth > prevDepth {
            dropCount += 1
        }
        prevDepth = currDepth
    }

    log.Infof("Number of measurements larger than previous measurments: %i", dropCount)
}

