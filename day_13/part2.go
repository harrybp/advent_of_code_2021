package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type point struct {
    x int
    y int
}

type fold struct {
    axis int
    loc int
}

func readData(fileName string)(map[point]bool, []fold) {
    // Open file
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    rPoint := regexp.MustCompile(`^(?P<p0>\d*),(?P<p1>\d*)$`)
    rFold  := regexp.MustCompile(`^fold\salong\s(?P<f0>[x|y])=(?P<f1>\d*)$`)

    var folds []fold
    points := make(map[point]bool)

    // Read file
    for scanner.Scan() {
        line := scanner.Text()

        // Extract points
        if strings.Contains(line, ",") {
            matches := rPoint.FindStringSubmatch(line)
            x, _ := strconv.Atoi(matches[1])
            y, _ := strconv.Atoi(matches[2])
            point := point{x, y}
            points[point] = true

        // Extract folds
        } else if strings.Contains(line, "fold") {
            matches := rFold.FindStringSubmatch(line)
            axis := 0
            if matches[1] == "y" {
                axis = 1
            }
            loc, _ := strconv.Atoi(matches[2])
            folds = append(folds, fold{axis, loc})
        }
    }

    return points, folds

}

func printPoints(points map[point]bool) {
    maxX := 0
    maxY := 0

    // Get max values for each axis
    for k, p := range points {
        if p {
            if k.x > maxX {
                maxX = k.x
            }
            if k.y > maxY {
                maxY = k.y
            }
        }
    }

    // Print points
    for i := 0; i <= maxY; i++ {
        for j := 0; j <= maxX; j++ {
            if points[point{j,i}] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func foldPoints(points map[point]bool, fold fold) map[point]bool {
    var newPoints []point
    for v, b := range points {
        if b {
            if fold.axis == 0 {
                if v.x > fold.loc {
                    newPoints = append(newPoints, point{fold.loc - (v.x-fold.loc), v.y})
                    points[v] = false
                }
            } else {
                if v.y > fold.loc {
                    newPoints = append(newPoints, point{v.x, fold.loc - (v.y-fold.loc)})
                    points[v] = false
                }
            }
        }
    }
    for _, v := range newPoints {
        points[v] = true
    }

    return points
}

func main() {
    points, folds := readData("input.txt")
    for _, fold := range folds {
        points = foldPoints(points, fold)
    }

    printPoints(points)
}
