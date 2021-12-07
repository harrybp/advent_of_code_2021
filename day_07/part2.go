package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
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

func addCost(crab int, costs [1955]int) [1955]int {
    for i := 0; i < len(costs); i++ {
        distance := 0
        if crab < i {
            distance = i - crab
        } else if crab > i {
            distance = crab - i
        }
        costs[i] += ((distance * distance) + distance) / 2
    }
    return costs
}

func main() {

    var costs [1955]int
    crabs := readCrabs("input.txt")

    // Sum up costs for each position for every crab
    for _, crab := range crabs {
        costs = addCost(crab, costs)
    }

    // Find the lowest cost
    minimumValue := int(^uint(0) >> 1)
    for _, value := range costs {
        if value < minimumValue {
            minimumValue = value
        }
    }

    fmt.Println(minimumValue)
}


