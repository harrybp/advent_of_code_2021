package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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

func check(line []string) (bool, string) {
    var stack []string

    // Parse line until we reach an error
    for _, bracket := range line {
        if isOpenBracket(bracket) {
            stack = append(stack, bracket)
        } else if isCloseBracket(bracket) && bracketsMatch(stack[len(stack)-1], bracket) {
            stack = stack[:len(stack)-1]
        } else {
            return false, bracket
        }
    }
    return true, ""
}


func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    scores := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
    total := 0

    // Read file
    for scanner.Scan() {

        // Read line into string array
        line := scanner.Text()
        var brackets []string
        for _, bracket := range strings.Split(line, "") {
            brackets = append(brackets, bracket)
        }

        okay, bracket := check(brackets)
        if !okay {
            total += scores[bracket]
        }

    }
    fmt.Println(total)
}
