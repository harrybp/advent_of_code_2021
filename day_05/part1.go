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

    if line.p0.x == line.p1.x {
        if line.p0.y < line.p1.y {
            for i := line.p0.y; i <= line.p1.y; i++ {
                points = append(points, point{line.p0.x, i})
            }
        } else {
            for i := line.p1.y; i <= line.p0.y; i++ {
                points = append(points, point{line.p0.x, i})
            }
        }
    } else if line.p0.y == line.p1.y {
        if line.p0.x < line.p1.x {
            for i := line.p0.x; i <= line.p1.x; i++ {
                points = append(points, point{i, line.p0.y})
            }
        } else {
            for i := line.p1.x; i <= line.p0.x; i++ {
                points = append(points, point{i, line.p0.y})
            }
        }
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
