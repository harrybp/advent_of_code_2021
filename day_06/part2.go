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

    var lanternFish [7]int

    // Read file
    for scanner.Scan() {

        // Extract fish from line
        line := scanner.Text()
        stringSlice := strings.Split(line, ",")
        for _, s := range stringSlice {
            i, _ := strconv.Atoi(s)
            lanternFish[i]++
        }
    }

    delay2 := 0
    delay1 := 0
    iterations := 256
    for i := 0; i < iterations; i++ {
        // lanternFish[index] <= delay1 <= delay2 <= lanternFish[index]
        index := i % 7
        temp := delay2
        delay2 = lanternFish[index]
        lanternFish[index] += delay1
        delay1 = temp
    }

    total := 0
    for _, l := range lanternFish {
        total  += l
    }

    fmt.Println(total + delay1 + delay2)
}
