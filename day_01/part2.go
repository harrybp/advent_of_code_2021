package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// Sum up a slice
func sum(slice []int) int {
    result := 0
    for _, v := range slice {
        result += v
    }
    return result
}

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Read file into array of ints
    var values []int
    for scanner.Scan() {

        // Convert line to int and add to array
        line := scanner.Text()
        value, _ := strconv.Atoi(line)
        values = append(values, value);
    }

    var total int = 0
    var previous int = int(^uint(0) >> 1)

    // Check each possible slice of length 3
    for x := 0; x < len(values)-2; x++ {
        value := sum(values[x:x+3])
        if value > previous {
            total++;
        }
        previous = value
    }

    fmt.Println(total)
}


