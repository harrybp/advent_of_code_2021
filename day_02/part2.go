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

    xPos := 0
    yPos := 0
    aim  := 0

    // Regexp for extracting direction and distance
    r := regexp.MustCompile(`(?P<direction>[a-z]*)\s(?P<distance>\d*)`)

    // Read file into array of ints
    for scanner.Scan() {

        // Extract direction and distance from line
        line := scanner.Text()
        matches := r.FindStringSubmatch(line)
        r.FindStringSubmatch(line)
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


