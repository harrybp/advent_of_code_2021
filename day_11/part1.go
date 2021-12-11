package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const rowSize int = 10
const colSize int = 10

func readOctopuses(fileName string) [rowSize][colSize]int {
    // Open file
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var heights [rowSize][colSize]int

    // Read file
    lineIndex := 0
    for scanner.Scan() {

        // Extract the row
        line := scanner.Text()
        for i, digit := range strings.Split(line, "") {
            num, _ := strconv.Atoi(digit)
            heights[lineIndex][i] = num
        }
        lineIndex++

    }

    return heights
}

func printOctopuses(octopuses *[rowSize][colSize]int){
    for i := 0; i < rowSize; i++ {
        for j := 0; j < colSize; j++ {
            fmt.Print(octopuses[i][j])
        }
        fmt.Println()
    }

    fmt.Println()
}

func increment(octopuses *[rowSize][colSize]int) {
    for i := 0; i < rowSize; i++ {
        for j := 0; j < colSize; j++ {
            octopuses[i][j]++
        }
    }
}

func flash(octopuses *[rowSize][colSize]int, row int, col int) int {
    flashes := 1
    octopuses[row][col] = -1
    for i := row - 1; i <= row + 1; i++ {
        for j := col - 1; j <= col + 1; j++ {
            if i < 0 || j < 0 || i >= rowSize || j >= colSize || octopuses[i][j] < 0 {
                continue
            }
            octopuses[i][j]++
            if octopuses[i][j] > 9 {
                flashes += flash(octopuses, i, j)
            }
        }
    }
    return flashes
}

func step(octopuses *[rowSize][colSize]int) int {
    flashes := 0

    // 1: Increment
    increment(octopuses)

    // 2: Flash
    for i := 0; i < rowSize; i++ {
        for j := 0; j < colSize; j++ {
            if octopuses[i][j] > 9 {
                flashes += flash(octopuses, i, j)
            }
        }
    }

    // 3: Reset flashes to 0
    for i := 0; i < rowSize; i++ {
        for j := 0; j < colSize; j++ {
            if octopuses[i][j] < 0 {
                octopuses[i][j] = 0
            }
        }
    }
    return flashes
}

func main() {
    octopuses := readOctopuses("input.txt")

    flashes := 0
    for i := 0; i < 100; i++ {
        flashes += step(&octopuses)
    }

    fmt.Println(flashes)
}
