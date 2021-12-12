package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
)

const rowSize int = 100
const colSize int = 100

func readHeights(fileName string) [rowSize][colSize]int {
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

// Recursive flood fill method, keep track of the size
func floodFill(heights [rowSize][colSize]int, size int, row int, col int) ([rowSize][colSize]int, int) {
    if heights[row][col] == 9 || heights[row][col] == -1 {
        return heights, size
    }

    heights[row][col] = -1
    size++

    if row > 0 {
        heights, size = floodFill(heights, size, row-1, col)
    }
    if row < rowSize-1 {
        heights, size = floodFill(heights, size, row+1, col)
    }
    if col > 0 {
        heights, size = floodFill(heights, size, row, col-1)
    }
    if col < colSize-1{
        heights, size = floodFill(heights, size, row, col+1)
    }
    return heights, size
}

func main() {
    heights := readHeights("input.txt")

    // Get the size of each basin
    var size int
    var sizes []int
    for i, _ := range heights {
        for j, _ := range heights[i] {
            heights, size = floodFill(heights, 0, i, j)
            if size > 0 {
                sizes = append(sizes, size)
            }
        }
    }

    // Multiply up the biggest 3
    totalSize := 1
    sort.Ints(sizes)
    for i := 1; i <= 3; i++ {
        totalSize *= sizes[len(sizes)-i]
    }

    fmt.Println(totalSize)
}
