package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Regexp for extracting all the digits
    r := regexp.MustCompile(`^(?P<d0>[a-z]*)\s(?P<d1>[a-z]*)\s(?P<d2>[a-z]*)\s(?P<d3>[a-z]*)\s(?P<d4>[a-z]*)\s(?P<d5>[a-z]*)\s(?P<d6>[a-z]*)\s(?P<d7>[a-z]*)\s(?P<d8>[a-z]*)\s(?P<d9>[a-z]*)\s.\s(?P<e0>[a-z]*)\s(?P<e1>[a-z]*)\s(?P<e2>[a-z]*)\s(?P<e3>[a-z]*)$`)

    total := 0

    // Read file
    for scanner.Scan() {

        // Extract the digits from each line
        line := scanner.Text()
        matches := r.FindStringSubmatch(line)
        for i := 11; i < 15; i++ {
            length := len(matches[i])
            if length == 2 || length == 4 || length == 3 || length == 7 {
                total++
            }
        }
    }

    fmt.Println(total)
}
