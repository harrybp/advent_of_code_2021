package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
)

func readCrabs(fileName string) []int {
    // Open file
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var crabs []int

    // Read file
    for scanner.Scan() {

        // Extract crabs from line
        line := scanner.Text()
        stringSlice := strings.Split(line, ",")
        for _, s := range stringSlice {
            i, _ := strconv.Atoi(s)
            crabs = append(crabs, i)
        }
    }

    return crabs
}

func main() {

    crabs := readCrabs("input.txt")
    sort.Ints(crabs)
    median := crabs[len(crabs)/2]

    // Sum up distance of each crab from the median
    total := 0
    for _, crab := range crabs {
        fuel := 0
        if crab < median {
            fuel = median - crab
        } else if crab > median {
            fuel = crab - median
        }
        total += fuel
    }

    fmt.Println(total)
}
