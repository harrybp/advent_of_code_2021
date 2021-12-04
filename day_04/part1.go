package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type cell struct {
    num int
    picked bool
}

type board struct {
    numbers [][]cell
    won bool
}

func printBoard(board [][]cell) {
    for _, line := range board {
        for _, cell := range line {
            if cell.picked {
                fmt.Print("## ")
            } else {
                fmt.Printf("%2d ", cell.num)
            }
        }
        fmt.Println()
    }
}

func checkBoard(board board) bool {
    for a := 0; a < len(board.numbers); a++ {
        rowTrue := true
        colTrue := true
        for b := 0; b < len(board.numbers); b++ {
            rowTrue = rowTrue && board.numbers[a][b].picked
            colTrue = colTrue && board.numbers[b][a].picked
        }
        if rowTrue || colTrue {
            return true;
        }
    }
    return false
}

func markBoard(board board, number int) board {
    for r, _ := range board.numbers {
        for c, _ := range board.numbers[r] {
            if board.numbers[r][c].num == number {
                board.numbers[r][c].picked = true;
            }
        }
    }
    return board
}

func scoreBoard(board board) int {
    total := 0
    for _, row := range board.numbers {
        for _, cell := range row {
            if !cell.picked {
                total += cell.num
            }
        }
    }
    return total
}

func readData(filename string) ([]board, []int) {
    // Open file
    file, _ := os.Open(filename)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Data to be read
    var boards []board
    var numbers []int

    first := true
    boardIndex := -1
    lineIndex := 0

    // Read file
    for scanner.Scan() {
        line := scanner.Text()

        // Read in random bingo numbers
        if first {
            numberStrings := strings.Split(line, ",")
            for _, n := range numberStrings {
                num, _ := strconv.Atoi(n)
                numbers = append(numbers, num);
            }
        }

        // Read in bingo board
        words := strings.Fields(line)
        if len(words) == 5 {

            // Initialise new board
            if lineIndex == 0 {
                boardIndex++
                var board board
                boards = append(boards, board)
            }

            // Initialise new line
            var line []cell
            boards[boardIndex].numbers = append(boards[boardIndex].numbers, line)

            // Read numbers into line
            for _, n := range words {
                num, _ := strconv.Atoi(n)
                boards[boardIndex].numbers[lineIndex] = append(boards[boardIndex].numbers[lineIndex], cell{num, false})
            }
            lineIndex = (lineIndex + 1) % 5
        }
        first = false
    }
    return boards, numbers
}

func main() {
   boards, numbers := readData("input.txt")

    for _, n := range numbers {
        for i, _ := range boards {
            boards[i] = markBoard(boards[i], n)
            if checkBoard(boards[i]) {
                fmt.Println(scoreBoard(boards[i]) * n)
                os.Exit(0)
            }
        }
    }
}
