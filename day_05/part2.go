package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "regexp"
)

type point struct {
    x int
    y int
}

type line struct {
    p0 point
    p1 point
}

func getPoints(line line) []point {
    var points []point
    startX := line.p0.x
    startY := line.p0.y
    xStep := 0
    yStep := 0
    if line.p0.x < line.p1.x {
        xStep = 1
    } else if line.p1.x < line.p0.x {
        xStep = -1
    }
    if line.p0.y < line.p1.y {
        yStep = 1
    } else if line.p1.y < line.p0.y {
        yStep = -1
    }
    points = append(points, point{startX,startY})
    for (startX != line.p1.x) || (startY != line.p1.y){
        startX += xStep
        startY += yStep
        points = append(points, point{startX,startY})
    }
    return points
}

func main() {
    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // Regexp for extracting direction and distance
    r := regexp.MustCompile(`(?P<x0>\d*),(?P<y0>\d*)\s->\s(?P<x1>\d*),(?P<y1>\d*)`)

    var m map[point]int
    m = make(map[point]int)
    doubles := 0

    // Read file
    for scanner.Scan() {

        // Extract points from line
        readLine := scanner.Text()
        matches := r.FindStringSubmatch(readLine)

        x0, _ := strconv.Atoi(matches[1])
        y0, _ := strconv.Atoi(matches[2])
        x1, _ := strconv.Atoi(matches[3])
        y1, _ := strconv.Atoi(matches[4])

        line := line{point{x0,y0},point{x1,y1}}
        points := getPoints(line)

        // Add to map
        for _, p := range points {
            if _, ok := m[p]; ok {
                m[p]++
            } else {
                m[p] = 1
            }
        }
    }

    // Count up all points which are part of multiple lines
    for _, v := range m {
        if v > 1 {
            doubles++
        }
    }

    fmt.Println(doubles)
}
