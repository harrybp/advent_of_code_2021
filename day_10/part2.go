package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sort"
)

func isOpenBracket(bracket string) bool {
    return bracket == "{" || bracket == "(" || bracket == "<" || bracket == "["
}

func isCloseBracket(bracket string) bool {
    return bracket == "}" || bracket == ")" || bracket == ">" || bracket == "]"
}

func bracketsMatch(bracketA string, bracketB string) bool {
    return (bracketA == "(" && bracketB == ")") ||
           (bracketA == "[" && bracketB == "]") ||
           (bracketA == "{" && bracketB == "}") ||
           (bracketA == "<" && bracketB == ">")
}

func check(line []string) int{
    var stack []string

    // Parse line until we reach an error
    for _, bracket := range line {
        if isOpenBracket(bracket) {
            stack = append(stack, bracket)
        } else if isCloseBracket(bracket) && bracketsMatch(stack[len(stack)-1], bracket) {
            stack = stack[:len(stack)-1]
        } else {
            return 0
        }
    }

    // Pop unclosed brackets off the stack and score them
    score := 0
    for i := len(stack)-1; i >= 0; i-- {
        score *= 5
        switch stack[i] {
            case "(": score += 1
            case "[": score += 2
            case "{": score += 3
            case "<": score += 4
        }
    }
    return score
}

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var scores []int

    // Read file
    for scanner.Scan() {

        // Read line into string array
        line := scanner.Text()
        var brackets []string
        for _, bracket := range strings.Split(line, "") {
            brackets = append(brackets, bracket)
        }

        score := check(brackets)
        if score > 0 {
            scores = append(scores, score)
        }
    }

    sort.Ints(scores)
    fmt.Println(scores[len(scores)/2])
}
