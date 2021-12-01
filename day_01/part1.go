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

    var total int = 0
    var previous int = int(^uint(0) >> 1)

    for scanner.Scan() {

        // Convert line to int
        line := scanner.Text()
        value, _ := strconv.Atoi(line)

        if(value > previous){
            total++;
        }
        previous = value
    }

    fmt.Println(total)
}


