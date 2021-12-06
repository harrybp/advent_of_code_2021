package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var lanternFish []int

    // Read file
    for scanner.Scan() {

        // Extract fish from line
        line := scanner.Text()
        stringSlice := strings.Split(line, ",")
        for _, s := range stringSlice {
            i, _ := strconv.Atoi(s)
            lanternFish = append(lanternFish, i)
        }

    }

    iterations := 80
    for i := 0; i < iterations; i++ {

        var newFish []int
        for l := 0; l < len(lanternFish); l++ {
            if lanternFish[l] == 0 {
                newFish = append(newFish, 8)
                lanternFish[l] = 7
            }
            lanternFish[l]--

        }
        lanternFish = append(lanternFish, newFish...)
    }

    fmt.Println(len(lanternFish))
}
