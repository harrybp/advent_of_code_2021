package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "regexp"
)

func main() {

    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Regexp for extracting direction and distance
    r := regexp.MustCompile(`(?P<direction>[a-z]*)\s(?P<distance>\d*)`)

    xPos := 0
    yPos := 0
    aim  := 0

    // Read file
    for scanner.Scan() {

        // Extract direction and distance from line
        line := scanner.Text()
        matches := r.FindStringSubmatch(line)
        direction := matches[1]
        distance, _ := strconv.Atoi(matches[2])

        // Update position
        switch direction {
            case "forward":
                xPos += distance
                yPos += distance * aim
            case "down":
                aim += distance
            case "up":
                aim -= distance
        }
    }

    fmt.Println(xPos * yPos)
}
