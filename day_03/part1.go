package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    count := 0
    var bitCount [12]int

    // Read file and count up 1's for each position
    for scanner.Scan() {
        line := scanner.Text()
        for i, c := range line {
            if c == '1' {
                bitCount[i]++
            }
        }
        count++
    }

    // Construct binary strings
    binaryString    := ""
    invBinaryString := ""
    for _, c := range bitCount {
        if c > (count/2) {
            binaryString    += "1"
            invBinaryString += "0"
        } else {
            binaryString    += "0"
            invBinaryString += "1"
        }
    }

    // Convert to int for result
    gamma, _   := strconv.ParseInt(binaryString, 2, 64)
    epsilon, _ := strconv.ParseInt(invBinaryString, 2, 64)
    fmt.Println(gamma * epsilon)
}
