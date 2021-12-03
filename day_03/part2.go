package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func findMostCommon(binaryStrings []string, position int, invert bool) string {

    // Count up number of ones in this position
    ones := 0.0
    for _, c := range binaryStrings {
        if string(c[position]) == "1" {
            ones++
        }
    }

    // Check if ones account for more than half of all
    halfLength := float64(len(binaryStrings))/2.0
    if (ones >= halfLength && !invert) || (ones < halfLength && invert) {
        return "1"
    } else {
        return "0"
    }
}

func reduce(binaryStrings []string, position int, invert bool) []string {

    // Base case
    if len(binaryStrings) < 2 {
        return binaryStrings
    }

    // Find most (or least if invert) common bit
    mostCommon := findMostCommon(binaryStrings, position, invert)

    // Filter by most common bit
    var newBinaryStrings []string
    for _, c := range binaryStrings {
        if string(c[position]) == mostCommon {
            newBinaryStrings = append(newBinaryStrings, c)
        }
    }

    return reduce(newBinaryStrings, position + 1, invert)
}

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Read file into string array
    var binaryStrings []string
    for scanner.Scan() {
        line := scanner.Text()
        binaryStrings = append(binaryStrings, line)
    }

    // Reduce
    oxygen := reduce(binaryStrings, 0, false)
    co2    := reduce(binaryStrings, 0, true)

    // Convert to ints
    oxygenDec, _ := strconv.ParseInt(oxygen[0], 2, 64)
    co2Dec, _    := strconv.ParseInt(co2[0], 2, 64)

    fmt.Println(oxygenDec * co2Dec)
}
