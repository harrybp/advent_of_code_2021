package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "math"
    "strings"
)

// Re-arrange a string into alphabetical order
func orderString(scrambled string) string {
    orderedList := strings.Split(scrambled, "")
    sort.Strings(orderedList)
    return strings.Join(orderedList, "")
}

// Calculate the digit mapping
func calculateDigits(unknown [10]string) [10]string {
    var ordered [10]string
    var c, f, b, d string

    // Get 1,4,7,8 trivially by length
    for _, digit := range unknown {
        if len(digit) == 2 {
            ordered[1] = digit
            c = string(digit[0])
            f = string(digit[1])
        } else if len(digit) == 3 {
            ordered[7] = digit
        } else if len(digit) == 4 {
            ordered[4] = digit
        } else if len(digit) == 7 {
            ordered[8] = digit
        }
    }

    // Get b/d (the two characters in 4 which are not c or f)
    foundB := false
    for _, char := range strings.Split(ordered[4], "") {
        if char != c && char != f {
            if foundB {
                d = char
            } else {
                b = char
                foundB = true
            }
        }
    }

    // Get 2,3,5 (only 2 contains c&f, only 5 contains b & d)
    no_e := ""
    for _, digit := range unknown {
        if len(digit) == 5 {
            if strings.Contains(digit, c) && strings.Contains(digit, f) {
                ordered[3] = digit;
                no_e += digit
            } else if strings.Contains(digit, b) && strings.Contains(digit, d) {
                ordered[5] = digit;
                no_e += digit
            } else {
                ordered[2] = digit;
            }
       }
    }

    // Get e (the only character not in 3 or 5
    e := "abcdefg"
    for _, char := range strings.Split(no_e, "") {
        e = strings.Replace(e, char, "", -1)
    }

    // Get 0,6,9 (only 0 does not contain d. Of 6 and 9, only 6 contains e)
    for _, digit := range unknown {
        if len(digit) == 6 {
            if !strings.Contains(digit, b) || !strings.Contains(digit, d) {
                ordered[0] = digit;
            } else if strings.Contains(digit, e) {
                ordered[6] = digit;
            } else {
                ordered[9] = digit;
            }
        }
    }

    return ordered
}

// Identify a digit using the mapping
func identify(ordered [10]string, unknown string) int {
    sortedUnknown := orderString(unknown)
    for i, digit := range ordered {
        sortedKnown := orderString(digit)
        if sortedKnown == sortedUnknown {
            return i
        }
    }
    return -1
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

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
        var unknown [10]string
        for i := 1; i <= 10; i++ {
            unknown[i-1] = orderString(matches[i])
        }

        // Calculate the mapping
        ordered := calculateDigits(unknown)

        // Identify the 4 digit number
        value := 0
        for i := 11; i <= 14; i++ {
            identity := identify(ordered, matches[i])
            value += identity * powInt(10, 14-i)
        }
        total += value
    }

    fmt.Println(total)
}
