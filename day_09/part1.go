package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func checkCell(heights [100][100]int, row int, col int) bool {
    cell := heights[row][col]
    if row > 0 && heights[row-1][col] <= cell {
        return false
    }
    if row < 99 && heights[row+1][col] <= cell {
        return false
    }
    if col > 0 && heights[row][col-1] <= cell {
        return false
    }
    if col < 99 && heights[row][col+1] <= cell {
        return false
    }
    return true
}

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var heights [100][100]int

    // Read file
    lineIndex := 0
    for scanner.Scan() {

        // Extract the digits from each line
        line := scanner.Text()

        for i, digit := range strings.Split(line, "") {
            num, _ := strconv.Atoi(digit)
            heights[lineIndex][i] = num
        }
        lineIndex++

    }

    total := 0
    for i, _ := range heights {
        for j, _ := range heights[i] {
            if checkCell(heights, i, j) {
                total += heights[i][j] + 1
            }
        }
    }

    fmt.Println(total)
}



