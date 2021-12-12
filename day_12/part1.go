package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "unicode"
)

type node struct {
    name string
    big bool
    children []*node
}

func checkExists(path []*node, node *node) bool {
    for _, item := range path {
        if item.name == node.name {
            return true
        }
    }
    return false
}

func readNodes(fileName string) map[string]*node {
    // Open file
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    r := regexp.MustCompile(`^(?P<p0>[a-zA-Z]*)-(?P<p1>[a-zA-Z]*)$`)
    allNodes := make(map[string]*node)

    // Read file
    for scanner.Scan() {

        // Extract the digits from each line
        line := scanner.Text()
        matches := r.FindStringSubmatch(line)

        // Add both nodes
        for i := 0; i < 2; i++ {
            if _, ok := allNodes[matches[i+1]]; !ok {
                newNode := new(node)
                newNode.name = matches[i+1]
                newNode.big = unicode.IsUpper(rune(matches[i+1][0]))
                allNodes[matches[i+1]] = newNode
            }
        }

        // Add links unless target is start
        if allNodes[matches[2]].name != "start" {
            allNodes[matches[1]].children = append(allNodes[matches[1]].children, allNodes[matches[2]])
        }
        if allNodes[matches[1]].name != "start" {
            allNodes[matches[2]].children = append(allNodes[matches[2]].children, allNodes[matches[1]])
        }
    }

    // No link from end node
    allNodes["end"].children = make([]*node, 0)
    return allNodes
}

func main() {

    allNodes := readNodes("input.txt")

    var finalPaths [][]*node;
    var pathStack [][]*node;
    var startPath []*node

    startPath = append(startPath, allNodes["start"])
    pathStack = append(pathStack, startPath)

    for len(pathStack) > 0 {

        // Pop off stack
        newPath := pathStack[len(pathStack)-1]
        pathStack = pathStack[:len(pathStack)-1]

        // For each child
        for _, item := range newPath[len(newPath)-1].children {

            // Create new path with child added
            childPath := make([]*node, len(newPath))
            copy(childPath, newPath)
            childPath = append(childPath, item)

            if item.name == "end" {
                finalPaths = append(finalPaths, childPath)
            } else if item.big || !checkExists(newPath, item) {

                // Create new path
                pathStack = append(pathStack, childPath)
            }
        }
    }

    fmt.Println(len(finalPaths))
}
